package response

import (
	"github.com/qianxia/blog/model"
)

type Index struct {
	Id          string      `json:"blogId"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	UpdatedAt   string      `json:"updatedAt"`
	TypeName    string      `json:"typeName"`
	Avatar      string      `json:"avatar"`
	Username    string      `json:"username"`
	Tags        []model.Tag `json:"tagNames"`
}
