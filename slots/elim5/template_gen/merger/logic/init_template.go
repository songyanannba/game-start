package logic

import (
	"elim5/logicPack/base"
	"elim5/utils"
	"elim5/utils/helper"
	"fmt"
	"strconv"
)

// InitTemplate 初始化模版
func (d *Dismantling) InitTemplate() error {
	d.Template = make(map[int][]*base.Tag)
	for _, gen := range d.TemGens {
		err := d.AppendTemplate(gen.Template)
		if err != nil {
			return err
		}
		if len(d.Template) != d.Config.Index {
			return fmt.Errorf("模版列数不正确")
		}
		// 检查模版是否正确
		TemLen := 0
		for _, tags := range d.Template {
			if len(tags) == 0 {
				return fmt.Errorf("模版长度为0")
			}
			if TemLen == 0 {
				TemLen = len(tags)
			} else if TemLen != len(tags) {
				return fmt.Errorf("模版长度不正确")
			}
		}
		d.TemplateLen = int(helper.MulToInt(TemLen, d.Req.Extra))
	}

	return nil
}

// AppendTemplate 追加模版
func (d *Dismantling) AppendTemplate(temStr string) error {
	colStrs := utils.FormatCommand(temStr)
	for _, colStr := range colStrs {
		if colStr == "" {
			continue
		}
		err := d.AppendTemplateCol(colStr)
		if err != nil {
			return err
		}
	}
	return nil
}

// AppendTemplateCol 追加模版列
func (d *Dismantling) AppendTemplateCol(str string) error {
	strs := helper.SplitStr(str, ":")
	if len(strs) != 2 {
		return fmt.Errorf("格式错误")
	}
	col, err := strconv.Atoi(strs[0])
	if err != nil {
		return err
	}
	tagIds := helper.SplitStr(strs[1], ",")
	for _, tagIdStr := range tagIds {
		if tagIdStr == "" {
			continue
		}
		tagId := 0
		tagId, err = strconv.Atoi(tagIdStr)
		if err != nil {
			return err
		}
		tag := d.Config.GetTagById(tagId)
		if tag.IsEmpty() {
			return fmt.Errorf("tagId:%d不存在", tagId)
		}
		d.Template[col] = append(d.Template[col], tag)
	}
	return nil
}
