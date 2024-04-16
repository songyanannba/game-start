package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type DataSurvivalSearch struct {
	business.DataSurvival
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	Sort        string `json:"sort" form:"sort"`
	Order       string `json:"order" form:"order"`
	BetweenDate []int  `json:"betweenDate[]" form:"betweenDate[]"`
}
