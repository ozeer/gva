package service

import (
	"github.com/ozeer/gva/service/designer"
	"github.com/ozeer/gva/service/example"
	"github.com/ozeer/gva/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	DesignerServiceGroup designer.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
