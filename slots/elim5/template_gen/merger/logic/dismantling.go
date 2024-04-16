package logic

import (
	"elim5/global"
	"elim5/logicPack/base"
	"elim5/logicPack/component"
	"elim5/logicPack/template"
	"elim5/model/business"
	businessReq "elim5/model/business/request"
	"elim5/service/template_gen/merger/classify"
	"fmt"
	"github.com/samber/lo"
)

type Dismantling struct {
	TemGens     []*business.SlotTemplateGen
	Req         *businessReq.MergerIncrease
	Config      *component.Config
	Template    map[int][]*base.Tag
	TemplateRes map[int][]uint16
	E           Eliminate
	SpinInfo    *template.SpinInfo
	TemplateLen int
}

const (
	NormalEliminateType = iota + 1
	CountingEliminateType
	MegaWayEliminateType
	SlotAllEliminateType
	SlotCoordsEliminateType
	MahjongEliminateType
)

// SetEliminate 设置合并类型
func (d *Dismantling) SetEliminate() (err error) {
	switch d.Req.EliminateType {
	case NormalEliminateType:
		d.E = classify.NewNormalEliminate(d.Req, d.SpinInfo)
	case CountingEliminateType:
		d.E = classify.NewCountingEliminate(d.Req, d.SpinInfo)
	case MegaWayEliminateType:
		d.E = classify.NewMegaWayEliminate(d.Req, d.SpinInfo)
	case SlotAllEliminateType:
		d.E = classify.NewSlotALlEliminate(d.Req, d.SpinInfo)
	case SlotCoordsEliminateType:
		d.E = classify.NewSlotCoordsEliminate(d.Req, d.SpinInfo)
	case MahjongEliminateType:
		d.E = classify.NewMahjongEliminate(d.Req, d.SpinInfo)

	default:
		err = fmt.Errorf("未知的合并类型")
	}
	return
}

// NewDismantling 新建拆分
func NewDismantling(req *businessReq.MergerIncrease) (dis *Dismantling, err error) {
	dis = &Dismantling{
		Req:         req,
		TemGens:     make([]*business.SlotTemplateGen, 0),
		TemplateRes: make(map[int][]uint16),
	}
	err = global.GVA_READ_DB.Model(&business.SlotTemplateGen{}).
		Select("slot_id,id,template,`type`,`rtp`,`special_config`,`which`").
		Where("id in ?", req.Ids).
		Find(&dis.TemGens).Error
	if err != nil {
		return nil, err
	}

	dis.Config, err = component.GetSlotConfig(uint(dis.TemGens[0].SlotId), false)
	if err != nil {
		return nil, err
	}

	if err = dis.Verify(); err != nil {
		return nil, err
	}

	err = dis.InitTemplate()
	if err != nil {
		return nil, err
	}

	dis.SpinInfo, err = template.NewGameInfo(dis, dis.GetGameType())
	if err != nil {
		return nil, err
	}
	err = dis.SetEliminate()
	return
}

// Verify 数据验证
func (d *Dismantling) Verify() error {
	if len(d.TemGens) == 0 {
		return fmt.Errorf("请选择需要合并的模板")
	}
	groups := lo.GroupBy(d.TemGens, func(item *business.SlotTemplateGen) string {
		return fmt.Sprintf("%d-%d-%d-%d", item.SlotId, int(item.Type), item.Rtp, item.Which)
	})
	if len(groups) > 1 || len(groups) == 0 {
		return fmt.Errorf("请选择相同类型的模版合并")
	}
	for _, extraTag := range d.Req.ExtraTags {
		tag := d.Config.GetTag(extraTag.Name)
		if tag.IsEmpty() {
			return fmt.Errorf("额外标签 %s 不存在", extraTag.Name)
		}
	}

	return nil
}
