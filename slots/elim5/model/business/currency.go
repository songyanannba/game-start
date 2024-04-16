// 自动生成模板Currency
package business

import (
	"elim5/global"
)

// Currency 结构体
type Currency struct {
	global.GVA_MODEL
	Name         string  `json:"name" form:"name" gorm:"column:name;comment:名称;size:50;"`
	ExchangeRate float64 `json:"exchangeRate" form:"exchangeRate" gorm:"type:decimal(14,3);column:exchange_rate;comment:汇率;"`
	Status       uint8   `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
}

// TableName Currency 表名
func (Currency) TableName() string {
	return "b_currency"
}
