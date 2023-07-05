package request

import (
	"time"
)

type SlotPayTableSearch struct {
	business.SlotPayTable
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
