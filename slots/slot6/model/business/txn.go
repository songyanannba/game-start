// 自动生成模板Txn
package business

import "slot6/global"

// Txn 结构体
type Txn struct {
	global.GVA_MODEL
	MerchantId uint   `json:"merchantId" form:"merchantId" gorm:"index;column:merchant_id;default:0;comment:Merchant ID;size:32;"`
	UserId     uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:User ID;size:32;"`
	PlayerName string `json:"playerName" form:"playerName" gorm:"column:player_name;comment:Player Name;size:100;"`

	GameId   int    `json:"gameId" form:"gameId" gorm:"column:game_id;default:0;comment:Game ID;size:32;"`
	TxnId    string `json:"txnId" form:"txnId" gorm:"column:txn_id;comment:Txn ID;size:50;"`
	Currency string `json:"currency" form:"currency" gorm:"column:currency;comment:Currency;size:10;"`

	Bet   float64 `json:"bet" form:"bet" gorm:"type:decimal(14,2);column:bet;comment:Bet;size:14;"`
	Raise float64 `json:"raise" form:"raise" gorm:"type:decimal(14,2);column:raise;comment:Raise;size:14;"`
	Win   float64 `json:"win" form:"win" gorm:"type:decimal(14,2);column:win;comment:Win;size:14;"`

	ChangeBal float64 `json:"changeBal" form:"changeBal" gorm:"type:decimal(14,2);column:change_bal;comment:Change Bal;size:14;"`
	BeforeBal float64 `json:"beforeBal" form:"beforeBal" gorm:"type:decimal(14,2);column:before_bal;comment:Before Bal;size:14;"`
	AfterBal  float64 `json:"afterBal" form:"afterBal" gorm:"type:decimal(14,2);column:after_bal;comment:After Bal;size:14;"`

	RealBal float64 `json:"realBal" form:"realBal" gorm:"type:decimal(14,2);column:real_bal;comment:Real Bal;size:14;"`

	PlatformTxnId string `json:"platformTxnId" form:"platformTxnId" gorm:"column:platform_txn_id;comment:Platform Txn ID;size:100;"`

	Status uint8 `json:"status" form:"status" gorm:"column:status;default:1;comment:Status;size:8;"`
}

// TableName Txn 表名
func (Txn) TableName() string {
	return "b_txn"
}
