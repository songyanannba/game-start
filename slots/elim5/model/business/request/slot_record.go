package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type SlotRecordSearch struct {
	business.SlotRecord
	PlayerName     string     `json:"playerName" form:"playerName"`
	DateChoice     *string    `json:"dateChoice" form:"dateChoice"`
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartGain      *float64   `json:"startGain" form:"startGain"`
	EndGain        *float64   `json:"endGain" form:"endGain"`
	StartChangeBal *float64   `json:"startChangeBal" form:"startChangeBal"`
	EndChangeBal   *float64   `json:"endChangeBal" form:"endChangeBal"`
	StartRate      *float64   `json:"startRate" form:"startRate"`
	EndRate        *float64   `json:"endRate" form:"endRate"`
	TxnNo          string     `json:"txnNo" form:"txnNo"`
	request.PageInfo
}

type SlotRecordPublicSearch struct {
	No string `json:"no" form:"no"`
}
