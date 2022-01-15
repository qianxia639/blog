package model

type Tag struct {
	// 主键
	Id int `json:"id" binding:"required"`
	// 标签名
	TagName string `json:"tag_name" binding:"required"`
}
