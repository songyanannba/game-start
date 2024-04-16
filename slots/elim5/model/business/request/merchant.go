package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type MerchantSearch struct {
	business.Merchant
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type AlterRtpParameter struct {
	Type        int    `json:"type"`
	Ids         []int  `json:"ids"`
	RtpOptional string `json:"rtp_optional"`
	RtpConfirm  int    `json:"rtp_confirm"`
}
