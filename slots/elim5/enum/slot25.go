package enum

var HorizontalColumn = map[int]int{
	1: 4,
	2: 3,
	3: 2,
	4: 1,
}

const (
	Unit25IsMergeWeight = iota //特殊玩法触发权重
	Unit25IsRankWeight         //选择标签权重
)
