package model

import "time"

type Comment struct {
	// 主键
	Id uint64 `json:"id"`
	// 博客id
	BlogId uint64 `json:"blogId" gorm:"NOT NULL"`
	// 父评论id
	ParentId uint64 `json:"parentId" gorm:"父评论Id"`
	// 用户名
	Username string `json:"username" gorm:"size:20;NOT NULL"`
	// 头像
	Avatar string `json:"avatar" gorm:"NOT NULL"`
	// 评论内容
	CommentContent string `json:"commentContent" gorm:"NOT NULL"`
	// 评论时间
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp"`
}
