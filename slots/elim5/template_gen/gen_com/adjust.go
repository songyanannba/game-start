package gen_com

import (
	"elim5/enum"
	"elim5/logicPack/base"
	"elim5/utils/helper"
	"github.com/samber/lo"
)

// AdjLarge  大调整
func AdjLarge(t *GenTemplate) {
	for i, intervals := range t.LargeScale {
		for _, interval := range intervals {
			if interval.MaxCount-interval.MinCount >= 0 {
				num := helper.RandInt(interval.MaxCount-interval.MinCount) + interval.MinCount
				t.InitialWeight[i][interval.Tag.Name] = num
			}
		}
	}
}

// AdjTrimDown 向下微调
func AdjTrimDown(t *GenTemplate) {
	rendCol := helper.RandInt(t.Config.Index)
	scales := t.TrimDown[rendCol]
	if scales == nil || len(scales) == 0 {
		return
	}
	randScales := scales[helper.RandInt(len(scales))]
	t.InitialWeight[rendCol][randScales.Tag.Name]--
	t.InitialWeight[rendCol][randScales.ReplaceTag.Name]++
}

// AdjTrimUp 向上微调
func AdjTrimUp(t *GenTemplate) {
	rendCol := helper.RandInt(t.Config.Index)
	scales := t.TrimUp[rendCol]
	if scales == nil || len(scales) == 0 {
		return
	}
	randScales := scales[helper.RandInt(len(scales))]
	t.InitialWeight[rendCol][randScales.Tag.Name]--
	t.InitialWeight[rendCol][randScales.ReplaceTag.Name]++
}

// AdjScatterUp Sca向上微调
func AdjScatterUp(t *GenTemplate) {
	col := 0
	num := 0
	for i, m := range t.InitialWeight {
		if m[enum.ScatterName] <= num {
			col = i
			num = m[enum.ScatterName]
		}
	}
	t.InitialWeight[col][enum.ScatterName]++
}

// AdjScatterDown Sca向下微调
func AdjScatterDown(t *GenTemplate) {
	col := 0
	num := 0
	for i, m := range t.InitialWeight {
		if m[enum.ScatterName] > num {
			col = i
			num = m[enum.ScatterName]
		}
	}
	t.InitialWeight[col][enum.ScatterName]--
}

// RemoveUp 连消向上微调
func RemoveUp(t *GenTemplate) {
	randCol := helper.RandInt(len(t.Adjacent))
	keys := t.GetAdjacentAllKeys(randCol)
	if len(keys) == 0 {
		return
	}
	randKey := keys[helper.RandInt(len(keys))]
	t.Adjacent[randCol][randKey]++
}

// RemoveDown 连消向下微调
func RemoveDown(t *GenTemplate) {
	randCol := helper.RandInt(len(t.Adjacent))
	keys := t.GetAdjacentMoreKeys(randCol)
	if len(keys) == 0 {
		return
	}
	randKey := keys[helper.RandInt(len(keys))]
	t.Adjacent[randCol][randKey]--
}

// TemAdjTrimDown 向下微调
func TemAdjTrimDown(t *GenTemplate) {
	randCol := helper.RandInt(t.Config.Index)
	scales := t.TrimDown[randCol]
	randScales := scales[helper.RandInt(len(scales))]
	tags := t.GetAllTemplateTags()
	modifyTags := lo.Filter(tags, func(tag *base.Tag, i int) bool {
		return tag.Name == randScales.Tag.Name && tag.Y == randCol
	}) //随机列 里面存在的被调整标签的集合
	if len(modifyTags) == 0 {
		return
	}
	randModify := modifyTags[helper.RandInt(len(modifyTags))]
	if t.InitialWeight[randCol][randModify.Name] == 1 {
		return
	}
	//模版中对应的标签被期望的标签所替代 然后权重对应列的标签权重做相应调整
	t.Template[randCol][randModify.X] = uint16(randScales.ReplaceTag.Id)
	t.InitialWeight[randCol][randModify.Name]--
	t.InitialWeight[randCol][randScales.ReplaceTag.Name]++
}

// TemAdjTrimUp 向上微调
func TemAdjTrimUp(t *GenTemplate) {
	randCol := helper.RandInt(t.Config.Index)
	scales := t.TrimUp[randCol]
	randScales := scales[helper.RandInt(len(scales))]

	tags := t.GetAllTemplateTags()
	modifyTags := lo.Filter(tags, func(tag *base.Tag, i int) bool {
		return tag.Name == randScales.Tag.Name && tag.Y == randCol
	})
	if len(modifyTags) == 0 {
		return
	}
	randModify := modifyTags[helper.RandInt(len(modifyTags))]
	if t.InitialWeight[randCol][randModify.Name] == 1 {
		return
	}
	//模版中对应的标签被期望的标签所替代 然后权重对应列的标签权重做相应调整
	t.Template[randCol][randModify.X] = uint16(randScales.ReplaceTag.Id)

	t.InitialWeight[randCol][randModify.Name]--
	t.InitialWeight[randCol][randScales.ReplaceTag.Name]++

}

// TemRemoveUp  连消向上微调
func TemRemoveUp(t *GenTemplate) {
	//randCol := helper.RandInt(t.Config.Index)
	//keys := t.GetAdjacentAllKeys(randCol)
	//if len(keys) == 0 {
	//	return
	//}
	//randKey := keys[helper.RandInt(len(keys))]
	//t.Adjacent[randCol][randKey]++
	//
	//allTags := t.Template[randCol]
	//tag := t.Config.GetTag(randKey)
	//
	//adjacency := GetRandAdjacency(allTags, tag)
	//if adjacency == -1 {
	//	return
	//}
	//t.Template[randCol] = ArrayInsert(allTags, adjacency, tag)

	keys := t.GetAdjacentMoreKeys(0)
	randKey := keys[helper.RandInt(len(keys))]
	for i, _ := range t.Adjacent {
		t.Adjacent[i][randKey] += 10
	}
}

// TemRemoveDown  连消向下微调
func TemRemoveDown(t *GenTemplate) {
	//randCol := helper.RandInt(t.Config.Index)
	//keys := t.GetAdjacentMoreKeys(randCol)
	//if len(keys) == 0 {
	//	return
	//}
	//randKey := keys[helper.RandInt(len(keys))]
	//if t.Adjacent[randCol][randKey] == 0 {
	//	return
	//}
	//t.Adjacent[randCol][randKey]--
	//
	//allTags := t.Template[randCol]
	//tag := t.Config.GetTag(randKey)
	//adjacency := GetRandAdjacencyMul(allTags, tag)
	//if adjacency == -1 {
	//	return
	//}
	//t.Template[randCol] = ArrayRemove(allTags, adjacency)
	keys := t.GetAdjacentMoreKeys(0)
	randKey := keys[helper.RandInt(len(keys))]
	for i, _ := range t.Adjacent {
		t.Adjacent[i][randKey] += 10
	}
}

// TemAdjScatterUp Sca向上微调
func TemAdjScatterUp(t *GenTemplate) {

	_, randCol := MapMinBy(t.InitialWeight, func(value1, value2 map[string]int) bool {
		if value1[enum.ScatterName] <= value2[enum.ScatterName] || value2 == nil {
			return true
		} else {
			return false
		}
	})
	randCol = GetScatRandCol(t.InitialWeight, int(t.Config.SlotId), randCol)

	t.InitialWeight[randCol][enum.ScatterName]++
	allTags := t.GetAllTemplateIndexTags(randCol)

	scaTags := lo.Filter(allTags, func(tag *base.Tag, i int) bool {
		return tag.Name == enum.ScatterName
	})
	fillTags := lo.Filter(allTags, func(tag *base.Tag, i int) bool {
		for _, scaTag := range scaTags {
			distance := helper.Abs(tag.X - scaTag.X)
			if distance <= 10 || distance >= len(allTags)-10 {
				return false
			}
		}
		return true
	})
	if len(fillTags) == 0 {
		return
	}
	randFill := fillTags[helper.RandInt(len(fillTags))]

	newTags := ArrayInsert(allTags, randFill.X, t.Config.GetTag(enum.ScatterName))
	ids := make([]uint16, 0)
	for _, tag := range newTags {
		ids = append(ids, uint16(tag.Id))
	}
	t.Template[randCol] = ids
}

// TemAdjScatterDown Sca向下微调
func TemAdjScatterDown(t *GenTemplate) {
	_, randCol := MapMinBy(t.InitialWeight, func(value1, value2 map[string]int) bool {
		if value1[enum.ScatterName] >= value2[enum.ScatterName] || value2 == nil {
			return true
		} else {
			return false
		}
	})
	randCol = GetScatRandCol(t.InitialWeight, int(t.Config.SlotId), randCol)

	if t.InitialWeight[randCol][enum.ScatterName] == 0 {
		return
	}
	t.InitialWeight[randCol][enum.ScatterName]--
	allTags := t.GetAllTemplateIndexTags(randCol)

	scaTags := lo.Filter(allTags, func(tag *base.Tag, i int) bool {
		return tag.Name == enum.ScatterName
	})
	if len(scaTags) == 0 {
		return
	}
	randSca := scaTags[helper.RandInt(len(scaTags))]

	newTags := ArrayRemove(allTags, randSca.X)
	ids := make([]uint16, 0)
	for _, tag := range newTags {
		ids = append(ids, uint16(tag.Id))
	}
	t.Template[randCol] = ids
}

func GetScatRandCol(InitialWeight map[int]map[string]int, slotId int, randCol int) int {
	if slotId == enum.SlotId9 { //第九台 只有 第1 3 5列 才能存在scat 增加权重只能在对应列添加
		cols := ExistScatCol(InitialWeight)
		randCol = cols[helper.RandInt(len(cols))]
	}

	return randCol
}

func ExistScatCol(InitialWeight map[int]map[string]int) []int {
	var cols []int
	for col, iw := range InitialWeight {
		if v, ok := iw[enum.ScatterName]; ok && v > 0 {
			cols = append(cols, col)
		}
	}
	return cols
}
