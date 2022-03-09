package example

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
	"gorm.io/gorm"
)

type BlogStrategy interface {
	GetBlog(*BlogContext) (map[string]interface{}, error)
}

type Blog struct {
	context  *BlogContext
	strategy BlogStrategy
}

type BlogContext struct {
	Id uint64
}

func NewBlog(id uint64, strategy BlogStrategy) *Blog {
	return &Blog{
		context: &BlogContext{
			Id: id,
		},
		strategy: strategy,
	}
}

func (c *Blog) GetBlog() (map[string]interface{}, error) {
	return c.strategy.GetBlog(c.context)
}

type Blog1 struct{}

func (*Blog1) GetBlog(ctx *BlogContext) (map[string]interface{}, error) {
	var b model.Blog
	if err := global.QX_DB.Debug().Select("id,user_id,type_id,title,content,description,flag,views,updated_at").Preload("Tags").Where("id = ?", ctx.Id).Find(&b).Error; err != nil {
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

	if err := global.QX_DB.Debug().Model(&model.Blog{Id: ctx.Id}).Update("views", gorm.Expr("views + 1")).Error; err != nil {
		global.QX_LOG.Error(err)
		return nil, err
	}
	m := make(map[string]interface{}, 11)
	m["id"] = ctx.Id
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

type Blog2 struct{}

func (*Blog2) GetBlog(ctx *BlogContext) (map[string]interface{}, error) {
	var b model.Blog
	if err := global.QX_DB.Debug().Select("id,title,content,description,flag").Where("id = ?", ctx.Id).Find(&b).Error; err != nil {
		global.QX_LOG.Error(err)
		return nil, errors.New("查询失败")
	}

	m := make(map[string]interface{}, 5)
	m["id"] = ctx.Id
	m["description"] = b.Description
	m["title"] = b.Title
	m["content"] = b.Content
	m["flag"] = b.Flag

	// 返回
	return m, nil
}
