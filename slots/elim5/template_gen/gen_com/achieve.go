package gen_com

import (
	"elim5/enum"
	"elim5/logicPack/base"
	"elim5/model/business"
	"elim5/utils/helper"
	"fmt"
)

func (t *GenTemplate) GetTagByName(name string) *base.Tag {
	return t.Config.GetTag(name)
}

func (t *GenTemplate) GetTagById(id int) *base.Tag {
	return t.Config.GetTagById(id)
}

func (t *GenTemplate) GetRow() int {
	return t.Config.Row
}

func (t *GenTemplate) GetCol() int {
	return t.Config.Index
}

func (t *GenTemplate) GetTemplate(spinType, which int) (map[int][]uint16, error) {
	if int(t.TemGen.Type) == spinType {
		return t.Template, nil
	} else {
		template, err := t.Config.GetTemplate(spinType, t.RatioConfirm, which)
		if err != nil {
			return t.Template, nil
		}
		return template, nil
	}
}

func (t *GenTemplate) GetIsBuy() int {
	if t.GetGameType() == enum.SpinAckType2FreeSpin {
		return enum.BuyFree
	}
	if t.GetGameType() == enum.SpinAckType1BranchBuyFree {
		return enum.BuyFree
	}
	return enum.NoBuy
}
func (t *GenTemplate) GetAllTags() []*base.Tag {
	return t.Config.GetAllTagQuote()
}
func (t *GenTemplate) GetPayTables() []*base.PayTable {
	return t.Config.PayTableList
}

func (t *GenTemplate) GetDebugConfigs() []*business.DebugConfig {
	return []*business.DebugConfig{}
}
func (t *GenTemplate) GetUserId() uint {
	return 0
}

func (t *GenTemplate) GetIsTest() bool {
	return true
}

func (t *GenTemplate) GetInitTemIndex(spinType, which int) (int, error) {
	return helper.RandInt(len(t.Template[0])), nil
}

func (t *GenTemplate) GetEvent(index int) (event any, err error) {
	var GetEvents *base.Event
	GetEvents, err = t.GetEvents()
	if err != nil {
		return nil, err
	}
	return GetEvents.Get(index)
}

func (t *GenTemplate) GetIsRaise() bool {
	return t.Raise
}

func (t *GenTemplate) GetSlotId() int {
	return int(t.Config.SlotId)
}

func (t *GenTemplate) GetEvents() (*base.Event, error) {
	event, ok := t.Config.Event[t.RatioConfirm]
	if !ok {
		return nil, fmt.Errorf("event not found, ratio: %d", t.RatioConfirm)
	}
	return event, nil
}

func (t *GenTemplate) GetCoords() [][]base.Coordinate {
	return t.Config.Coords
}

func (t *GenTemplate) GetWhich() int {
	return t.TemGen.Which
}
