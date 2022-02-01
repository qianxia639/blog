package model

import "time"

type Blog struct {
	// 主键
	Id int64 `json:"id" binding:"required"`
	// 用户id，
	UserId int64 `json:"userId" binding:"required" gorm:"NOT NULL"`
	// 分类id
	TypeId int `json:"typeId" binding:"required" gorm:"NOT NULL"`
	// 标题
	Title string `json:"title" binding:"required" gorm:"size:20;NOT NULL"`
	// 内容
	Content string `json:"content" binding:"required" gorm:"size:255;NOT NULL"`
	// 标记
	Flag string `json:"flag" gorm:"type:varchar(10);comment:标记"`
	// 是否开启点赞
	Were bool `json:"were" gorm:"comment:是否开启点赞"`
	// 是否显示转载声明
	ShareStatement bool `json:"shareStatement" gorm:"comment:是否显示转载声明"`
	// 是否开启评论
	EnableComment bool `json:"enableComment" gorm:"comment:是否开启评论"`
	// 点赞数
	Likes int `json:"likes" gorm:"comment:点赞数"`
	// 浏览次数
	Views int `json:"views" gorm:"comment:浏览次数" `
	// 创建时间
	// CreatedAt Time `json:"createdAt" gorm:"type:timestamp;comment:创建时间"`
	CreatedAt time.Time `json:"createdAt" `
	// 更新时间
	// UpdatedAt Time   `json:"updatedAt" gorm:"type:timestamp;comment:更新时间"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tags      []*Tag    `gorm:"many2many:blog_tag;"`
}
