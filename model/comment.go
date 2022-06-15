package model

import (
	"time"
)

// 评论表
type Comment struct {
	Id        uint64    `json:"id,omitempty" gorm:"comment:评论Id"`                       // 评论Id
	BlogId    uint64    `json:"blogId,omitempty" gorm:"comment:博客Id"`                   // 博客id
	ParentId  uint64    `json:"parentId,omitempty" gorm:"comment:父评论Id"`                // 父评论id
	Nickname  string    `json:"nickname,omitempty" gorm:"size:20;comment:用户名"`          // 用户名
	Avatar    string    `json:"avatar,omitempty" gorm:"comment:头像"`                     // 头像
	Content   string    `json:"content,omitempty" gorm:"comment:评论内容"`                  // 评论内容
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"type:timestamp;comment:评论时间"` // 评论时间
}
