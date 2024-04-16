package template

import (
	"elim5/global"
	"elim5/logicPack/base"
	"elim5/logicPack/template/flow"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
)

//#region 通用游戏逻辑

// GetTemplateTag 从模版获取Tag
func (s *SpinInfo) GetTemplateTag(x, y int) *base.Tag {
	defer func() {
		helper.PanicRecover()
	}()

	// 获取模版（列）对应的x
	temX := s.GetTemX(y)

	// 获取模版对应的tag
	fillTag := s.GetTemTag(y, temX)

	// 如果模版对应的tag为空,则报错
	if fillTag == nil || fillTag.Name == "" {
		global.GVA_LOG.Error(fmt.Sprintf("x:%d,y:%d,TemX:%d", x, y, temX))
	}
	// 设置tag的坐标
	fillTag.X = x
	fillTag.Y = y

	return s.GetCustomTag(fillTag)
}

// GetTemX 获取模版Y对应的X值并修改
// y 是列的意思; 获取每一列的一个值（相当于x轴坐标） 然后x轴向上移动一步
func (s *SpinInfo) GetTemX(y int) int {
	temX := s.templateRowMap[y]
	if temX < 0 {
		temX = len(s.GetColTemplate(y)) + temX
	}
	if temX < 0 {
		return 0
	}
	s.templateRowMap[y] = temX
	s.templateRowMap[y]--
	return temX
}

// FindAllTags 获取当前显示窗口所有Tag
func (s *SpinInfo) FindAllTags() []*base.Tag {
	return helper.ListToArr(s.Display)
}

// FindAllTagsQuote 获取当前显示窗口所有Tag,直接引用
func (s *SpinInfo) FindAllTagsQuote() []*base.Tag {
	var arr []*base.Tag
	for _, v := range s.Display {
		for _, t := range v {
			arr = append(arr, t)
		}
	}
	return arr
}

// FindTagsByName 查找指定名称的tag
func (s *SpinInfo) FindTagsByName(names ...string) []*base.Tag {
	mapByName := s.GetTagMapByName()
	tags := make([]*base.Tag, 0)
	for _, name := range names {
		tags = append(tags, lo.Filter(mapByName[name], func(tag *base.Tag, index int) bool {
			return s.CustomEffective(tag.X, tag.Y)
		})...)
	}
	return tags
}

func (s *SpinInfo) FindLocationEmpty(list [][2]int) []*base.Tag {
	tags := make([]*base.Tag, 0)
	for _, ints := range list {
		if s.Display[ints[0]][ints[1]].IsEmpty() {
			tags = append(tags, s.Display[ints[0]][ints[1]].Copy())
		}
	}
	return tags
}

//#endregion

//#region 数个数游戏逻辑

// FindCountLine 查找当前窗口满足指定数量的tag
func (s *SpinInfo) FindCountLine(number int) [][]*base.Tag {
	tagsMap := s.GetTagMapByName() //标签名字=>对应的标签个数
	tagList := make([][]*base.Tag, 0)
	for _, tags := range tagsMap {
		if len(tags) >= number {
			tagList = append(tagList, helper.CopyList(tags))
		}
	}
	return tagList
}

func (s *SpinInfo) FindSpecifyCountLine(spTag *base.Tag, number int) []*base.Tag {
	tagsMap := s.GetTagMapByName() //标签名字=>对应的标签个数
	if _, ok := tagsMap[spTag.Name]; ok {
		if len(tagsMap[spTag.Name]) >= number {
			return tagsMap[spTag.Name]
		}
	}
	return []*base.Tag{}
}

//#endregion

//#region 相邻消除游戏逻辑

// FindUDLR 查找上下左右相邻的tag
func (s *SpinInfo) FindUDLR(tag *base.Tag) []*base.Tag {
	var adjacent []*base.Tag
	x := tag.X
	y := tag.Y
	if x > 0 {
		adjacent = append(adjacent, s.Display[x-1][y].Copy())
	}
	if x < s.Config.GetRow()-1 {
		adjacent = append(adjacent, s.Display[x+1][y].Copy())
	}
	if y > 0 {
		adjacent = append(adjacent, s.Display[x][y-1].Copy())
	}
	if y < s.Config.GetCol()-1 {
		adjacent = append(adjacent, s.Display[x][y+1].Copy())
	}
	return adjacent
}

// FindBiasUDLR 查找斜角相邻的tag
func (s *SpinInfo) FindBiasUDLR(tag *base.Tag) []*base.Tag {
	var adjacent []*base.Tag
	x := tag.X
	y := tag.Y
	col := s.Config.GetCol()
	row := s.Config.GetRow()
	if x > 0 && y > 0 {
		adjacent = append(adjacent, s.Display[x-1][y-1].Copy())
	}
	if x > 0 && y < col-1 {
		adjacent = append(adjacent, s.Display[x-1][y+1].Copy())
	}
	if x < row-1 && y > 0 {
		adjacent = append(adjacent, s.Display[x+1][y-1].Copy())
	}
	if x < row-1 && y < col-1 {
		adjacent = append(adjacent, s.Display[x+1][y+1].Copy())
	}

	return adjacent
}

// FindAllUDLR  获取附近8个方向的tag
func (s *SpinInfo) FindAllUDLR(tag *base.Tag) []*base.Tag {
	var adjacent []*base.Tag
	adjacent = append(adjacent, s.FindUDLR(tag)...)
	adjacent = append(adjacent, s.FindBiasUDLR(tag)...)
	return adjacent
}

// FindAdjacentLine 查找相邻划线
func (s *SpinInfo) FindAdjacentLine(number int) [][]*base.Tag {
	tagList := make([][]*base.Tag, 0)
	v := NewVerify()
	for _, tags := range s.Display {
		for _, tag := range tags {
			v.Restart()
			v.ResetVerify(s)
			if v.SetSite(tag) {
				s.FindErgodic(v)
			} else {
				continue
			}
			tagLine := v.GetSites()
			if len(tagLine) >= number {
				tagList = append(tagList, tagLine)
			}
		}
	}
	return tagList
}

// FindSpecifyLine 查找指定位置相邻划线
func (s *SpinInfo) FindSpecifyLine(spTag *base.Tag, number int) [][]*base.Tag {
	tagList := make([][]*base.Tag, 0)
	v := NewVerify()
	v.Restart()
	if v.SetSite(spTag) {
		s.FindErgodic(v)
	}
	tagLine := v.GetSites()
	if len(tagLine) >= number {
		tagList = append(tagList, tagLine)
	}
	return tagList
}

// FindErgodic 遍历所有相邻
func (s *SpinInfo) FindErgodic(v *Verify) {
	tags := s.FindUDLR(v.site)
	for _, tag := range tags {
		if v.SetSite(tag) {
			s.FindErgodic(v)
		}
	}
}

//#endregion

// #region slot消除掉落逻辑

// FindSlotLine 划线规则
func (s *SpinInfo) FindSlotLine() []*flow.WinLine {
	winline := make([]*flow.WinLine, 0)
	resData := make([][]*base.Tag, 0)
	for _, coords := range s.Config.GetCoords() {
		var tags []*base.Tag
		for _, coord := range coords {
			// 从初始数据中获取结果Tag
			tag := s.Display[coord.Y][coord.X].Copy()
			tag.IsLine = true
			tags = append(tags, tag)
		}
		resData = append(resData, tags)
	}

	for _, tags := range resData {
		for _, table := range s.Config.GetPayTables() {
			if ok, newTable := table.Match(tags); ok {
				winline = append(winline, &flow.WinLine{
					Tags: newTable.Tags,
					Mul:  newTable.Multiple,
				})
				break
			}
		}
	}
	return winline
}

// FindSlotAllLine 全线划线规则
func (s *SpinInfo) FindSlotAllLine(length int, fa func(tags []*base.Tag) []*base.Tag) []*flow.WinLine {
	mathMap := helper.ArrToMapListValue(fa(s.FindAllTags()), func(item *base.Tag) (vs []string, ok bool) {
		if !s.CustomEffective(item.X, item.Y) {
			return nil, false
		}
		return item.Match(), true
	})
	lines := make([]*flow.WinLine, 0)
	for name, tags := range mathMap {
		colMap := lo.GroupBy(tags, func(tag *base.Tag) int {
			return tag.Y
		})
		if len(colMap) < length {
			continue
		}
		line := flow.WinLine{
			Tags:     make([]*base.Tag, 0),
			Name:     name,
			MaxCount: 0,
			LineNum:  1,
		}
		for i := 0; i < s.Config.GetCol(); i++ {
			if len(colMap[i]) == 0 {
				break
			}
			line.Tags = append(line.Tags, colMap[i]...)
			line.MaxCount++
			line.LineNum = int(helper.MulToInt(line.LineNum, len(colMap[i])))
		}
		if line.MaxCount >= length {
			s.PayTableMatchLine(&line)
			lines = append(lines, &line)
		}
	}
	return lines
}

func (s *SpinInfo) FindMegaWayAllLine(length int, f func(tags []*base.Tag) []*base.Tag) []*flow.WinLine {
	mathMap := helper.ArrToMapListValue(f(s.FindAllTags()), func(item *base.Tag) (vs []string, ok bool) {
		return item.Match(), true
	})
	lines := make([]*flow.WinLine, 0)
	for name, tags := range mathMap {
		colMap := lo.GroupBy(tags, func(tag *base.Tag) int {
			if tag.Y == s.Config.GetCol()-1 {
				return s.Config.GetRow() - tag.X
			} else {
				return tag.Y
			}
		})
		if len(colMap) < length {
			continue
		}
		line := flow.WinLine{
			Tags:     make([]*base.Tag, 0),
			Name:     name,
			MaxCount: 0,
			LineNum:  1,
		}
		for i := 0; i < s.Config.GetCol()-1; i++ {
			if len(colMap[i]) == 0 {
				break
			}
			line.Tags = append(line.Tags, colMap[i]...)
			line.MaxCount++
			line.LineNum = int(helper.MulToInt(line.LineNum, len(colMap[i])))
		}
		if line.MaxCount >= length {
			s.PayTableMatchLine(&line)
			lines = append(lines, &line)
		}
	}
	return lines
}

func (s *SpinInfo) FindMegaWayAllLineDisplacement(length int, f func(tags []*base.Tag) []*base.Tag) []*flow.WinLine {
	mathMap := helper.ArrToMapListValue(f(s.FindAllTags()), func(item *base.Tag) (vs []string, ok bool) {
		return item.Match(), true
	})
	lines := make([]*flow.WinLine, 0)
	for name, tags := range mathMap {
		colMap := lo.GroupBy(tags, func(tag *base.Tag) int {
			if tag.Y == s.Config.GetCol()-1 {
				return s.Config.GetRow() - (tag.X - 1)
			} else {
				return tag.Y
			}
		})
		if len(colMap) < length {
			continue
		}
		line := flow.WinLine{
			Tags:     make([]*base.Tag, 0),
			Name:     name,
			MaxCount: 0,
			LineNum:  1,
		}
		for i := 0; i < s.Config.GetCol()-1; i++ {
			if len(colMap[i]) == 0 {
				break
			}
			line.Tags = append(line.Tags, colMap[i]...)
			line.MaxCount++
			line.LineNum = int(helper.MulToInt(line.LineNum, len(colMap[i])))
		}
		if line.MaxCount >= length {
			s.PayTableMatchLine(&line)
			lines = append(lines, &line)
		}
	}
	return lines
}

//#endregion

// FindMahjongAllLine 麻将类型划线
