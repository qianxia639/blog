package model

import "time"

type User struct {
	Id int64 `gorm:"primarykey"` // 主键
	// 用户名
	Username string `gorm:"type:varchar(20)" form:"username" json:"username" binding:"required"`
	// 密码
	Password string `gorm:"type:varchar(80)" form:"password" json:"password" binding:"required"`
	// 邮箱
	Email string `gorm:"type:varchar(20)" form:"email" json:"email" binding:"required"`
	// 头像
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	Blogs     []Blog
}
