package example

type ExampleGroup struct {
	BlogService
	TypeService
	TagService
	ArchiveService
}

var ExampleGroups = new(ExampleGroup)
