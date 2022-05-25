package model

// 分类表
type Type struct {
	Id       uint16 `json:"id,omitempty" gorm:"comment:分类Id"`                    // 分类Id
	TypeName string `json:"typeName,omitempty" gorm:"size:10;comment:分类名"`       // 分类名称
	Amount   uint32 `json:"amount,omitempty" gorm:"default:0;comment:分类对应的博客数量"` // 分类对应的博客数量
	Blogs    []Blog
}
