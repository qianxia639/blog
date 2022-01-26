package model

type Type struct {
	Id       int    `json:"id" binding:"required"`
	TypeName string `json:"typeName" binding:"required" gorm:"size:10;NOT NULL"`
	Amount   int    `json:"amount"`
	Blogs    []Blog
}
