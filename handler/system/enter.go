package system

import "github.com/qianxia/blog/service"

type SystemRouterGroup struct {
	UserHandler
	SearchHandler
	UploadHandler
	CaptchaHandler
	CommentHandler
}

var SystemRouterGroups = new(SystemRouterGroup)

var (
	userService    = service.ServiceGroups.SystemGroup.UserService
	searchService  = service.ServiceGroups.SystemGroup.SearchService
	commentService = service.ServiceGroups.SystemGroup.CommentService
)
