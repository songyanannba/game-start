// 自动生成模板ReportWinLoss
package business

import (
	"elim5/global"
	"elim5/utils/helper"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/hints"
)

// ReportWinLoss 结构体
type ReportWinLoss struct {
	global.GVA_MODEL
	Date       uint    `json:"date" form:"date" gorm:"column:date;default:0;comment:Date;size:32;"`
	MerchantId uint    `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;default:0;comment:Merchant ID;size:32;"`
	Agent      string  `json:"agent" form:"agent" gorm:"column:agent;comment:Agent;size:50;"`
	Name       string  `json:"name" form:"name" gorm:"column:name;comment:Name;size:50;"`
	Currency   string  `json:"currency" form:"currency" gorm:"column:currency;comment:Currency;size:50;"`
	GameId     int     `json:"gameId" form:"gameId" gorm:"column:game_id;default:0;comment:Game ID;size:32;"`
	UserId     int     `json:"userId" form:"userId" gorm:"column:user_id;default:0;comment:User ID;size:64;"`
	Player     int     `json:"player" form:"player" gorm:"column:player;default:0;comment:Player;size:64;"`
	Tickets    int     `json:"tickets" form:"tickets" gorm:"column:tickets;default:0;comment:Tickets;size:64;"`
	Bet        int     `json:"bet" form:"bet" gorm:"column:bet;default:0;comment:Bet;size:64;"`
	Win        int     `json:"win" form:"win" gorm:"column:win;default:0;comment:Win;size:64;"`
	Margin     float64 `json:"margin" form:"margin" gorm:"column:margin;comment:Margin;size:14;"`
	WinNum     int64   `json:"winNum" form:"winNum" gorm:"column:win_num"`
	UserName   string  `json:"userName" form:"userName" gorm:"column:user_name"`
	IsCache    bool    `json:"isCache" form:"isCache" gorm:"-"`
}

// TableName ReportWinLoss 表名
func (ReportWinLoss) TableName() string {
	return "b_report_win_loss"
}

func SumMerchantReportWinLossListByDate(date uint, merchantId uint) (list []*ReportWinLoss, err error) {
	columns := []string{
		"sum(total_bet) as bet",
		"sum(gain) as win",
		"count(*) as tickets",
		"count(DISTINCT user_id) as player",
		"sum(case when gain > 0 then 1 else 0 end) as win_num",
		"currency",
		"merchant_id",
	}
	db := global.GVA_READ_DB.Model(&SlotRecord{}).
		Clauses(hints.UseIndex("idx_b_slot_record_created_at", "idx_b_slot_record_merchant_id", "idx_b_slot_record_date")).
		Group("currency, merchant_id").
		Select(columns).
		Where("date = ?", date)

	if merchantId != 0 {
		db = db.Where("merchant_id = ?", merchantId)
	}

	err = db.Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询当天数据失败", zap.Error(err))
		return
	}
	for _, report := range list {
		report.Date = date
		report.Name = helper.Itoa(date)
	}
	return list, nil
}
