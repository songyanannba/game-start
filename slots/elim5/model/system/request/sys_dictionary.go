package request

import (
	"elim5/model/common/request"
	"elim5/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
