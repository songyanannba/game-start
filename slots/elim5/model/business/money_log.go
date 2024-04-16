// 自动生成模板MoneyLog
package business

import (
	"elim5/enum"
	"elim5/global"
	"elim5/utils/helper"
	"gorm.io/gorm"
	"time"
)

// MoneyLog 结构体
type MoneyLog struct {
	global.GVA_MODEL
	Date        uint   `json:"date" form:"date" gorm:"index;default:0;column:date;comment:日期;size:32;"`
	UserId      uint   `json:"userId" form:"userId" gorm:"index;column:user_id;default:0;comment:用户编号;size:32;"`
	Action      uint   `json:"action" form:"action" gorm:"index;column:action;default:0;comment:操作类型;size:32;"`
	ActionType  uint   `json:"actionType" form:"actionType" gorm:"column:action_type;default:0;comment:子类型;size:32;"`
	CoinInitial int64  `json:"coinInitial" form:"coinInitial" gorm:"column:coin_initial;default:0;comment:初始金币;size:64;"`
	CoinChange  int64  `json:"coinChange" form:"coinChange" gorm:"column:coin_change;default:0;comment:金币变化;size:64;"`
	CoinResult  int64  `json:"coinResult" form:"coinResult" gorm:"column:coin_result;default:0;comment:金币结果;size:64;"`
	GameId      uint   `json:"gameId" form:"gameId" gorm:"index;column:game_id;default:0;comment:游戏编号;size:64;"`
	TxnId       string `json:"txnId" form:"txnId" gorm:"column:txn_id;comment:三方交易编号;size:50;"`
}

// TableName MoneyLog 表名
func (MoneyLog) TableName() string {
	return "b_money_log"
}

type MoneyLogOption func(*MoneyLog)

func MoneyLogWithRecharge() MoneyLogOption {
	return func(l *MoneyLog) {
		l.Action = enum.MoneyAction2Cash
		l.ActionType = enum.MoneyType200Recharge
	}
}

func MoneyLogWithGive() MoneyLogOption {
	return func(l *MoneyLog) {
		l.Action = enum.MoneyAction3System
		l.ActionType = enum.MoneyType300Give
	}
}

func MoneyLogWithTxnId(slotId uint, txnId string, actionType uint) MoneyLogOption {
	return func(l *MoneyLog) {
		l.TxnId = txnId
		l.GameId = slotId
		l.Action = enum.MoneyAction1Play
		l.ActionType = actionType
	}
}

func CreateMoneyLog(tx *gorm.DB, userId uint, gameId uint, b, c, a int64, action, actionType uint) (moneyLog *MoneyLog, err error) {
	moneyLog = &MoneyLog{
		Date:        helper.IntDate(time.Now()),
		UserId:      userId,
		GameId:      gameId,
		Action:      action,
		ActionType:  actionType,
		CoinInitial: b,
		CoinChange:  c,
		CoinResult:  a,
	}
	if tx == nil {
		tx = global.GVA_DB
	}
	err = tx.Create(moneyLog).Error
	return
}
