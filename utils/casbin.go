package utils

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/qianxia/blog/global"
)

func Casbin() *casbin.SyncedEnforcer {
	a, _ := gormadapter.NewAdapterByDB(global.QX_DB)
	syncEnforcer, _ := casbin.NewSyncedEnforcer("./config/rbac_model.conf", a)
	syncEnforcer.LoadPolicy()

	return syncEnforcer
}
