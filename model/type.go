package model

type Type struct {
	Id       int    `json:"id" binding:"required"`
	TypeName string `json:"type_name" binding:"required"`
	Amount   int    `json:"amount"`
	Blogs    []Blog
}
