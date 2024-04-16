package component

import (
	"elim5/enum"
	"elim5/logicPack/base"
	"elim5/model/business"
	"elim5/utils/helper"
	"fmt"
)

func (s *Spin) GetTagByName(name string) *base.Tag {
	return s.Config.GetTag(name)
}

func (s *Spin) GetTagById(id int) *base.Tag {
	return s.Config.GetTagById(id)
}

func (s *Spin) GetRow() int {
	return s.Config.Row
}

func (s *Spin) GetCol() int {
	return s.Config.Index
}

func (s *Spin) GetTemplate(typ, which int) (tem map[int][]uint16, err error) {
	key := fmt.Sprintf("%d-%d-%d", s.RatioConfirm, typ, which)
	var ok bool
	if tem, ok = s.Config.Templates[key]; ok {
		return tem, nil
	}
	return nil, fmt.Errorf("template not found typ:%d rtp:%d which:%d", typ, s.RatioConfirm, which)
}

func (s *Spin) GetIsBuy() int {
	if s.IsMustFree {
		return enum.BuyFree
	} else if s.IsMustRes {
		return enum.BuyRe
	} else {
		return enum.NoBuy
	}
}
func (s *Spin) GetAllTags() []*base.Tag {
	return s.Config.GetAllTagQuote()
}
func (s *Spin) GetPayTables() []*base.PayTable {
	return s.Config.PayTableList
}

func (s *Spin) GetDebugConfigs() []*business.DebugConfig {
	return s.Config.Debugs
}

func (s *Spin) GetUserId() uint {
	return s.Options.UserId
}

func (s *Spin) GetIsTest() bool {
	return s.Options.IsTest
}

func (s *Spin) GetInitTemIndex(typ, which int) (int, error) {
	var (
		ok     bool
		colTem []uint16
		tem    map[int][]uint16
		err    error
	)
	tem, err = s.GetTemplate(typ, which)
	if err != nil {
		return 0, err
	}
	colTem, ok = tem[0]
	if !ok || len(colTem) == 0 {
		return 0, fmt.Errorf("template not found col:%d", 0)
	}

	//randIndex := helper.RandInt(len(tem[0]))
	//if spinType == 2 {
	//	i := len(tem[0])
	//
	//	if randIndex <= int(i/3)+5 { //最下面2/3
	//		randIndex = 5700 + 5 + helper.RandInt(len(tem[0])-5700-5)
	//		return randIndex, nil
	//	}
	//
	//	//if randIndex > int(i/3) { //最上面 1/3
	//	//	randIndex = helper.RandInt(int(i / 3))
	//	//}
	//
	//	//
	//	//if randIndex > int(i/3)*2 { //最上面 2/3
	//	//	randIndex = helper.RandInt(int(i/3) * 2)
	//	//}
	//
	//	//if randIndex > int(i/2) { //最上面 1/2
	//	//	randIndex = helper.RandInt(int(i / 2))
	//	//}
	//
	//}
	////
	//return randIndex, nil
	return helper.RandInt(len(tem[0])), nil
}

func (s *Spin) GetIsRaise() bool {
	return s.Options.Raise > 0
}

func (s *Spin) GetSlotId() int {
	return int(s.Config.SlotId)
}

func (s *Spin) GetReels() ([]*Reel, error) {
	if reels, ok := s.Config.GetReelMap()[s.RatioConfirm]; ok {
		return reels, nil
	}
	return nil, fmt.Errorf("not found reel ratio:%d", s.RatioConfirm)
}

func (s *Spin) GetReel(col int) (*Reel, error) {
	var (
		reels []*Reel
		err   error
	)
	reels, err = s.GetReels()
	if err != nil {
		return nil, err
	}
	if col >= len(reels) {
		return nil, fmt.Errorf("not found reel which:%d", col)
	}
	return reels[col], nil

}

func (s *Spin) GetReelData(typ, col, which int) (*ReelData, error) {
	var (
		reel *Reel
		err  error
	)
	reel, err = s.GetReel(col)
	if err != nil {
		return nil, err
	}
	return reel.GetReelData(typ, which)
}

func (s *Spin) GetEvents() (*base.Event, error) {
	if events, ok := s.Config.Event[s.RatioConfirm]; ok {
		return events, nil
	}
	return nil, fmt.Errorf("not found event ratio:%d", s.RatioConfirm)
}

func (s *Spin) GetEvent(index int) (event any, err error) {
	var (
		events *base.Event
	)
	events, err = s.GetEvents()
	if err != nil {
		return nil, err
	}
	return events.Get(index)
}

func (s *Spin) GetCoords() [][]base.Coordinate {
	return s.Config.Coords
}

func (s *Spin) GetWhich() int {
	return s.Which
}
