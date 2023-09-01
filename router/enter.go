package router

import (
	"github.com/ozeer/gva/router/example"
	"github.com/ozeer/gva/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
