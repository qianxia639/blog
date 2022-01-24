package model

import "time"

type Blog struct {
	// 主键
	Id int64 `json:"id" binding:"required"`
	// 用户id，
	UserId int64 `json:"user_id" `
	// 分类id
	TypeId int `json:"type_id" `
	// 标题
	Title string `json:"title" binding:"required"`
	// 内容
	Content string `json:"content" binding:"required"`
	// 标记
	Flag string `json:"flag"`
	// 是否开启点赞
	Were bool `json:"were"`
	// 是否显示转载声明
	ShareStatement bool `json:"share_statement"`
	// 是否开启评论
	Commentabled bool `json:"commentabled"`
	// 点赞数
	Likes int `json:"likes"`
	// 浏览次数
	Views int `json:"views"`
	// 创建时间
	CreateTime time.Time `gorm:"type:timestamp" json:"create_time"`
	// 更新时间
	UpdateTime time.Time `gorm:"type:timestamp" json:"update_time"`
	Tags       []*Tag    `gorm:"many2many:ry_blog_tag;"`
}
