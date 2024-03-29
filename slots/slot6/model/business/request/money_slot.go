package request

import (
	"slot6/model/business"
	"slot6/model/common/request"
	"time"
)

type MoneySlotSearch struct {
	business.MoneySlot
	StartCreatedAt   *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt     *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartRecentSpins *int       `json:"startRecentSpins" form:"startRecentSpins"`
	EndRecentSpins   *int       `json:"endRecentSpins" form:"endRecentSpins"`
	StartAvgSpins    *int       `json:"startAvgSpins" form:"startAvgSpins"`
	EndAvgSpins      *int       `json:"endAvgSpins" form:"endAvgSpins"`
	request.PageInfo
}
