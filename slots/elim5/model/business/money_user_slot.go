// 自动生成模板MoneyUserSlot
package business

import (
	"elim5/global"
)

// MoneyUserSlot 结构体
type MoneyUserSlot struct {
	global.GVA_MODEL
	Type         int     `json:"type" form:"type" gorm:"uniqueIndex:idx_unique;column:type;default:0;comment:类型;size:32;"`
	Date         uint    `json:"date" form:"date" gorm:"uniqueIndex:idx_unique;column:date;default:0;comment:日期;size:32;"`
	UserId       int     `json:"userId" form:"userId" gorm:"uniqueIndex:idx_unique;column:user_id;default:0;comment:用户ID;size:32;"`
	Currency     string  `json:"currency" form:"currency" gorm:"uniqueIndex:idx_unique;column:currency;comment:货币;size:10;"`
	MerchantId   int     `json:"merchantId" form:"merchantId" gorm:"uniqueIndex:idx_unique;column:merchant_id;default:0;comment:商户ID;size:32;"`
	SlotId       int     `json:"slotId" form:"slotId" gorm:"uniqueIndex:idx_unique;column:slot_id;default:0;comment:游戏ID;size:32;"`
	BetAmount    int     `json:"betAmount" form:"betAmount" gorm:"column:bet_amount;default:0;comment:押注金额;size:32;"`
	BetNum       int     `json:"betNum" form:"betNum" gorm:"column:bet_num;default:0;comment:押注次数;size:32;"`
	AvgBet       float64 `json:"avgBet" form:"avgBet" gorm:"column:avg_bet;default:0;comment:平均押注;size:32;"`
	WinAmount    int     `json:"winAmount" form:"winAmount" gorm:"column:win_amount;default:0;comment:赢钱金额;size:32;"`
	WinNum       int     `json:"winNum" form:"winNum" gorm:"column:win_num;default:0;comment:赢钱次数;size:32;"`
	Win          int     `json:"win" form:"win" gorm:"column:win;default:0;comment:赢钱比;size:32;"`
	Rtp          float64 `json:"rtp" form:"rtp" gorm:"column:rtp;comment:返还比;size:10;"`
	WinRate      float64 `json:"winRate" form:"winRate" gorm:"column:win_rate;comment:获奖率;size:10;"`
	BigWinNum    int     `json:"bigWinNum" form:"bigWinNum" gorm:"column:big_win_num;default:0;comment:大赢钱次数;size:32;"`
	BigWinAmount int     `json:"bigWinAmount" form:"bigWinAmount" gorm:"column:big_win_amount;default:0;comment:大赢钱金额;size:32;"`
	BigWinRate   float64 `json:"bigWinRate" form:"bigWinRate" gorm:"column:big_win_rate;default:0;comment:大赢钱比例;size:32;"`
	NBetNum      int     `json:"nBetNum" form:"nBetNum" gorm:"column:n_bet_num;default:0;comment:普通转投注次数;size:32;"`
	NBetAmount   int     `json:"nBetAmount" form:"nBetAmount" gorm:"column:n_bet_amount;default:0;comment:普通转投注额;size:32;"`
	FBetNum      int     `json:"fBetNum" form:"fBetNum" gorm:"column:f_bet_num;default:0;comment:购买FS投注次数;size:32;"`
	FBetAmount   int     `json:"fBetAmount" form:"fBetAmount" gorm:"column:f_bet_amount;default:0;comment:购买FS投注额;size:32;"`
	RBetNum      int     `json:"rBetNum" form:"rBetNum" gorm:"column:r_bet_num;default:0;comment:购买RS投注次数;size:32;"`
	RBetAmount   int     `json:"rBetAmount" form:"rBetAmount" gorm:"column:r_bet_amount;default:0;comment:购买RS投注额;size:32;"`
	Bk           int     `json:"bk" form:"bk" gorm:"column:bk;default:0;comment:破产次数;size:32;"`
	Rt           int     `json:"rt" form:"rt" gorm:"column:rt;default:0;comment:破产充值次数;size:32;"`
	Rta          float32 `json:"rta" form:"rta" gorm:"column:rta;default:0;comment:破产充值金额;size:32;"`
	Arta         float64 `json:"arta" form:"arta" gorm:"column:arta;default:0;comment:平均破产充值;size:32;"`
}

// TableName MoneyUserSlot 表名
func (MoneyUserSlot) TableName() string {
	return "b_money_user_slot"
}
