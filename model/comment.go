package model

import "time"

type Comment struct {
	// 主键
	Id int64 `json:"id"`
	// 博客id
	BlogId int64 `json:"blogId" gorm:"NOT NULL"`
	// 评论id
	ParentCommentId int64 `json:"parentCommentId"`
	// 用户名
	Username string `json:"username" gorm:"size:20;NOT NULL"`
	// 头像
	Avatar string `json:"avatar" gorm:"NOT NULL"`
	// 评论内容
	Comment string `json:"comment" gorm:"NOT NULL"`
	// 评论时间
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp"`
}
