package model

type Type struct {
	Id       uint32 `json:"id" binding:"required"`
	TypeName string `json:"typeName" binding:"required" gorm:"size:10;NOT NULL"`
	Amount   uint32 `json:"amount"`
	Blogs    []Blog
}
