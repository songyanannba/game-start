// 自动生成模板DataPlayer
package business

import (
	"elim5/global"
	"elim5/utils/helper"
	"fmt"
)

// DataPlayer 结构体
type DataPlayer struct {
	global.GVA_MODEL
	Type         uint8   `json:"type" form:"type" gorm:"uniqueIndex:idx_unique;column:type;default:1;comment:;size:8;"`
	Date         int     `json:"date" form:"date" gorm:"uniqueIndex:idx_unique;column:date;default:0;comment:;size:32;"`
	UserId       int     `json:"userId" form:"userId" gorm:"uniqueIndex:idx_unique;column:user_id;default:0;comment:;size:32;"`
	SetId        int     `json:"setId" form:"setId" gorm:"uniqueIndex:idx_unique;column:set_id;default:0;comment:;size:32;"`
	MerchantId   int     `json:"merchantId" form:"merchantId" gorm:"uniqueIndex:idx_unique;column:merchant_id;default:0;comment:;size:32;"`
	SlotId       int     `json:"slotId" form:"slotId" gorm:"uniqueIndex:idx_unique;column:slot_id;default:0;comment:;size:32;"`
	Currency     string  `json:"currency" form:"currency" gorm:"uniqueIndex:idx_unique;column:currency;comment:;size:255;"`
	SpinCount    int     `json:"spinCount" form:"spinCount" gorm:"column:spin_count;default:0;comment:;size:64;"`
	Bk           int     `json:"bk" form:"bk" gorm:"column:bk;default:0;comment:;size:64;"`
	Rt           int     `json:"rt" form:"rt" gorm:"column:rt;default:0;comment:;size:64;"`
	Rta          int64   `json:"rta" form:"rta" gorm:"column:rta;comment:;size:64;"`
	RBetCount    int     `json:"rBetCount" form:"rBetCount" gorm:"column:r_bet_count;default:0;comment:;size:64;"`
	FBetCount    int     `json:"fBetCount" form:"fBetCount" gorm:"column:f_bet_count;default:0;comment:;size:64;"`
	Arta         float64 `json:"arta" form:"arta" gorm:"column:arta;default:0;comment:;size:64;"`
	Bet          int64   `json:"bet" form:"bet" gorm:"column:bet;default:0;comment:;size:64;"`
	BetAvg       float64 `json:"betAvg" form:"betAvg" gorm:"column:bet_avg;default:0;comment:;size:64;"`
	Payout       int64   `json:"payout" form:"payout" gorm:"column:payout;default:0;comment:;size:64;"`
	Win          int64   `json:"win" form:"win" gorm:"column:win;default:0;comment:;size:64;"`
	BkPlayerData []byte  `json:"bkPlayerData" form:"bkPlayerData" gorm:"type:longblob;column:bk_player_data;comment:;"`

	PlayerName          []uint64         `json:"playerName" form:"playerName" gorm:"-"`
	BkPlayerDataList    map[uint64]int64 `gorm:"-"`
	BkPlayerDataListArr string           `json:"bkPlayerDataList" form:"bkPlayerDataList" gorm:"-"`
}

// TableName DataPlayer 表名
func (DataPlayer) TableName() string {
	return "b_data_player"
}

func (d *DataPlayer) Conversion() {
	d.BkPlayerDataList = helper.ProtoToBkPlayer(d.BkPlayerData)
}

func (d *DataPlayer) Reverse() {
	d.BkPlayerData = helper.BkPlayerToProto(d.BkPlayerDataList)
}

func (d *DataPlayer) ConversionArr() {
	d.BkPlayerDataListArr = fmt.Sprintf("%+v", helper.ProtoToBkPlayer(d.BkPlayerData))
}

func (d *DataPlayer) GetPreviousUsers() DataPlayer {
	var dataPlayer DataPlayer
	global.NOLOG_DB.Model(&DataPlayer{}).
		Where("`type` =? ",
			d.Type).
		Where("date < ?", d.Date).
		Order("date desc").
		First(&dataPlayer)
	return dataPlayer
}
