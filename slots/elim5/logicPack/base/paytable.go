package base

import (
	"elim5/enum"
	"elim5/utils/helper"
	"fmt"

	"strings"
)

type PayTable struct {
	Id        uint
	Type      int
	Tags      []*Tag
	Num       int
	Tags2     []*Tag
	Num2      int
	Multiple  float64
	SingleMap map[*Tag]bool // 已使用的单出
}

func (p PayTable) String() string {
	return fmt.Sprintf("%v  mul: %f", p.Tags, p.Multiple)
}

func (p PayTable) Dump() string {
	var arr []string
	for _, tag := range p.Tags {
		arr = append(arr, tag.Name)
	}
	return strings.Join(arr, ",")
}

func NewCommonPayTable(id uint, tags []*Tag, multiple float64, single ...*Tag) *PayTable {
	singleMap := make(map[*Tag]bool)
	for _, tag := range single {
		singleMap[tag] = true
	}
	return &PayTable{
		Id:        id,
		Type:      enum.SlotPayTableType1Common,
		Tags:      tags,
		Multiple:  multiple,
		SingleMap: singleMap,
	}
}

func NewAnyPayTable(id uint, tags []*Tag, num int, tags2 []*Tag, num2 int, multiple float64) *PayTable {
	return &PayTable{
		Id:       id,
		Type:     enum.SlotPayTableType2Any,
		Tags:     tags,
		Num:      num,
		Tags2:    tags2,
		Num2:     num2,
		Multiple: multiple,
	}
}

// Match 匹配并返回新的payTable结果
func (p PayTable) Match(tags []*Tag) (bool, *PayTable) {
	if p.Type == enum.SlotPayTableType1Common {
		return p.matchCommon(tags)
	} else {
		return p.matchAny(tags)
	}
}

func (p PayTable) matchCommon(tags []*Tag) (bool, *PayTable) {
	pLen := len(p.Tags)
	tLen := len(tags)
	if pLen == 0 || tLen == 0 || pLen > tLen {
		return false, nil
	}
	var matchKey []int
	for i, tag := range p.Tags {
		if !tags[i].IsMatchName(tag.Name) {
			matchKey = []int{}
			return false, nil
		}
		matchKey = append(matchKey, i)
	}
	return true, setMatchTags(p.Id, tags, matchKey, p.Multiple)
}

func (p PayTable) matchAny(tags []*Tag) (bool, *PayTable) {
	if len(tags) == 0 {
		return false, nil
	}
	var (
		matchKey []int
	)
	if len(p.Tags) != 0 && p.Num != 0 {
		matchNum := 0
		for i, tag := range tags {
			if tag.InTags(p.Tags) {
				matchNum++
				matchKey = append(matchKey, i)
				if matchNum >= p.Num {
					break
				}
			}
		}
		if matchNum < p.Num {
			return false, nil
		}
	}
	if len(p.Tags2) != 0 && p.Num2 != 0 {
		matchNum := 0
		for i, tag := range tags {
			if tag.InTags(p.Tags2) {
				matchNum++
				if !helper.InArr(i, matchKey) {
					matchKey = append(matchKey, i)
				}
				if matchNum >= p.Num2 {
					break
				}
			}
		}
		if matchNum < p.Num2 {
			return false, nil
		}
	}
	return true, setMatchTags(p.Id, tags, matchKey, p.Multiple)
}

// 将匹配payTable的tag设为true 并返回payTable列表
func setMatchTags(id uint, tags []*Tag, matchKey []int, multiple float64) *PayTable {
	var (
		res        []*Tag
		singleTags []*Tag
	)
	for _, value := range matchKey {
		t := tags[value]
		t.IsPayTable = true
		if t.Multiple > 1 {
			multiple *= t.Multiple
		}
		if t.IsSingle {
			singleTags = append(singleTags, t)
		}
		res = append(res, t)
	}
	return NewCommonPayTable(id, res, multiple, singleTags...)
}
