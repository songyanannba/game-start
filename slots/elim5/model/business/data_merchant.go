// 自动生成模板DataMerchant
package business

import (
	"elim5/enum"
	"elim5/global"
	"elim5/pbs/common"
	"elim5/utils/helper"
	"fmt"
)

// DataMerchant 结构体
type DataMerchant struct {
	global.GVA_MODEL
	Type              uint8   `json:"type" form:"type" gorm:"uniqueIndex:idx_unique;column:type;default:1;comment:;size:8;"`
	Date              int     `json:"date" form:"date" gorm:"uniqueIndex:idx_unique;column:date;default:0;comment:;size:32;"`
	SetId             int     `json:"setId" form:"setId" gorm:"uniqueIndex:idx_unique;column:set_id;default:0;comment:;size:32;"`
	MerchantId        int     `json:"merchantId" form:"merchantId" gorm:"uniqueIndex:idx_unique;column:merchant_id;default:0;comment:;size:32;"`
	SlotId            int     `json:"slotId" form:"slotId" gorm:"uniqueIndex:idx_unique;column:slot_id;default:0;comment:;size:32;"`
	Currency          string  `json:"currency" form:"currency" gorm:"uniqueIndex:idx_unique;column:currency;comment:;size:255;"`
	AddPlayer         int     `json:"addPlayer" form:"addPlayer" gorm:"column:add_player;default:0;comment:;size:64;"`
	AddPlayerData     []byte  `json:"addPlayerData" form:"addPlayerData" gorm:"type:longblob;column:add_player_data;comment:;"`
	Player            int     `json:"player" form:"player" gorm:"column:player;default:0;comment:;size:64;"`
	PlayerData        []byte  `json:"playerData" form:"playerData" gorm:"type:longblob;column:player_data;comment:;"`
	HistoryPlayerData []byte  `json:"historyPlayerData" form:"historyPlayerData" gorm:"type:longblob;column:history_player_data;comment:;"`
	SpinCount         int     `json:"spinCount" form:"spinCount" gorm:"column:spin_count;default:0;comment:;size:64;"`
	SpinAvg           float64 `json:"spinAvg" form:"spinAvg" gorm:"type:float;column:spin_avg;default:0;comment:;size:64;"`
	Bk                int     `json:"bk" form:"bk" gorm:"column:bk;default:0;comment:;size:64;"`
	BkPlayerData      []byte  `json:"bkPlayerData" form:"bkPlayerData" gorm:"type:longblob;column:bk_player_data;comment:;"`
	Rt                int     `json:"rt" form:"rt" gorm:"column:rt;default:0;comment:;size:64;"`
	Rta               int64   `json:"rta" form:"rta" gorm:"column:rta;default:0;comment:;size:64;"`
	Bet               int64   `json:"bet" form:"bet" gorm:"column:bet;default:0;comment:;size:64;"`
	Payout            int64   `json:"payout" form:"payout" gorm:"column:payout;default:0;comment:;size:64;"`
	Win               int64   `json:"win" form:"win" gorm:"column:win;default:0;comment:;size:64;"`

	AddPlayerList     map[uint64]bool  `gorm:"-"`
	PlayerList        map[uint64]bool  `gorm:"-"`
	HistoryPlayerList map[uint64]bool  `gorm:"-"`
	BkPlayerDataList  map[uint64]int64 `gorm:"-"`

	AddPlayerListArr     []uint64 `json:"addPlayerList" form:"addPlayerList" gorm:"-"`
	PlayerListArr        []uint64 `json:"playerList" form:"playerList" gorm:"-"`
	HistoryPlayerListArr []uint64 `json:"historyPlayerList" form:"historyPlayerList" gorm:"-"`
	BkPlayerDataListArr  string   `json:"bkPlayerDataList" form:"bkPlayerDataList" gorm:"-"`
}

// TableName DataMerchant 表名
func (DataMerchant) TableName() string {
	return "b_data_merchant"
}

func (d *DataMerchant) Conversion() {
	d.AddPlayerList = helper.MapKeysToBoolMap(helper.ProtoToIds(d.AddPlayerData))
	d.PlayerList = helper.MapKeysToBoolMap(helper.ProtoToIds(d.PlayerData))
	d.HistoryPlayerList = helper.MapKeysToBoolMap(helper.ProtoToIds(d.HistoryPlayerData))
	d.BkPlayerDataList = helper.ProtoToBkPlayer(d.BkPlayerData)
}

func (d *DataMerchant) Reverse() {
	d.AddPlayerData = helper.IdsToProto(helper.MapGetKeys(d.AddPlayerList))
	d.PlayerData = helper.IdsToProto(helper.MapGetKeys(d.PlayerList))
	d.HistoryPlayerData = helper.IdsToProto(helper.MapGetKeys(d.HistoryPlayerList))
	d.BkPlayerData = helper.BkPlayerToProto(d.BkPlayerDataList)
}

func (d *DataMerchant) ConversionArr() {
	d.AddPlayerListArr = helper.ProtoToIds(d.AddPlayerData)
	d.PlayerListArr = helper.ProtoToIds(d.PlayerData)
	d.HistoryPlayerListArr = helper.ProtoToIds(d.HistoryPlayerData)
	d.BkPlayerDataListArr = fmt.Sprintf("%+v", helper.ProtoToBkPlayer(d.BkPlayerData))
}

func (d *DataMerchant) GetPreviousUsers() DataMerchant {
	var dataMerchant DataMerchant
	global.NOLOG_DB.Model(&DataMerchant{}).
		Where("`type` =? and set_id = ? and merchant_id =?",
			d.Type, d.SetId, d.MerchantId).
		Where("date < ?", d.Date).
		Order("date desc").
		First(&dataMerchant)
	return dataMerchant
}

func (d *DataMerchant) GetDataSurvival() *DataSurvival {
	return &DataSurvival{
		Type:       enum.SurvivalTypeMerchant,
		Date:       d.Date,
		MerchantId: d.MerchantId,
		SlotId:     0,
		Currency:   d.Currency,
		SurvivalMap: map[uint32]*common.SurvivalData{
			0: {
				Player: 0,
				Bk:     0,
				Rt:     0,
				Rta:    0,
			},
		},
	}
}
