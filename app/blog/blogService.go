package app

import (
	"errors"

	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/request"
	"github.com/qianxia/blog/response"
	"github.com/qianxia/blog/utils"
	"github.com/qianxia/blog/vo"
	"gorm.io/gorm"
)

type BlogService struct {
}

/**
新增博客
根据前端所选的额外选项进行相应的变化，
进行博客新增的数据插入时，不仅要在博客表中新增数据，还要在博客标签表中进行数据的插入
*/

func (bs BlogService) Save(post request.Post) error {
	var were, shareStatment, enableComment bool

	// 修改 点赞、转载声明、评论的值
	for i := 0; i < len(post.Selected); i++ {
		if post.Selected[i] == "点赞" {
			were = true
		} else if post.Selected[i] == "转载声明" {
			shareStatment = true
		} else if post.Selected[i] == "评论" {
			enableComment = true
		}
	}
	// 根据post.tags[]的值查询对应的id
	// tags := make([]model.Tag, 4)
	tags := make([]model.Tag, 4)

	if err := global.RY_DB.Debug().Select("id").Where("tag_name in (?)", post.Tags).Find(&tags).Error; err != nil {
		return errors.New("数据查询失败")
	}

	// 构建数据
	blog := model.Blog{
		Id:             utils.NextId(),
		UserId:         post.UserId,
		TypeId:         post.TypeId,
		Title:          post.Title,
		Description:    post.Description,
		Content:        post.Content,
		Flag:           post.Flag,
		Were:           were,
		ShareStatement: shareStatment,
		EnableComment:  enableComment,
		Tags:           tags,
	}

	// 开启事务
	tx := global.RY_DB.Begin()
	// 插入博客表数据以及博客标签中间表数据
	if err := tx.Debug().Create(&blog).Error; err != nil {
		tx.Rollback()
		return errors.New("数据插入失败")
	}
	// 更新分类表中amount字段的值
	if err := tx.Model(&model.Type{Id: blog.TypeId}).Debug().Update("amount", gorm.Expr("amount + ?", 1)).Error; err != nil {
		tx.Rollback()
		return errors.New("数据插入失败")
	}
	// 提交事务
	tx.Commit()
	return nil
}

// 个人博客列表展示
func (bs BlogService) List(id int64) ([]response.Blog, error) {
	blogs := make([]response.Blog, 10)
	// type Blog struct {
	// 	Id         string
	// 	Title      string
	// 	UpdateTime model.Time
	// }
	// blogs := make([]Blog, 10)
	if err := global.RY_DB.Debug().Select("id,title,updated_at").Where("user_id = ?", id).Find(&blogs).Error; err != nil {
		return nil, errors.New("查询失败")
	}

	return blogs, nil
}

// 首页博客展示及分页
func (bs BlogService) PageList(pageMap map[string]int) (*vo.PageListVO, error) {
	// 获取total
	var total int64
	if err := global.RY_DB.Table(command.DBBlog).Count(&total).Error; err != nil {
		return nil, errors.New("操作失败")
	}
	// 获取dataList
	var blogs []vo.IndexVO
	if err := global.RY_DB.Raw(`SELECT b.id,b.title,b.content,b.update_time,t.type_name,u.avatar,u.username
						FROM t_blog b JOIN t_user u ON u.id = b.user_id JOIN t_type t ON b.type_id = t.id LIMIT ?,?`, pageMap["skipCount"], pageMap["pageSize"]).Scan(&blogs).Error; err != nil {
		return nil, errors.New("查询失败")
	}
	for k, v := range blogs {
		if err := global.RY_DB.Raw(`SELECT t.id,t.tag_name FROM t_tag t JOIN
					(SELECT DISTINCT(bt.tag_id) FROM t_blog_tag bt JOIN t_blog b ON bt.blog_id = ?) as tag
					ON t.id = tag.tag_id`, v.Id).Scan(&blogs[k].TagNames).Error; err != nil {
			return nil, errors.New("查询失败")
		}
	}
	// 将total和dataList封装到pageListVO中
	vo := vo.PageListVO{}
	vo.Total = total
	vo.DataList = blogs
	// 返回vo
	return &vo, nil
}

/**
博客删除
删除时除了要删除博客表中的数据以外，还要删除博客标签表中对应的数据
*/

func (bs BlogService) Delete(id int64) error {
	blog := new(model.Blog)

	if err := global.RY_DB.Raw("SELECT id,type_id FROM "+command.DBBlog+" WHERE id = ?", id).Scan(&blog).Error; err != nil {
		return errors.New("操作失败")
	}

	// 开启事务
	tx := global.RY_DB.Begin()

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

	// 分类表中对应的条数要-1
	if err := tx.Exec("UPDATE "+command.DBType+" SET amount = amount - 1 WHERE id = ?", blog.TypeId).Error; err != nil {
		tx.Rollback() // 事务回滚
		return errors.New("数据更新失败")
	}

	// 提交事务
	tx.Commit()

	return nil
}
