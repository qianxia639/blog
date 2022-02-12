package model

type User struct {
	// 主键
	Id int64 `gorm:"primaryKey;comment:主键"`
	// 用户名
	Username string `json:"username" gorm:"size:20;NOT NULL"`
	// 密码
	Password string `json:"password" gorm:"size:80;NOT NULL"`
	// 邮箱
	Email string `json:"email" gorm:"size:20;NOT NULL"`
	// 头像
	Avatar string `json:"avatar" gorm:"default:https://picsum.photos/30/30/?image=41"`
	// 创建时间
	CreatedAt uint64 `json:"createdAt" gorm:"autoCreateTine:milli;comment:创建时间"`
	// 更新时间
	UpdatedAt uint64 `json:"updatedAt" gorm:"autoUpdateTine:milli;comment:更新时间"`
	Blogs     []Blog
}
