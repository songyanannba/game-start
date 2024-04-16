package unit5

import (
	"elim5/enum"
	"elim5/logicPack/base"
	"elim5/logicPack/component"
	"elim5/logicPack/template"
	"elim5/logicPack/template/flow"
	"elim5/pbs/game"
	"elim5/utils/helper"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// NewSpin 创建Spin
func (m *Machine) NewSpin(isFree bool, getFreeNum int, id, parentId int) *component.Spin {
	newSpin := &component.Spin{}
	Option := &component.Options{
		IsTest:       m.BaseSpin.Options.IsTest,
		RatioConfirm: m.BaseSpin.RatioConfirm,
		IsFree:       isFree,
		Raise:        m.BaseSpin.Raise,
		IsMustFree:   m.BaseSpin.IsMustFree,
		IsMustRes:    m.BaseSpin.IsMustRes,
	}
	newSpin.Options = Option
	newSpin.Bet = m.BaseSpin.Bet

	newSpin.FreeSpinParams = component.FreeSpinParams{
		FreeNum: getFreeNum,
	}
	newSpin.SpinInfo = m.SpinInfo.Copy()
	m.SumGain()
	newSpin.Gain = m.SpinInfo.Gain
	newSpin.Id = id
	newSpin.ParentId = parentId
	newSpin.Config = m.BaseSpin.Config
	if m.BaseSpin.IsMustFree && id == 0 {
		newSpin.BuyFreeCoin = decimal.NewFromInt(int64(m.BaseSpin.Bet)).Mul(decimal.NewFromFloat(m.BaseSpin.Config.BuyFee)).IntPart()
		//newSpin.Bet = int(newSpin.BuyFreeCoin)
	}
	return newSpin
}

// GetSpinInfo 获取SpinInfo
func (m *Machine) GetSpinInfo(spinType int, needDebug bool) (spinInfo *template.SpinInfo, err error) {
	spinInfo = &template.SpinInfo{}
	if m.BaseSpin.IsTemGen && m.Tem != nil {
		spinInfo, err = template.NewGameInfo(m.Tem, spinType)
	} else {
		spinInfo, err = template.NewGameInfo(m.BaseSpin, spinType)
	}
	if err != nil {
		return nil, err
	}

	spinInfo.CustomFill[enum.BuyFree] = FillByFreeDisplay
	if needDebug {
		err = spinInfo.SetDebugConfig()
		if err != nil {
			return nil, err
		}
	}
	spinInfo.FillInit()
	return
}

func GetFreeNum(ScNum int) int {
	if ScNum < 3 {
		return 0
	} else if ScNum == 3 {
		return 10
	} else if ScNum == 4 {
		return 11
	} else if ScNum == 5 {
		return 12
	} else if ScNum == 6 {
		return 13
	} else {
		return 14
	}
}

func (m *Machine) ModifyWinLines(winLines []*flow.WinLine) []*flow.WinLine {
	lines := make([]*flow.WinLine, 0)
	gameType := m.SpinInfo.GetGameType()
	switch gameType {
	case enum.SpinAckType2FreeSpin:
		for _, line := range winLines {
			x2Num := 0
			if v, ok := m.BaseSpin.GetEvent(enum.Slot5EventFreeX2); ok == nil {
				x2Num = v.(*base.ChangeTableEvent).Fetch()
			}
			helper.SliceShuffle(line.Tags)
			for i := 0; i < x2Num; i++ {
				if i < len(line.Tags) {
					line.Tags[i].Multiple = 2
					line.Mul = helper.Mul(line.Mul, 2)
				}
			}
			lines = append(lines, line)
			//winLines[i] = line
		}
	default:
		for _, line := range winLines {
			lines = append(lines, line)
		}
	}
	return lo.Filter(lines, func(item *flow.WinLine, index int) bool {
		return item.Mul > 0
	})
}

// FillByFreeDisplay 填充购买免费展示
func FillByFreeDisplay(s *template.SpinInfo) error {
	fillTags := s.GetNormalTags()
	cols := make([]int, 0)
	for i := 0; i < s.Config.GetCol(); i++ {
		cols = append(cols, i)
	}
	helper.SliceShuffle(cols)
	count := 0
	for _, col := range cols {
		rows := lo.Filter(helper.ListToArr(s.Display), func(item *base.Tag, index int) bool {
			return item.Y == col && item.IsEmpty()
		})
		if len(rows) < 1 {
			continue
		}
		row := rows[helper.RandInt(len(rows))].X
		fillTag := s.Config.GetTagByName(enum.ScatterName)
		fillTag.X = row
		fillTag.Y = col
		s.Display[row][col] = fillTag.Copy()
		count++
		if count >= 3 {
			break
		}
	}
	emptyTags := s.GetEmptyTags()
	for _, tag := range emptyTags {
		count = 0
		for {
			count++
			if count > 100 {
				break
			}
			fillTag := fillTags[helper.RandInt(len(fillTags))].Copy()
			fillTag.X = tag.X
			fillTag.Y = tag.Y
			s.Display[tag.X][tag.Y] = fillTag
			if len(s.FindSpecifyLine(fillTag, 5)) > 0 {
				continue
			} else {
				break
			}
		}
	}
	return nil
}

func (m *Machine) GetAck(base *game.BaseSpinAck) protoreflect.ProtoMessage {
	return nil
}

// AckToInfo ack转译
func (m *Machine) AckToInfo(bytes []byte) string {
	return "I haven't written it yet. I'm waiting for it to be perfect."
}
