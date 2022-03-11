package example

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
)

type TagService struct{}

func (ts *TagService) List() ([]string, error) {
	tags := make([]model.Tag, 0, 10)
	var tagNames []string

	err := global.QX_DB.Debug().Select("id,tag_name").Find(&tags).Error

	for _, v := range tags {
		tagNames = append(tagNames, v.TagName)
	}

	return tagNames, err
}
