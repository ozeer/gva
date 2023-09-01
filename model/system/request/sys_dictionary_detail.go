package request

import (
	"github.com/ozeer/gva/model/common/request"
	"github.com/ozeer/gva/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
