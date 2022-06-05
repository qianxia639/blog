package example

type exampleGroup struct {
	BlogService
	TypeService
	TagService
	ArchiveService
}

var ExampleGroups = new(exampleGroup)
