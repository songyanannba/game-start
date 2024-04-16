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

// MegaWayEliminate Megaway消除类型 类型25台
type MegaWayEliminate struct {
	req      *businessReq.MergerIncrease
	spinInfo *template.SpinInfo
}

func NewMegaWayEliminate(req *businessReq.MergerIncrease, spinInfo *template.SpinInfo) *MegaWayEliminate {
	return &MegaWayEliminate{
		req:      req,
		spinInfo: spinInfo,
	}
}

func (n MegaWayEliminate) Merger(index int) error {
	var lines []*flow.WinLine
	fmt.Println(n.spinInfo.PrintTable("index:" + strconv.Itoa(index) + " start"))
	alterTags := make([]*base.Tag, 0)
	for i := 0; i < 1000; i++ {
		lines = n.spinInfo.FindMegaWayAllLine(n.req.TriggerLineNum, func(tags []*base.Tag) []*base.Tag {
			return tags
		})
		lines = append(lines, n.spinInfo.FindMegaWayAllLineDisplacement(n.req.TriggerLineNum, func(tags []*base.Tag) []*base.Tag {
			return tags
		})...)
		scaLine := n.spinInfo.FindTagsByName(enum.ScatterName)
		if len(lines) == 0 && len(scaLine) < n.req.TriggerSca {
			fmt.Printf("AlterTags:%v\n", alterTags)
			fmt.Println(n.spinInfo.PrintTable("index:" + strconv.Itoa(index) + " end"))
			return nil
		}
		tagsList := make([][]*base.Tag, 0)
		for _, line := range lines {
			tagsList = append(tagsList, line.Tags)
		}
		tags := helper.ListToArr(tagsList)
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
