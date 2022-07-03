package model

type Role struct {
	Id              uint32 `json:"id,omitempty" gorm:"primaryKey;comment:角色Id"`
	RoleName        string `json:"roleName,omitempty" gorm:"size:20;not null;comment:角色名"`
	RoleDescription string `json:"-" gorm:"size:40;not null;comment:角色描述"`
	User            []User `gorm:"many2many:t_user_role"`
}

func (*Role) TableName() string {
	return "t_role"
}
