package model

type Type struct {
	Id       uint16 `json:"id"`
	TypeName string `json:"typeName" gorm:"size:10;NOT NULL;comment:分类名"`
	Amount   uint32 `json:"amount" gorm:"default:0;comment:分类对应的博客数量"`
	Blogs    []Blog
}
