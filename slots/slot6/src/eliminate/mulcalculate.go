package eliminate

import (
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"slot6/src/base"
	"slot6/utils/helper"
	"sort"
)

// WinMatchList 赢钱划线匹配
func (t *Table) WinMatchList(tagList [][]*base.Tag, isFree bool) ([]*base.Eliminate, float64) {
	var eliminates []*base.Eliminate
	sumMul := float64(0)
	for _, tags := range tagList {
		//mul := t.WinMatch(tags)
		//获取匹配画线的赔率
		mul := t.WinMatchV2(tags)
		doubleTags := t.DoubleProc(tags, isFree)

		for _, tag := range doubleTags {
			if tag.Multiple > 0 {
				mul = mul.Mul(decimal.NewFromFloat(tag.Multiple))
			}
		}
		//wild倍率
		for _, tag := range tags {
			if tag.IsWild {
				mul = mul.Mul(decimal.NewFromFloat(tag.Multiple))
			}
		}
		sumMul, _ = decimal.NewFromFloat(sumMul).Add(mul).Float64()
		elMul, _ := mul.Float64()

		eliminate := &base.Eliminate{
			RemoveList: doubleTags,
			Mul:        elMul,
		}

		eliminates = append(eliminates, eliminate)
	}
	return eliminates, sumMul
}

// DoubleProc FreeSpin随机翻倍标签
func (t *Table) DoubleProc(tags []*base.Tag, isFree bool) []*base.Tag {
	rTags := make([]*base.Tag, len(tags))
	for i, i2 := range tags {
		tag := i2.Copy()
		tag.Multiple = 0
		rTags[i] = tag
	}
	if !isFree {
		return rTags
	} else {
		mulNum := t.Target.MulNumEvent.Fetch()
		if mulNum >= len(rTags) {
			for c, _ := range rTags {
				rTags[c].Multiple = 2
			}
		} else {
			v := make(map[int]int)
			count := 0
			for count < mulNum {
				index := helper.RandInt(len(rTags))
				if _, ok := v[index]; !ok {
					v[index] = index
					count++
				}
			}
			for c, _ := range rTags {
				if _, ok := v[c]; ok {
					rTags[c].Multiple = 2
				} else {
					rTags[c].Multiple = 0
				}
			}
		}
	}
	return rTags
}

// WinMatch 赢钱划线匹配
func (t *Table) WinMatch(tags []*base.Tag) decimal.Decimal {
	mul := decimal.NewFromInt(0)
	payTables := lo.Filter(t.PayTableList, func(item *base.PayTable, index int) bool {
		return item.Tags[0].Name == tags[0].Name
	})
	if len(payTables) == 0 {
		return mul
	}

	sort.Slice(payTables, func(i, j int) bool {
		return len(payTables[i].Tags) > len(payTables[j].Tags)
	})

	payTable := lo.Filter(payTables, func(item *base.PayTable, index int) bool {
		return len(item.Tags) == len(tags)
	})

	if len(payTable) > 0 {
		mul = mul.Add(decimal.NewFromFloat(payTable[0].Multiple))
	} else if len(tags) > len(payTables[0].Tags) {
		mul = mul.Add(decimal.NewFromFloat(payTables[0].Multiple))
	}

	return mul
}

// WinMatchV2 单次 经测试 WinMatchV2 比  WinMatch 快 0 -20 微妙
// 没经过反射；map o1 操作
func (t *Table) WinMatchV2(tags []*base.Tag) decimal.Decimal {
	mul := decimal.NewFromInt(0)
	payTableListMaps := t.PayTableListMaps
	if payTableListMaps == nil || len(payTableListMaps) == 0 {
		return mul
	}
	payTableLists, ok := payTableListMaps[tags[0].Name]
	if !ok {
		return mul
	}
	// 和每种画线的标签数量进行比较 数量相等 即匹配成功
	if payTable, ok := payTableLists[len(tags)]; ok {
		mul = mul.Add(decimal.NewFromFloat(payTable.Multiple))
	} else {
		//如果没有匹配到，选择最大画线的标签数量
		// 找到相同标签 数量（划线tags）最大的划线 获取最大 Multiple
		var multiple float64
		for pk, pv := range payTableLists {
			if pk <= len(tags) && multiple < pv.Multiple {
				multiple = pv.Multiple
			}
		}
		mul = mul.Add(decimal.NewFromFloat(multiple))
	}
	return mul
}

func (t *Table) WinMatchName(name string, l int) decimal.Decimal {
	mul := decimal.NewFromInt(0)
	payTables := lo.Filter(t.PayTableList, func(item *base.PayTable, index int) bool {
		return item.Tags[0].Name == name
	})
	if len(payTables) == 0 {
		return mul
	}

	sort.Slice(payTables, func(i, j int) bool {
		return len(payTables[i].Tags) > len(payTables[j].Tags)
	})

	payTable := lo.Filter(payTables, func(item *base.PayTable, index int) bool {
		return len(item.Tags) == l
	})

	if len(payTable) > 0 {
		mul = mul.Add(decimal.NewFromFloat(payTable[0].Multiple))
	} else if l > len(payTables[0].Tags) {
		mul = mul.Add(decimal.NewFromFloat(payTables[0].Multiple))
	}

	return mul
}
