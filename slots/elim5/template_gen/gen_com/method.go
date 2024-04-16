package gen_com

import (
	"elim5/enum"
	"elim5/global"
	"elim5/logicPack/base"
	"elim5/model/business"
	"elim5/utils/helper"
	"fmt"
)

// CreateTem 创建模版
func (t *GenTemplate) CreateTem(tem *business.SlotTemplateGen) {
	var items []*business.SlotTemplate
	for i, tags := range t.Template {
		item := &business.SlotTemplate{
			SlotId: tem.SlotId,
			Type:   tem.Type,
			Column: i,
			GenId:  int(tem.ID),
			Rtp:    tem.Rtp,
			Which:  tem.Which,
		}
		temStr := ""
		for _, tag := range tags {
			temStr += fmt.Sprintf("%d,", tag)
		}
		temStr += "\n"
		item.Layout = temStr
		items = append(items, item)
	}
	global.NOLOG_DB.Create(&items)
}

// GetFinalWeight 获取最后模版
func (t *GenTemplate) GetFinalWeight(weStr string) string {
	finalWeight := ""
	colMap, _ := GetColMap(weStr)
	for col := 0; col < t.Config.Index; col++ {
		str := ""
		if colMap[col] != "" {
			str = colMap[col]
		} else {
			str = colMap[0]
		}

		finalWeightCol := ""
		if str == "" {
			continue
		}
		tagsStr := helper.SplitStr(str, "@")[0]
		finalWeightCol += fmt.Sprintf("%d:%s@0", col, tagsStr)
		tags := helper.SplitStr(tagsStr, "&")

		count := 0
		for _, tag := range tags {
			if tag == "" {
				continue
			}
			count += t.InitialWeight[col][tag]
			finalWeightCol += fmt.Sprintf("&%d%s", count, helper.If(t.Adjacent[col][tag] > 0, fmt.Sprintf(",%d", t.Adjacent[col][tag]), ""))
		}
		finalWeight += fmt.Sprintf("%s\n", finalWeightCol)
	}

	return finalWeight
}

// GetFinalTemplate 获取最后模版
func (t *GenTemplate) GetFinalTemplate() string {
	temStr := ""
	for i := 0; i < t.Config.Index; i++ {
		tags := t.Template[i]
		temStr += fmt.Sprintf("%d:", i)
		for _, tag := range tags {
			temStr += fmt.Sprintf("%d,", tag)
		}
		temStr += "\n"
	}
	return temStr
}

func (t *GenTemplate) GetAdjacentAllKeys(col int) []string {
	var keys []string
	for key := range t.Adjacent[col] {
		if key == enum.ScatterName {
			continue
		}
		keys = append(keys, key)
	}
	return keys
}

// GetAllTemplateTags 获取所有标签
func (t *GenTemplate) GetAllTemplateTags() []*base.Tag {
	var reTags []*base.Tag
	for c, ids := range t.Template {
		for r, id := range ids {
			newTag := t.Config.GetTagById(int(id)).Copy()
			newTag.X = r
			newTag.Y = c
			reTags = append(reTags, newTag)
		}
	}
	return reTags
}

// GetAllTemplateIndexTags 获取所有标签
func (t *GenTemplate) GetAllTemplateIndexTags(index int) []*base.Tag {
	var tags []*base.Tag
	for r, id := range t.Template[index] {
		newTag := t.Config.GetTagById(int(id)).Copy()
		newTag.X = r
		newTag.Y = index
		tags = append(tags, newTag)
	}
	return tags
}

func (t *GenTemplate) GetAdjacentMoreKeys(col int) []string {
	var keys []string
	for key := range t.Adjacent[col] {
		if key == enum.ScatterName || key == enum.MultiplierName {
			continue
		}
		keys = append(keys, key)
	}
	return keys
}

func (t *GenTemplate) GetDisparity(res *LogicResult) float64 {
	totalInterval := 0.0
	for s, cond := range t.CondMap {
		switch s {
		case ScaTriggerCond:
			totalInterval += cond.GetDisparity(res.ScatterRatio)
		case GainRatioCond:
			totalInterval += cond.GetDisparity(res.GainRatio)
		case WinRateCond:
			totalInterval += cond.GetDisparity(res.WinRatio)
		case RemoveRateCond:
			totalInterval += cond.GetDisparity(res.RemoveRate)
		}
	}
	return totalInterval
}
