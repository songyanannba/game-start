package template

import (
	"elim5/enum"
	"elim5/logicPack/base"
	"elim5/logicPack/template/flow"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
)

func (s *SpinInfo) PayTableMatch(winLine [][]*base.Tag) (sum float64, flowWinLines []*flow.WinLine, computeFormula string) {
	if len(winLine) <= 0 {
		return
	}
	flowWinLines = s.FlowWinLine(winLine)
	if len(flowWinLines) <= 0 {
		return
	}
	fls := lo.FilterMap(flowWinLines, func(line *flow.WinLine, i int) (float64, bool) {
		return line.Mul, true
	})

	sum = helper.Sum(fls...)
	sum = helper.Mul(sum, s.Counter.LeftCount)

	computeFormula = s.LeftComputeFormula(s.Counter.LeftCount, sum)
	return
}

func (s *SpinInfo) PayTableMatchLine(line *flow.WinLine) {
	if table, ok := s.payTable[line.Name]; ok {
		for _, mul := range table {
			if line.MaxCount >= mul.Num {
				//line.Mul = mul.Mul
				line.TableMul = mul.Mul
			}
		}
		line.Mul = helper.Mul(line.TableMul, float64(line.LineNum))
	}
}

func (s *SpinInfo) LeftComputeFormula(LeftCount int32, sum float64) string {
	return fmt.Sprintf("左边计数 %d * 赔率和 %f: \n", LeftCount, sum)
}

func (s *SpinInfo) RightComputeFormula5(sum float64, counterRightCount float64) string {
	return fmt.Sprintf("五个相连 :左边计数 %d + 右边计数 %f * 单个标签赔率 %f: \n", s.Counter.LeftCount, counterRightCount, sum)
}

func (s *SpinInfo) RightComputeFormula(sum float64, counterRightCount float64) string {
	return fmt.Sprintf(" 右边计数 %f * 单个标签赔率 %f: \n", counterRightCount, sum)
}

func (s *SpinInfo) PayTableMatchRight(rightWinLine [][]*base.Tag, counterRightCount, CounterLiftCount int32) (sum float64, flowWinLines []*flow.WinLine, computeFormula string) {
	if len(rightWinLine) <= 0 {
		return
	}
	flowWinLines = s.FlowWinLine(rightWinLine)
	if len(flowWinLines) <= 0 {
		return
	}

	for _, wLine := range flowWinLines {
		if len(wLine.Tags) == enum.MatchLine5 { //当连线上相同标签达到 5 个，赢钱计算=标签价值*左侧翻倍计数*右侧翻倍计数
			sum += helper.Mul(wLine.Mul, CounterLiftCount+counterRightCount) * 2
			computeFormula += s.RightComputeFormula5(wLine.Mul, float64(CounterLiftCount+counterRightCount))
		} else {
			sum += wLine.Mul * float64(counterRightCount)
			computeFormula += s.RightComputeFormula(wLine.Mul, float64(counterRightCount))
		}
	}
	return
}

func (s *SpinInfo) FreeSpinPayTableMatch(leftEinLine, rightEinLine [][]*base.Tag, counterRightCount, CounterLiftCount int32) (sumL, sumR float64, flowWinLineL, flowWinLinR []*flow.WinLine, computeFormulaStr string) {
	//过滤掉 从左至右五个连续匹配的，防止重复计算
	newLeftEinLine := lo.Filter(leftEinLine, func(item []*base.Tag, index int) bool {
		if len(item) == enum.MatchLine5 {
			return false
		}
		return true
	})

	flowWinLines := make([]*flow.WinLine, 0)
	//左边
	sumL, flowWinLineL, computeFormulaL := s.PayTableMatchLift(newLeftEinLine, CounterLiftCount)
	flowWinLines = append(flowWinLines, flowWinLineL...)
	computeFormulaStr += computeFormulaL
	//右边
	sumR, flowWinLinR, computeFormulaR := s.PayTableMatchRight(rightEinLine, counterRightCount, CounterLiftCount)
	flowWinLines = append(flowWinLines, flowWinLinR...)
	computeFormulaStr += computeFormulaR

	return sumL, sumR, flowWinLineL, flowWinLinR, computeFormulaStr
}

func (s *SpinInfo) PayTableMatchLift(winLine [][]*base.Tag, CounterLiftCount int32) (sum float64, flowWinLines []*flow.WinLine, computeFormula string) {
	if len(winLine) <= 0 {
		return
	}
	flowWinLines = s.FlowWinLine(winLine)
	if len(flowWinLines) <= 0 {
		return
	}
	fls := lo.FilterMap(flowWinLines, func(line *flow.WinLine, i int) (float64, bool) {
		return line.Mul, true
	})

	sum = helper.Sum(fls...)
	sum = helper.Mul(sum, CounterLiftCount)

	computeFormula = s.LeftComputeFormula(CounterLiftCount, sum)
	return
}

func (s *SpinInfo) FlowWinLine(winLine [][]*base.Tag) []*flow.WinLine {
	fWin := make([]*flow.WinLine, 0)
	for _, tags := range winLine {
		for _, table := range s.Config.GetPayTables() {
			if ok, newTable := table.Match(tags); ok {
				line := s.MatchWinLine(newTable.Tags, newTable.Multiple)
				line.Mul = newTable.Multiple
				fWin = append(fWin, line)
				//s.PayTables = append(s.PayTables, newTable)
				break
			}
		}
	}

	return fWin
}

func (s *SpinInfo) MatchWinLine(tags []*base.Tag, mul float64) *flow.WinLine {
	if len(tags) == 0 {
		return &flow.WinLine{
			Tags: make([]*base.Tag, 0),
			Mul:  0,
		}
	}
	return &flow.WinLine{
		Tags: helper.CopyList(tags),
		Mul:  mul,
	}
}

func (s *SpinInfo) FindWinMatch() [][]*base.Tag {

	list := s.DisplayConvertToColumnLeftToRight()
	matchResult := base.MatchSameTagList(list, enum.SameTagLen)
	return matchResult

}

func (s *SpinInfo) DisplayConvertToColumnLeftToRight() [][]*base.Tag {
	//map  列 - 列tag name
	disPlays := helper.CopyListArr(s.Display)
	var tagList [][]*base.Tag
	colDisplay := make(map[int][]*base.Tag)
	for _, disPlay := range disPlays {
		for y, tag := range disPlay {
			colDisplay[y] = append(colDisplay[y], tag)
		}
	}
	for i := 0; i < len(colDisplay); i++ {
		if colTags, ok := colDisplay[i]; ok {
			tagList = append(tagList, colTags)
		}
	}
	return tagList
}

func (s *SpinInfo) FindWinMatchByRightToLeft() [][]*base.Tag {

	list := s.DisplayConvertToColByRightToLeft()
	matchResult := base.MatchSameTagList(list, enum.SameTagLen)
	return matchResult
}

func (s *SpinInfo) DisplayConvertToColByRightToLeft() [][]*base.Tag {
	//map  列 - 列tag name
	disPlays := helper.CopyListArr(s.Display)
	var tagList [][]*base.Tag
	colDisplay := make(map[int][]*base.Tag)
	for _, disPlay := range disPlays {
		for y, tag := range disPlay {
			colDisplay[y] = append(colDisplay[y], tag)
		}
	}
	for i := len(colDisplay) - 1; i >= 0; i-- {
		if colTags, ok := colDisplay[i]; ok {
			tagList = append(tagList, colTags)
		}
	}
	return tagList
}

func (s *SpinInfo) InterTagsChangeWild(turnIntoWoldTag []*base.Tag, count int) {
	if len(turnIntoWoldTag) <= 0 {
		return
	}
	for _, wTag := range turnIntoWoldTag {
		if s.Display[wTag.X][wTag.Y].Name == "" {
			wt := s.Config.GetTagByName(enum.SlotWild).Copy()
			wt.X = wTag.X
			wt.Y = wTag.Y
			s.Display[wTag.X][wTag.Y] = wt

			singleTag := SingleTag{
				Tag:       wt,
				FlowIndex: count, //在那一次产生的wild
				PosOder:   make([]*base.Tag, 0),
			}
			singleTag.PosOder = append(singleTag.PosOder, wt)
			s.SingleTags = append(s.SingleTags, singleTag)
		}
	}
}

type FlowWinLineMerge struct {
	Name   string
	TagMap map[*base.Tag]struct{}
	Win    float64
}

func NewFlowWinLineMerge(name string) *FlowWinLineMerge {
	return &FlowWinLineMerge{
		Name:   name,
		TagMap: map[*base.Tag]struct{}{},
	}
}

func (s *SpinInfo) FlowWinLineMergeSame(fWls []*flow.WinLine) []*flow.WinLine {
	var (
		lineMap = map[string]*FlowWinLineMerge{}
		newLine []*flow.WinLine
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
		lineMap[name] = merge
	}
	// 重新生成winLine
	for name, merge := range lineMap {
		var tags []*base.Tag
		for tag := range merge.TagMap {
			tags = append(tags, tag)
		}
		newLine = append(newLine, &flow.WinLine{
			Name: name,
			Tags: tags,
			Mul:  merge.Win,
		})
	}
	return newLine
}

func (s *SpinInfo) FindWinMatchForTem() [][]*base.Tag {
	matchResult := s.FindWinMatch()
	var newMatchResult [][]*base.Tag
	for _, results := range matchResult {
		isMatch := true
		for _, tag := range results {
			if len(tag.Name) == 0 {
				isMatch = false
				break
			}
		}
		if isMatch {
			newMatchResult = append(newMatchResult, results)
		}
	}
	return newMatchResult

}

func (s *SpinInfo) FindWinMatchForTemRightToLeft() [][]*base.Tag {
	var winLine [][]*base.Tag
	winLineLR := s.FindWinMatch() //赢钱划线 左到右
	if len(winLineLR) > 0 {
		winLine = append(winLine, winLineLR...)
	}
	winLineRL := s.FindWinMatchByRightToLeft() //赢钱划线 右->左
	if len(winLineRL) > 0 {
		winLine = append(winLine, winLineRL...)
	}

	var newMatchResult [][]*base.Tag
	for _, results := range winLine {
		isMatch := true
		for _, tag := range results {
			if len(tag.Name) == 0 {
				isMatch = false
				break
			}
		}
		if isMatch {
			newMatchResult = append(newMatchResult, results)
		}
	}
	return newMatchResult

}
