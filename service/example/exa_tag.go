package example

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
)

type TagService struct{}

func (ts *TagService) List() ([]model.Tag, error) {
	tags := make([]model.Tag, 10)
	global.QX_DB.Debug().Find(&tags)
	return tags, nil
}

func (ts *TagService) CreateTag(tagName string) error {
	var tag model.Tag
	global.QX_DB.Debug().Where("tag_name = ?", tagName).First(&tag)
	if tag.TagName == tagName {
		return errors.New("该标签已存在")
	}

	tag.TagName = tagName
	return global.QX_DB.Debug().Create(&tag).Error
}
