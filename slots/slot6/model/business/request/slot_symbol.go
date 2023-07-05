package request

import (
	"time"
)

type SlotSymbolSearch struct {
	business.SlotSymbol
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
