package app

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/request"
	"github.com/qianxia/blog/response"
	"github.com/qianxia/blog/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if err := global.RY_DB.Debug().Select("id,title,updated_at").Where("user_id = ?", id).Find(&blogs).Error; err != nil {
		return nil, errors.New("查询失败")
	}

	return blogs, nil
}

// 首页博客展示及分页
func (bs BlogService) PageList(page map[string]int) (*response.PageList, error) {
	var (
		// 获取total
		total int64
		b     []model.Blog
		// 获取dataList
		blogs []response.Index
	)
	if err := global.RY_DB.Debug().Select("id,user_id,type_id,title,description,updated_at").Preload("Tags").Limit(page["pageSize"]).Offset(page["skipCount"]).Find(&b).Count(&total).Error; err != nil {
		return nil, errors.New("查询失败")
	}

	for _, v := range b {
		var users model.User
		if err := global.RY_DB.Debug().Select("username,avatar").Where("id = ?", v.UserId).Find(&users).Error; err != nil {
			return nil, errors.New("查询失败")
		}
		var types model.Type
		if err := global.RY_DB.Debug().Select("type_name").Where("id = ?", v.TypeId).Find(&types).Error; err != nil {
			return nil, errors.New("查询失败")
		}
		index := response.Index{
			Title:       v.Title,
			Description: v.Description,
			UpdatedAt:   utils.TomestampToTime(v.UpdatedAt),
			TypeName:    types.TypeName,
			Avatar:      users.Avatar,
			Username:    users.Username,
			Tags:        v.Tags,
		}
		blogs = append(blogs, index)
	}

	// 将total和dataList封装到pageList中
	var pageList response.PageList
	pageList.Total = total
	pageList.DataList = blogs
	// 返回vo
	return &pageList, nil
}

/**
博客删除
删除时除了要删除博客表中的数据以外，还要删除博客标签表中对应的数据
*/
func (bs BlogService) Delete(id int64) error {
	var blog model.Blog

	if err := global.RY_DB.Debug().Select("id,type_id").Where("id = ?", id).Find(&blog).Error; err != nil {
		return errors.New("操作失败")
	}

	// 开启事务
	tx := global.RY_DB.Begin()

	if err := tx.Debug().Select(clause.Associations).Delete(&blog).Error; err != nil {
		tx.Rollback() // // 事务回滚
		return errors.New("操作失败")
	}

	if err := tx.Debug().Model(&model.Type{Id: blog.TypeId}).Update("amount", gorm.Expr("amount - ?", 1)).Error; err != nil {
		tx.Rollback()
		return errors.New("操作失败")
	}

	// 提交事务
	tx.Commit()

	return nil
}
