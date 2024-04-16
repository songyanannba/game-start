package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type SlotTemplateSearch struct {
	business.SlotTemplate
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	IsExport bool  `json:"isExport" form:"isExport"`
	Ids      []int `json:"ids[]" form:"ids[]"`
}
