package dto

import (
	"fmt"
	"time"

	"github.com/qianxia/blog/model"
)

type BlogDto struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	UpdateTime time.Time `json:"update_time"`
}

func ToBlogDto(blog model.Blog) BlogDto {
	return BlogDto{
		Id:         fmt.Sprintf("%v", blog.Id),
		Title:      blog.Title,
		UpdateTime: blog.UpdateTime,
	}
}
