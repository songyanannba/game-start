package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type ReportWinLossSearch struct {
	business.ReportWinLoss
	GroupBy        string     `json:"groupBy" form:"groupBy"`
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	TimeOffset     int        `json:"timeOffset" form:"timeOffset"`
	request.PageInfo
}
