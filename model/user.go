package model

import uuid "github.com/satori/go.uuid"

type User struct {
	// 主键
	Id uint64 `gorm:"primaryKey;comment:主键"`
	// 用户UUID
	UUID uuid.UUID `json:"uuid" gorm:"size:36;comment:用户UUID"`
	// 用户名
	Username string `json:"username" gorm:"size:20;comment:用户名(默认为邮箱号)"`
	// 密码
	Password string `json:"-" gorm:"size:80comment:密码"`
	// 邮箱 (做登录账户)
	Email string `json:"email" gorm:"index;size:20;comment:邮箱"`
	// 头像
	Avatar string `json:"avatar" gorm:"default:https://picsum.photos/30/30/?image=41;comment:头像"`
	// 创建时间
	CreatedAt uint64 `json:"createdAt" gorm:"autoCreateTine:milli;comment:创建时间"`
	// 更新时间
	UpdatedAt uint64 `json:"updatedAt" gorm:"autoUpdateTine:milli;comment:更新时间"`
	Blogs     []Blog
}
