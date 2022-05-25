package model

// 标签表
type Tag struct {
	Id      uint16 `json:"id,omitempty" gorm:"comment:标签Id"`             // 主键
	TagName string `json:"tagName,omitempty" gorm:"size:20;comment:标签名"` // 标签名
	Blogs   []Blog `gorm:"many2many:blog_tag;"`
}
