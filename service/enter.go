package service

import (
	"github.com/qianxia/blog/service/example"
	"github.com/qianxia/blog/service/system"
)

type ServiceGroup struct {
	example.ExampleGroup
	system.SystemGroup
}

var ServiceGroups = new(ServiceGroup)
