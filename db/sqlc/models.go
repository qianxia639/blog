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
	Views int32 `json:"views"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 修改时间
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	// 主键
	ID int64 `json:"id"`
	// 博客Id
	BlogID int64 `json:"blog_id"`
	// 父评论Id
	CommentID int64 `json:"comment_id"`
	// 昵称
	Nickname string `json:"nickname"`
	// 头像
	Avatar string `json:"avatar"`
	// 评论内容
	Content string `json:"content"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
}

type RequestLog struct {
	// 主键
	ID int64 `json:"id"`
	// 请求方式
	Method string `json:"method"`
	// 路由
	Path string `json:"path"`
	// 状态码
	StatusCode int32 `json:"status_code"`
	// 访问ip
	Ip string `json:"ip"`
	// 主机名
	Hostname string `json:"hostname"`
	// 请求体
	RequestBody sql.NullString `json:"request_body"`
	// 响应时间/ms
	ResponseTime int64 `json:"response_time"`
	// 请求时间
	RequestTime time.Time `json:"request_time"`
	// 请求数据类型
	ContentType string `json:"content_type"`
	// ua
	UserAgent string `json:"user_agent"`
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
