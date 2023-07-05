// 自动生成模板TxnSub
package business

import "slot6/global"

// TxnSub 结构体
type TxnSub struct {
	global.GVA_MODEL
	MerchantId uint    `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;default:0;comment:Merchant ID;size:32;"`
	Pid        uint    `json:"pid" form:"pid" gorm:"column:pid;default:0;comment:PID;size:32;"`
	Type       uint8   `json:"type" form:"type" gorm:"column:type;default:1;comment:Type;size:8;"`
	Bet        float64 `json:"bet" form:"bet" gorm:"column:bet;comment:Bet;size:14;"`
	Raise      float64 `json:"raise" form:"raise" gorm:"column:raise;comment:Raise;size:14;"`
	Win        float64 `json:"win" form:"win" gorm:"column:win;comment:Win;size:14;"`
	BeforeBal  float64 `json:"beforeBal" form:"beforeBal" gorm:"column:before_bal;comment:Before Bal;size:14;"`
	ChangeBal  float64 `json:"changeBal" form:"changeBal" gorm:"column:change_bal;comment:Change Bal;size:14;"`
	AfterBal   float64 `json:"afterBal" form:"afterBal" gorm:"column:after_bal;comment:After Bal;size:14;"`
}

// TableName TxnSub 表名
func (TxnSub) TableName() string {
	return "b_txn_sub"
}
