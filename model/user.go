package model

import (
	"encoding/json"
	"time"
)

// 用户表
type User struct {
	Id        uint64    `json:"id,omitempty" gorm:"primaryKey;comment:用户Id"`                     // 用户Id
	UUID      string    `json:"-" gorm:"type:varchar(36);not null;comment:用户UUID"`               // 用户UUID
	Username  string    `json:"username,omitempty" gorm:"type:varchar(20);not null;comment:用户名"` // 用户名
	Nickname  string    `json:"nickname,omitempty" gorm:"type:varchar(20);not null;comment:昵称"`  // 昵称
	Password  string    `json:"-" gorm:"type:varchar(80);not null;comment:密码"`                   // 密码
	Avatar    string    `json:"avatar,omitempty" gorm:"comment:头像"`                              // 头像
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"type:timestamp;comment:注册时间"`          // 创建时间
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"type:timestamp;comment:更新时间"`          // 更新时间
	Blogs     []Blog    `json:"Blogs,omitempty"`
}

func (u *User) TableName() string {
	return "t_user"
}

func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(&u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
