package enum

const (
	TriggerLinkFreeSpinSucc = iota + 1 // 1 触发 linkFreeSpin 成功
	TriggerLinkFreeSpinFail            // 2 触发 linkFreeSpin 失败
	TriggerSpin                        // 3 普通转
)

const (
	IsFreeSpin   = true  //
	IsNoFreeSpin = false //
)

const (
	TriggerEventT  = iota //link 标签倍率
	TriggerEventT1        // 那种模式
	TriggerEventT2        // 那种模式
)

const (
	MinSelectLinkNum = 5 //link 至少多少个link数量才能被收集
	MaxLinkFreeSpin  = 8 // 触发link 成功最大次数
)
