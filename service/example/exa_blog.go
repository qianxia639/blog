package example

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/request"
	"github.com/qianxia/blog/model/response"
	"github.com/qianxia/blog/utils"
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

	err := global.QX_DB.Debug().Select("id").Where("tag_name in (?)", post.Tags).Find(&tags).Error

	// 构建数据
	blog := model.Blog{
		UserId:      post.UserId,
		TypeId:      post.TypeId,
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
	return err
}

/**
* 个人博客列表展示
 */
func (bs BlogService) List(id uint64, page map[string]int) (*response.PageList, error) {
	var blogs []response.Blog
	var total int64

	err := global.QX_DB.Debug().Select("id,title,updated_at,views").Where("user_id = ?", id).Offset(page["offset"]).Limit(page["pageSize"]).Find(&blogs).Error

	global.QX_DB.Debug().Model(&model.Blog{}).Where("user_id = ?", id).Count(&total)

	var pageList response.PageList

	pageList.Total = total
	pageList.PageNum = page["pageNum"]
	pageList.PageSize = page["pageSize"]

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
func (bs BlogService) PageList(page map[string]int) (pageList response.PageList, err error) {
	var (
		// 获取total
		total int64
		b     []model.Blog
		// 获取dataList
		blogs []response.Index
	)
	err = global.QX_DB.Debug().Select("id,user_id,type_id,title,description,updated_at").Preload("Tags").Offset(page["offset"]).Limit(page["pageSize"]).Find(&b).Error

	for _, v := range b {
		var users model.User
		if err = global.QX_DB.Debug().Select("username,avatar").Where("id = ?", v.UserId).Find(&users).Error; err != nil {
			return
		}
		var types model.Type
		if err = global.QX_DB.Debug().Select("type_name").Where("id = ?", v.TypeId).Find(&types).Error; err != nil {
			return
		}

		index := response.Index{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			UpdatedAt:   utils.TimestampToString(v.UpdatedAt),
			TypeName:    types.TypeName,
			Avatar:      users.Avatar,
			Username:    users.Username,
			Tags:        v.Tags,
		}
		blogs = append(blogs, index)
	}

	global.QX_DB.Model(&model.Blog{}).Count(&total)
	// 将total和dataList封装到pageList中
	pageList.Total = total
	pageList.PageNum = page["pageNum"]
	pageList.PageSize = page["pageSize"]

	pageList.DataList = blogs

	return
}

/**
* 博客删除
 */
func (bs BlogService) Delete(id int64) error {
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

	return err
}

/**
* 修改博客
 */
func (*BlogService) Update(post request.Post) error {

	return global.QX_DB.Debug().Model(&model.Blog{Id: post.Id}).Omit("id,user_id,views").Updates(&model.Blog{
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		Flag:        post.Flag,
	}).Error
}

/**
* 获取博客信息
 */
func (bs BlogService) GetBlog(id uint64) (map[string]interface{}, error) {

	var b model.Blog
	if err := global.QX_DB.Debug().Select("id,user_id,type_id,title,content,description,flag,views,updated_at").Preload("Tags").Where("id = ?", id).Find(&b).Error; err != nil {
		return nil, err
	}

	var users model.User
	if err := global.QX_DB.Debug().Select("username,avatar").Where("id = ?", b.UserId).Find(&users).Error; err != nil {
		return nil, err
	}
	var types model.Type
	if err := global.QX_DB.Debug().Select("type_name").Where("id = ?", b.TypeId).Find(&types).Error; err != nil {
		return nil, err
	}

	if err := global.QX_DB.Debug().Model(&model.Blog{Id: id}).UpdateColumn("views", gorm.Expr("views + 1")).Error; err != nil {
		return nil, err
	}
	m := make(map[string]interface{}, 11)
	m["id"] = id
	m["description"] = b.Description
	m["title"] = b.Title
	m["content"] = b.Content
	m["flag"] = b.Flag
	m["views"] = b.Views
	m["updatedAt"] = utils.TimestampToString(b.UpdatedAt)
	m["username"] = users.Username
	m["avatar"] = users.Avatar
	m["typeName"] = types.TypeName
	m["tagNames"] = b.Tags

	// 返回
	return m, nil
}

/**
* 获取要编辑的博客信息
 */
func (*BlogService) GetUpdateBlog(id uint64) (map[string]interface{}, error) {
	var b model.Blog
	err := global.QX_DB.Debug().Select("id,title,content,description,flag").Where("id = ?", id).Find(&b).Error

	m := make(map[string]interface{}, 6)
	m["id"] = id
	m["description"] = b.Description
	m["title"] = b.Title
	m["content"] = b.Content
	m["flag"] = b.Flag

	// 返回
	return m, err
}
