package system

import "github.com/qianxia/blog/service/system"

type SystemRouterGroup struct {
	UserHandler
	SearchHandler
	UploadHandler
	CaptchaHandler
}

var SystemRouterGroups = new(SystemRouterGroup)

var (
	userService   = system.SystemGroups.UserService
	searchService = system.SystemGroups.SearchService
)
