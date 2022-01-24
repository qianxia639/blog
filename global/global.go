package global

import (
	"gorm.io/gorm"
)

var (
	RY_DB      *gorm.DB
	RY_JWT_Key = []byte("l_ruo_yu_y_y")
)
