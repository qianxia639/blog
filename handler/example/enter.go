package example

import "github.com/qianxia/blog/service/example"

type ExampleRouterGroup struct {
	BlogHandler
	TypeHandler
	TagHandler
	ArchiveHandler
}

var ExampleRouterGroups = new(ExampleRouterGroup)

var (
	blogService    = example.ExampleGroups.BlogService
	typeService    = example.ExampleGroups.TypeService
	tagService     = example.ExampleGroups.TagService
	archiveService = example.ExampleGroups.ArchiveService
)
