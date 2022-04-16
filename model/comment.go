package model

import "time"

type Comment struct {
	// 主键
	Id uint64 `json:"id"`
	// 博客id
	BlogId uint64 `json:"blogId" gorm:"comment:博客Id"`
	// 父评论id
	ParentId uint64 `json:"parentId" gorm:"comment:父评论Id"`
	// 用户名
	Username string `json:"username" gorm:"size:20;comment:用户名"`
	// 头像
	Avatar string `json:"avatar" gorm:"comment:头像"`
	// 评论内容
	Content string `json:"content" gorm:"comment:评论内容"`
	// 评论时间
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;comment:评论时间"`
}
