package request

import (
	"github.com/ozeer/gva/model/common/request"
	"github.com/ozeer/gva/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
