// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"time"
)

type Blog struct {
	// 主键
	ID int64 `json:"id"`
	// 创建者Id
	OwnerID int64 `json:"owner_id"`
	// 分类Id
	TypeID int64 `json:"type_id"`
	// 标题
	Title string `json:"title"`
	// 内容
	Content string `json:"content"`
	// 图片链接
	Image string `json:"image"`
	// 浏览次数
	Views sql.NullInt32 `json:"views"`
	// 创建时间
	CreatedAt sql.NullTime `json:"created_at"`
	// 修改时间
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Type struct {
	// 主键
	ID int64 `json:"id"`
	// 类别
	TypeName string `json:"type_name"`
}

type User struct {
	// 主键Id
	ID int64 `json:"id"`
	// 用户名
	Username string `json:"username"`
	// 用户邮箱
	Email string `json:"email"`
	// 用户昵称
	Nickname string `json:"nickname"`
	// 密码
	Password string `json:"password"`
	// 用户头像
	Avatar string `json:"avatar"`
	// 注册时间
	RegisterTime time.Time `json:"register_time"`
}
