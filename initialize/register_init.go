package initialize

import (
	_ "github.com/ozeer/gva/source/example"
	_ "github.com/ozeer/gva/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
