package unit

import (
	"elim5/enum"
	"elim5/logicPack"
	"elim5/logicPack/component"
	"elim5/service/template_gen/gen_com"
	"elim5/service/test/handler/repeat/unit6"
	"fmt"

	"elim5/utils/helper"
)

type Slot6 struct {
	SlotFaceImp
	gTem      *gen_com.GenTemplate
	testSlot6 *unit6.Unit
}

func GetUnit6(t *gen_com.GenTemplate) *Slot6 {
	return &Slot6{
		testSlot6: unit6.NewUnit(),
		gTem:      t,
	}

}

func (s *Slot6) RunTem() ([]*component.Spin, error) {
	sp := &component.Spin{
		Options: &component.Options{
			IsFree:       helper.If(s.gTem.TemGen.Type == enum.SpinAckType2FreeSpin, true, false),
			IsTemGen:     true,
			RatioConfirm: s.gTem.RatioConfirm,
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

func (s *Slot6) Calculate(spins []*component.Spin) {
	s.testSlot6.Calculate(spins)
}

func (s *Slot6) GetStatus() (float64, string, bool) {
	r := s.GetLogicResult()
	totalInterval := s.gTem.GetDisparity(r) //6
	str, ok := s.Adjust(r)
	return totalInterval, str, ok
}

func (s *Slot6) GetLogicResult() *gen_com.LogicResult {
	res := s.testSlot6.Result
	gainRatio := float64(res.N1) / float64(res.N2)
	return &gen_com.LogicResult{
		ScatterRatio: 0,
		GainRatio:    gainRatio,
		WinRatio:     float64(res.N4) / float64(res.N3),
		RemoveRate:   float64(res.N33+res.N35+res.N37) / float64(res.N3),
	}
}

// Adjust slot8 调整数据
func (s *Slot6) Adjust(genc *gen_com.LogicResult) (str string, ok bool) {
	//condKeys := []string{gen_com.ScaTriggerCond, gen_com.GainRatioCond}
	condKeys := []string{gen_com.GainRatioCond, gen_com.WinRateCond, gen_com.RemoveRateCond}
	var adjust string
	for _, key := range condKeys {
		cond := s.gTem.GetCond(key)
		switch key {
		case gen_com.GainRatioCond: //向上 向下 微调
			adjust = cond.Adjust(genc.GainRatio, s.gTem)
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
