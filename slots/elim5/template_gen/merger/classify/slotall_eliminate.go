package classify

import (
	"elim5/enum"
	"elim5/logicPack/base"
	"elim5/logicPack/template"
	businessReq "elim5/model/business/request"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
)

// SlotALlEliminate slot 全线逻辑
type SlotALlEliminate struct {
	req      *businessReq.MergerIncrease
	spinInfo *template.SpinInfo
}

func NewSlotALlEliminate(req *businessReq.MergerIncrease, spinInfo *template.SpinInfo) *SlotALlEliminate {
	return &SlotALlEliminate{
		req:      req,
		spinInfo: spinInfo,
	}
}

func (n *SlotALlEliminate) Merger(index int) error {
	for i := 0; i < 1000; i++ {
		lines := n.spinInfo.FindSlotAllLine(n.req.TriggerLineNum, func(tags []*base.Tag) []*base.Tag {
			return tags
		})
		scaLine := n.spinInfo.FindTagsByName(enum.ScatterName)
		if len(lines) == 0 && len(scaLine) < n.req.TriggerSca {
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
		}

		if len(scaLine) > 0 {
			randSca := scaLine[helper.RandInt(len(scaLine))]
			n.spinInfo.Display[randSca.X][randSca.Y] = n.spinInfo.GetRandNormalTags(randSca.X, randSca.Y).Copy()
		}
	}
	return fmt.Errorf("normal eliminate error index:%d \n%s", index, n.spinInfo.PrintTable("normal eliminate"))
}
