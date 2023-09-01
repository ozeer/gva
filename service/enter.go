package service

import (
	"github.com/ozeer/gva/service/example"
	"github.com/ozeer/gva/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
