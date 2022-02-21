package system

type SystemRouterGroup struct {
	UserHandler
	SearchHandler
}

var systemRouterGroups *SystemRouterGroup

// 单例对象(懒加载)
func GetInstance() *SystemRouterGroup {
	if systemRouterGroups == nil {
		systemRouterGroups = &SystemRouterGroup{}
	}
	return systemRouterGroups
}
