package unit

import (
	"elim5/enum"
	"elim5/logicPack"
	"elim5/logicPack/base"
	"elim5/logicPack/component"
	"elim5/logicPack/template"
	"elim5/service/template_gen/gen_com"
	"elim5/service/test/handler/repeat/unit43"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
)

type Slot43 struct {
	gTem     *gen_com.GenTemplate
	testSlot *unit43.Unit
}

func GetUnit43(t *gen_com.GenTemplate) *Slot43 {
	return &Slot43{
		testSlot: unit43.NewUnit(),
		gTem:     t,
	}
}

func (s *Slot43) RunTem() ([]*component.Spin, error) {
	sp := &component.Spin{
		Options: &component.Options{
			IsFree:       helper.If(s.gTem.TemGen.Type == enum.SpinAckType2FreeSpin, true, false),
			IsTemGen:     true,
			RatioConfirm: s.gTem.RatioConfirm,
			Raise:        helper.If(s.gTem.TemGen.Type == enum.SpinAckType1BranchRaise, helper.MulToInt(s.gTem.Config.Raise, 100), 0),
			IsTest:       true,
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
func (s *Slot43) Calculate(spins []*component.Spin) {
	s.testSlot.Calculate(spins)
}

// GetLogicResult slot33 计算数据
func (s *Slot43) GetLogicResult() *gen_com.LogicResult {
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

	default:
		return &gen_com.LogicResult{}
	}

}

// GetStatus slot15 获取状态
func (s *Slot43) GetStatus() (totalInterval float64, str string, ok bool) {
	r := s.GetLogicResult()
	totalInterval = s.gTem.GetDisparity(r)
	str, ok = s.Adjust(r)
	return
}

// Adjust slot33 调整数据
func (s *Slot43) Adjust(genc *gen_com.LogicResult) (str string, ok bool) {
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

func GetNoLine(s *template.SpinInfo, specifyTags map[string]int) (err error) {
	fillTags := s.GetNormalTags()
	tagsMap := s.GetTagMapByName()

	scaCount := helper.RandInt(helper.Min(specifyTags[enum.ScatterName], helper.If(len(tagsMap[enum.ScatterName]) > 0, 0, 3)))
	specifyTags[enum.ScatterName] -= scaCount

	allTag := s.FindAllTagsQuote()
	col0EmpTags := lo.Filter(allTag, func(item *base.Tag, index int) bool {
		return item.Y == 0
	})

	nowClo2Tags := lo.Filter(allTag, func(item *base.Tag, index int) bool {
		return item.Y == 2 || (item.Y == 6 && item.X == 3)
	})

	for _, tag := range col0EmpTags {
		if !tag.IsEmpty() {
			continue
		}
		var fillTag *base.Tag
		for i := 0; i < 1000; i++ {
			if i == 999 {
				return fmt.Errorf("填充失败")
			}
			if scaCount > 0 {
				fillTag = s.Config.GetTagByName(enum.ScatterName)
				scaCount--
			} else {
				fillTag = fillTags[helper.RandInt(len(fillTags))].Copy()
			}
			fillTag.X = tag.X
			fillTag.Y = tag.Y
			s.Display[tag.X][tag.Y] = fillTag
			if len(lo.Filter(nowClo2Tags, func(item *base.Tag, index int) bool {
				return item.Name == fillTag.Name
			})) > 0 {
				continue
			} else {
				break
			}
		}
	}

	allTag = s.FindAllTagsQuote()
	col0EmpTags = lo.Filter(allTag, func(item *base.Tag, index int) bool {
		return item.Y == 0
	})

	nowClo2Tags = lo.Filter(allTag, func(item *base.Tag, index int) bool {
		return item.Y == 2 || (item.Y == 6 && item.X == 3)
	})

	for _, tag := range nowClo2Tags {
		if !tag.IsEmpty() {
			continue
		}
		var fillTag *base.Tag
		for i := 0; i < 1000; i++ {
			if i == 999 {
				return fmt.Errorf("填充失败")
			}
			if scaCount > 0 {
				fillTag = s.Config.GetTagByName(enum.ScatterName)
				scaCount--
			} else {
				fillTag = fillTags[helper.RandInt(len(fillTags))].Copy()
			}
			fillTag.X = tag.X
			fillTag.Y = tag.Y
			s.Display[tag.X][tag.Y] = fillTag
			if len(lo.Filter(col0EmpTags, func(item *base.Tag, index int) bool {
				return item.Name == fillTag.Name
			})) > 0 {
				continue
			} else {
				break
			}
		}
	}

	emptyTags := s.GetEmptyTags()
	for _, tag := range emptyTags {
		var fillTag *base.Tag
		if scaCount > 0 {
			fillTag = s.Config.GetTagByName(enum.ScatterName)
			scaCount--
		} else {
			fillTag = fillTags[helper.RandInt(len(fillTags))].Copy()
		}
		fillTag.X = tag.X
		fillTag.Y = tag.Y
		s.Display[tag.X][tag.Y] = fillTag
	}

	return nil
}
