package eliminate

import (
	"elim5/logicPack/base"
)

type SpecialTag struct {
	Scatter base.Tag // scatter
	//Wild    component.Tag // wild
	//Mul     component.Tag // 翻倍
}

func SetScatterName(name string) func(s *SpecialTag) {
	return func(s *SpecialTag) {
		s.Scatter.Name = name
	}
}

type SingleTag struct {
	Tag       *base.Tag
	FlowIndex int
	PosOder   []*base.Tag
	SubId     int
}
