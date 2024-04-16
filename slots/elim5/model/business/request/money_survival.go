package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type MoneySurvivalSearch struct {
	business.MoneySurvival
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartDate      *int       `json:"startDate" form:"startDate"`
	EndDate        *int       `json:"endDate" form:"endDate"`
	request.PageInfo
	Sort        string `json:"sort" form:"sort"`
	Order       string `json:"order" form:"order"`
	BetweenDate []int  `json:"betweenDate[]" form:"betweenDate[]"`
}
