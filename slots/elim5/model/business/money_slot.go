package business

import (
	"elim5/enum"
	"elim5/global"
	"elim5/utils/helper"
)

// MoneySlot 结构体
type MoneySlot struct {
	global.GVA_MODEL
	Type     int    `json:"type" form:"type" gorm:"uniqueIndex:idx_unique;default:0;column:type;comment:类型;size:32;"`
	Date     uint   `json:"date" form:"date" gorm:"uniqueIndex:idx_unique;default:0;column:date;comment:日期;size:32;"`
	SlotId   uint   `json:"slotId" form:"slotId" gorm:"uniqueIndex:idx_unique;column:slot_id;default:0;comment:机器编号;size:32;"`
	Currency string `json:"currency" form:"currency" gorm:"uniqueIndex:idx_unique;column:currency;comment:货币;size:10;"`
	Rtp      int    `json:"rtp" form:"rtp" gorm:"uniqueIndex:idx_unique;column:rtp;comment:Rtp;size:16;default:1;"`

	LaunchDate uint    `json:"launchDate" form:"launchDate" gorm:"column:launch_date;comment:上线日期;size:32;"`
	Bets       int64   `json:"bets" form:"bets" gorm:"column:bets;default:0;comment:消耗金币;size:64;"`
	Wins       int64   `json:"wins" form:"wins" gorm:"column:wins;default:0;comment:产出金币;size:64;"`
	Rate       float64 `json:"rate" form:"rate" gorm:"column:rate;comment:返还比;size:14;"`
	PeopleNum  int     `json:"peopleNum" form:"peopleNum" gorm:"column:people_num;default:0;comment:游玩人数;size:32;"`

	SpinNum       int `json:"spinNum" form:"spinNum" gorm:"column:spin_num;default:0;comment:Spin次数;size:32;"`
	SpinNumCommon int `json:"spinNumCommon" form:"spinNumCommon" gorm:"column:spin_num_common;default:0;comment:Spin次数(普通);size:32;"`
	SpinNumFree   int `json:"spinNumFree" form:"spinNumFree" gorm:"column:spin_num_free;default:0;comment:Spin次数(购买免费);size:32;"`
	SpinNumRe     int `json:"spinNumRe" form:"spinNumRe" gorm:"column:spin_num_re;default:0;comment:Spin次数(购买重转);size:32;"`

	SpinAvg float64 `json:"spinAvg" form:"spinAvg" gorm:"column:spin_avg;comment:平均Spin次数;size:14;"`

	WinNum int `json:"winNum" form:"winNum" gorm:"column:win_num;default:0;comment:中奖次数;size:32;"`

	Win10Num        int `json:"win10Num" form:"win10Num" gorm:"column:win10_num;default:0;comment:10倍中奖次数;size:32;"`
	Win10NumCommon  int `json:"win10NumCommon" form:"win10NumCommon" gorm:"column:win10_num_common;default:0;comment:10倍中奖次数(普通);size:32;"`
	Win10NumBuyFree int `json:"win10NumFree" form:"win10NumFree" gorm:"column:win10_num_free;default:0;comment:10倍中奖次数(购买免费);size:32;"`
	Win10NumBuyRe   int `json:"win10NumRe" form:"win10NumRe" gorm:"column:win10_num_re;default:0;comment:10倍中奖次数(购买重转);size:32;"`

	UserIdList    []byte   `json:"userIdList" form:"userIdList" gorm:"type:blob;column:user_id_list;comment:用户ID列表;"`
	UserIdListArr []uint64 `json:"userIdListArr" form:"userIdListArr" gorm:"-"`

	MerchantId uint `json:"merchantId" form:"merchantId" gorm:"-"`
}

// TableName MoneySlot 表名
func (MoneySlot) TableName() string {
	return "b_money_slot"
}

// SumRecordsToDay 通过一组记录计算天统计数据
func SumRecordsToDay(date uint, arr []*SlotRecord) *MoneySlot {
	s := &MoneySlot{
		Type:    enum.MoneySlotTypeDay,
		Date:    date,
		SpinNum: len(arr),
	}
	if s.SpinNum == 0 {
		return s
	}
	s.SlotId = arr[0].SlotId
	s.Currency = arr[0].Currency
	s.Rtp = arr[0].Rtp

	var (
		m       = map[uint]struct{}{}
		userIds []uint64
	)
	// 计算合计数值
	for _, v := range arr {
		// 累加押注 赢钱 赢钱数
		s.Bets += v.TotalBet
		s.Wins += int64(v.Gain)
		if v.Gain > 0 {
			s.WinNum++
		}

		// 累加各种类型的spin次数
		switch v.BetType {
		case enum.BetTypeBuyFree:
			s.SpinNumFree++
		case enum.BetTypeBuyRe:
			s.SpinNumRe++
		default:
			s.SpinNumCommon++
		}

		// 累加10倍以上中奖次数
		if int64(v.Gain) >= v.TotalBet*10 {
			s.Win10Num++
			switch v.BetType {
			case enum.BetTypeCommon, enum.BetTypeRaise:
				s.Win10NumCommon++
			case enum.BetTypeBuyFree:
				s.Win10NumBuyFree++
			case enum.BetTypeBuyRe:
				s.Win10NumBuyRe++
			}
		}

		// 统计用户数量
		if _, ok := m[v.UserId]; !ok {
			m[v.UserId] = struct{}{}
			userIds = append(userIds, uint64(v.UserId))
		}
	}
	s.PeopleNum = len(userIds)
	s.Rate = helper.Div(s.Wins, s.Bets)
	s.SpinAvg = helper.Div(s.SpinNum, s.PeopleNum)
	s.UserIdList = helper.IdsToProto(userIds)
	return s
}

func MoneySlotSumTotal(arr []*MoneySlot, typ int) *MoneySlot {
	s := &MoneySlot{
		Type: typ,
	}
	if len(arr) == 0 {
		return s
	}
	if typ == enum.MoneySlotTypeMonth {
		date := helper.Itoa(arr[0].Date)
		if len(date) > 6 {
			date = date[0:6]
		}
		s.Date = uint(helper.Atoi(date + "01"))
	}
	var (
		m       = map[uint]struct{}{}
		userIds []uint64
	)
	s.SlotId = arr[0].SlotId
	s.Currency = arr[0].Currency
	s.Rtp = arr[0].Rtp
	for _, data := range arr {
		s.Wins += data.Wins
		s.Bets += data.Bets
		s.SpinNum += data.SpinNum
		s.SpinNumCommon += data.SpinNumCommon
		s.SpinNumFree += data.SpinNumFree
		s.SpinNumRe += data.SpinNumRe
		s.WinNum += data.WinNum
		s.Win10Num += data.Win10Num
		s.Win10NumCommon += data.Win10NumCommon
		s.Win10NumBuyFree += data.Win10NumBuyFree
		s.Win10NumBuyRe += data.Win10NumBuyRe
		for _, id := range helper.ProtoToIds(data.UserIdList) {
			if _, ok := m[uint(id)]; !ok {
				m[uint(id)] = struct{}{}
				userIds = append(userIds, id)
			}
		}
	}
	s.UserIdList = helper.IdsToProto(userIds)
	s.PeopleNum = len(userIds)
	s.Rate = helper.Div(s.Wins, s.Bets)
	return s
}
