package system

type systemGroup struct {
	UserService
	SearchService
	ElasticSearchService
}

var SystemGroups = new(systemGroup)
