package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type SlotReelSearch struct {
	business.SlotReel
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
