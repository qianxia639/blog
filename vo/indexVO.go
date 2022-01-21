package vo

import (
	"github.com/qianxia/blog/model"
)

type IndexVO struct {
	Id         string      `json:"blog_id"`
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	UpdateTime model.Time  `json:"update_time"`
	TypeName   string      `json:"type_name"`
	Avatar     string      `json:"avatar"`
	Username   string      `json:"username"`
	TagNames   []model.Tag `json:"tag_names"`
}
