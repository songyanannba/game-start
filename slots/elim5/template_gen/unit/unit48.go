package unit

import (
	"elim5/enum"
	"elim5/logicPack"
	"elim5/logicPack/component"
	"elim5/service/template_gen/gen_com"
	"elim5/service/test/handler/repeat/unit48"
	"elim5/utils/helper"
	"fmt"
)

type Slot48 struct {
	SlotFaceImp
	gTem       *gen_com.GenTemplate
	testSlot48 *unit48.Unit
}

func GetUnit48(t *gen_com.GenTemplate) *Slot48 {
	return &Slot48{
		testSlot48: unit48.NewUnit(),
		gTem:       t,
	}
}

func (s *Slot48) RunTem() ([]*component.Spin, error) {
	sp := &component.Spin{
		Options: &component.Options{
			IsFree:       helper.If(s.gTem.TemGen.Type == enum.SpinAckType2FreeSpin, true, false),
			IsTemGen:     true,
			RatioConfirm: s.gTem.RatioConfirm,
			//IsMustFree: true,
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

func (s *Slot48) Calculate(spins []*component.Spin) {
	s.testSlot48.Calculate(spins) //
}

func (s *Slot48) GetStatus() (float64, string, bool) {
	r := s.GetLogicResult()
	totalInterval := s.gTem.GetDisparity(r)
	str, ok := s.Adjust(r)
	return totalInterval, str, ok
}

func (s *Slot48) GetLogicResult() *gen_com.LogicResult {
	res := s.testSlot48.Result

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
	default:
		return &gen_com.LogicResult{}
	}
}

// Adjust  调整数据
func (s *Slot48) Adjust(genc *gen_com.LogicResult) (str string, ok bool) {
	condKeys := []string{gen_com.GainRatioCond, gen_com.ScaTriggerCond, gen_com.WinRateCond, gen_com.RemoveRateCond}
	//condKeys := []string{gen_com.GainRatioCond, gen_com.ScaTriggerCond}
	var adjust string
	for _, key := range condKeys {
		cond := s.gTem.GetCond(key)
		switch key {
		case gen_com.GainRatioCond: //向上 向下 微调
			adjust = cond.Adjust(genc.GainRatio, s.gTem)
		case gen_com.ScaTriggerCond: //调整scat
			adjust = cond.Adjust(genc.ScatterRatio, s.gTem)
		case gen_com.RemoveRateCond: //连消向上/下微调
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
