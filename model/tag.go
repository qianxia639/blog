package model

// 标签表
type Tag struct {
	Id      uint16 `json:"id,omitempty" gorm:"primaryKey,comment:标签Id"`                    // 主键
	TagName string `json:"tagName,omitempty" gorm:"type:varchar(20);not null;comment:标签名"` // 标签名
	Blogs   []Blog `json:"Blogs,omitempty" gorm:"many2many:t_blog_tag"`
}

func (*Tag) TableName() string {
	return "t_tag"
}
