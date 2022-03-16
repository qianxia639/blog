package example

import "sync"

type exampleRouterGroup struct {
	BlogHandler
	TypeHandler
	TagHandler
	ArchiveHandler
}

var exampleRouterGroups *exampleRouterGroup
var once sync.Once

func GetInstance() *exampleRouterGroup {
	once.Do(func() {
		exampleRouterGroups = new(exampleRouterGroup)
	})
	return exampleRouterGroups
}
