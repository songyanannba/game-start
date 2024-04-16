package template

import (
	"elim5/logicPack/base"
	"elim5/logicPack/template/flow"
	"elim5/utils/helper"
	"github.com/samber/lo"
)

//第八台游戏逻辑

// GetWinLine 获取一条赢钱划线
func (s *SpinInfo) GetWinLine(tags []*base.Tag) *flow.WinLine {
	if len(tags) == 0 {
		return &flow.WinLine{
			Tags: make([]*base.Tag, 0),
			Mul:  0,
		}
	}
	return &flow.WinLine{
		Tags:     helper.CopyList(tags),
		Mul:      s.MatchTagsWin(tags),
		Integral: len(tags),
	}
}

// GetWinLines 获取多条赢钱划线
func (s *SpinInfo) GetWinLines(tagList [][]*base.Tag) []*flow.WinLine {
	return lo.FilterMap(tagList, func(tags []*base.Tag, i int) (*flow.WinLine, bool) {
		line := s.GetWinLine(tags)
		if line == nil {
			return nil, false
		}
		return line, true
	})
}

// MatchTagsWin 获取一条划线的倍率
func (s *SpinInfo) MatchTagsWin(tags []*base.Tag) (mul float64) {
	nowPayTable := s.payTable[tags[0].Name]
	mul = 0.0
	for _, table := range nowPayTable {
		if len(tags) >= table.Num {
			mul = table.Mul
		} else {
			return mul
		}
	}
	return mul
}

// MatchTagsCountWin 获取一条划线的倍率
func (s *SpinInfo) MatchTagsCountWin(name string, count int) (mul float64) {
	nowPayTable := s.payTable[name]
	mul = 0.0
	for _, table := range nowPayTable {
		if count >= table.Num {
			mul = table.Mul
		} else {
			return mul
		}
	}
	return mul
}
