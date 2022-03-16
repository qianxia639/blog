package system

import "sync"

type systemRouterGroup struct {
	UserHandler
	SearchHandler
}

var systemRouterGroups *systemRouterGroup
var once sync.Once

// 单例对象(懒加载)
func GetInstance() *systemRouterGroup {
	once.Do(func() {
		systemRouterGroups = new(systemRouterGroup)
	})
	return systemRouterGroups
}
