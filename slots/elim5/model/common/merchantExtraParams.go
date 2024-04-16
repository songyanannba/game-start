package common

type MerchantExtraParams struct {
	RtpGear string `json:"rtpGear"  example:"1" gorm:"column:rtp_gear;comment:rtp档位;size:20;"` // RTP档位
}
