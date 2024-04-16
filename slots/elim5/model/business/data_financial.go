// 自动生成模板DataFinancial
package business

import (
	"elim5/global"
	"gorm.io/gorm/clause"
)

// DataFinancial 结构体
type DataFinancial struct {
	global.GVA_MODEL
	Type       uint8  `json:"type" form:"type" gorm:"uniqueIndex:idx_unique;column:type;default:1;comment:;size:8;"`
	Date       int    `json:"date" form:"date" gorm:"uniqueIndex:idx_unique;column:date;default:0;comment:;size:32;"`
	SetId      int    `json:"setId" form:"setId" gorm:"uniqueIndex:idx_unique;column:set_id;default:0;comment:;size:32;"`
	MerchantId int    `json:"merchantId" form:"merchantId" gorm:"uniqueIndex:idx_unique;column:merchant_id;default:0;comment:;size:32;"`
	SlotId     int    `json:"slotId" form:"slotId" gorm:"uniqueIndex:idx_unique;column:slot_id;default:0;comment:;size:32;"`
	Currency   string `json:"currency" form:"currency" gorm:"uniqueIndex:idx_unique;column:currency;comment:;size:255;"`
	SpinCount  int    `json:"spinCount" form:"spinCount" gorm:"column:spin_count;default:0;comment:;size:64;"`
	Bet        int64  `json:"bet" form:"bet" gorm:"column:bet;default:0;comment:;size:64;"`
	Payout     int64  `json:"payout" form:"payout" gorm:"column:payout;default:0;comment:;size:64;"`
	Win        int64  `json:"win" form:"win" gorm:"column:win;default:0;comment:;size:64;"`
}

// TableName DataFinancial 表名
func (DataFinancial) TableName() string {
	return "b_data_financial"
}

func (DataFinancial) GetUniqueIndex() []clause.Column {
	return []clause.Column{{Name: "type"}, {Name: "date"}, {Name: "set_id"}, {Name: "merchant_id"}, {Name: "slot_id"}, {Name: "currency"}}
}
