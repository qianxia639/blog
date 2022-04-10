package system

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
)

type CommentService struct{}

func (*CommentService) Save(m map[string]interface{}) error {
	var pid = m["parent_id"]
	var c model.Comment
	if pid != "0" {
		global.QX_DB.Debug().Where("parent_id = ?", pid).First(&c)
		m["parent_id"] = c.ParentId
	}

	return global.QX_DB.Debug().Model(&model.Comment{}).Create(&m).Error
}
