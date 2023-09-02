package v1

import (
	"github.com/ozeer/gva/api/v1/designer"
	"github.com/ozeer/gva/api/v1/example"
	"github.com/ozeer/gva/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	DesignerApiGroup designer.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
