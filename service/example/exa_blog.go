package example

import (
	"errors"
	"fmt"

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

	if err := global.QX_DB.Debug().Select("id").Where("tag_name in (?)", post.Tags).Find(&tags).Error; err != nil {
		global.QX_LOG.Error(err)
		return errors.New("数据查询失败")
	}

	// 构建数据
	blog := model.Blog{
		UserId:      post.UserId,
		TypeId:      post.TypeId,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		Flag:        post.Flag,
		Publish:     true,
		Tags:        tags,
	}

	// 开启事务
	tx := global.QX_DB.Begin()
	// 插入博客表数据以及博客标签中间表数据
	if err := tx.Debug().Create(&blog).Error; err != nil {
		global.QX_LOG.Error(err)
		tx.Rollback()
		return errors.New("数据插入失败")
	}
	// 更新分类表中amount字段的值
	if err := tx.Model(&model.Type{Id: blog.TypeId}).Debug().Update("amount", gorm.Expr("amount + ?", 1)).Error; err != nil {
		global.QX_LOG.Error(err)
		tx.Rollback()
		return errors.New("数据插入失败")
	}
	// 提交事务
	tx.Commit()
	return nil
}

/**
* 保存博客
 */
func (bs BlogService) SaveBlog(post request.Post) error {
	// 根据post.tags[]的值查询对应的id
	tags := make([]model.Tag, 0, 4)

	if err := global.QX_DB.Debug().Select("id").Where("tag_name in (?)", post.Tags).Find(&tags).Error; err != nil {
		global.QX_LOG.Error(err)
		return errors.New("数据查询失败")
	}

	// 构建数据
	blog := model.Blog{
		UserId:      post.UserId,
		TypeId:      post.TypeId,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		Flag:        post.Flag,
		Publish:     false,
		Tags:        tags,
	}

	// 插入博客表数据以及博客标签中间表数据
	if err := global.QX_DB.Debug().Create(&blog).Error; err != nil {
		global.QX_LOG.Error(err)
		return errors.New("数据插入失败")
	}
	return nil
}

/**
* 个人博客列表展示
 */
func (bs BlogService) List(id uint64, page map[string]int) (*response.PageList, error) {
	var blogs []response.Blog
	var total int64
	if err := global.QX_DB.Debug().Select("id,title,publish,updated_at").Offset(page["offset"]).Limit(page["pageSize"]).Find(&blogs).Error; err != nil {
		global.QX_LOG.Error(err)
		return nil, errors.New("查询失败")
	}

	global.QX_DB.Debug().Model(&model.Blog{}).Count(&total)

	var pageList response.PageList

	pageList.Pagination.Total = total
	pageList.Pagination.CurrentPage = page["pageNo"]
	pageList.Pagination.PerPage = page["pageSize"]
	if int(total)/page["pageSize"] == 0 {
		pageList.Pagination.LastPage = int(total) / page["pageSize"]
	} else {
		pageList.Pagination.LastPage = int(total)/page["pageSize"] + 1
	}

	pageList.DataList = blogs

	return &pageList, nil
}

/**
* 最新推荐展示
 */
func (bs BlogService) LatestList() ([]model.Blog, error) {
	list := make([]model.Blog, 0, 4)
	if err := global.QX_DB.Debug().Select("id,title").Where("publish = ?", true).Order("updated_at DESC").Limit(5).Offset(-1).Find(&list).Error; err != nil {
		global.QX_LOG.Error(err)
		return nil, errors.New("查询失败")
	}
	return list, nil
}

/**
* 首页博客展示及分页
 */
func (bs BlogService) PageList(page map[string]int) (*response.PageList, error) {
	var (
		// 获取total
		total int64
		b     []model.Blog
		// 获取dataList
		blogs []response.Index
	)
	if err := global.QX_DB.Debug().Select("id,user_id,type_id,title,description,updated_at").Where("publish = ?", true).Preload("Tags").Offset(page["offset"]).Limit(page["pageSize"]).Find(&b).Error; err != nil {
		global.QX_LOG.Error(err)
		return nil, errors.New("查询失败")
	}

	for _, v := range b {
		var users model.User
		if err := global.QX_DB.Debug().Select("username,avatar").Where("id = ?", v.UserId).Find(&users).Error; err != nil {
			global.QX_LOG.Error(err)
			return nil, errors.New("查询失败")
		}
		var types model.Type
		if err := global.QX_DB.Debug().Select("type_name").Where("id = ?", v.TypeId).Find(&types).Error; err != nil {
			global.QX_LOG.Error(err)
			return nil, errors.New("查询失败")
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
	var pageList response.PageList
	pageList.Pagination.Total = total
	pageList.Pagination.CurrentPage = page["pageNo"]
	pageList.Pagination.PerPage = page["pageSize"]
	if int(total)/page["pageSize"] == 0 {
		pageList.Pagination.LastPage = int(total) / page["pageSize"]
	} else {
		pageList.Pagination.LastPage = int(total)/page["pageSize"] + 1
	}

	pageList.DataList = blogs

	return &pageList, nil
}

/**
* 博客删除
 */
func (bs BlogService) Delete(id int64) error {
	var blog model.Blog

	if err := global.QX_DB.Debug().Select("id,type_id").Where("id = ?", id).Find(&blog).Error; err != nil {
		global.QX_LOG.Error(err)
		return errors.New("操作失败")
	}

	// 开启事务
	tx := global.QX_DB.Begin()

	if err := tx.Debug().Select(clause.Associations).Delete(&blog).Error; err != nil {
		global.QX_LOG.Error(err)
		tx.Rollback() // // 事务回滚
		return errors.New("操作失败")
	}

	if err := tx.Debug().Model(&model.Type{Id: blog.TypeId}).Update("amount", gorm.Expr("amount - ?", 1)).Error; err != nil {
		global.QX_LOG.Error(err)
		tx.Rollback()
		return errors.New("操作失败")
	}

	// 提交事务
	tx.Commit()

	return nil
}

/**
* 修改博客
 */
func (*BlogService) Update(blog model.Blog) error {
	return nil
}

/**
* 获取博客信息
 */
func (bs BlogService) GetBlog(id uint64) (map[string]interface{}, error) {

	var b model.Blog
	if err := global.QX_DB.Debug().Select("id,user_id,type_id,title,content,description,flag,views,updated_at").Preload("Tags").Where("id = ?", id).Find(&b).Error; err != nil {
		global.QX_LOG.Error(err)
		return nil, errors.New("查询失败")
	}

	var users model.User
	if err := global.QX_DB.Debug().Select("username,avatar").Where("id = ?", b.UserId).Find(&users).Error; err != nil {
		global.QX_LOG.Error(err)
		return nil, errors.New("查询失败")
	}
	var types model.Type
	if err := global.QX_DB.Debug().Select("type_name").Where("id = ?", b.TypeId).Find(&types).Error; err != nil {
		global.QX_LOG.Error(err)
		return nil, errors.New("查询失败")
	}

	if err := global.QX_DB.Debug().Model(&model.Blog{Id: id}).Update("views", gorm.Expr("views + 1")).Error; err != nil {
		global.QX_LOG.Error(err)
		return nil, err
	}
	m := make(map[string]interface{}, 11)
	m["id"] = fmt.Sprintf("%v", id)
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
