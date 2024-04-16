package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type MoneyUserSlotSearch struct {
	business.MoneyUserSlot
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
