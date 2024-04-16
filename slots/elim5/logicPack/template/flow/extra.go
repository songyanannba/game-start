package flow

import "elim5/logicPack/base"

type SFlowExt struct {
	SkillCenter []*base.Tag
	IsUserSkill int
	RankId      int
	RemoveCount int
	//slot49:
	//1 土龙模式  (所有的low标签将被移除)
	//2 水龙模式 （第二列和第四列将变成随机且一列相同的high标签）
	//3 火龙模式 （一个百搭符号将分别添加到卷轴1、2、4、5列的随机位置上）
	//4 巨龙模式 （所有low标签将随机变为high标签，并且机台四个角落都将添加一个wild标签，并保留到本次旋转结束为止）
	//slot51:
	//1 土龙模式  (将卷轴上所有的低价值标签清除)
	//2 水龙模式 （将卷轴上四个中心标签变成wild）
	//3 火龙模式 （随机选择一个标签（wild除外）按图片上去变换）
	//4 龙王技能 （将卷轴上所有的低价值标签变成高价值标签或wild标签）
	Pattern    int         // slot49 and slot51 在那种模式下（1/2/3/4）
	ChangeTags []*base.Tag // slot49 and slot51 （被改变的标签）
}

type MateSlot21 struct {
	InitLists []*base.Tag // 初始列表
	Gamble    int
}

type MateSlot50 struct {
	MattsTag [][]*base.Tag //slot50 田字格标签
	AddWild  []*base.Tag   //添加的wild标签
	Num      int
}

type MatchTable struct {
	LeftResList     [][]*base.Tag    // 结果数据
	RightResList    [][]*base.Tag    // 结果数据
	PayTables       []*base.PayTable // 最终payTable
	TurnIntoWoldTag []*base.Tag      //匹配到的划线 中间将变成wild的点
	RightOmitList   []*WinLine       // 消除列表
	LeftSumMul      float64
	RightSumMul     float64
	ComputeFormula  string //计算公式
}

func NewMatchTable() MatchTable {
	return MatchTable{
		LeftResList: make([][]*base.Tag, 0),
		PayTables:   make([]*base.PayTable, 0),
	}
}
