package example

type exampleRouterGroup struct {
	BlogHandler
	TypeHandler
	TagHandler
	ArchiveHandler
}

var ExampleRouterGroups = new(exampleRouterGroup)
