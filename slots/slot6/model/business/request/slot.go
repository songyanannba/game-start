package request

import (
	"slot6/model/business"
	"slot6/model/common/request"
	"time"
)

type SlotSearch struct {
	business.Slot
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
