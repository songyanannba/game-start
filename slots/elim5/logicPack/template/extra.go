package template

import "elim5/logicPack/base"

type Slot43Data struct {
	RankExtraMuls []int32 //额外倍数
	ChoiceList    []int32 //选择列表
	Choice        int32   //选择 是ChoiceList 的key 不是value
}

type Counter struct {
	LeftCount  int32 //翻倍计数
	RightCount int32
}

type SingleTag struct {
	Tag       *base.Tag
	FlowIndex int
	PosOder   []*base.Tag
	SubId     int
}
