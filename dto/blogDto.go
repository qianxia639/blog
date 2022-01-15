package dto

import "github.com/qianxia/blog/model"

type BlogDto struct {
	Id         int64      `json:"id"`
	Title      string     `json:"title"`
	UpdateTime model.Time `json:"update_time"`
}

func ToBlogDto(blog model.Blog) BlogDto {
	return BlogDto{
		Id:         blog.Id,
		Title:      blog.Title,
		UpdateTime: blog.UpdateTime,
	}
}
