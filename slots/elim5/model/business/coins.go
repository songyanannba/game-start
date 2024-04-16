// 自动生成模板Coins
package business

import (
	"elim5/global"
)

// Coins 结构体
type Coins struct {
	global.GVA_MODEL
	CoinType uint   `json:"coinType" form:"coinType" gorm:"column:coin_type;default:0;comment:代币类型;size:32;"`
	BetNum   string `json:"betNum" form:"betNum" gorm:"column:bet_num;comment:机器押注;"`
	Remark   string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
}

// TableName Coins 表名
func (Coins) TableName() string {
	return "b_coins"
}
