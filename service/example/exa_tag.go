package example

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"gorm.io/gorm"
)

type TagService struct{}

// @function List
// @description tag列表
// @param {}
// @return[]model.Tag, error
func (ts *TagService) List() ([]model.Tag, error) {
	tags := make([]model.Tag, 10)
	global.DB.Debug().Find(&tags)
	return tags, nil
}

// @function CreateTag
// @description 新增标签
// @param tagName string
// @return error
func (ts *TagService) CreateTag(tagName string) error {
	var tag model.Tag
	err := global.DB.Debug().Where("tag_name = ?", tagName).First(&tag).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该标签已存在")
	}

	tag.TagName = tagName
	return global.DB.Debug().Create(&tag).Error
}
