package example

import (
	"fmt"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/request"
	"github.com/qianxia/blog/model/response"
	"github.com/qianxia/blog/service/system"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BlogService struct{}

/**
* 新增博客
 */
func (bs BlogService) Save(post request.Post) error {
	// 根据post.tags[]的值查询对应的id
	tags := make([]model.Tag, 0, 4)
	var tp []string
	err := global.QX_DB.Debug().Select("id,tag_name").Where("tag_name in (?)", post.Tags).Find(&tags).Error
	global.QX_DB.Debug().Model(&model.Type{}).Select("type_name").Where("id = ?", post.TypeId).Find(&tp)

	// 构建数据
	blog := model.Blog{
		UserId:      post.UserId,
		Nickname:    post.Username,
		TypeId:      post.TypeId,
		TypeName:    tp[0],
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		Flag:        post.Flag,
		Tags:        tags,
	}

	// 开启事务
	tx := global.QX_DB.Begin()
	// 插入博客表数据以及博客标签中间表数据
	if err := tx.Debug().Create(&blog).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 更新分类表中amount字段的值
	if err := tx.Model(&model.Type{Id: blog.TypeId}).Debug().Update("amount", gorm.Expr("amount + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()

	res, err := system.SystemGroups.ElasticSearchService.Insert("blog", fmt.Sprintf("%v", blog.Id), &blog)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	return err
}

/**
* 个人博客列表展示
 */
func (bs BlogService) List(id uint64, pageNo, pageSize int) (*response.PageList, error) {
	var blogs []response.Blog
	var total int64

	offset := (pageNo - 1) * pageSize
	err := global.QX_DB.Debug().Select("id,title,updated_at,views").Where("user_id = ?", id).Offset(offset).Limit(pageSize).Find(&blogs).Error

	global.QX_DB.Debug().Model(&model.Blog{}).Where("user_id = ?", id).Count(&total)

	var pageList response.PageList

	pageList.Total = total
	pageList.PageNo = pageNo
	pageList.PageSize = pageSize

	pageList.DataList = blogs

	return &pageList, err
}

/**
* 最新推荐展示
 */
func (bs BlogService) LatestList() ([]model.Blog, error) {
	list := make([]model.Blog, 0, 4)
	err := global.QX_DB.Debug().Select("id,title").Order("updated_at DESC").Limit(5).Offset(-1).Find(&list).Error

	return list, err
}

/**
* 首页博客展示及分页
 */
func (bs BlogService) PageList(pageSize, pageNo int) (pageList response.PageList, err error) {
	var (
		total int64
		blogs []model.Blog
	)
	offset := (pageNo - 1) * pageSize
	err = global.QX_DB.Debug().Select("id,user_id,type_id,username,type_name,title,description,updated_at").Preload("Tags").
		Offset(offset).Limit(pageSize).Find(&blogs).Error

	global.QX_DB.Model(&model.Blog{}).Count(&total)

	// 将分页信息和dataList封装到pageList中
	pageList.Total = total
	pageList.PageNo = pageNo
	pageList.PageSize = pageSize

	pageList.DataList = blogs

	return
}

/**
* 博客删除
 */
func (bs BlogService) Delete(id uint64) error {
	var blog model.Blog

	err := global.QX_DB.Debug().Select("id,type_id").Where("id = ?", id).Find(&blog).Error

	// 开启事务
	tx := global.QX_DB.Begin()

	if err := tx.Debug().Select(clause.Associations).Delete(&blog).Error; err != nil {
		tx.Rollback() // // 事务回滚
		return err
	}

	if err := tx.Debug().Model(&model.Type{Id: blog.TypeId}).Update("amount", gorm.Expr("amount - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	tx.Commit()

	res, err := system.SystemGroups.ElasticSearchService.Delete("blog", fmt.Sprintf("%v", id))
	defer res.Body.Close()

	return err
}

/**
* 修改博客
 */
func (*BlogService) Update(post request.Post) error {

	err := global.QX_DB.Debug().Model(&model.Blog{Id: post.Id}).Omit("id,user_id,views").Updates(&model.Blog{
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		Flag:        post.Flag,
	}).Error

	if err != nil {
		return err
	}

	doc := map[string]interface{}{
		"doc": map[string]interface{}{
			"title":       post.Title,
			"description": post.Description,
			"content":     post.Content,
			"flag":        post.Flag,
		},
	}
	system.SystemGroups.ElasticSearchService.Update("blog", fmt.Sprintf("%v", post.Id), doc)

	return err
}

/**
* 获取博客信息
 */
func (bs BlogService) GetBlog(id uint64, avatar string) (map[string]interface{}, error) {

	var b model.Blog
	if err := global.QX_DB.Debug().Select("id,username,type_name,title,content,description,flag,views,updated_at").Preload("Tags").Where("id = ?", id).Find(&b).Error; err != nil {
		return nil, err
	}

	if err := global.QX_DB.Debug().Model(&model.Blog{Id: id}).UpdateColumn("views", gorm.Expr("views + 1")).Error; err != nil {
		return nil, err
	}

	doc := map[string]interface{}{
		"doc": map[string]interface{}{
			"views": b.Views + 1,
		},
	}
	system.SystemGroups.ElasticSearchService.Update("blog", fmt.Sprintf("%v", b.Id), doc)

	m := make(map[string]interface{}, 11)
	m["id"] = id
	m["description"] = b.Description
	m["title"] = b.Title
	m["content"] = b.Content
	m["flag"] = b.Flag
	m["views"] = b.Views
	m["updatedAt"] = b.UpdatedAt
	m["username"] = b.Nickname
	m["avatar"] = avatar
	m["typeName"] = b.TypeName
	m["tagNames"] = b.Tags

	// 返回
	return m, nil
}

/**
* 获取要编辑的博客信息
 */
func (*BlogService) GetUpdateBlog(id uint64) (model.Blog, error) {
	var b model.Blog
	err := global.QX_DB.Debug().Select("id,title,content,description,flag").Where("id = ?", id).Find(&b).Error

	// 返回
	return b, err
}
