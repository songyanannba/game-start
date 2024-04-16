package classify

import (
	"elim5/enum"
	"elim5/logicPack/base"
	"elim5/logicPack/template"
	"elim5/logicPack/template/flow"
	businessReq "elim5/model/business/request"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
	"strconv"
)

// MahjongEliminate Megaway消除类型 类型25台
type MahjongEliminate struct {
	req      *businessReq.MergerIncrease
	spinInfo *template.SpinInfo
}

func NewMahjongEliminate(req *businessReq.MergerIncrease, spinInfo *template.SpinInfo) *MahjongEliminate {
	return &MahjongEliminate{
		req:      req,
		spinInfo: spinInfo,
	}
}

func (n MahjongEliminate) Merger(index int) error {
	var lines *flow.WinLine
	fmt.Println(n.spinInfo.PrintTable("index:" + strconv.Itoa(index) + " start"))
	alterTags := make([]*base.Tag, 0)
	for i := 0; i < 1000; i++ {
		lines = FindMahjongAllLine(n.spinInfo, n.req.TriggerLineNum)
		scaLine := n.spinInfo.FindTagsByName(enum.ScatterName)
		if len(lines.Tags) < n.req.TriggerLineNum && len(scaLine) < n.req.TriggerSca {
			fmt.Printf("AlterTags:%v\n", alterTags)
			fmt.Println(n.spinInfo.PrintTable("index:" + strconv.Itoa(index) + " end"))
			return nil
		}
		//tagsList := make([][]*base.Tag, 0)
		//for _, line := range lines {
		//	tagsList = append(tagsList, line.Tags)
		//}
		tags := helper.CopyList(lines.Tags)
		if index != 0 {
			tags = lo.Filter(tags, func(item *base.Tag, index int) bool {
				return item.X == n.spinInfo.Config.GetRow()-1
			})

			scaLine = lo.Filter(scaLine, func(item *base.Tag, index int) bool {
				return item.X == n.spinInfo.Config.GetRow()-1
			})
		}

		if len(tags) > 0 {
			randTag := tags[helper.RandInt(len(tags))]
			n.spinInfo.Display[randTag.X][randTag.Y] = n.spinInfo.GetRandNormalTags(randTag.X, randTag.Y).Copy()
			alterTags = append(alterTags, randTag.Copy())
		}

		if len(scaLine) > 0 {
			randSca := scaLine[helper.RandInt(len(scaLine))]
			alterTags = append(alterTags, randSca.Copy())
			n.spinInfo.Display[randSca.X][randSca.Y] = n.spinInfo.GetRandNormalTags(randSca.X, randSca.Y).Copy()
		}

	}
	return fmt.Errorf("normal eliminate error index:%d \n%s", index, n.spinInfo.PrintTable("normal eliminate"))
}

func FindMahjongAllLine(s *template.SpinInfo, length int) *flow.WinLine {
	var positionMap = make(map[[2]int]bool)
	startingPoint := lo.Filter(s.FindAllTags(), func(tag *base.Tag, index int) bool {
		return tag.Y == 0 && winTag(tag)
	})
	for _, tag := range startingPoint {
		derived(s, tag, positionMap)
	}

	winLine := &flow.WinLine{
		Tags: helper.MapFilter(positionMap, func(v [2]int, t bool) (*base.Tag, bool) {
			return s.Display[v[0]][v[1]].Copy(), t
		}),
	}
	//if len(winLine.Tags) >= length {
	//	return winLine
	//}
	return winLine
}

func derived(s *template.SpinInfo, tag *base.Tag, positionMap map[[2]int]bool) {
	positionMap[[2]int{tag.X, tag.Y}] = true
	derivedTags := lo.Filter(s.FindAllTags(), func(item *base.Tag, i int) bool {
		if positionMap[[2]int{item.X, item.Y}] {
			return false
		}
		if !winTag(item) {
			return false
		}

		if item.Y == tag.Y && helper.Abs(item.X-tag.X) == 1 {
			return true
		}

		if item.X == tag.X && helper.Abs(item.Y-tag.Y) == 1 {
			return true
		}
		return false
	})
	for _, derivedTag := range derivedTags {
		derived(s, derivedTag, positionMap)
	}
}

func winTag(tag *base.Tag) bool {
	return !tag.IsEmpty() && !tag.IsFixed() && tag.Name != "scatter" && tag.Name != "dragon" && tag.Name != "null_2" && tag.Name != "null_1"
}
