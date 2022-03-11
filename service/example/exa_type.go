package example

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
	"github.com/qianxia/blog/utils"
)

type TypeService struct{}

// 查詢type列表，按amount降序排列
func (ts *TypeService) ListOrderByAmountDesc() ([]model.Type, error) {
	types := make([]model.Type, 0, 4)
	err := global.QX_DB.Debug().Select("id,type_name,amount").Preload("Blogs").Order("amount DESC").Find(&types).Error

	return types, err
}

// 只显示分类列表不排序
func (ts *TypeService) List() ([]model.Type, error) {
	types := make([]model.Type, 0, 10)
	err := global.QX_DB.Debug().Select("id,type_name,amount").Find(&types).Error

	return types, err
}

// 点击分类进行查询并分页
func (ts *TypeService) TypeList(id, pageSize, pageNo int) (*response.PageList, error) {
	var (
		// 获取total
		total int64
		b     []model.Blog
		// 获取dataList
		blogs []response.Index
	)

	if err := global.QX_DB.Debug().Select("id,user_id,type_id,title,description,updated_at").Preload("Tags").Where("type_id = ? AND publish = ?", id, true).
		Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&b).Count(&total).Error; err != nil {
		return nil, err
	}

	for _, v := range b {
		var users model.User
		if err := global.QX_DB.Debug().Select("username,avatar").Where("id = ?", v.UserId).Find(&users).Error; err != nil {
			return nil, err
		}
		var types model.Type
		if err := global.QX_DB.Debug().Select("type_name").Where("id = ?", v.TypeId).Find(&types).Error; err != nil {
			return nil, err
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

	// 将total和dataList封装到pageList中
	var pageList response.PageList
	pageList.Pagination.Total = total
	pageList.Pagination.PerPage = pageSize
	pageList.Pagination.CurrentPage = pageNo

	if int(total)/pageSize == 0 {
		pageList.Pagination.LastPage = int(total) / pageSize
	} else {
		pageList.Pagination.LastPage = int(total)/pageSize + 1
	}

	pageList.DataList = blogs

	return &pageList, nil
}
