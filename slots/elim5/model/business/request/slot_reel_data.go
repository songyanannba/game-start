package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type SlotReelDataSearch struct {
	business.SlotReelData
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
