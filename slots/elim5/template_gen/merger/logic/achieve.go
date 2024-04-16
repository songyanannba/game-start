package logic

import (
	"elim5/enum"
	"elim5/logicPack/base"
	"elim5/model/business"
	"elim5/utils/helper"
	"fmt"
)

func (d *Dismantling) GetTagByName(name string) *base.Tag {
	return d.Config.GetTag(name)
}

func (d *Dismantling) GetTagById(id int) *base.Tag {
	return d.Config.GetTagById(id)
}

func (d *Dismantling) GetRow() int {
	return d.Config.Row
}

func (d *Dismantling) GetCol() int {
	return d.Config.Index
}

func (d *Dismantling) GetTemplate(spinType, which int) (map[int][]uint16, error) {
	return nil, nil
}

func (d *Dismantling) GetIsBuy() int {
	if d.GetGameType() == enum.SpinAckType2FreeSpin {
		return enum.BuyFree
	}
	if d.GetGameType() == enum.SpinAckType1BranchBuyFree {
		return enum.BuyFree
	}
	return enum.NoBuy
}

func (d *Dismantling) GetAllTags() []*base.Tag {
	return d.Config.GetAllTagQuote()
}

func (d *Dismantling) GetPayTables() []*base.PayTable {
	return d.Config.PayTableList
}

func (d *Dismantling) GetDebugConfigs() []*business.DebugConfig {
	return []*business.DebugConfig{}
}

func (d *Dismantling) GetUserId() uint {
	return 0
}

func (d *Dismantling) GetIsTest() bool {
	return true
}

func (d *Dismantling) GetInitTemIndex(spinType, which int) (int, error) {
	return helper.RandInt(len(d.Template[0])), nil
}

func (d *Dismantling) GetEvent(index int) (event any, err error) {
	var GetEvents *base.Event
	GetEvents, err = d.GetEvents()
	if err != nil {
		return nil, err
	}
	return GetEvents.Get(index)
}

func (d *Dismantling) GetIsRaise() bool {
	if d.GetGameType() == enum.SpinAckType1BranchRaise {
		return true
	}
	return false
}

func (d *Dismantling) GetSlotId() int {
	return int(d.Config.SlotId)
}

func (d *Dismantling) GetEvents() (*base.Event, error) {
	event, ok := d.Config.Event[d.GetRtp()]
	if !ok {
		return nil, fmt.Errorf("event not found, ratio: %d", d.GetRtp())
	}
	return event, nil
}

func (d *Dismantling) GetCoords() [][]base.Coordinate {
	return d.Config.Coords
}

func (d *Dismantling) GetWhich() int {
	return d.TemGens[0].Which
}
