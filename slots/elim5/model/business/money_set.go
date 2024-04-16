// 自动生成模板MoneySet
package business

import (
	"elim5/enum"
	"elim5/global"
	"elim5/utils/helper"
)

// MoneySet 结构体
type MoneySet struct {
	global.GVA_MODEL
	Type       uint8   `json:"type" form:"type" gorm:"uniqueIndex:idx_unique;column:type;default:1;comment:类型;size:8;"`
	Date       uint    `json:"date" form:"date" gorm:"uniqueIndex:idx_unique;column:date;default:0;comment:日期;size:32;"`
	SetId      uint    `json:"setId" form:"setId" gorm:"uniqueIndex:idx_unique;column:set_id;default:0;comment:集合编号;size:32;"`
	MerchantId uint    `json:"merchantId" form:"merchantId" gorm:"uniqueIndex:idx_unique;column:merchant_id;default:0;comment:商户编号;size:32;"`
	Currency   string  `json:"currency" form:"currency" gorm:"uniqueIndex:idx_unique;column:currency;comment:货币;size:10;"`
	SlotId     uint    `json:"slotId" form:"slotId" gorm:"uniqueIndex:idx_unique;column:slot_id;default:0;comment:机台编号;size:32;"`
	AddPlayer  int     `json:"addPlayer" form:"addPlayer" gorm:"column:add_player;default:0;comment:新增人数;size:32;"`
	Player     int     `json:"player" form:"player" gorm:"column:player;default:0;comment:活跃人数;size:32;"`
	SpinCount  int     `json:"spinCount" form:"spinCount" gorm:"column:spin_count;default:0;comment:游戏次数;size:32;"`
	SpinAvg    float64 `json:"spinAvg" form:"spinAvg" gorm:"column:spin_avg;comment:人均游戏次数;size:22;"`
	Bk         int     `json:"bk" form:"bk" gorm:"column:bk;default:0;comment:破产次数;size:32;"`
	Rt         int     `json:"rt" form:"rt" gorm:"column:rt;default:0;comment:破产充值次数;size:32;"`
	Rta        float32 `json:"rta" form:"rta" gorm:"column:rta;default:0;comment:破产充值金额;size:64;"`
	Bet        int     `json:"bet" form:"bet" gorm:"column:bet;default:0;comment:押注金额;size:64;"`
	Payout     int     `json:"payout" form:"payout" gorm:"column:payout;default:0;comment:获奖金额;size:64;"`
	Win        int     `json:"win" form:"win" gorm:"column:win;default:0;comment:赢钱金额;size:64;"`
	GroupBy    uint8   `json:"groupBy" form:"groupBy" gorm:"uniqueIndex:idx_unique;column:group_by;default:1;comment:分组;size:8;"`

	NewUserIdList    []byte `json:"newUserIdList" form:"newUserIdList" gorm:"type:longblob;column:new_user_id_list;comment:新增用户ID列表;"`
	ActiveUserIdList []byte `json:"activeUserIdList" form:"activeUserIdList" gorm:"type:longblob;column:active_user_id_list;comment:活跃用户ID列表;"`
	SumUserIdList    []byte `json:"sumUserIdList" form:"sumUserIdList" gorm:"type:longblob;column:sum_user_id_list;comment:累计用户ID列表;"`

	NewUserIdListArr    []uint64 `json:"newUserIdListArr" form:"newUserIdListArr" gorm:"-"`
	ActiveUserIdListArr []uint64 `json:"activeUserIdListArr" form:"activeUserIdListArr" gorm:"-"`
	SumUserIdListArr    []uint64 `json:"sumUserIdListArr" form:"sumUserIdListArr" gorm:"-"`
}

// TableName MoneySet 表名
func (MoneySet) TableName() string {
	return "b_money_set"
}

func (m *MoneySet) GetBeforeMoneySet() []uint64 {
	beforeDate := helper.BeforeDate(m.Date, m.Type)
	if m.Date == 0 {
		return []uint64{}
	}
	beforeMoneyMerchant := &MoneySet{
		SumUserIdListArr: []uint64{},
	}
	global.GVA_READ_DB.Model(&MoneySet{}).
		Where(" merchant_id = ? and set_id = ? and slot_id = ? and currency = ?  and `type` = ? and `group_by` = ?",
			m.MerchantId, m.SetId, m.SlotId, m.Currency, enum.MoneySlotTypeDay, m.GroupBy).
		Where("date <= ?", beforeDate).
		Order("date desc").
		First(&beforeMoneyMerchant)
	return helper.ProtoToIds(beforeMoneyMerchant.SumUserIdList)
}

func (m *MoneySet) Init() {
	m.NewUserIdListArr = helper.ProtoToIds(m.NewUserIdList)
	m.ActiveUserIdListArr = helper.ProtoToIds(m.ActiveUserIdList)
	m.SumUserIdListArr = helper.ProtoToIds(m.SumUserIdList)
}
