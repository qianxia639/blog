package model

import (
	"time"
)

// 用户表
type User struct {
	Id        uint64    `json:"id,omitempty" gorm:"primaryKey;comment:用户Id"`                                      // 用户Id
	UUID      string    `json:"-" gorm:"type:varchar(36);not null;index;comment:用户UUID"`                          // 用户UUID
	Username  string    `json:"username,omitempty" gorm:"type:varchar(20);not null;unique;comment:用户名"`           // 用户名
	Nickname  string    `json:"nickname,omitempty" gorm:"type:varchar(20);not null;unique;comment:昵称"`            // 昵称
	Signer    string    `json:"-" gorm:"type:varchar(10);not null;comment:签名"`                                    // 签名
	Password  string    `json:"-" gorm:"type:varchar(80);not null;comment:密码"`                                    // 密码
	Avatar    string    `json:"avatar,omitempty" gorm:"default:https://picsum.photos/30/30/?image=41;comment:头像"` // 头像
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"type:timestamp;comment:注册时间"`                           // 创建时间
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"type:timestamp;comment:更新时间"`                           // 更新时间
	Blogs     []Blog    `json:"Blogs,omitempty"`
}
