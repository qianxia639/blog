package model

type User struct {
	Id int64 `binding:"required" gorm:"primaryKey;comment:主键"` // 主键
	// 用户名
	Username string `form:"username" json:"username" binding:"required" gorm:"size:20;NOT NULL"`
	// 密码
	Password string `form:"password" json:"password" binding:"required" gorm:"size:80;NOT NULL"`
	// 邮箱
	Email string `form:"email" json:"email" binding:"required" gorm:"size:20;NOT NULL"`
	// 头像
	Avatar string `json:"avatar" gorm:"default:https://picsum.photos/30/30/?image=41"`
	// 创建时间
	CreatedAt Time `json:"createdAt" gorm:"type:timestamp;comment:创建时间"`
	// 更新时间
	UpdatedAt Time `json:"updatedAt" gorm:"type:timestamp;comment:更新时间"`
	Blogs     []Blog
}
