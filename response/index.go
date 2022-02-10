package response

import (
	"time"

	"github.com/qianxia/blog/model"
)

type Index struct {
	Id        string      `json:"blogId"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	UpdatedAt time.Time   `json:"updatedAt"`
	TypeName  string      `json:"typeName"`
	Avatar    string      `json:"avatar"`
	Username  string      `json:"username"`
	TagNames  []model.Tag `json:"tagNames"`
}
