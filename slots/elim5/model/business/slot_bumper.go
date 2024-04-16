// 自动生成模板SlotBumper
package business

import (
	"elim5/global"
)

// SlotBumper 结构体
type SlotBumper struct {
	global.GVA_MODEL
	Currency   string  `json:"currency" form:"currency" gorm:"column:currency;comment:货币;size:50;"`
	MerchantId uint    `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;comment:商户id;size:32;"`
	Open       int64   `json:"open" form:"open" gorm:"column:open;default:0;comment:上限开启条件;size:64;"`
	Target     float64 `json:"target" form:"target" gorm:"column:target;comment:上限触发条件;size:10;"`
	OpenLow    int64   `json:"openLow" form:"openLow" gorm:"column:open_low;default:0;comment:下限开启条件;size:64;"`
	TargetLow1 float64 `json:"targetLow1" form:"targetLow1" gorm:"type:decimal(14,2);column:target_low1;comment:下限触发条件1;"`
	Low1Num    int     `json:"low1Num" form:"low1Num" gorm:"column:low1_num;comment:下限触发次数1;size:32;"`
	TargetLow2 float64 `json:"targetLow2" form:"targetLow2" gorm:"type:decimal(14,2);column:target_low2;comment:下限触发条件2;"`
	Low2Num    int     `json:"low2Num" form:"low2Num" gorm:"column:low2_num;comment:下限触发次数2;size:32;"`
}

// TableName SlotBumper 表名
func (SlotBumper) TableName() string {
	return "b_slot_bumper"
}
