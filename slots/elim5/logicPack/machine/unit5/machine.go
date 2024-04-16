package unit5

import (
	"elim5/enum"
	"elim5/global"
	"elim5/logicPack/base"
	"elim5/logicPack/component"
	"elim5/logicPack/template"
	"elim5/logicPack/template/flow"
	"elim5/pbs/game"
	"elim5/template_gen/gen_com"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
)

// Machine 划线 + 特殊免费玩
type Machine struct {
	Spin     *component.Spin    `json:"-"`
	Spins    []*component.Spin  `json:"-"`
	SpinInfo *template.SpinInfo `json:"-"`
	BaseSpin *component.Spin    `json:"-"`

	Tem *gen_com.GenTemplate
}

// NewMachine 游戏入口
func NewMachine(s *component.Spin) *Machine {
	return &Machine{Spin: &component.Spin{},
		Spins:    []*component.Spin{},
		BaseSpin: s,
	}
}

// NewTemMachine 模版生成入口
func NewTemMachine(s *component.Spin, tem *gen_com.GenTemplate) *Machine {
	return &Machine{Spin: &component.Spin{},
		Spins:    []*component.Spin{},
		BaseSpin: s,
		Tem:      tem,
	}
}

func (m *Machine) GetInitData() {}

func (m *Machine) GetResData() {}

func (m *Machine) GetSpin() *component.Spin {
	return m.Spin
}
func (m *Machine) EndTheGame(recordId int32) (err error) { return nil }
func (m *Machine) GetSpins() []*component.Spin {
	return m.Spins
}

// SumGain 合计赢得的金额
func (m *Machine) SumGain() {
	s := m.SpinInfo
	m.SpinInfo.Gain = lo.SumBy(s.SpinFlow, func(i flow.SpinFlow) int {
		return int(i.Gain)
	})
}

// Exec Spin运行
func (m *Machine) Exec() (err error) {
	verify := base.Verify{
		Count: 0,
	}
	m.SpinInfo, err = m.GetSpinInfo(enum.SpinAckType1NormalSpin, true)
	if err != nil {
		return err
	}
	err = m.PlayGame(0)
	if err != nil {
		return err
	} //根据模版去消除数据的玩法

	scatters := m.SpinInfo.FindTagsByName(enum.ScatterName)
	scatterLine := m.SpinInfo.GetWinLine(scatters)
	m.SpinInfo.Scatter = scatterLine

	datum := GetFreeNum(len(scatters))
	m.Spin = m.NewSpin(false, datum, 0, 0)

	for i := 0; i < datum; i++ {
		if verify.Count >= enum.SlotMaxSpinNum {
			return fmt.Errorf(enum.SlotMaxFreeSpinErr)
		}
		err := m.FreeSpinExec(&verify, 0, helper.If(i == 0, true, false))
		if err != nil {
			return err
		}
	}
	return nil
}

// FreeSpinExec 免费游戏
func (m *Machine) FreeSpinExec(verify *base.Verify, parentId int, needDebug bool) (err error) {
	verify.Count++
	m.SpinInfo, err = m.GetSpinInfo(enum.SpinAckType2FreeSpin, needDebug)
	if err != nil {
		return err
	}
	err = m.PlayGame(0)
	if err != nil {
		return err
	}

	scatters := m.SpinInfo.FindTagsByName(enum.ScatterName)
	scatterLine := m.SpinInfo.GetWinLine(scatters)
	m.SpinInfo.Scatter = scatterLine

	datum := GetFreeNum(len(scatters))

	m.Spins = append(m.Spins, m.NewSpin(true, datum, verify.Count, parentId))
	parentId = verify.Count

	for i := 0; i < datum; i++ {
		if verify.Count >= enum.SlotMaxSpinNum {
			return fmt.Errorf(enum.SlotMaxFreeSpinErr)
		}
		err := m.FreeSpinExec(verify, parentId, false)
		if err != nil {
			return err
		}
	}
	return nil
}

// PlayGame 游戏流程
func (m *Machine) PlayGame(count int) error {
	SpinFlow := flow.NewSpinFlow(count)
	SpinFlow.AddList = m.SpinInfo.Drop()
	SpinFlow.FlowMap += m.SpinInfo.PrintTable("初始")
	if len(m.SpinInfo.GetEmptyTags()) > 0 {
		global.GVA_LOG.Error("模版缺失还有标签没有填充")
	}
	SpinFlow.InitList = helper.CopyListArr(m.SpinInfo.Display)
	//获取划线
	lines := m.SpinInfo.FindAdjacentLine(enum.Slot5FreeLineNum)
	winLines := m.ModifyWinLines(m.SpinInfo.GetWinLines(lines))
	if len(winLines) == 0 {
		m.SpinInfo.SpinFlow = append(m.SpinInfo.SpinFlow, SpinFlow)
		return nil
	}
	//获取划线赢钱
	SpinFlow.AddOmitList(winLines...)
	SpinFlow.Gain = helper.MulToInt(SpinFlow.SumMul, m.BaseSpin.Bet)
	//删除标签
	m.SpinInfo.DeleteTagList(helper.ArrConversion(winLines, func(t *flow.WinLine) ([]*base.Tag, bool) {
		return lo.Filter(t.Tags, func(item *base.Tag, index int) bool {
			if item.Name == enum.ScatterName || item.IsWild {
				return false
			}
			return true
		}), true
	}))
	SpinFlow.FlowMap += m.SpinInfo.PrintTable("删除")
	//掉落标签

	m.SpinInfo.SpinFlow = append(m.SpinInfo.SpinFlow, SpinFlow)
	if count > enum.SlotMaxSpinNum {
		return fmt.Errorf(enum.SlotMaxSpinStr)
	}
	return m.PlayGame(count + 1)
}

// GetOptionAck 获取OptionAck
func (m *Machine) GetOptionAck(req *game.OptionsReq, recordId int32, ack *game.OptionsAck) (err error) {
	//TODO 个别机台,需要用户操作,用户操作后根据用户操作,返回特定ack
	return enum.ErrUnrealized
}
