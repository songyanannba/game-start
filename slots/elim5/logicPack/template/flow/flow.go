package flow

import (
	"elim5/enum"
	"elim5/global"
	"elim5/logicPack/base"
	"elim5/logicPack/template/slot_skill"
	"elim5/pbs/common"
	"elim5/pbs/game"
	"fmt"
	"github.com/samber/lo"
	"strconv"

	"elim5/utils/helper"
)

type SpinFlow struct {
	Id        int
	InitList  [][]*base.Tag // 初始列表
	AlterList []*WinLine    // 改变列表
	OmitList  []*WinLine    // 消除列表
	AddList   []*base.Tag   // 填充列表
	WinList   []*WinLine    // 赢钱列表
	FlowMap   string        // 流程图
	Gain      int64         // 获得
	SumMul    float64       // 总倍数

	EmitList   [][]*base.Tag // 发射列表
	Level      int           // 等级 (21 // 1 翻倍 2 收集 3 转换)
	Multiplier *WinLine      // multiplier数量
	Points     int
	InitSkills []slot_skill.Skill[[2]int]
	Rank       int
	ExtraMul   int

	//机台特殊需要字段
	SFlowExt
	MatchTable //slot9再用
	MateSlot21
	MateSlot50
}

// GetPoints 获取总积分
func (s *SpinFlow) GetPoints() int {
	return lo.SumBy(s.OmitList, func(line *WinLine) int {
		return line.Integral
	})
}

func NewSpinFlow(id int) SpinFlow {
	spinFlow := SpinFlow{
		Id:       id,
		InitList: make([][]*base.Tag, 0),
		OmitList: make([]*WinLine, 0),
		AddList:  make([]*base.Tag, 0),
		Gain:     0,
		SumMul:   0,
		SFlowExt: SFlowExt{
			SkillCenter: make([]*base.Tag, 0),
		},
		Multiplier: &WinLine{
			Tags: make([]*base.Tag, 0),
			Mul:  0,
		},
	}
	spinFlow.MatchTable = NewMatchTable()
	return spinFlow
}

func (s *SpinFlow) AddOmitList(lines ...*WinLine) {
	for _, line := range lines {
		if line == nil {
			continue
		}
		s.OmitList = append(s.OmitList, line)
	}
	sum := lo.SumBy(lines, func(line *WinLine) float64 {
		return line.Mul
	})
	s.SumMul += sum
}

type View struct {
	Data string `json:"data"`
	Text string `json:"text"`
}

type Label struct {
	Name        string `json:"name"`
	Mul         int    `json:"mul"`
	Colour      string `json:"colour"` //Default Primary Success Into Warning Danger
	Hide        string `json:"hide"`
	BorderColor string `json:"borderColor"`
	Life        int    `json:"life"`
	Visible     bool   `json:"visible"` //是否v
}

const (
	BorderColorGold   = "#ffd700"
	BorderColorSilver = "#C0C0C0"
)

func (s *SpinFlow) String() string {
	list := make([][]*Label, 0)
	for _, tags := range s.InitList {
		rows := make([]*Label, 0)
		for _, tag := range tags {
			isRemove := s.GetIsRemove(tag.X, tag.Y)
			Mul := s.GetRemoveMul(tag.X, tag.Y)
			IsEmit := s.GetIsEmit(tag.X, tag.Y)
			//skill := s.GetInitSkill(tag.X, tag.Y)
			label := &Label{
				Name:   tag.Name,
				Colour: enum.VueTypeInfo,
				Hide:   helper.If(tag.IsMegastack, "M", ""),
			}
			slotIdReds := helper.SplitStr(tag.Redundancy, "-")
			slotId, err := strconv.Atoi(slotIdReds[0])
			if err == nil {
				switch slotId {
				case 7, 16, 22:
					label.Hide = slotIdReds[1]
				}
			}
			if isRemove {
				label.Colour = enum.VueTypePrimary
			}
			if IsEmit {
				label.Colour = enum.VueTypeSuccess
			}
			if tag.IsWild {
				label.Colour = enum.VueTypeDanger
			}
			if Mul > 1 || tag.Multiple > 1 {
				label.Name = fmt.Sprintf("%s*%d", label.Name, int(tag.Multiple))
			}
			if tag.Gold {
				label.BorderColor = BorderColorGold
			}

			rows = append(rows, label)
		}
		list = append(list, rows)
	}
	detail, err := global.Json.MarshalToString(list)
	if err != nil {
		return ""
	}
	return detail
}

func (s *SpinFlow) QueryElimination(x, y int) (int, int, int) {
	var a1, a2, a3 int
	a2 = int(s.InitList[x][y].Multiple)
	for _, eliminate := range s.OmitList {
		for _, tag := range eliminate.Tags {
			if tag.X == x && tag.Y == y {
				a1 = 1
				if tag.Multiple >= 2 {
					a2 = int(tag.Multiple)
				}
				break
			}
		}
		if a1 != 0 {
			break
		}
	}
	for _, tags := range s.EmitList {
		for _, tag := range tags {
			if tag.X == x && tag.Y == y {
				a3 = 1
				break
			}
		}
	}

	return a1, a2, a3
}

func (s *SpinFlow) GetInitSkill(x, y int) int {
	for _, skill := range s.InitSkills {
		if skill.SkillInfo == [2]int{x, y} {
			return skill.Id
		}
	}
	return 0
}

func (s *SpinFlow) GetIsRemove(x, y int) bool {
	for _, eliminate := range s.OmitList {
		for _, tag := range eliminate.Tags {
			if tag.X == x && tag.Y == y {
				return true
			}
		}
	}
	return false
}

func (s *SpinFlow) GetRemoveMul(x, y int) int {
	for _, eliminate := range s.OmitList {
		for _, tag := range eliminate.Tags {
			if tag.X == x && tag.Y == y {
				return int(tag.Multiple)
			}
		}
	}
	return 0
}

func (s *SpinFlow) GetIsEmit(x, y int) bool {
	for _, tags := range s.EmitList {
		for _, tag := range tags {
			if tag.X == x && tag.Y == y {
				return true
			}
		}
	}
	for _, line := range s.AlterList {
		for _, tag := range line.Tags {
			if tag.X == x && tag.Y == y {
				if tag.X == x && tag.Y == y {
					return true
				}
			}
		}
	}
	return false
}

func (s *SpinFlow) GetInformation() string {
	str := ""
	str += fmt.Sprintf("编号:%d, 赢钱:%g, 倍率:%g \n", s.Id, s.Gain, s.SumMul)
	for _, eliminate := range s.OmitList {
		str += fmt.Sprintf("消除 %d 个 %s 倍率: %g \n", len(eliminate.Tags), eliminate.Tags[0].Name, eliminate.Mul)
	}
	for _, add := range s.AddList {
		str += fmt.Sprintf("增加 %s  %d:%d \n", add.Name, add.X, add.Y)
	}
	str += fmt.Sprintf("扩散位置 %v 个 \n", s.EmitList)
	str += fmt.Sprintf("特殊倍率 %v  \n", helper.ArrConversion(s.Multiplier.Tags, func(t *base.Tag) (int32, bool) {
		return int32(t.Multiple), true
	}))
	return str
}

func (s *SpinFlow) SaveMatchTable(leftWinLine, rightWinLine [][]*base.Tag) {
	s.LeftResList = leftWinLine
	s.RightResList = rightWinLine

	var winLine [][]*base.Tag
	winLine = append(winLine, leftWinLine...)
	winLine = append(winLine, rightWinLine...)

	recordTagSite := map[[2]int]struct{}{} //被记录过的 将变成wild 的 tag
	for _, tags := range winLine {
		intoWild := tags[2]
		if _, ok := recordTagSite[[2]int{intoWild.X, intoWild.Y}]; ok {
			continue
		}
		recordTagSite[[2]int{intoWild.X, intoWild.Y}] = struct{}{}
		s.TurnIntoWoldTag = append(s.TurnIntoWoldTag, intoWild)
	}
}

func (s *SpinFlow) FindCanTurnWildLine(winLine [][]*base.Tag) [][]*base.Tag {

	var newWinLine [][]*base.Tag
	for _, wTags := range winLine {
		//先找到在同一行
		var xLineMap map[int][]*base.Tag
		xLineMap = make(map[int][]*base.Tag)
		for _, tag := range wTags {
			if tag.IsWild {
				continue
			}
			xLineMap[tag.X] = append(xLineMap[tag.X], tag)
		}

		//xLineMap 每行 相同的标签
		for _, tags := range xLineMap {
			if len(tags) < 3 {
				continue
			}
			//从左到右匹配
			var yAdjoin1 []*base.Tag
			for i := 0; i < len(tags); i++ {
				if i == 0 {
					yAdjoin1 = append(yAdjoin1, tags[i])
					continue
				}
				if tags[i].Y-1 == tags[i-1].Y || tags[i].Y+1 == tags[i-1].Y {
					yAdjoin1 = append(yAdjoin1, tags[i])
				} else {
					if len(yAdjoin1) >= 3 {
						newWinLine = append(newWinLine, yAdjoin1)
					}
					yAdjoin1 = make([]*base.Tag, 0)
					yAdjoin1 = append(yAdjoin1, tags[i])
				}
				//
				if len(yAdjoin1) < 3 {
					continue
				}
				newWinLine = append(newWinLine, yAdjoin1)
			}
		}
	}

	return newWinLine
}

func (s *SpinFlow) TurnIntoWildTags(winLine [][]*base.Tag) {
	//s.LeftResList = winLine
	//for _, wls := range winLine { 测试
	//	for _, val := range wls {
	//		val.X = wls[0].X
	//	}
	//}

	s.TurnIntoWoldTag = make([]*base.Tag, 0)
	recordTagSite := map[[2]int]struct{}{} //被记录过的 将变成wild 的 tag

	lines := s.FindCanTurnWildLine(winLine) //再次过滤 找到可以转换wild的划线

	for _, line := range lines {
		tags := helper.CopyList(line)
		turnWildTags := tags[1 : len(tags)-1]
		for _, intoWild := range turnWildTags {
			if intoWild.IsWild {
				continue
			}
			if _, ok := recordTagSite[[2]int{intoWild.X, intoWild.Y}]; ok {
				continue
			}
			recordTagSite[[2]int{intoWild.X, intoWild.Y}] = struct{}{}
			s.TurnIntoWoldTag = append(s.TurnIntoWoldTag, intoWild)
		}
	}

	//global.GVA_LOG.Info("TurnIntoWildTags..")
}

// SaveMatchTableLeftV1 仅仅把 中间(最中间一列:2)的标签变成wild
func (s *SpinFlow) SaveMatchTableLeftV1(winLine [][]*base.Tag) {
	s.LeftResList = winLine

	recordTagSite := map[[2]int]struct{}{} //被记录过的 将变成wild 的 tag
	for _, tags := range winLine {
		intoWild := tags[2]
		if _, ok := recordTagSite[[2]int{intoWild.X, intoWild.Y}]; ok {
			continue
		}
		recordTagSite[[2]int{intoWild.X, intoWild.Y}] = struct{}{}
		s.TurnIntoWoldTag = append(s.TurnIntoWoldTag, intoWild)
	}
}

func (s *SpinFlow) PayTableMatch(PayTableList []*base.PayTable) (sum float64) {
	fWin := make([]*WinLine, 0)
	for _, tags := range s.MatchTable.LeftResList {
		for _, table := range PayTableList {
			if ok, newTable := table.Match(tags); ok {
				//line := s.GetWinLine(newTable.Tags)
				line := s.MatchWinLine(newTable.Tags, newTable.Multiple)
				//line.Mul = newTable.Multiple
				fWin = append(fWin, line)
				s.PayTables = append(s.PayTables, newTable)
				s.OmitList = append(s.OmitList, &WinLine{
					Tags: newTable.Tags,
					Mul:  newTable.Multiple,
				})
				break
			}
		}
	}
	if len(fWin) <= 0 {
		return
	}
	fls := lo.FilterMap(fWin, func(line *WinLine, i int) (float64, bool) {
		return line.Mul, true
	})
	sum = lo.SumBy(fls, func(item float64) float64 {
		return item
	})
	return
}

func (s *SpinFlow) MatchWinLine(tags []*base.Tag, mul float64) *WinLine {
	if len(tags) == 0 {
		return &WinLine{
			Tags: make([]*base.Tag, 0),
			Mul:  0,
		}
	}
	return &WinLine{
		Tags: helper.CopyList(tags),
		Mul:  mul,
	}
}

func (s *SpinFlow) ToAck(bet int) (flowAck *common.StepFlow) {
	flowAck = &common.StepFlow{
		Index:      int64(s.Id),
		Gain:       s.Gain,
		RemoveList: s.GetRemoveList(bet),
		AddList: &common.Tags{
			Tags: s.GetAddList(),
		},
		Points:      int32(s.GetPoints()),
		RemoveCause: int32(s.IsUserSkill),
		SpecialMul:  s.GetSpecialMul(),
		SkillCenter: s.getSkillCenter(), // 选择的中心点
		AlterList:   s.GetAlterList(),
		Accumulate:  s.Accumulate(),
	}

	return
}

func (s *SpinFlow) GetSpecialMul() []int32 {
	return helper.ArrConversion(s.Multiplier.Tags, func(t *base.Tag) (int32, bool) {
		return int32(t.Multiple), true
	})
}

func (s *SpinFlow) GetRemoveList(bet int) []*common.Tags {
	return helper.ArrConversion(s.OmitList, func(w *WinLine) (*common.Tags, bool) {
		return w.ToAck(bet), true
	})
}

func (s *SpinFlow) GetAlterList() []*common.Tags {
	return helper.ArrConversion(s.AlterList, func(w *WinLine) (*common.Tags, bool) {
		return w.ToAck(0), true
	})
}

func (s *SpinFlow) GetAddList() []*common.Tag {
	return helper.ArrConversion(s.AddList, func(t *base.Tag) (*common.Tag, bool) {
		return t.ToTagAck(), true
	})
}

func (s *SpinFlow) getSkillCenter() []*common.Pos {
	var pos []*common.Pos
	for _, sCenters := range s.SkillCenter {
		cPos := &common.Pos{
			X: int32(sCenters.X),
			Y: int32(sCenters.Y),
		}
		pos = append(pos, cPos)
	}
	return pos
}

func (s *SpinFlow) Accumulate() *common.Accumulate {
	ac := &common.Accumulate{
		General: 0,
		Left:    int32(s.AccumulateLeft()),
		Right:   int32(s.AccumulateRight()),
		Points:  int32(s.GetPoints()),
	}
	return ac
}

func (s *SpinFlow) AccumulateLeft() int {
	if s.MatchTable.LeftResList != nil {
		if len(s.MatchTable.LeftResList) > 0 {
			return 1
		}
	}
	return 0
}

func (s *SpinFlow) AccumulateRight() int {
	if s.MatchTable.RightResList != nil {
		if len(s.MatchTable.RightResList) > 0 {
			return 1
		}
	}
	return 0
}

func (s *SpinFlow) WinLineOrientation(flow []*WinLine, orientation int) {
	for _, omi := range flow {
		omi.Orientation = orientation
	}
}

type WinLineMerge struct {
	Name        string
	TagMap      map[*base.Tag]struct{}
	Win         float64
	Orientation int
}

func NewFlowWinLineMerge(name string) *WinLineMerge {
	return &WinLineMerge{
		Name:   name,
		TagMap: map[*base.Tag]struct{}{},
	}
}

func (s *SpinFlow) FlowWinLineMergeSame(fWls []*WinLine, count int32, side int) []*WinLine {
	var (
		lineMap = map[string]*WinLineMerge{}
		newLine []*WinLine
	)
	// 合并相同name的线
	for _, line := range fWls {
		// 获取首个不是wild的tag
		name := base.GetTagsNameByFunc(line.Tags, func(tag *base.Tag) bool {
			return !tag.IsWild
		})
		line.Name = name
		// 使用这个tag的name做为key
		merge, ok := lineMap[name]
		if !ok {
			lineMap[name] = NewFlowWinLineMerge(name)
			merge = lineMap[name]
		}
		for _, tag := range line.Tags {
			merge.TagMap[tag] = struct{}{}
		}
		merge.Win += line.Mul
		merge.Orientation = line.Orientation
		lineMap[name] = merge
	}
	// 重新生成winLine
	for name, merge := range lineMap {
		var tags []*base.Tag
		for tag := range merge.TagMap {
			tags = append(tags, tag)
		}
		mul := helper.Mul(merge.Win, count)
		newLine = append(newLine, &WinLine{
			Name:        name,
			Tags:        tags,
			Mul:         helper.Mul(mul, side),
			Orientation: merge.Orientation,
		})
	}
	return newLine
}

func (s *SpinFlow) GetInfo(i int) string {
	str := fmt.Sprintf("\n--------------------Flow%d Gain: %d---------------------------\n", i, s.Gain)
	var (
		rel  string
		addl string
		altl string
	)
	str += fmt.Sprintf("特殊翻倍:%v\n", s.ExtraMul)

	if len(s.AlterList) > 0 {
		altl = "alterList:\n"

		for _, line := range s.AlterList {
			groups := lo.GroupBy(line.Tags, func(item *base.Tag) string {
				return item.Name
			})
			for i2, tags := range groups {
				altl += fmt.Sprintf("%d个%v\n", len(tags), i2)
			}

		}
	}
	//fmt.Println(s)
	if len(s.OmitList) > 0 {
		rel = "removeList:"
		for _, tags := range s.OmitList {
			if len(tags.Tags) == 0 {
				continue
			}
			//fmt.Println(flow)
			rel += fmt.Sprintf("%d个%v len:%d payMul:%g count:%d sumMul:%g\n", len(tags.Tags), tags.Name, tags.MaxCount, tags.TableMul, tags.LineNum, tags.Mul)
			//rel += fmt.Sprintf("\n%d个%v ", len(tags.Tags), base.GetTagsName(tags.Tags, enum.WildName))
			//rel += fmt.Sprintf("赢钱:%g\n", tags.Mul)
			mulTags := lo.Filter(tags.Tags, func(item *base.Tag, index int) bool {
				return item.Multiple > 1
			})
			if len(mulTags) > 0 {
				rel += fmt.Sprintf("翻倍标签:")
				for _, tag := range tags.Tags {
					rel += fmt.Sprintf("%v", helper.If(tag.Multiple > 1, fmt.Sprintf("%g,", tag.Multiple), ""))
				}
			}
		}
	}
	if len(s.AddList) > 0 {
		addl = "addlist:\n"
		//slices.SortFunc(s.AddList, func(i, j *base.Tag) bool {
		//	return i.Y < j.Y || (i.Y == j.Y && i.X < j.X)
		//})
		for _, tag := range s.AddList {
			addl += fmt.Sprintf("%v{%d:%d%v} ", tag.Name, tag.X, tag.Y, helper.If(tag.Multiple > 1, fmt.Sprintf("*%g", tag.Multiple), ""))
		}
	}

	str += fmt.Sprintf("\n%v\n%v\n%v\n", rel, addl, altl)
	return str
}

func (s *SpinFlow) GetInfoSkillCenter() string {
	str := ""
	if len(s.SFlowExt.SkillCenter) > 0 {
		str = "SkillCenter:\n"
		str += fmt.Sprintf("%d个%v\n", len(s.SFlowExt.SkillCenter), s.SFlowExt.SkillCenter)
	}

	str += "\n level == " + helper.Itoa(s.Level)

	str += "\n level == " + s.FlowMap
	return str
}

func (s *SpinFlow) GetBaseStepFlow() *game.BaseStepFlow {
	return &game.BaseStepFlow{
		Index: int32(s.Id),
		Gain:  s.Gain,
	}
}

func (s *SpinFlow) GetInfo50() string {
	str := ""
	if len(s.MateSlot50.MattsTag) > 0 {
		str = "田字格 MattsTag (上面消除的标签倍率已经 *2):\n"
		str += fmt.Sprintf("%d个%v\n", len(s.MateSlot50.MattsTag), s.MateSlot50.MattsTag)
	}

	if len(s.MateSlot50.AddWild) > 0 {
		str = "掉落之前的 wild填充 信息:\n"
		str += fmt.Sprintf("%d个%v\n", len(s.MateSlot50.AddWild), s.MateSlot50.AddWild)
	}

	str += "\n 第几次转 == " + helper.Itoa(s.MateSlot50.Num)

	return str
}

func (s *SpinFlow) GetInfo49() string {
	var str string

	if len(s.SFlowExt.ChangeTags) > 0 {
		patternStr := ""
		if s.SFlowExt.Pattern == 1 {
			patternStr = "土龙模式"
		} else if s.SFlowExt.Pattern == 2 {
			patternStr = "水龙模式"
		} else if s.SFlowExt.Pattern == 3 {
			patternStr = "火龙模式"
		} else if s.SFlowExt.Pattern == 4 {
			patternStr = "巨龙模式"
		}

		str = "\n"
		str += fmt.Sprintf("%s模式:%d个%v\n", patternStr, s.SFlowExt.Pattern, s.SFlowExt.ChangeTags)
	}

	return str
}
