package component

import (
	"elim5/enum"
	"elim5/global"
	"elim5/logicPack/base"
	. "elim5/model/business"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
)

// DbRawData 数据库原始数据
type DbRawData struct {
	Slot      *Slot
	Reel      []*SlotReelData // 滚轮
	PayTable  []*SlotPayTable // 赢钱组合
	Payline   []*SlotPayline  // 划线规格
	Symbol    []*SlotSymbol   // 图标
	Jackpot   []*Jackpot      // 奖池规则
	Event     []*SlotEvent    // 特殊事件
	Fake      []*SlotFake     // 特殊事件
	Debugs    []*DebugConfig  // 调试配置
	Templates []*SlotTemplate // 模板类型=>列号=>标签
}

// NewDbRawDataBySlotId 创建数据库原始数据集
func NewDbRawDataBySlotId(slotId uint) (rawData *DbRawData, err error) {
	rawData = &DbRawData{}
	err = global.GVA_DB.First(&rawData.Slot, "id = ?", slotId).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Reel, "slot_id = ?", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	// payTable按倍数由高到低排序
	err = global.GVA_DB.Order("win_multiple desc").Find(&rawData.PayTable, "slot_id = ?", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Payline, "no = ?", rawData.Slot.PaylineNo).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Symbol, "slot_id = ?", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	rawData.Jackpot, err = GetJackpotListBySlot(rawData.Slot)
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Event, "slot_id = ?", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Fake, "slot_id = ?", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Debugs, "slot_id = ? and start = 1", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Templates, "slot_id = ? and `b_slot_template`.`lock`=? ", rawData.Slot.ID, 1).Error
	if err != nil {
		return
	}
	return
}

// NewSlotConfig 从原始数据创建slotConfig
func (d *DbRawData) NewSlotConfig() (*Config, *Config) {
	c := &Config{
		SlotId:       d.Slot.ID,
		Raise:        d.Slot.Raise,
		BetMap:       base.NewBetMap(d.Slot.BetNum),
		BuyFee:       d.Slot.BuyFreeSpin,
		BuyRes:       d.Slot.BuyReSpin,
		Status:       d.Slot.Status,
		tagMap:       map[string]*base.Tag{},
		tagIdMap:     map[int]*base.Tag{},
		TopMul:       d.Slot.TopMul,
		Templates:    map[string]map[int][]uint16{},
		Event:        map[int]*base.Event{},
		IsOpenABTest: d.Slot.OpenABTest,
		AbTestIds:    GetAbTestIds(d.Slot.ID),
	}

	// 解析规格
	c.Row, c.Index = d.Payline[0].ParseSpec()

	// 解析坐标
	for _, payline := range d.Payline {
		if payline.Position != "" {
			c.Coords = append(c.Coords, ParseCoordinate(payline.Position))
		}
	}

	reel, reelDemo := helper.Apart(d.Reel, func(v *SlotReelData) bool {
		return v.Demo == enum.No
	})

	// 解析滚轮
	c.reel = parseReelData(reel)

	// 解析图标
	for _, symbol := range d.Symbol {
		var include []string
		if symbol.IsWild == enum.Yes {
			include = symbol.ParseInclude()
		}
		if symbol.Multiple < 1 {
			symbol.Multiple = 1
		}
		tag := base.NewTag(int(symbol.ID), symbol.Name, float64(symbol.Multiple), symbol.IsSingleWin, include...)
		c.tagMap[symbol.Name] = tag
		c.tagIdMap[int(symbol.ID)] = tag
	}

	// 解析paytable
	for _, payTable := range d.PayTable {
		combine1, combine2 := payTable.ParseCombine()
		tags := c.GetTags(combine1...)
		if payTable.Type == enum.SlotPayTableType1Common {
			c.PayTableList = append(c.PayTableList, base.NewCommonPayTable(payTable.ID, tags, payTable.WinMultiple))
		} else {
			tags2 := c.GetTags(combine2...)
			c.PayTableList = append(c.PayTableList, base.NewAnyPayTable(payTable.ID, tags, payTable.CombineNum1, tags2, payTable.CombineNum2, payTable.WinMultiple))
		}
	}

	// 解析奖池
	for _, jackpot := range d.Jackpot {
		c.JackpotList = append(c.JackpotList, NewJackpotData(jackpot.ID, 0, 0, jackpot.End, jackpot.ParseCombine()))
	}
	c.place = make([]int, c.Index)
	c.freePlace = make([]int, c.Index)
	// 解析额外配置
	//xConfig, _ := GetXConfigCacheByName(enum.ConfigNameSlotDefaultTag + strconv.Itoa(int(c.SlotId)))
	//if xConfig.Status == enum.Yes {
	//	common, free, _ := strings.Cut(xConfig.Value, "&")
	//	c.place = ParseDefaultTag(common, c.Index)
	//	c.freePlace = ParseDefaultTag(free, c.Index)
	//}

	// 解析事件  event 不是试玩的事件  eventDemo 试玩的事件
	event, eventDemo := helper.Apart(d.Event, func(v *SlotEvent) bool {
		return v.Demo == enum.No
	})
	parseEventData(c, event)

	// 解析假数据
	c.Fakes = newFakes(d.Fake)
	c.Debugs = d.Debugs

	// 解析demo
	demo := *c
	// 解析滚轮
	demo.reel = parseReelData(reelDemo)
	demo.Event = map[int]*base.Event{}
	// 解析事件
	parseEventData(&demo, eventDemo)

	templatesGroup := lo.GroupBy(d.Templates, func(i *SlotTemplate) string {
		key := fmt.Sprintf("%d-%d-%d", i.Rtp, i.Type, i.Which)
		return key
	})
	for group, templates := range templatesGroup {
		c.Templates[group] = map[int][]uint16{}
		for _, template := range templates {
			c.Templates[group][template.Column] = c.GetTagsLayout(template.Layout)
		}
	}
	return c, &demo
}

func parseReelData(reelData []*SlotReelData) map[int][]*Reel {
	reelsMap := make(map[int][]*Reel)
	ratioMap := lo.GroupBy(reelData, func(i *SlotReelData) int {
		return i.Rtp
	})
	for i, data := range ratioMap {
		var reels []*Reel
		// 解析滚轮
		reelMap := lo.GroupBy(data, func(i *SlotReelData) int {
			return i.Group
		})

		for i := 1; i <= len(reelMap); i++ {
			reels = append(reels, ParseReel(reelMap[i]))
		}
		reelsMap[i] = reels
	}

	return reelsMap
}

func parseEventData(c *Config, eventData []*SlotEvent) {
	ratios := lo.GroupBy(eventData, func(i *SlotEvent) int {
		return i.Rtp
	})
	for ratio, events := range ratios {
		c.Event[ratio] = base.NewEvent()
		slotId := c.SlotId
		if c.SlotId > enum.AbTestMinSlotId {
			id, err := DecodeSlotId(c.SlotId)
			slotId = id
			if err != nil {
				global.GVA_LOG.Error("parseEventData DecodeSlotId err")
				return
			}
		}
		switch slotId {
		case 1, 2, 3, 4:
			for i, ev := range events {
				c.Event[ratio].Add(i+1, ev.Event1)
			}
		case 6:
			c.Event[ratio].Unit6NewEvent(events)
		case 7:
			c.Event[ratio].Unit7NewEvent(events)
		case 8, 15:
			c.Event[ratio].Unit8NewEvent(events)
		case 51:
			for i, event := range events {
				c.Event[ratio].AddEventJoinLayout(i, event.Event1)
			}
		default:
			for i, ev := range events {
				c.Event[ratio].AddDefault(i, ev.Event1)
			}
		}
	}
}
