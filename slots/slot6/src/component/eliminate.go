package component

import (
	"github.com/samber/lo"
	"slot6/model/business"
	"slot6/src/base"
	"slot6/src/eliminate"
	"slot6/utils/helper"
)

func NewGraph(spin *Spin, isFree bool) *eliminate.Table {
	config := spin.Config
	t := &eliminate.Table{}
	t.Target = NewTarget(spin, isFree)
	t.Row = config.Row
	t.Col = config.Index
	t.PayTableList = config.PayTableList
	t.AlterFlows = make([]*base.AlterFlow, 0)
	for _, tag := range config.GetAllTag() {
		// 取出scatter
		if tag.Name == "scatter" {
			t.Scatter = tag.Copy()
			continue
		}
		if tag.IsWild {
			t.WildTags = append(t.WildTags, tag.Copy())
			continue
		}
		t.Tags = append(t.Tags, tag)
	}

	t.TableReset()
	t.FillTest(spin.IsSetResult)
	debugStr := GetDebugData(spin, isFree)
	if debugStr != "" {
		t.FillDebug(debugStr)
	} else {
		t.InitFill()
	}
	//fmt.Print(t.PrintTable("测试"))
	t.FillAll()
	t.InitTable = t.GetGraph()
	t.PayTableListMap()
	return t
}

func NewTarget(spin *Spin, IsFree bool) *eliminate.Target {
	comp := spin.Config
	var (
		minMul      float64
		maxMul      float64
		initNum     int
		scatterNum  int
		randMul     int
		mulNumEvent *base.ChangeTableEvent
	)

	if !IsFree {
		if spin.IsMustFree {
			minMul = 0
			maxMul = 0
			initNum = 0
			scatterNum = 3
			mulNumEvent = comp.Event.M[19].(*base.ChangeTableEvent)
		} else {
			ints, key := comp.Event.M[0].(*base.IntervalRatioEvent).Fetch()
			maxMul = float64(ints[1])
			randMul = helper.RandInt(ints[1]-ints[0]) + ints[0]
			minMul = float64(randMul)
			initNum = comp.Event.M[key+1].(*base.ChangeTableEvent).Fetch()
			switch spin.Config.SlotId {
			case 5:
				scatterNum = comp.Event.M[16].(*base.ChangeTableEvent).Fetch()
				mulNumEvent = comp.Event.M[19].(*base.ChangeTableEvent)
			case 6:
				scatterNum = 0
				mulNumEvent = &base.ChangeTableEvent{}
			}
		}
	} else {
		ints, key := comp.Event.M[18].(*base.IntervalRatioEvent).Fetch()
		maxMul = float64(ints[1])
		randMul = helper.RandInt(ints[1]-ints[0]) + ints[0]
		minMul = float64(randMul)
		initNum = comp.Event.M[key+1].(*base.ChangeTableEvent).Fetch()
		scatterNum = comp.Event.M[20].(*base.ChangeTableEvent).Fetch()
		mulNumEvent = comp.Event.M[19].(*base.ChangeTableEvent)
	}

	//minMul = 100
	//maxMul = 200
	//initNum = 10
	//scatterNum = 0
	//mulNumEvent = comp.Event.M[19].(*base.ChangeTableEvent)
	return &eliminate.Target{
		MinMul:      minMul,
		MaxMul:      maxMul,
		InitNum:     initNum,
		ScatterNum:  scatterNum,
		MulNumEvent: mulNumEvent,
	}

}

func GetDebugData(spin *Spin, isFree bool) string {
	debugType := uint8(1)
	playType := uint8(1)
	if spin.IsTest {
		debugType = 2
	}
	if isFree {
		playType = 3
	}
	debugs := lo.Filter(spin.Config.Debugs, func(item *business.DebugConfig, index int) bool {
		return item.DebugType == debugType && item.PalyType == playType
	})
	if len(debugs) == 0 {
		return ""
	}
	return debugs[0].ResultData
}
