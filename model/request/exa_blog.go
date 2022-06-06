package request

import "github.com/qianxia/blog/model"

type Post struct {
	Id       uint64
	TypeId   uint16
	UserId   uint64
	Nickname string
	Title    string
	Content  string
	Flag     string
	Tags     []model.Tag
}
