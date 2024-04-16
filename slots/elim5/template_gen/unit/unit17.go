package unit

import (
	"elim5/enum"
	"elim5/logicPack"
	"elim5/logicPack/component"
	"elim5/service/template_gen/gen_com"
	"elim5/service/test/handler/repeat/unit17"
	"elim5/utils/helper"
	"fmt"
)

type Slot17 struct {
	gTem     *gen_com.GenTemplate
	testSlot *unit17.Unit
}

func GetUnit17(t *gen_com.GenTemplate) *Slot17 {
	return &Slot17{
		testSlot: unit17.NewUnit(),
		gTem:     t,
	}
}

func (s *Slot17) RunTem() ([]*component.Spin, error) {
	sp := &component.Spin{
		Options: &component.Options{
			IsFree:     s.gTem.TemGen.Type == enum.SpinAckType2FreeSpin,
			IsTemGen:   true,
			Raise:      helper.If(s.gTem.TemGen.Type == enum.SpinAckType1BranchRaise, helper.MulToInt(s.gTem.Config.Raise, 100), 0),
			IsMustFree: s.gTem.TemGen.Type == enum.SpinAckType1BranchBuyFree,
		},
		Bet:  100,
		Gain: 0,
	}
	sp.Config = s.gTem.Config
	tem, err := slot.GetMachineTem(sp, s.gTem)
	if err != nil {
		return nil, err
	}
	err = tem.Exec()
	if err != nil {
		return nil, err
	}
	spins := []*component.Spin{tem.GetSpin()}
	spins = append(spins, tem.GetSpins()...)
	return spins, nil
}

// Calculate 累加数据
func (s *Slot17) Calculate(spins []*component.Spin) {
	s.testSlot.Calculate(spins)
}

// GetLogicResult slot8 计算数据
func (s *Slot17) GetLogicResult() *gen_com.LogicResult {
	res := s.testSlot.Result
	str := s.testSlot.GetDetail()
	fmt.Println(str)

	if s.gTem.TemGen.Type == enum.SpinAckType1NormalSpin {

	}
	switch s.gTem.TemGen.Type {
	case enum.SpinAckType1NormalSpin:
		return &gen_com.LogicResult{
			ScatterRatio: float64(res.NormalSca) / float64(res.NormalCount),
			GainRatio:    float64(res.NormalWin) / float64(res.N2),
			WinRatio:     float64(res.NormalWinCount) / float64(res.NormalCount),
			RemoveRate:   float64(res.N33+res.N35+res.N37) / float64(res.NormalCount),
		}
	case enum.SpinAckType2FreeSpin:
		return &gen_com.LogicResult{
			ScatterRatio: float64(res.FreeSca) / float64(res.FreeCount),
			GainRatio:    float64(res.FreeWin) / float64(res.N2),
			WinRatio:     float64(res.FreeWinCount) / float64(res.FreeCount),
			RemoveRate:   float64(res.N79+res.N81+res.N83) / float64(res.N97),
		}
	case enum.SpinAckType1BranchRaise:
		return &gen_com.LogicResult{
			ScatterRatio: float64(res.NormalSca) / float64(res.NormalCount),
			GainRatio:    float64(res.NormalWin) / float64(res.N2) * (s.gTem.Config.Raise + 1),
			WinRatio:     float64(res.NormalWinCount) / float64(res.NormalCount),
			RemoveRate:   float64(res.N33+res.N35+res.N37) / float64(res.NormalCount),
		}
	case enum.SpinAckType1BranchBuyFree:
		return &gen_com.LogicResult{
			ScatterRatio: float64(res.FreeSca) / float64(res.FreeCount),
			GainRatio:    (float64(res.FreeWin) + float64(res.NormalWin)) / float64(res.N2) * (s.gTem.Config.BuyFee + 1),
			WinRatio:     float64(res.FreeWinCount) / float64(res.FreeCount),
			RemoveRate:   float64(res.N79+res.N81+res.N83) / float64(res.N97),
		}
	default:
		return &gen_com.LogicResult{}
	}

}

// GetStatus slot15 获取状态
func (s *Slot17) GetStatus() (totalInterval float64, str string, ok bool) {
	r := s.GetLogicResult()
	totalInterval = s.gTem.GetDisparity(r)
	str, ok = s.Adjust(r)
	return
}

// Adjust slot8 调整数据
func (s *Slot17) Adjust(genc *gen_com.LogicResult) (str string, ok bool) {
	//global.GVA_LOG.Info(fmt.Sprintf("%v", genc))
	condKeys := []string{gen_com.RemoveRateCond, gen_com.ScaTriggerCond, gen_com.WinRateCond, gen_com.GainRatioCond}
	var adjust string
	for _, key := range condKeys {
		cond := s.gTem.GetCond(key)
		switch key {
		case gen_com.GainRatioCond:
			adjust = cond.Adjust(genc.GainRatio, s.gTem)
		case gen_com.ScaTriggerCond:
			adjust = cond.Adjust(genc.ScatterRatio, s.gTem)
		case gen_com.RemoveRateCond:
			adjust = cond.Adjust(genc.RemoveRate, s.gTem)
		case gen_com.WinRateCond:
			adjust = cond.Adjust(genc.WinRatio, s.gTem)
		}
		if adjust != "ok" {
			return fmt.Sprintf("%v %v %v \n", genc, key, adjust), false
		}
	}
	return fmt.Sprintf("%v ok", genc), true
}
