package example

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
	"gorm.io/gorm"
)

type TypeService struct{}

// 查詢type列表，按amount降序排列
func (ts *TypeService) ListOrderByAmountDesc() ([]model.Type, error) {
	types := make([]model.Type, 5)
	global.QX_DB.Debug().Order("amount DESC").Limit(5).Find(&types)
	return types, nil
}

// 只显示分类列表不排序
func (ts *TypeService) List() ([]model.Type, error) {
	types := make([]model.Type, 10)
	global.QX_DB.Debug().Find(&types)

	return types, nil
}

// 按分类查询博客并分页
func (ts *TypeService) TypePageList(id, pageSize, pageNo int) (response.PageList, error) {

	var total int64
	var blogs []model.Blog

	offset := (pageNo - 1) * pageSize
	err := global.QX_DB.Debug().Select("id,user_id,type_id,nickname,type_name,title,updated_at").Preload("Tags").Where("type_id = ?", id).
		Offset(offset).Limit(pageSize).Find(&blogs).Count(&total).Error

	// 将分页信息和dataList封装到pageList中
	var pageList response.PageList
	pageList.Total = total
	pageList.PageNo = pageNo
	pageList.PageSize = pageSize

	pageList.DataList = blogs

	return pageList, err
}

// 新增分类
func (ts *TypeService) CreateType(typeName string) error {

	var tp model.Type
	err := global.QX_DB.Debug().Where("type_name = ?", typeName).First(&tp).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该分类已存在")
	}

	tp.TypeName = typeName
	return global.QX_DB.Debug().Create(&tp).Error
}
