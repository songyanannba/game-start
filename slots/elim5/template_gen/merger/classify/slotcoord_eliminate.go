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

// SlotCoordsEliminate Slot 坐标划线类型
type SlotCoordsEliminate struct {
	req      *businessReq.MergerIncrease
	spinInfo *template.SpinInfo
}

func NewSlotCoordsEliminate(req *businessReq.MergerIncrease, spinInfo *template.SpinInfo) *SlotCoordsEliminate {
	return &SlotCoordsEliminate{
		req:      req,
		spinInfo: spinInfo,
	}
}

func (n *SlotCoordsEliminate) Merger(index int) error {
	for i := 0; i < 1000; i++ {
		lines := n.spinInfo.FindSlotLine()
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
