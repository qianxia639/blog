package example

import "github.com/qianxia/blog/service/example"

type exampleRouterGroup struct {
	BlogHandler
	TypeHandler
	TagHandler
	ArchiveHandler
}

var ExampleRouterGroups = new(exampleRouterGroup)

var (
	blogService    = example.ExampleGroups.BlogService
	typeService    = example.ExampleGroups.TypeService
	tagService     = example.ExampleGroups.TagService
	archiveService = example.ExampleGroups.ArchiveService
)
