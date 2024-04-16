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

// CountingEliminate 计数消除类型 类似第8台
type CountingEliminate struct {
	req      *businessReq.MergerIncrease
	spinInfo *template.SpinInfo
}

func NewCountingEliminate(req *businessReq.MergerIncrease, spinInfo *template.SpinInfo) *CountingEliminate {
	return &CountingEliminate{
		req:      req,
		spinInfo: spinInfo,
	}
}

func (n CountingEliminate) Merger(index int) error {
	//fmt.Println(spinInfo.PrintTable("index:" + strconv.Itoa(index) + " start"))
	//alterTags := make([]*base.Tag, 0)
	for i := 0; i < 1000; i++ {
		lines := n.spinInfo.FindCountLine(n.req.TriggerLineNum)
		scaLine := n.spinInfo.FindTagsByName(enum.ScatterName)
		if len(lines) == 0 && len(scaLine) < n.req.TriggerSca {
			//fmt.Printf("AlterTags:%v\n", alterTags)
			//fmt.Println(n.spinInfo.PrintTable("index:" + strconv.Itoa(index) + " end"))
			return nil
		}

		tags := helper.ListToArr(lines)
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
			//alterTags = append(alterTags, randTag.Copy())
		}

		if len(scaLine) > 0 {
			randSca := scaLine[helper.RandInt(len(scaLine))]
			//alterTags = append(alterTags, randSca.Copy())
			n.spinInfo.Display[randSca.X][randSca.Y] = n.spinInfo.GetRandNormalTags(randSca.X, randSca.Y).Copy()
		}

	}
	return fmt.Errorf("normal eliminate error index:%d \n%s", index, n.spinInfo.PrintTable("normal eliminate"))
}
