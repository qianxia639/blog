package response

import (
	"time"

	"github.com/qianxia/blog/model"
)

type Blog struct {
	Id        uint64 `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Views     uint32 `json:"views,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

type BlogResult struct {
	Id        uint64      `json:"id,omitempty"`
	Views     uint32      `json:"views,omitempty"`
	Nickname  string      `json:"nickname,omitempty"`
	TypeName  string      `json:"typeName,omitempty"`
	Content   string      `json:"content,omitempty"`
	Title     string      `json:"title,omitempty"`
	Avatar    string      `json:"avatar,omitempty"`
	Flag      string      `json:"flag,omitempty"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty"`
	Tags      []model.Tag `json:"Tags,omitempty"`
}
