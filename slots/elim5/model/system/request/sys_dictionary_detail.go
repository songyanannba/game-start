package request

import (
	"elim5/model/common/request"
	"elim5/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
