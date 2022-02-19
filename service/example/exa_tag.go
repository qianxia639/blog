package example

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
)

type TagService struct{}

func (ts *TagService) List() ([]string, error) {
	tags := make([]model.Tag, 0, 10)
	var tagNames []string

	if err := global.RY_DB.Debug().Select("id,tag_name").Find(&tags).Error; err != nil {
		return nil, errors.New("查询失败")
	}

	for _, v := range tags {
		tagNames = append(tagNames, v.TagName)
	}

	return tagNames, nil
}
