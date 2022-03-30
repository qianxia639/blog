package system

import (
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/qianxia/blog/global"
)

type AuthorityService struct{}

var (
	one          sync.Once
	syncEnforcer *casbin.SyncedEnforcer
)

func (*AuthorityService) Casbin() *casbin.SyncedEnforcer {
	one.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.QX_DB)
		syncEnforcer, _ = casbin.NewSyncedEnforcer("./config/rbac_model.conf", a)
		syncEnforcer.LoadPolicy()
	})
	return syncEnforcer
}
