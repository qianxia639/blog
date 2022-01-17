package dto

import (
	"github.com/qianxia/blog/model"
)

type IndexDto struct {
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	UpdateTime model.Time `json:"update_time"`
}
