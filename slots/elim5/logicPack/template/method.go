package template

import (
	"elim5/enum"
	"elim5/logicPack/base"
	"elim5/logicPack/template/flow"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
	"strconv"
)

// GetTagMapByName 获取当前显示窗口所有Tag名称的Map ;display结构中: 标签名字=>对应的标签个数(>=8)
func (s *SpinInfo) GetTagMapByName() map[string][]*base.Tag {
	tags := s.FindAllTags() //Display
	return lo.GroupBy(tags, func(item *base.Tag) string {
		return item.Name
	})
}

// GetEmptyTags 获取窗口空白标签
func (s *SpinInfo) GetEmptyTags() []*base.Tag {
	return lo.Filter(s.FindAllTags(), func(tag *base.Tag, i int) bool {
		return tag.IsEmpty()
	})
}

// SetAllLocation 重新设置所有标签的位置
func (s *SpinInfo) SetAllLocation() {
	for x, tags := range s.Display {
		for y, tag := range tags {
			tag.X = x
			tag.Y = y
		}
	}
}

// PrintTable 打印table
func (s *SpinInfo) PrintTable(str string) string {
	str += ":\n"
	for _, row := range s.Display {
		for _, col := range row {
			str += fmt.Sprintf("%s\t", strconv.Itoa(col.X)+":"+strconv.Itoa(col.Y)+" "+helper.If(col.Name == "", "   🀆", col.Name))
		}
		str += "\r\n"
	}
	//fmt.Println(str)
	return str + "\r\n"
}

// GetLayout 获取窗口布局
func (s *SpinInfo) GetLayout() (row, col int) {
	return s.Config.GetRow(), s.Config.GetCol()
}

// GetNormalTags 获取普通tag集合s
func (s *SpinInfo) GetNormalTags() []*base.Tag {
	return lo.Filter(s.Config.GetAllTags(), func(tag *base.Tag, i int) bool {
		return tag.Name != "scatter" && tag.Name != "multiplier" && tag.IsWild == false && tag.Name != enum.MegastackName && tag.Name != enum.GoldName && tag.Name != "dragon"
	})
}

func (s *SpinInfo) GetRandNormalTags(x, y int) *base.Tag {
	tags := s.GetNormalTags()
	fillTag := tags[helper.RandInt(len(tags))].Copy()
	fillTag.X = x
	fillTag.Y = y
	return fillTag
}

// GetColTemplate  获取模版
func (s *SpinInfo) GetColTemplate(col int) []uint16 {
	tem, err := s.Config.GetTemplate(s.infoType, s.Which)
	if err != nil {
		return []uint16{}
	}
	return tem[col]
}

func (s *SpinInfo) GetTemTag(col, index int) *base.Tag {
	tem := s.GetColTemplate(col)
	if len(tem) == 0 || index >= len(tem) {
		return &base.Tag{}
	}
	tag := s.Config.GetTagById(int(tem[index]))
	return tag.Copy()
}

// ResetIndex   获取模版
func (s *SpinInfo) ResetIndex() {
	for i, tags := range s.Display {
		for i2, _ := range tags {
			s.Display[i][i2].X = i
			s.Display[i][i2].Y = i2
		}
	}
}

// SetIndexMap  设置模版起始点
func (s *SpinInfo) SetIndexMap() error {
	intRow, err := s.Config.GetInitTemIndex(s.infoType, s.Which)
	if err != nil {
		return err
	}
	col := s.Config.GetCol()
	s.temInit = make([]int32, 0)
	for y := 0; y < col; y++ {
		s.templateRowMap[y] = intRow
		s.temInit = append(s.temInit, int32(intRow))
	}
	return nil
}

// GetCustomTag  获取tag自定义处理
func (s *SpinInfo) GetCustomTag(tag *base.Tag) *base.Tag {
	if _, ok := s.CustomTag[tag.Name]; ok {
		return s.CustomTag[tag.Name](tag, s).Copy()
	} else {
		return tag
	}
}

// SetDisparityUint16 将Id排布转换为标签
func (s *SpinInfo) SetDisparityUint16(list [][]uint16) {
	s.Display = helper.ListConversion(list, func(item uint16) *base.Tag {
		return s.Config.GetTagById(int(item)).Copy()
	})
	s.SetAllLocation()
}

// GetTagPayTable 查找指定标签名字的赔付表
func (s *SpinInfo) GetTagPayTable(tagName string) []*NumMul {
	return s.payTable[tagName]
}

func (s *SpinInfo) GetCountTagPayTable(tagName string, count int) float64 {
	mul := 0.0
	for _, i2 := range s.payTable[tagName] {
		if i2.Num <= count {
			mul = i2.Mul
		}
	}
	return mul
}

// GetRandPosition 随机获取一个位置
func (s *SpinInfo) GetRandPosition() *base.Tag {
	tags := s.FindAllTags()
	return tags[helper.RandInt(len(tags))].Copy()
}

// GetPositionDir 获取指定方向的位置
func (s *SpinInfo) GetPositionDir(dir int, pos *base.Tag) []*base.Tag {
	switch dir {
	case enum.SiteUP:
		return s.GetPositionUp(pos)
	case enum.SiteDown:
		return s.GetPositionDown(pos)
	case enum.SiteLeft:
		return s.GetPositionLeft(pos)
	case enum.SiteRight:
		return s.GetPositionRight(pos)
	default:
		return []*base.Tag{}
	}
}

// GetPositionUp 获取上方位置
func (s *SpinInfo) GetPositionUp(pos *base.Tag) []*base.Tag {
	tags := s.FindAllTags()
	return helper.CopyList(lo.Filter(tags, func(tag *base.Tag, i int) bool {
		return tag.Y == pos.Y && tag.X < pos.X
	}))
}

// GetPositionDown 获取下方位置
func (s *SpinInfo) GetPositionDown(pos *base.Tag) []*base.Tag {
	tags := s.FindAllTags()
	return helper.CopyList(lo.Filter(tags, func(tag *base.Tag, i int) bool {
		return tag.Y == pos.Y && tag.X > pos.X
	}))
}

// GetPositionLeft 获取左方位置
func (s *SpinInfo) GetPositionLeft(pos *base.Tag) []*base.Tag {
	tags := s.FindAllTags()
	return helper.CopyList(lo.Filter(tags, func(tag *base.Tag, i int) bool {
		return tag.X == pos.X && tag.Y < pos.Y
	}))
}

// GetPositionRight 获取右方位置
func (s *SpinInfo) GetPositionRight(pos *base.Tag) []*base.Tag {
	tags := s.FindAllTags()
	return helper.CopyList(lo.Filter(tags, func(tag *base.Tag, i int) bool {
		return tag.X == pos.X && tag.Y > pos.Y
	}))
}

// GetSumRemoveCount 获取消除标签个数
func (s *SpinInfo) GetSumRemoveCount() int {
	return lo.SumBy(s.SpinFlow, func(i flow.SpinFlow) int {
		return lo.SumBy(i.OmitList, func(item *flow.WinLine) int {
			return len(item.Tags)
		})
	})
}
