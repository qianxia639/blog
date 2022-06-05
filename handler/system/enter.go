package system

import "github.com/qianxia/blog/service/system"

type systemRouterGroup struct {
	UserHandler
	SearchHandler
	UploadHandler
	CaptchaHandler
}

var SystemRouterGroups = new(systemRouterGroup)

var (
	userService   = system.SystemGroups.UserService
	searchService = system.SystemGroups.SearchService
)
