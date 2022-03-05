package model

type Tag struct {
	// 主键
	Id uint16 `json:"id"`
	// 标签名
	TagName string `json:"tagName" gorm:"size:20;NOT NULL;comment:标签名"`
	Blogs   []Blog `gorm:"many2many:blog_tag;"`
}
