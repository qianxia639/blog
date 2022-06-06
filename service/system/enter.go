package system

type SystemGroup struct {
	UserService
	SearchService
	ElasticSearchService
}

var SystemGroups = new(SystemGroup)
