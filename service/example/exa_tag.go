package example

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
)

type TagService struct{}

func (ts *TagService) List() ([]model.Tag, error) {
	tags := make([]model.Tag, 10)

	global.QX_DB.Debug().Find(&tags)
	return tags, nil
}
