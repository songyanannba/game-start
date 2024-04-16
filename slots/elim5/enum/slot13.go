package enum

const (
	Event0  = iota
	ENullSp //中到null_sp时，移动次数@权重
	EWild   //中到wild时，移动次数@权重
	EWild2  //中到wild_2时，移动次数@权重
	EWild5  //中到wild_5时，移动次数@权重
)

const (
	Move     = iota // 检测中间是否是 wild_reSpin
	MoveUp          // 检测是否赢钱
	MoveDown        // 检测是否赢钱
)

const (
	ColLen     = 3
	ColSpeLen  = 11
	LongTagLen = 9
	MinStep    = 4
)
