package model

type User struct {
	// 主键
	Id int64 `binding:"required"`
	// 用户名
	Username string `form:"username" json:"username" binding:"required"`
	// 密码
	Password string `form:"password" json:"password" binding:"required"`
	// 邮箱
	Email string `form:"email" json:"email" binding:"required"`
	// 头像
	Avatar string `json:"avatar"`
	// 注册时间
	CreateTime Time `json:"create_time" `
	// 修改时间
	UpdateTime Time `json:"update_time"`
	// 一对多关系，User表作为关系维护方
	Blogs []Blog
}
