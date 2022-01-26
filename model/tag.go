package model

type Tag struct {
	// 主键
	Id int `json:"id" binding:"required"`
	// 标签名
	TagName string  `json:"tagName" binding:"required" gorm:"size:20;NOT NULL"`
	Blogs   []*Blog `gorm:"many2many:blog_tag;"`
}
