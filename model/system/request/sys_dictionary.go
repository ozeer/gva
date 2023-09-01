package request

import (
	"github.com/ozeer/gva/model/common/request"
	"github.com/ozeer/gva/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
