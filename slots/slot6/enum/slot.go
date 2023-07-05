package enum

const (
	SlotMaxSpinNum = 10000 // slot最大转动
)

// Slot 第六台等级
const (
	Rand1 = iota + 1
	Rand2
	Rand3
	Rand4
	Rand5
)

const (
	SlotWild = "wild"
)

const (
	CoreTagMinNum = 4 //中心点周围 最少填充的tag个数
	CoreTagMaxNum = 8 //中心点周围 最多填充的tag个数

	GetLine = 5 //默认匹配连续标签的最少个数
)
