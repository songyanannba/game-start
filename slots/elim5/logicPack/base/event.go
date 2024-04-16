package base

import (
	"elim5/enum"
	"elim5/global"
	"elim5/utils"
	"elim5/utils/helper"
	"fmt"
	"strconv"
	"strings"
)

type Event struct {
	m map[int]any
}

func NewEvent() *Event {
	return &Event{m: map[int]any{}}
}

func (e *Event) Add(typ int, params string) {
	var event any
	switch typ {
	// 目前换表和免费转事件参数一致
	case enum.SlotEvent1ChangeTable:
		event = ParseWeightData(params)
	default:
		event = ParseWeightData(params)
	}
	if event != nil {
		e.m[typ] = event
	}
}

func (e *Event) AddDefault(typ int, params string) {
	var event any
	strs := helper.SplitStr(params, "&")
	_, err := strconv.Atoi(strs[0])
	if err != nil {
		event = ParseWeightDataStr(params)
	} else {
		event = ParseWeightData(params)
	}

	if event != nil {
		e.m[typ] = event
	}
}

func (e *Event) GetChangeTableEvent(typ int) *ChangeTableEvent {
	res, ok := e.m[typ]
	if !ok {
		return nil
	}
	event := res.(*ChangeTableEvent)
	return event
}

func (e *Event) Get(typ int) (any, error) {
	res, ok := e.m[typ]
	if ok {
		return res, nil
	} else {
		return nil, fmt.Errorf("event not found %d", typ)
	}
}

// ChangeTableEvent 数值换表事件
type ChangeTableEvent struct {
	Data   []int
	Weight []int
}

// ChangeTableStrEvent 字符换表事件
type ChangeTableStrEvent struct {
	Data   []string
	Weight []int
}

// IntervalRatioEvent 倍率区间
type IntervalRatioEvent struct {
	Data   [][2]int
	weight []int
}

type LevelInfo struct {
	Level    int
	Count    int
	Multiple int
}

type Unit6LevelEvent struct {
	Collect   int               // 收集的数量
	CoreCount int               // 核心数量
	EmitEvent *ChangeTableEvent // 发射的数量
	WildEvent *ChangeTableEvent // wild的数量
}

// Fetch 根据权重获取一个值
func (e ChangeTableEvent) Fetch() int {
	k := helper.RandomLongWeight(e.Weight)
	return helper.SliceVal(e.Data, k)
}

func (s ChangeTableStrEvent) Fetch() string {
	k := helper.RandomLongWeight(s.Weight)
	return helper.SliceVal(s.Data, k)
}

func (s ChangeTableStrEvent) Copy() *ChangeTableStrEvent {
	data := make([]string, len(s.Data))
	copy(data, s.Data)
	weight := make([]int, len(s.Weight))
	copy(weight, s.Weight)
	return &ChangeTableStrEvent{
		Data:   data,
		Weight: weight,
	}
}

func (s ChangeTableStrEvent) GetSection(i int) int {
	if len(s.Weight) > i+1 && i >= 0 {
		return s.Weight[i+1] - s.Weight[i]
	}
	return 0
}

func (s IntervalRatioEvent) Fetch() ([2]int, int) {
	k := helper.RandomLongWeight(s.weight)
	return helper.SliceVal(s.Data, k), k
}

func newChangeTableEvent(data, weight []int) *ChangeTableEvent {
	return &ChangeTableEvent{
		Data:   data,
		Weight: weight,
	}
}

// ParseWeightData 解析权重数据的格式
func ParseWeightData(s string) *ChangeTableEvent {
	if s == "" {
		return nil
	}
	s = utils.FormatCommandStr(s)
	tagStr, weightStr, _ := strings.Cut(s, "@")
	tags := helper.SplitInt[int](tagStr, "&")
	weights := helper.SplitInt[int](weightStr, "&")
	if len(weights)-1 != len(tags) {
		return nil
	}
	return newChangeTableEvent(tags, weights)
}

func ParseWeightDataStr(s string) *ChangeTableStrEvent {
	if s == "" {
		return nil
	}
	tagStr, weightStr, _ := strings.Cut(s, "@")
	tags := helper.SplitStr(tagStr, "&")
	weights := helper.SplitInt[int](weightStr, "&")
	if len(weights)-1 != len(tags) {
		return nil
	}
	return &ChangeTableStrEvent{
		Data:   tags,
		Weight: weights,
	}
}

func ParseEventLevel(s string) map[int]*LevelInfo {
	if s == "" {
		return nil
	}
	str := utils.FormatCommandStr(s)
	strs := helper.SplitStr(str, "@")
	if len(strs) != 3 {
		global.GVA_LOG.Error("ParseEventLevel error " + s)
		return nil
	}
	rMap := map[int]*LevelInfo{}
	level := helper.SplitInt[int](strs[0], "&")
	collect := helper.SplitInt[int](strs[1], "&")
	multiple := helper.SplitInt[int](strs[2], "&")
	for i, i2 := range level {
		if i >= len(collect) || i >= len(multiple) {
			continue
		}
		rMap[i2] = &LevelInfo{
			Level:    i2,
			Count:    collect[i],
			Multiple: multiple[i],
		}
	}
	return rMap
}

type Unit51LevelEvent struct {
	Data   []int
	Weight []int
	MutArr []int
}

func (e Unit51LevelEvent) Fetch() int {
	k := helper.RandomLongWeight(e.Weight)
	return helper.SliceVal(e.Data, k)
}
