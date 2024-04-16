package logicPack

import (
	. "elim5/enum"
	"elim5/logicPack/component"
	"elim5/logicPack/machine/unit5"
	"elim5/pbs/common"
	"elim5/pbs/game"
	"github.com/aws/aws-sdk-go/aws/session"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Machine interface {
	GetSpin() *component.Spin
	Exec() error
	GetSpins() []*component.Spin
	GetAck(ack *game.BaseSpinAck) protoreflect.ProtoMessage
	AckToInfo(bytes []byte) string
	GetOptionAck(req *game.OptionsReq, recordId int32, ack *game.OptionsAck) (err error)
	EndTheGame(recordId int32) (err error)
}

type DemoMachine interface {
	GetDomeOptionAck(req *game.OptionsReq, s *session.Session, ack *game.OptionsAck) (err error)
}

func Play(slotId uint, amount int, options ...component.Option) (m Machine, err error) {
	var s *component.Spin
	s, err = component.NewSpin(slotId, amount, options...)
	if err != nil {
		return nil, err
	}
	return RunSpin(s)
}

func RunSpin(s *component.Spin) (m Machine, err error) {
	m, err = GetMachine(s)
	if err != nil {
		return
	}
	err = m.Exec()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetSlotMachine(slotId uint, options ...component.Option) (m Machine, err error) {
	var s *component.Spin
	s, err = component.NewSpin(slotId, 0, options...)
	if err != nil {
		return nil, err
	}
	m, err = GetMachine(s)
	return
}

// GetMachine 游戏入口
func GetMachine(s *component.Spin) (m Machine, err error) {
	if s.RatioConfirm == 0 {
		s.RatioConfirm = 1
	}

	//slotId := s.Config.SlotId
	m = unit5.NewMachine(s)
	return m, err

	// (如果开启AB测 且 编号不为0) 或 (机台id大于等于100000)
	//if s.Config.IsOpenABTest == OpenABTest && s.IndicateNum != 0 || s.Config.SlotId > AbTestMinSlotId {
	//	slotId, err = business.DecodeSlotId(s.Config.SlotId)
	//	if err != nil {
	//		return nil, ErrNoServer
	//	}
	//}

	//switch slotId {
	//case SlotId1:
	//	m = unit1.NewMachine(s)
	//case SlotId2:
	//	m = unit2.NewMachine(s)
	//case SlotId3:
	//	m = unit3.NewMachine(s)
	//case SlotId4:
	//	m = unit4.NewMachine(s)
	//case SlotId5:
	//	m = unit5.NewMachine(s)
	//case SlotId6:
	//	m = unit6.NewMachine(s)
	//case SlotId7:
	//	m = unit7.NewMachine(s)
	//case SlotId8, SlotId15:
	//	m = unit8.NewMachine(s)
	//case SlotId9:
	//	m = unit9.NewMachine(s)
	//case SlotId10:
	//	m = unit10.NewMachine(s)
	//case SlotId14:
	//	m = unit14.NewMachine(s)
	//case SlotId12:
	//	m = unit12.NewMachine(s)
	//case SlotId13:
	//	m = unit13.NewMachine(s)
	//case SlotId16:
	//	m = unit16.NewMachine(s)
	//case SlotId17:
	//	m = unit17.NewMachine(s)
	//case SlotId18:
	//	m, err = unit18.NewMachine(s)
	//case SlotId19:
	//	m, err = unit19.NewMachine(s)
	//case SlotId21:
	//	m = unit21.NewMachine(s)
	//case SlotId22:
	//	m = unit22.NewMachine(s)
	//case SlotId23:
	//	m, err = unit23.NewMachine(s)
	//case SlotId24:
	//	m = unit24.NewMachine(s)
	//case SlotId25:
	//	m = unit25.NewMachine(s)
	//case SlotId26:
	//	m = unit26.NewMachine(s)
	//case SlotId31:
	//	m = unit31.NewMachine(s)
	//case SlotId30:
	//	m = unit30.NewMachine(s)
	//case SlotId28:
	//	m = unit28.NewMachine(s)
	//case SlotId29:
	//	m = unit29.NewMachine(s)
	//case SlotId27:
	//	m = unit27.NewMachine(s)
	//case SlotId32:
	//	m = unit32.NewMachine(s)
	//case SlotId33:
	//	m = unit33.NewMachine(s)
	//case SlotId43:
	//	m = unit43.NewMachine(s)
	//case SlotId46:
	//	m, err = unit46.NewMachine(s)
	//case SlotId44:
	//	m = unit44.NewMachine(s)
	//case SlotId49:
	//	m = unit49.NewMachine(s)
	//case SlotId50:
	//	m = unit50.NewMachine(s)
	//case SlotId51:
	//	m, err = unit51.NewMachine(s)
	//case SlotId52:
	//	m = unit52.NewMachine(s)
	//case SlotId57:
	//	m, err = unit57.NewMachine(s)
	//case SlotId47:
	//	m = unit47.NewMachine(s)
	//case SlotId48:
	//	m = unit48.NewMachine(s)
	//case SlotId53:
	//	m = unit53.NewMachine(s)
	//case SlotId58:
	//	m, err = unit58.NewMachine(s)
	//case SlotId54:
	//	m = unit54.NewMachine(s)
	//
	//default:
	//	return nil, errors.New("slotId not found")
	//}
	//return m, err
}

// GetMachineTem 跑模版
//func GetMachineTem(s *component.Spin, tem *gen_com.GenTemplate) (m Machine, err error) {
//	switch s.Config.SlotId {
//	case SlotId8, SlotId15:
//		m = unit8.NewTemMachine(s, tem)
//	case SlotId17:
//		m = unit17.NewTemMachine(s, tem)
//	case SlotId6:
//		m = unit6.NewTemMachine(s, tem)
//	case SlotId9:
//		m = unit9.NewTemMachine(s, tem)
//	case SlotId7:
//		m = unit7.NewTemMachine(s, tem)
//	case SlotId5:
//		m = unit5.NewTemMachine(s, tem)
//	case SlotId14:
//		m = unit14.NewTemMachine(s, tem)
//	case SlotId16:
//		m = unit16.NewTemMachine(s, tem)
//	case SlotId22:
//		m = unit22.NewTemMachine(s, tem)
//	case SlotId25:
//		m = unit25.NewTemMachine(s, tem)
//	case SlotId33:
//		m = unit33.NewTemMachine(s, tem)
//	case SlotId43:
//		m = unit43.NewTemMachine(s, tem)
//	case SlotId44:
//		m = unit44.NewTemMachine(s, tem)
//	case SlotId49:
//		m = unit49.NewTemMachine(s, tem)
//	case SlotId50:
//		m = unit50.NewTemMachine(s, tem)
//	case SlotId52:
//		m = unit52.NewTemMachine(s, tem)
//	case SlotId47:
//		m = unit47.NewTemMachine(s, tem)
//	case SlotId48:
//		m = unit48.NewTemMachine(s, tem)
//	case SlotId53:
//		m = unit53.NewTemMachine(s, tem)
//	case SlotId54:
//		m = unit54.NewTemMachine(s, tem)
//
//	default:
//		return nil, errors.New("slotId not found")
//	}
//	return m, nil
//}
//
//func GetMachineNoLine(s *template.SpinInfo, specifyTags map[string]int, fillWeight *base.ChangeTableStrEvent) (err error) {
//	switch s.Config.GetSlotId() {
//	case SlotId8, SlotId15:
//		err = unit8.GetNoLine(s, specifyTags)
//	case SlotId17:
//		err = unit17.GetNoLine(s, specifyTags, fillWeight)
//	case SlotId9:
//		err = unit9.GetNoLine(s, specifyTags)
//	case SlotId25:
//		err = unit25.GetNoLine(s, specifyTags)
//	case SlotId16:
//		err = unit16.GetNoLine(s, specifyTags)
//	case SlotId33:
//		err = unit33.GetNoLine(s, specifyTags)
//	case SlotId43:
//		err = unit43.GetNoLine(s, specifyTags)
//	case SlotId44:
//		err = unit44.GetNoLine(s, specifyTags)
//	case SlotId52:
//		err = unit52.GetNoLine(s, specifyTags)
//	case SlotId47:
//		err = unit47.GetNoLine(s, specifyTags)
//	case SlotId48, SlotId49:
//		s.FillNoLine(4)
//	case SlotId53:
//		err = unit53.GetNoLine(s, specifyTags)
//	case SlotId54:
//		err = unit54.GetNoLine(s, specifyTags)
//	default:
//		if GetMatchSlotInfo(s.Config.GetSlotId()) == MatchSlotType2Count {
//			s.FillCountNoLine()
//		} else {
//			s.FillNoLine(GetLine)
//		}
//	}
//	return err
//}

func GetAckBySlotId(id int32) any {
	switch id {
	case SlotId1, SlotId2, SlotId3, SlotId4, SlotId10, SlotId12, SlotId13:
		return &common.SpinAck{Head: &common.AckHead{}}
	case SlotId5, SlotId6, SlotId7, SlotId8, SlotId9, SlotId14, SlotId16, SlotId17, SlotId15:
		return &common.MatchSpinAck{Head: &common.AckHead{}}
	//case SlotId19:
	//	return &game19.SpinAck{Head: &common.AckHead{}}
	//case SlotId18:
	//	return &game18.SpinAck{Head: &common.AckHead{}}
	//case SlotId21:
	//	return &game21.SpinAck{Head: &common.AckHead{}}
	//case SlotId22:
	//	return &game22.SpinAck{Head: &common.AckHead{}}
	//case SlotId23:
	//	return &game23.SpinAck{Head: &common.AckHead{}}
	//case SlotId24:
	//	return &game24.SpinAck{Head: &common.AckHead{}}
	//case SlotId25:
	//	return &game25.SpinAck{Head: &common.AckHead{}}
	//case SlotId11, SlotId20:
	//	// 还不知道的类型
	//	return nil
	//case SlotId29:
	//	return &game29.SpinAck{Head: &common.AckHead{}}
	//case SlotId27:
	//	return &game27.SpinAck{Head: &common.AckHead{}}
	//case SlotId32:
	//	return &game32.SpinAck{Head: &common.AckHead{}}
	//case SlotId33:
	//	return &game33.SpinAck{Head: &common.AckHead{}}
	//case SlotId43:
	//	return &game43.SpinAck{Head: &common.AckHead{}}
	//case SlotId46:
	//	return &game46.SpinAck{Head: &common.AckHead{}}
	//case SlotId44:
	//	return &game44.SpinAck{Head: &common.AckHead{}}
	//case SlotId52:
	//	return &game52.SpinAck{Head: &common.AckHead{}}
	//case SlotId49:
	//	return &game49.SpinAck{Head: &common.AckHead{}}
	//case SlotId50:
	//	return &game50.SpinAck{Head: &common.AckHead{}}
	//case SlotId51:
	//	return &game51.SpinAck{Head: &common.AckHead{}}
	//case SlotId57:
	//	return &game57.SpinAck{Head: &common.AckHead{}}
	//case SlotId47:
	//	return &game47.SpinAck{Head: &common.AckHead{}}
	//case SlotId48:
	//	return &game48.SpinAck{Head: &common.AckHead{}}
	//case SlotId53:
	//	return &game53.SpinAck{Head: &common.AckHead{}}
	//case SlotId54:
	//	return &game54.SpinAck{Head: &common.AckHead{}}
	//case SlotId58:
	//	return &game58.SpinAck{Head: &common.AckHead{}}
	default:
		return nil
	}
}
