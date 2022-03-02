package example

import (
	"errors"
	"fmt"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/response"
	"github.com/qianxia/blog/utils"
)

type TypeService struct{}

// 查詢type列表，按amount降序排列
func (ts *TypeService) ListOrderByAmountDesc() ([]model.Type, error) {
	types := make([]model.Type, 0, 4)
	if err := global.RY_DB.Debug().Select("id,type_name,amount").Order("amount DESC").Find(&types).Error; err != nil {
		global.RY_LOG.Error(err)
		return nil, errors.New("查询失败")
	}
	return types, nil
}

// 只显示分类列表不排序
func (ts *TypeService) List() ([]model.Type, error) {
	types := make([]model.Type, 0, 10)
	if err := global.RY_DB.Debug().Select("id,type_name,amount").Find(&types).Error; err != nil {
		global.RY_LOG.Error(err)
		return nil, errors.New("查询失败")
	}
	return types, nil
}

func (ts *TypeService) TypeList(id int) ([]response.Index, error) {
	var (
		// 获取total
		total int64
		b     []model.Blog
		// 获取dataList
		blogs []response.Index
	)
	if err := global.RY_DB.Debug().Select("id,user_id,type_id,title,description,updated_at").Preload("Tags").Where("type_id = ?", id).Find(&b).Count(&total).Error; err != nil {
		global.RY_LOG.Error(err)
		return nil, errors.New("查询失败")
	}

	for _, v := range b {
		var users model.User
		if err := global.RY_DB.Debug().Select("username,avatar").Where("id = ?", v.UserId).Find(&users).Error; err != nil {
			global.RY_LOG.Error(err)
			return nil, errors.New("查询失败")
		}
		var types model.Type
		if err := global.RY_DB.Debug().Select("type_name").Where("id = ?", v.TypeId).Find(&types).Error; err != nil {
			global.RY_LOG.Error(err)
			return nil, errors.New("查询失败")
		}
		index := response.Index{
			Id:          fmt.Sprintf("%v", v.Id),
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

	// 将total和dataList封装到pageList中
	var pageList response.PageList
	pageList.Total = total
	pageList.DataList = blogs

	return blogs, nil
}
