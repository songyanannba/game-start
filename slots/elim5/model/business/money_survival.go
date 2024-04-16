// 自动生成模板MoneySurvival
package business

import (
	"elim5/enum"
	"elim5/global"
	"elim5/pbs/common"
	"elim5/utils/helper"
)

// MoneySurvival 结构体
type MoneySurvival struct {
	global.GVA_MODEL
	Type             uint8  `json:"type" form:"type" gorm:"uniqueIndex:idx_unique;column:type;default:1;comment:类型;size:8;"`
	Date             uint   `json:"date" form:"date" gorm:"uniqueIndex:idx_unique;column:date;default:0;comment:日期;size:32;"`
	MerchantId       uint   `json:"merchantId" form:"merchantId" gorm:"uniqueIndex:idx_unique;column:merchant_id;default:0;comment:商户id;size:32;"`
	Currency         string `json:"currency" form:"currency" gorm:"uniqueIndex:idx_unique;column:currency;comment:货币;size:255;"`
	AddPlayer        int    `json:"addPlayer" form:"addPlayer" gorm:"column:add_player;default:0;comment:新增人数;size:32;"`
	Player           int    `json:"player" form:"player" gorm:"column:player;default:0;comment:活跃人数;size:32;"`
	GroupBy          uint8  `json:"groupBy" form:"groupBy" gorm:"uniqueIndex:idx_unique;column:group_by;default:1;comment:分组;size:8;"`
	Survival         []byte `json:"survival" form:"survival" gorm:"type:blob;column:survival;comment:存活率;"`
	NewUserIdList    []byte `json:"newUserIdList" form:"newUserIdList" gorm:"type:longblob;column:new_user_id_list;comment:新增用户ID列表;"`
	ActiveUserIdList []byte `json:"activeUserIdList" form:"activeUserIdList" gorm:"type:longblob;column:active_user_id_list;comment:活跃用户ID列表;"`
	SumUserIdList    []byte `json:"sumUserIdList" form:"sumUserIdList" gorm:"type:longblob;column:sum_user_id_list;comment:累计用户ID列表;"`

	SurvivalInfo        *common.SurvivalMap `json:"survivalInfo" form:"survivalInfo" gorm:"-"`
	NewUserIdListArr    []uint64            `json:"newUserIdListArr" form:"newUserIdListArr" gorm:"-"`
	ActiveUserIdListArr []uint64            `json:"activeUserIdListArr" form:"activeUserIdListArr" gorm:"-"`
	SumUserIdListArr    []uint64            `json:"sumUserIdListArr" form:"sumUserIdListArr" gorm:"-"`
}

// TableName MoneySurvival 表名
func (MoneySurvival) TableName() string {
	return "b_money_survival"
}

func (m *MoneySurvival) GetBeforeMoneySurvival() []uint64 {
	beforeDate := helper.BeforeDate(m.Date, m.Type)
	beforeMoneyMerchant := &MoneySurvival{
		SumUserIdListArr: []uint64{},
	}
	global.GVA_READ_DB.Model(&MoneySurvival{}).
		Where("merchant_id = ?  and currency = ?  and `type` = ? and `group_by` = ?",
			m.MerchantId, m.Currency, enum.MoneySlotTypeDay, m.GroupBy).
		Where("date <= ?", beforeDate).
		Order("date desc").
		First(&beforeMoneyMerchant)
	return helper.ProtoToIds(beforeMoneyMerchant.SumUserIdList)
}

func (m *MoneySurvival) Init() {
	m.NewUserIdListArr = helper.ProtoToIds(m.NewUserIdList)
	m.ActiveUserIdListArr = helper.ProtoToIds(m.ActiveUserIdList)
	m.SumUserIdListArr = helper.ProtoToIds(m.SumUserIdList)
	m.SurvivalInfo = helper.ProtoToSurvivalMap(m.Survival)
}
