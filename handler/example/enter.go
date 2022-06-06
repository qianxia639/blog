package example

import "github.com/qianxia/blog/service"

type exampleRouterGroup struct {
	BlogHandler
	TypeHandler
	TagHandler
	ArchiveHandler
}

var ExampleRouterGroups = new(exampleRouterGroup)

var (
	blogService    = service.ServiceGroups.ExampleGroup.BlogService
	typeService    = service.ServiceGroups.ExampleGroup.TypeService
	tagService     = service.ServiceGroups.ExampleGroup.TagService
	archiveService = service.ServiceGroups.ExampleGroup.ArchiveService
)
