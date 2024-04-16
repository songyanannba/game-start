// 自动生成模板SlotAmountLimit
package business

import (
	"elim5/global"
)

// SlotAmountLimit 结构体
type SlotAmountLimit struct {
	global.GVA_MODEL
	Agent    string `json:"agent" form:"agent" gorm:"column:agent;comment:agent;size:30;"`
	Currency string `json:"currency" form:"currency" gorm:"column:currency;comment:货币;size:30;"`
	Limit    int64  `json:"limit" form:"limit" gorm:"column:limit;default:0;comment:余额限制;size:64;"`
}

// TableName SlotAmountLimit 表名
func (SlotAmountLimit) TableName() string {
	return "b_slot_amount_limit"
}
