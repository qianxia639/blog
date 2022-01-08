package model

type User struct {
	// 主键
	Id int64 `gorm:"primary_key" binding:"required"`
	// 用户名
	Username string `gorm:"size:40" form:"username" json:"username" binding:"required"`
	// 密码 md5加密
	Password string `gorm:"size:80" form:"password" json:"password" binding:"required"`
	// 邮箱
	Email string `gorm:"size:40;default:null" form:"email" json:"email" binding:"required"`
	// 头像
	// Face string
	// 注册时间
	CreateTime Time `json:"create_time" gorm:"type:timestamp"`
	// 修改时间
	EditTime Time `json:"edit_time" gorm:"type:timestamp"`
}
