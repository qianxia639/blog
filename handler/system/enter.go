package system

import "github.com/qianxia/blog/service"

type SystemRouterGroup struct {
	UserHandler
	SearchHandler
	UploadHandler
	CaptchaHandler
}

var SystemRouterGroups = new(SystemRouterGroup)

var (
	userService   = service.ServiceGroups.SystemGroup.UserService
	searchService = service.ServiceGroups.SystemGroup.SearchService
)
