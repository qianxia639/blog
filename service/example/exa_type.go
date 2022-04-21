package example

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
)

type TypeService struct{}

// 查詢type列表，按amount降序排列
func (ts *TypeService) ListOrderByAmountDesc() ([]model.Type, error) {
	types := make([]model.Type, 0, 4)
	val, err := global.QX_REDIS.Get(context.Background(), "type").Result()
	if err != nil {
		err = global.QX_DB.Debug().Select("id,type_name,amount").Order("amount DESC").Limit(4).Find(&types).Error
		if err != nil {
			return nil, err
		}
		by, err := json.Marshal(&types)
		if err != nil {
			return nil, err
		}

		err = global.QX_REDIS.Set(context.Background(), "type", by, 30*time.Second).Err()
	}
	json.Unmarshal([]byte(val), &types)
	// err := global.QX_DB.Debug().Select("id,type_name,amount").Order("amount DESC").Limit(4).Find(&types).Error
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

// 新增分类
func (ts *TypeService) CreateType(typeName string) error {

	var t []string

	global.QX_DB.Debug().Model(&model.Type{}).Select("type_name").Where("type_name = ?", typeName).Find(&t)

	if len(t) != 0 {
		return errors.New("该分类已存在")
	}

	return global.QX_DB.Debug().Model(&model.Type{}).Create(&model.Type{TypeName: typeName}).Error
}
