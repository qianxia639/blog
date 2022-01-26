package response

import (
	"github.com/qianxia/blog/model"
)

type Index struct {
	Id        string      `json:"blog_id"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	UpdatedAt model.Time  `json:"updated_at"`
	TypeName  string      `json:"type_name"`
	Avatar    string      `json:"avatar"`
	Username  string      `json:"username"`
	TagNames  []model.Tag `json:"tag_names"`
}
