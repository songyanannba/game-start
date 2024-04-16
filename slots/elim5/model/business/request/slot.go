package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type SlotSearch struct {
	business.Slot
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
