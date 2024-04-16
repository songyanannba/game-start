package request

import (
	"elim5/model/common/request"
	"elim5/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
