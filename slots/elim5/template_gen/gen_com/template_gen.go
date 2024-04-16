package gen_com

import (
	"elim5/logicPack/base"
	"elim5/logicPack/component"
	"elim5/model/business"
	"sync"
)

type GenTemplate struct {
	Config        *component.Config         // 配置
	TemGen        *business.SlotTemplateGen // 模版
	InitialWeight map[int]map[string]int    // 列=>(标签=>数量); 意思每一列对应多少不同标签种类 每个标签根据权重有多少个;
	LargeScale    map[int][]*WeightInterval // 列=>权重区间; 意思每一列对应多少不同标签种类 每个标签根据权重有多少个;
	Interval      []*WeightInterval         // 权重区间
	TrimDown      map[int][]*Scale          // 列=>(标签=>替换标签); 意思每一列对应多少不同标签种类 每个标签根据权重有多少个;
	TrimUp        map[int][]*Scale          // 列=>(标签=>替换标签); 意思每一列对应多少不同标签种类 每个标签根据权重有多少个;
	Template      map[int][]uint16          //key 代表列; val 代表列生成的标签个数（个数根据标签权重计算:标签权重的后一个位置减去前一个位置）
	ExtraTemplate map[int][]uint16          //key 代表列; val 代表列生成的标签个数（个数根据标签权重计算:标签权重的后一个位置减去前一个位置）
	SpecialWeight map[int]any               // 特殊权重
	Adjacent      map[int]map[string]int    // 列=>(标签=>数量); 意思每一列对应多少不同标签种类 每个标签根据权重有多少个;
	CondMap       map[string]*Cond          // 条件
	AdjRecords    int                       // 相邻记录
	Closest       float64                   // 最接近的值
	TemInitIndex  int                       // 模版初始索引
	Raise         bool                      // 是否上升
	mu            sync.Mutex                // 互斥锁
	RatioConfirm  int                       // 返奖率确认
	Which         int                       // 选择
}

func (t *GenTemplate) GetGameType() int {
	return int(t.TemGen.Type)
}

type TagCount struct {
	Tag   *base.Tag
	Count int
}

type WeightInterval struct {
	Tag      *base.Tag
	MinCount int
	MaxCount int
}

type Scale struct {
	Tag        *base.Tag
	ReplaceTag *base.Tag
}

func (t *GenTemplate) GetCond(str string) *Cond {
	if t.CondMap == nil {
		t.CondMap = map[string]*Cond{}
	}
	if _, ok := t.CondMap[str]; !ok {
		t.CondMap[str] = &Cond{}
	}
	return t.CondMap[str]
}

func (t *GenTemplate) SetCond(str string, cond *Cond) {
	if t.CondMap == nil {
		t.CondMap = map[string]*Cond{}
	}
	t.CondMap[str] = cond
}
