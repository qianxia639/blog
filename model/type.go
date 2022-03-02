package model

type Type struct {
	Id       uint16 `json:"id"`
	TypeName string `json:"typeName" gorm:"size:10;NOT NULL"`
	Amount   uint32 `json:"amount"`
	Blogs    []Blog
}
