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
		exampleRouterGroups = &ExampleRouterGroup{}
	})
	return exampleRouterGroups
}
