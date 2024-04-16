package logic

import (
	"elim5/global"
	"elim5/logicPack/base"
	"elim5/model/business"
	"elim5/service/cache"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
)

// Merger 合并主逻辑
func (d *Dismantling) Merger() error {
	num := d.TemplateLen - d.GetRow()

	for i := 0; i < num; i++ {
		d.GetDisplay(i)
		d.AddExtraTags(i)
		err := d.E.Merger(i)
		if err != nil {
			return err
		}
		d.UpdateDisplay(i)
		if i%(num/20) == 0 && i < (num-1) {
			WritingProgress(num, i, "合并中")
		}
	}
	return nil
}

// Save 保存合并后的模版
func (d *Dismantling) Save() error {
	tems := make([]business.SlotTemplate, 0)
	toString, err := global.Json.MarshalToString(d.Req)
	if err != nil {
		return err
	}
	tem := d.TemGens[0]
	for col, _ := range d.Template {
		item := &business.SlotTemplate{
			SlotId:    tem.SlotId,
			Type:      tem.Type,
			Column:    col,
			GenId:     int(tem.ID),
			MergeInfo: toString,
			Rtp:       tem.Rtp,
			Which:     tem.Which,
		}
		item.Layout = d.GetTemLayout(col)
		resLayout := d.GetTemResLayout(col)
		if resLayout != "" {
			item.Layout += fmt.Sprintf(",%s", resLayout)
		}
		tems = append(tems, *item)
	}
	return global.NOLOG_DB.Create(&tems).Error
}

// SaveRes 只保存不赢钱模版用于测试
func (d *Dismantling) SaveRes() error {
	tems := make([]business.SlotTemplate, 0)
	toString, err := global.Json.MarshalToString(d.Req)
	if err != nil {
		return err
	}
	tem := d.TemGens[0]
	for col, _ := range d.Template {
		item := &business.SlotTemplate{
			SlotId:    tem.SlotId,
			Type:      tem.Type,
			Column:    col,
			GenId:     int(tem.ID),
			MergeInfo: toString,
			Rtp:       tem.Rtp,
			Which:     tem.Which,
		}
		//item.Layout = d.GetTemLayout(col)
		//resLayout := d.GetTemResLayout(col)
		//if resLayout != "" {
		//	item.Layout += fmt.Sprintf(",%s", resLayout)
		//}
		item.Layout = d.GetTemResLayout(col)
		tems = append(tems, *item)
	}
	return global.NOLOG_DB.Create(&tems).Error
}

// WritingProgress 写入进度
func WritingProgress(count, num int, info string) {
	f := float64(0)
	if count != 0 {
		f = float64(num) / float64(count) * 100
	}
	cache.SetTaskProgress(cache.TemplateMergerKey, &cache.TaskProgress{
		Type:     cache.TemplateMergerKey,
		Progress: int(f),
		Info:     info,
	})
}

// GetDisplay 填充需要拆解的模版 （index 总模版的 行索引）
func (d *Dismantling) GetDisplay(index int) {
	if index == 0 {
		d.SpinInfo.Display = helper.NewTable(d.GetCol(), d.GetRow(), func(x, y int) *base.Tag {
			fillTag := d.GetTemTag(x+index, y)
			fillTag.X = x
			fillTag.Y = y
			return fillTag
		})
	} else {
		//窗口组成 （从结果模版取四行 填充display前4行 + 从总模版取一行 填充到display最后一行）
		for row, tags := range d.SpinInfo.Display {
			for col, _ := range tags {
				if row == d.GetRow()-1 {
					fillTag := d.GetTemTag(row+index, col)
					fillTag.X = row
					fillTag.Y = col
					d.SpinInfo.Display[row][col] = fillTag
				} else {
					fillTag := d.Config.GetTagById(int(d.TemplateRes[col][row+index]))
					fillTag.X = row
					fillTag.Y = col
					d.SpinInfo.Display[row][col] = fillTag
				}
			}
		}
	}

}

// AddExtraTags 填充额外的tag
func (d *Dismantling) AddExtraTags(index int) {
	fillTags := make([]*base.Tag, 0)
	for i, tag := range d.Req.ExtraTags {
		tags := d.SpinInfo.FindTagsByName(tag.Name)
		if len(tags) > 0 {
			continue
		}
		randCount := helper.RandInt(tag.Max - len(tags))
		for count := 0; count < randCount; count++ {
			fillTags = append(fillTags, d.SpinInfo.Config.GetTagByName(tag.Name).Copy())
		}
		d.Req.ExtraTags[i].Count -= randCount
	}

	if len(fillTags) == 0 {
		return
	}

	needTags := d.SpinInfo.FindAllTags()
	if index != 0 {
		needTags = lo.Filter(needTags, func(item *base.Tag, index int) bool {
			return item.X == d.SpinInfo.Config.GetRow()-1
		})
	}
	helper.SliceShuffle(needTags)

	for i, tag := range fillTags {
		if i >= len(needTags) {
			break
		}
		fillTag := tag.Copy()
		fillTag.X = needTags[i].X
		fillTag.Y = needTags[i].Y
		d.SpinInfo.Display[needTags[i].X][needTags[i].Y] = fillTag
	}
}

// UpdateDisplay 更新拆解后的模版
func (d *Dismantling) UpdateDisplay(index int) {
	if index == 0 {
		for _, tags := range d.SpinInfo.Display {
			for col, tag := range tags {
				d.TemplateRes[col] = append(d.TemplateRes[col], uint16(tag.Id))
			}
		}
	} else {
		tags := lo.Filter(d.SpinInfo.FindAllTags(), func(item *base.Tag, index int) bool {
			return item.X == d.GetRow()-1
		})
		for _, tag := range tags {
			d.TemplateRes[tag.Y] = append(d.TemplateRes[tag.Y], uint16(tag.Id))
		}
	}
}
