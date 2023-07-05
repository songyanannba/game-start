package request

import (
	"slot6/model/business"
	"slot6/model/common/request"
	"time"
)

type SlotFakeSearch struct {
	business.SlotFake
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
