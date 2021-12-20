package model

import "time"

type User struct {
	// 主键
	Id int64 `gorm:"primary_key" binding:"required"`
	// 用户名
	Username string `gorm:"size:40" form:"username" json:"username" binding:"required"`
	// 密码 md5加密
	Password string `gorm:"size:80" form:"password" json:"password" binding:"required"`
	// 邮箱
	Email string `gorm:"size:40;default:null" form:"email" json:"email"`
	// 头像
	// Face string
	// 注册时间
	CreateTime time.Time `time_format:"2006-01-02" `
	// 修改时间
	EditTime time.Time `time_format:"2006-01-02" `
}
