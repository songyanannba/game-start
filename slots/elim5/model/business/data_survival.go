// 自动生成模板DataSurvival
package business

import (
	"elim5/global"
	"elim5/pbs/common"
	"elim5/utils/helper"
)

// DataSurvival 结构体
type DataSurvival struct {
	global.GVA_MODEL
	Type             uint8                           `json:"type" form:"type" gorm:"uniqueIndex:idx_unique;column:type;default:1;comment:类型;size:8;"`
	Date             int                             `json:"date" form:"date" gorm:"uniqueIndex:idx_unique;column:date;default:0;comment:日期;size:32;"`
	MerchantId       int                             `json:"merchantId" form:"merchantId" gorm:"uniqueIndex:idx_unique;column:merchant_id;default:0;comment:商户id;size:32;"`
	SlotId           int                             `json:"slotId" form:"slotId" gorm:"uniqueIndex:idx_unique;column:slot_id;default:0;comment:;size:32;"`
	Currency         string                          `json:"currency" form:"currency" gorm:"uniqueIndex:idx_unique;column:currency;comment:货币;size:255;"`
	Survival         []byte                          `json:"survival" form:"survival" gorm:"type:longblob;column:survival;comment:存活率;"`
	AddPlayerData    []byte                          `json:"addPlayerData" form:"addPlayerData" gorm:"type:longblob;column:add_player_data;comment:;"`
	Player           int                             `json:"player" form:"player" gorm:"column:player;default:0;comment:;size:64;"`
	AddPlayerListArr []uint64                        `json:"addPlayerList" form:"addPlayerList" gorm:"-"`
	PlayerListArr    []uint64                        `json:"playerList" form:"playerList" gorm:"-"`
	SurvivalMap      map[uint32]*common.SurvivalData `json:"survivalMap" form:"survivalMap" gorm:"-"`
}

// TableName DataSurvival 表名
func (DataSurvival) TableName() string {
	return "b_data_survival"
}

func (d *DataSurvival) Conversion() {
	//d.AddPlayerList = helper.MapKeysToBoolMap(helper.ProtoToIds(d.AddPlayerData))
	//d.PlayerList = helper.MapKeysToBoolMap(helper.ProtoToIds(d.PlayerData))
	//d.HistoryPlayerList = helper.MapKeysToBoolMap(helper.ProtoToIds(d.HistoryPlayerData))
	//d.BkPlayerDataList = helper.ProtoToBkPlayer(d.BkPlayerData)
}

func (d *DataSurvival) Reverse() {
	d.AddPlayerData = helper.IdsToProto(d.AddPlayerListArr)
	d.Survival = helper.SurvivalToProto(&common.SurvivalMap{DaySurvival: d.SurvivalMap})
}

func (d *DataSurvival) ConversionArr() {
	d.AddPlayerListArr = helper.ProtoToIds(d.AddPlayerData)
	d.SurvivalMap = helper.ProtoToSurvivalMap(d.Survival).DaySurvival
}
