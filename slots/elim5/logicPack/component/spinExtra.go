package component

import (
	"elim5/logicPack/base"
	"elim5/logicPack/template/flow"
)

type MateDataSlot21 struct {
	FinalGamble        int //
	Gamble             int //todo
	ScatMult           int //scat 数量 对应的倍数
	TriggerFreeSpin    int
	MysteryNum         int // 神秘标签的数量
	NudgeNum           int
	TriggerRe          int //是否触发 re Spin
	TriggerLink        bool
	LinkFlowMultSum    float64     //link标签的全部赢钱
	MysteryColSite     map[int]int //神秘标签的列坐标
	MysteryColDisplace map[int][]*base.Tag
	ScatNumNul         map[string]int
	TriggerLinkEvent   map[string]int
	BehindTag          *base.Tag     //神秘标签将要变成的真实标签
	DataList           [][]*base.Tag //神秘标签转换后的窗口值 进行计算
	ScatLinks          []*base.Tag   //（由神秘标签变成金色招财猫 ，再有金色招财猫变成link标签的集合）link标签的集合
	Display            []*base.Tag
	LinkFlow           []*flow.SpinFlow
}

func NewMateDataSlot21() MateDataSlot21 {
	mateDataSlot21 := MateDataSlot21{
		Gamble:             0,
		ScatMult:           0,
		TriggerRe:          0,
		TriggerLink:        false,
		MysteryNum:         0,
		BehindTag:          nil,
		DataList:           nil,
		ScatLinks:          nil,
		LinkFlowMultSum:    0,
		NudgeNum:           0,
		TriggerLinkEvent:   make(map[string]int),
		MysteryColSite:     make(map[int]int),
		ScatNumNul:         make(map[string]int),
		MysteryColDisplace: make(map[int][]*base.Tag),
	}
	return mateDataSlot21
}

type MoveDataList struct {
	*base.Tag
	Start    int
	Col      int
	Dir      int         //那个方向  up=1 down=2
	Step     int         //移动几步
	Which    int         //那个模版
	InitSite int         //转轮开始位置
	AfterTag *base.Tag   //移动后的标签
	Tags     []*base.Tag //转轮标签数据
	IniTags  []*base.Tag
}

func NewMoveDataList(which, InitSite, startFirst int) *MoveDataList {
	m := &MoveDataList{
		Tag:      nil,
		Dir:      0,
		Step:     0,
		Which:    0,
		InitSite: 0,
		AfterTag: nil,
		Tags:     make([]*base.Tag, 0),
		IniTags:  make([]*base.Tag, 0),
	}
	m.InitSite = InitSite
	m.Which = which
	m.Start = startFirst
	return m
}

type MateDataSlot26 struct {
	IsMod    bool //是否修改
	DataList [][]*base.Tag
	ModList  []*MoveDataList
}

type MateDataSlot23 struct {
	LinkTags []*base.Tag
	AllMult  int
}

type MateDataSlot51 struct {
	MutArr []int
}

type MateDataSlot57 struct {
	Mut            int         //当前局 倍率 *N
	WildProgress   int         //进度
	HistoryWildMun int         //历史wild收集的数量
	WildTags       []*base.Tag //当前局 的wild标签
	MutIsTop       bool
}

type MateDataSlot58 struct {
	//1 如果中间卷轴上有一个标签，而最左卷轴和最右卷轴上的标签都是coin标签，但不相同，则最左卷轴和最右卷轴将重新旋转一次
	Trigger int           // 1 重转模式
	ColList [][]*base.Tag `json:"-"`
}

type ExpandTagData struct {
	ExpDataList [][]*base.Tag //列 拓展标签转换后的 窗口数据
	ExpCols     []int
	ExpandTag   *base.Tag //拓展标签（logicPack 10）
}

func NewScatLinkStep(id int, initLinks, addTags []*base.Tag, fTag *base.Tag, LkType int, initList [][]*base.Tag) *flow.SpinFlow {
	flow := &flow.SpinFlow{
		Id:       id,
		InitList: initList,
		OmitList: nil,
		AddList:  addTags,
		EmitList: nil,
		Level:    LkType,
		MateSlot21: flow.MateSlot21{
			InitLists: initLinks,
		},
	}
	if fTag != nil {
		flow.SkillCenter = append(flow.SkillCenter, fTag)
	}
	return flow
}

//type MoveResData struct {
//	MoveResDataList [][]*base.Tag
//}
