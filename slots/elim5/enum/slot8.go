package enum

const (
	TriggerFreeGet = 4
	TriggerFreeAdd = 3
)

const (
	FreeGetNum = 15
	FreeAddNum = 5
)

const (
	Slot8LineLength = 8
)

const (
	LargeScale = 0.1
	Trim       = 0.05
	Ok         = 0.02
)

const (
	Unit8NormalMulWeight = iota //普通转权重
	Unit8RaiseMulWeight         //加注转权重
	Unit8FreeMulWeight          //免费转权重

	Unit17MegastackWeight    //特殊标签变化权重
	Unit17MegastackIsMul     //特殊标签变成翻倍权重       //特殊标签变化
	Unit17MegastackIsMulFree //特殊标签变成wild权重
)
