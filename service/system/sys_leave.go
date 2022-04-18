package system

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
)

type LeaveService struct{}

// 显示所有留言记录
func (*LeaveService) All() ([]model.Leave, error) {
	var leaves []model.Leave
	err := global.QX_DB.Debug().Select("id,name,content").Order("created_at DESC").Find(&leaves).Error
	return leaves, err
}

// 新增留言记录
func (*LeaveService) Insert(l model.Leave) error {
	return global.QX_DB.Create(&l).Error
}

// 删除留言记录
func (*LeaveService) Delete(id uint64) error {
	return global.QX_DB.Debug().Delete(&model.Leave{}, id).Error
}
