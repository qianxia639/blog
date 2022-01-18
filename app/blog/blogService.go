package app

import (
	"errors"
	"time"

	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/dto"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
	"github.com/qianxia/blog/vo"
)

type BlogService struct {
}

func (bs BlogService) Save(post vo.Post) error {
	Db := utils.GetDB()
	var were, shareStatment, commentabled bool

	// 修改 点赞、转载声明、评论的值
	for i := 0; i < len(post.Selected); i++ {
		if post.Selected[i] == "点赞" {
			were = true
		} else if post.Selected[i] == "转载声明" {
			shareStatment = true
		} else if post.Selected[i] == "评论" {
			commentabled = true
		}
	}
	// 根据post.tags[]的值查询对应的id
	// tags := make([]model.Tag, 4)
	var tags []model.Tag
	var tagIds []int
	if err := Db.Raw("SELECT id FROM "+command.DBTag+" WHERE tag_name in (?) FOR UPDATE", post.Tags).Scan(&tags).Error; err != nil {
		return errors.New("数据查询失败")
	}

	for _, v := range tags {
		tagIds = append(tagIds, v.Id)
	}

	// 构建数据
	blog := model.Blog{
		Id:             utils.NextId(),
		UserId:         post.UserId,
		TypeId:         post.TypeId,
		Title:          post.Title,
		Content:        post.Content,
		Flag:           post.Flag,
		Were:           were,
		ShareStatement: shareStatment,
		Commentabled:   commentabled,
		CreateTime:     model.Time(time.Now()),
		UpdateTime:     model.Time(time.Now()),
	}
	// 开启事务
	tx := Db.Begin()
	// 插入博客数据
	if err := tx.Exec("INSERT INTO "+command.DBBlog+"(id,user_id,type_id,title,content,flag,were,share_statement,commentabled,create_time,update_time) VALUES(?,?,?,?,?,?,?,?,?,?,?)",
		blog.Id, blog.UserId, blog.TypeId, blog.Title, blog.Content, blog.Flag, blog.Were, blog.ShareStatement, blog.Commentabled, blog.CreateTime, blog.UpdateTime).Error; err != nil {
		tx.Rollback() // 事务回滚
		return errors.New("数据插入失败")
	}

	//插入博客与标签关系表中的对应数据
	for i := 0; i < len(tagIds); i++ {
		if err := Db.Exec("INSERT INTO "+command.DBBlogTag+"(id,blog_id,tag_id) VALUES(?,?,?)",
			utils.NextId(), blog.Id, tagIds[i]).Error; err != nil {
			tx.Rollback() // 事务回滚
			return errors.New("数据插入失败")
		}
	}
	// 提交事务
	tx.Commit()

	return nil
}

func (bs BlogService) List(id int64) ([]dto.BlogDto, error) {
	Db := utils.GetDB()
	blogs := make([]dto.BlogDto, 10)

	if err := Db.Raw("SELECT id,title,update_time FROM "+command.DBBlog+" WHERE user_id = ?", id).Scan(&blogs).Error; err != nil {
		return nil, errors.New("查询失败")
	}

	return blogs, nil
}

func (bs BlogService) Show() ([]dto.IndexDto, error) {
	Db := utils.GetDB()

	var blogs []dto.IndexDto
	if err := Db.Raw(`SELECT b.id,b.title,b.content,b.update_time,t.type_name,u.avatar,u.username FROM t_blog b JOIN t_user u ON u.id = b.user_id JOIN t_type t ON b.type_id = t.id`).Scan(&blogs).Error; err != nil {
		return nil, errors.New("查询失败")
	}
	// var tagName []string
	// bs.TagService.Get()
	// Db.Raw("SELECT DISTINCT(t.tag_name) FROM t_blog_tag bt JOIN t_blog b ON bt.blog_id = ? JOIN t_tag t on t.id = bt.tag_id", blogs[len(blogs)-1].Id).Scan(&tagName)

	// var tagNames []model.Tag
	// Db.Raw("SELECT DISTINCT(t.tag_name) FROM t_blog_tag bt JOIN t_blog b ON bt.blog_id = ? JOIN t_tag t on t.id = bt.tag_id", id).Scan(&tagNames)
	for k, v := range blogs {
		// tagNames := bs.TagService.Get(v.Id)
		// // fmt.Println("tagNames ===> ", tagNames)
		// Db.Raw("SELECT t.id,t.tag_name FROM t_tag t JOIN t_blog_tag bt ON t.id = bt.tag_id JOIN t_blog b ON bt.blog_id = ?", v.Id).Scan(&blogs[len(blogs)-1].TagNames)
		if err := Db.Raw(`select t.id,t.tag_name from t_tag t JOIN
					(select DISTINCT(bt.tag_id) from t_blog_tag bt JOIN t_blog b ON bt.blog_id = ?) as tag
					ON t.id = tag.tag_id`, v.Id).Scan(&blogs[k].TagNames).Error; err != nil {
			return nil, errors.New("查询失败")
		}
		// // v.TagNames = tagNames
		// for _, t := range tagNames {
		// 	v.TagNames = append(v.TagNames, t.TagName)
		// }
		// fmt.Println("v.TagNames ===> ", v.TagNames)
	}
	return blogs, nil
}

func (bs BlogService) Latest() ([]model.Blog, error) {
	Db := utils.GetDB()
	var blogs []model.Blog
	if err := Db.Raw("SELECT title FROM " + command.DBBlog + " ORDER BY update_time DESC LIMIT 4").Scan(&blogs).Error; err != nil {
		return nil, errors.New("查询失败")
	}
	return blogs, nil
}

func (bs BlogService) Delete(id int64) error {
	Db := utils.GetDB()

	if err := Db.Raw("SELECT id FROM "+command.DBBlog+" WHERE id = ?", id).Scan(&model.Blog{}).Error; err != nil {
		return errors.New("操作失败")
	}

	// 开启事务
	tx := Db.Begin()

	// 删除blog表中的数据
	if err := tx.Exec("DELETE FROM "+command.DBBlog+" WHERE id = ?", id).Error; err != nil {
		tx.Rollback() // 事务回滚
		return errors.New("操作失败")
	}

	// 删除blog_tag中对应的数据
	if err := tx.Exec("DELETE FROM "+command.DBBlogTag+" WHERE blog_id = ?", id).Error; err != nil {
		tx.Rollback() // 事务回滚
		return errors.New("操作失败")
	}

	// 提交事务
	tx.Commit()

	return nil
}
