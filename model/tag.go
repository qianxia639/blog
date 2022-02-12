package model

type Tag struct {
	// 主键
	Id uint16 `json:"id"`
	// 标签名
	TagName string `json:"tagName" binding:"required" gorm:"size:20;NOT NULL"`
	Blogs   []Blog `gorm:"many2many:blog_tag;"`
}
