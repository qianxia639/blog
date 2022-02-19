package system

import "sync"

type SystemRouterGroup struct {
	UserHandler
	SearchHandler
}

var systemRouterGroups *SystemRouterGroup
var once sync.Once

// 单例对象(懒加载)
func GetInstance() *SystemRouterGroup {
	once.Do(func() {
		systemRouterGroups = &SystemRouterGroup{}
	})
	return systemRouterGroups
}
