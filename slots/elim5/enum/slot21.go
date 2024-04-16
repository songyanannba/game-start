package enum

import "time"

const (
	Unit21TriggerReSpin = iota + 1 //成功触发特殊模式

)

const RedisNil = "redis: nil" //没有对应的key 返回的错误

const (
	ScatNum3 = 3
	ScatNum4 = 4
	ScatNum5 = 5
)

const (
	LinkType          = iota
	LinkTypeMult      //翻倍
	LinkTypeColl      // 收集
	LinkTypeConverter //转换
	LinkTypeUp
)

const (
	LinkWhich = 2
)

const (
	ComSaveTime  = time.Hour * 24 * 15
	TestSaveTime = time.Hour
)

const (
	Mul          = 1
	Mul5         = 5
	Mul10        = 10
	Mul25        = 25
	Mul50        = 50
	MaxGambleMul = 100 //gamble 最大值
)

var ScatNumMul = map[int]int{
	ScatNum3: Mul,
	ScatNum4: Mul5,
	ScatNum5: Mul25,
}

var MulEvent = map[int]int{
	Mul:   2,
	Mul5:  3,
	Mul10: 4,
	Mul25: 5,
	Mul50: 6,
}

const (
	Save  = iota // all
	Save1 = iota //redis
	Save2 = iota //mysql
)

// ：link金钱标签、multipier翻倍标签、link_collect收集标签、converter转换器标签、scatter标签
const (
	Mystery       = "mystery"      //神秘标签
	ScatLink      = "scatter_link" //金色招财猫 -- 可以变成link玩法
	LinkCoin      = "link_coin"
	LinkConverter = "link_converter"
	LinkCollect   = "link_collect"
	LinkNudgeUp   = "link_nudge_up"
	LinkMult      = "link_mult" //翻倍标签
	Scatter       = "scatter"
	Scatter5      = "scatter_5"
	Scatter25     = "scatter_25"
)

// //转换 > 翻倍 > 收集 > 上移
var LinkName = map[string]string{
	LinkCollect:   LinkCollect,   // 收集
	LinkMult:      LinkMult,      // 翻倍
	LinkNudgeUp:   LinkNudgeUp,   //
	LinkConverter: LinkConverter, //
}

var SpecialTagName = map[string]string{
	LinkCollect:   LinkCollect,
	LinkMult:      LinkMult,
	LinkNudgeUp:   LinkNudgeUp,
	LinkConverter: LinkConverter,
	LinkCoin:      LinkCoin,
	ScatLink:      ScatLink,
	Mystery:       Mystery,
	Scatter:       Scatter,
	Scatter5:      Scatter5,
	Scatter25:     Scatter25,
}

var ScatterTags = []string{Scatter, Scatter5, Scatter25}

var ComTags = []string{"low_1", "low_2", "low_3", "low_4", "high_1", "high_2", "high_3", "high_4"}

var ScatterNumTags = map[int]string{
	1: Scatter,
	2: Scatter5,
	3: Scatter25,
	4: "",
	5: Mystery,
}

var ScatterStr = map[string]int{
	Scatter:   3,
	Scatter5:  4,
	Scatter25: 5,
}

var OptMap = map[int]int{
	1: Mul,
	2: Mul5,
	3: Mul25,
	4: Mul, //todo
	5: 0,
}

const (
	OptIdx = iota
	OptIdx1
	OptIdx2
	OptIdx3
	OptIdx4
	OptIdx5 = 5
)

const (
	GamBleUnderway = iota //gamble 进行中
	GambleFinish          //gamble状态 1 的时候说明gamble 结束 开始游戏
)

var SkillLeave = map[string]int{
	LinkMult:      Leave1,
	LinkCollect:   Leave2,
	LinkConverter: Leave3,
	LinkNudgeUp:   Leave4,
}

const (
	Leave1 = iota + 1
	Leave2
	Leave3
	Leave4
)

var LeaveSkill = map[int]string{
	Leave1: LinkMult,
	Leave2: LinkCollect,
	Leave3: LinkConverter,
	Leave4: LinkNudgeUp,
}

// OptionScatMulti 选项对应的Scatter 标签基本倍率
var OptionScatMulti = map[int]int{
	1: 1,
	2: 5,
	3: 25,
	4: 1, //todo
}
