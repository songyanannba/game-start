// 自动生成模板Tracking
package business

import (
	"elim5/global"
)

// Tracking 结构体
type Tracking struct {
	global.GVA_MODEL
	Date       uint   `json:"date" form:"date" gorm:"index;default:0;column:date;comment:日期;size:32;"`
	Type       uint8  `json:"type" form:"type" gorm:"column:type;default:0;comment:类型;size:8;"`
	SlotId     uint   `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器id;size:32;"`
	UserId     uint   `json:"userId" form:"userId" gorm:"column:user_id;default:0;comment:用户id;size:32;"`
	Val        string `json:"val" form:"val" gorm:"column:val;comment:值;"`
	Ip         string `json:"ip" form:"ip" gorm:"column:ip;comment:ip;size:32;"`
	MerchantId uint   `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;default:0;comment:商户id;size:32;"`
}

// TableName Tracking 表名
func (Tracking) TableName() string {
	return "b_tracking"
}
