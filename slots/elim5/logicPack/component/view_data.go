package component

import (
	"elim5/enum"
	"elim5/global"
	"elim5/logicPack/base"
	"elim5/utils/helper"
	"strings"
)

type ViewData struct {
	Data         string    `json:"data"`
	Bet          int       `json:"bet"`          // 压注
	Gain         int       `json:"gain"`         // 赢钱
	PayTableId   []uint    `json:"payTableId"`   // payTableId
	PayTableMuls []float64 `json:"payTableMuls"` // payTable倍数
	WildMuls     []float64 `json:"wildMuls"`     // 百搭倍数
	JackpotMul   float64   `json:"jackpotMul"`   // 奖池倍数
	FreeSpin     int       `json:"freeSpin"`     // 免费转次数
}

func (s *Spin) DumpPayTable() string {
	var arr []string
	for _, table := range s.PayTables {
		arr = append(arr, table.Dump())
	}
	return strings.Join(arr, "\n")
}

//func (s *Spin) Dump() *ViewData {
//	var view = &ViewData{}
//	view.Data = s.DumpGameData()
//	view.Bet = s.Bet
//	view.Gain = s.Gain
//	for _, table := range s.PayTables {
//		view.PayTableId = append(view.PayTableId, table.Id)
//		view.PayTableMuls = append(view.PayTableMuls, table.Multiple)
//	}
//	for _, wild := range s.WildList {
//		view.WildMuls = append(view.WildMuls, wild.Multiple)
//	}
//	if s.Jackpot != nil {
//		view.JackpotMul = s.Jackpot.End
//	}
//	return view
//}

func (s *Spin) DumpGameData() string {
	data := helper.ListConversion(s.InitDataList, func(tag *base.Tag) *base.Label {
		return tag.Dump()
	})
	vertical := helper.ArrVertical(data)

	var slot []int
	slot = []int{enum.SlotId23, enum.SlotId18} //这两台是中间列 多一个标签 需要特殊处理
	if helper.InArr(s.GetSlotId(), slot) && len(s.InitDataList[0]) != len(s.InitDataList[1]) {
		vertical = append(vertical, []*base.Label{&base.Label{}, data[1][3], &base.Label{}})
	}
	detail, err := global.Json.MarshalToString(vertical)
	if err != nil {
		return ""
	}
	return detail
}

func dataStr[T any](arr [][]T, f func(T) string, sep string) string {
	arr = helper.ArrVertical(arr)
	str := ""
	for i, row := range arr {
		for ii, col := range row {
			str += f(col)
			if ii < len(row)-1 {
				str += ","
			}
		}
		if i < len(arr)-1 {
			str += sep
		}
	}
	return str
}
