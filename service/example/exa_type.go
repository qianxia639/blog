package example

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
)

type TypeService struct{}

// 查詢type列表，按amount降序排列
func (ts *TypeService) ListOrderByAmountDesc() ([]model.Type, error) {
	types := make([]model.Type, 0, 4)

	err := global.QX_DB.Debug().Select("id,type_name,amount").Order("amount DESC").Find(&types).Error
	return types, err
}

// 只显示分类列表不排序
func (ts *TypeService) List() ([]model.Type, error) {
	types := make([]model.Type, 0, 10)
	err := global.QX_DB.Debug().Select("id,type_name,amount").Find(&types).Error

	return types, err
}

// 点击分类进行查询并分页
func (ts *TypeService) TypeList(id, pageSize, pageNum int) (response.PageList, error) {
	var (
		total int64
		blogs []model.Blog
	)

	err := global.QX_DB.Debug().Select("id,user_id,type_id,username,type_name,title,description,updated_at").Preload("Tags").Where("type_id = ?", id).
		Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&blogs).Count(&total).Error

	// 将分页信息和dataList封装到pageList中
	var pageList response.PageList
	pageList.Total = total
	pageList.PageNum = pageNum
	pageList.PageSize = pageSize

	pageList.DataList = blogs

	return pageList, err
}
