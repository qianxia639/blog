package request

import "github.com/qianxia/blog/model"

type SaveBlog struct {
	TypeId  uint16
	Title   string
	Content string
	Flag    string
	Tags    []model.Tag
}

type UpdateBlog struct {
	Id      uint64
	TypeId  uint16
	Title   string
	Content string
	Flag    string
	Tags    []model.Tag
}
