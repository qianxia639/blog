package system

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
)

type LeaveService struct{}

// 显示所有留言记录
func (*LeaveService) All() ([]model.Leave, error) {
	var leaves []model.Leave
	err := global.QX_DB.Debug().Select("id,content").Find(&leaves).Error
	return leaves, err
}

// 新增留言记录
func (*LeaveService) Insert(l model.Leave) error {
	return global.QX_DB.Create(&l).Error
}
