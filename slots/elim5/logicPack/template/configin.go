package template

import (
	"elim5/logicPack/base"
	"elim5/model/business"
)

type TemConfig interface {
	GetTagByName(name string) *base.Tag
	GetTagById(id int) *base.Tag
	GetRow() int
	GetCol() int
	GetTemplate(typ, which int) (map[int][]uint16, error)
	GetIsBuy() int
	GetAllTags() []*base.Tag
	GetPayTables() []*base.PayTable
	GetDebugConfigs() []*business.DebugConfig
	GetUserId() uint
	GetIsTest() bool
	GetInitTemIndex(typ, which int) (int, error)
	GetEvent(index int) (event any, err error)
	GetIsRaise() bool
	GetSlotId() int
	GetCoords() [][]base.Coordinate
	GetWhich() int
}
