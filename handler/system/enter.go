package system

import "sync"

type SystemRouterGroup struct {
	UserHandler
	SearchHandler
}

var systemRouterGroups = new(SystemRouterGroup)
var once sync.Once

// 单例对象(懒加载)
func GetInstance() *SystemRouterGroup {
	once.Do(func() {
		if systemRouterGroups == nil {
			systemRouterGroups = &SystemRouterGroup{}
		}
	})
	return systemRouterGroups
}
