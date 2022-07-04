package data

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitCasbinRuleData(db *gorm.DB) error {
	casbinData := []gormadapter.CasbinRule{
		{Ptype: "g", V0: "777", V1: "999"},

		{Ptype: "p", V0: "999", V1: "/user/info", V2: "GET"},
		{Ptype: "p", V0: "999", V1: "/user/name", V2: "PUT"},
		{Ptype: "p", V0: "999", V1: "/user/pwd", V2: "PUT"},
		{Ptype: "p", V0: "999", V1: "/user/avatar", V2: "PUT"},
		{Ptype: "p", V0: "777", V1: "/user/list", V2: "GET"},

		{Ptype: "p", V0: "999", V1: "/upload/mdFile", V2: "POST"},

		{Ptype: "p", V0: "999", V1: "/blog/save", V2: "POST"},
		{Ptype: "p", V0: "999", V1: "/blog/update", V2: "PUT"},
		{Ptype: "p", V0: "999", V1: "/blog/list", V2: "GET"},
		{Ptype: "p", V0: "999", V1: "/blog/:id", V2: "DELETE"},
		{Ptype: "p", V0: "777", V1: "/blog/all", V2: "GET"},
		{Ptype: "p", V0: "777", V1: "/blog/flag/list", V2: "GET"},

		{Ptype: "p", V0: "777", V1: "/type/save", V2: "POST"},

		{Ptype: "p", V0: "777", V1: "/tag/save", V2: "POST"},
	}

	return db.Create(&casbinData).Error

}
