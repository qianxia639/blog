package system

type SystemRouterGroup struct {
	UserHandler
	SearchHandler
	UploadHandler
	CaptchaHandler
	CommentHandler
}

var SystemRouterGroups = new(SystemRouterGroup)
