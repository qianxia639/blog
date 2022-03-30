package model

import "time"

type Role struct {
	Id        uint32 `gorm:"primaryKey;comment:主键"`
	RoleName  string `json:"roleName" gorm:"comment:角色名"`
	Users     []User `gorm:"many2many:user_role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
