package model

type Type struct {
	Id       int    `binding:"required"`
	TypeName string `json:"type_name" binding:"required"`
}
