package system

import (
	"sync"
)

type systemRouterGroup struct {
	UserHandler
	SearchHandler
	UploadHandler
	CommentHandler
	CaptchaHandler
	EmailHandler
}

var systemRouterGroups *systemRouterGroup
var once sync.Once

func GetInstance() *systemRouterGroup {
	once.Do(func() {
		systemRouterGroups = new(systemRouterGroup)
	})
	return systemRouterGroups
}
