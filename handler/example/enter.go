package example

type ExampleRouterGroup struct {
	BlogHandler
	TypeHandler
	TagHandler
	ArchiveHandler
}

var exampleRouterGroups *ExampleRouterGroup

func GetInstance() *ExampleRouterGroup {
	if exampleRouterGroups == nil {
		exampleRouterGroups = &ExampleRouterGroup{}
	}
	return exampleRouterGroups
}
