package system

import (
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/qianxia/blog/global"
)

type CasbinService struct{}

var CasbinServices = new(CasbinService)

var (
	once     sync.Once
	enforcer *casbin.Enforcer
)

func (cs *CasbinService) Casbin() *casbin.Enforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.DB)
		text := `
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			global.LOG.Error("字符串模型加载失败!", err)
			return
		}
		enforcer, _ = casbin.NewEnforcer(m, a)
	})
	_ = enforcer.LoadPolicy()
	return enforcer
}
