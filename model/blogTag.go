package model

type BlogTag struct {
	Id     int   `json:"id"`
	BlogId int64 `json:"blog_id"`
	TagId  int   `json:"tag_id"`
}
