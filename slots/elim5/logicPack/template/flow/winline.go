package flow

import (
	"elim5/logicPack/base"
	"elim5/pbs/common"
	"elim5/utils/helper"
)

type WinLine struct {
	Tags        []*base.Tag
	Mul         float64
	Integral    int //标签个数（积分）
	Skill       int //技能
	Orientation int //0 左 1 右
	Name        string
	Reason      int //理由 (3 消除)
	GroupCount  int
	MaxCount    int
	LineNum     int
	TableMul    float64
}

func (w *WinLine) ToAck(bet int) (winLineAck *common.Tags) {
	winLineAck = &common.Tags{
		Tags:        make([]*common.Tag, 0),
		Amount:      helper.MulToInt(bet, w.Mul),
		Point:       int32(w.Integral),
		Orientation: int32(w.Orientation),
	}
	for _, tag := range w.Tags {
		winLineAck.Tags = append(winLineAck.Tags, tag.ToTagAck())
	}
	return
}

func (w *WinLine) TransformIntegral() {
	// 当前是一个标签一个积分
	w.Integral = len(w.Tags)
}

func NewWinLine(tags []*base.Tag, mul float64, name string, integral, skill, orientation, reason int) *WinLine {
	return &WinLine{
		Tags:        tags,
		Mul:         mul,
		Integral:    integral,
		Skill:       skill,
		Orientation: orientation,
		Name:        name,
		Reason:      reason,
	}
}
