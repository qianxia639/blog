package example

import "sync"

type ExampleRouterGroup struct {
	BlogHandler
	TypeHandler
	TagHandler
	ArchiveHandler
}

var exampleRouterGroups = new(ExampleRouterGroup)
var once sync.Once

func GetInstance() *ExampleRouterGroup {
	once.Do(func() {
		if exampleRouterGroups == nil {
			exampleRouterGroups = &ExampleRouterGroup{}
		}
	})
	return exampleRouterGroups
}
