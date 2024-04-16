package component

import (
	"elim5/enum"
	"elim5/logicPack/base"
	. "elim5/logicPack/cache"
	"elim5/model/business"
	"elim5/utils"
	"elim5/utils/helper"
	"fmt"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"strings"
)

// Config 请勿在初始化后修改config中的数据 该数据为机台的全局数据
type Config struct {
	SlotId uint // 机器id

	Status uint8       // 状态
	Index  int         // 列数
	Row    int         // 行数
	BetMap base.BetMap // 押注区间

	Raise  float64 // 加注倍率
	BuyFee float64 // 购买费用
	BuyRes float64 // 购买资源

	reel         map[int][]*Reel      // Map[返还比]滚轮数据
	Coords       [][]base.Coordinate  // 划线坐标
	tagMap       map[string]*base.Tag // 所有标签的Map
	tagIdMap     map[int]*base.Tag    // 所有标签的Map
	PayTableList []*base.PayTable     // 赢钱组合
	JackpotList  []*JackpotData       // 奖池
	Event        map[int]*base.Event  // Map[返还比]特殊事件
	Fakes        *Fakes               // 假数据

	place     []int                   //后台配置指定的排布索引
	freePlace []int                   //后台配置指定的排布索引
	Debugs    []*business.DebugConfig // 调试配置

	TopMul       int // 最高倍数
	IsOpenABTest uint8
	Templates    map[string]map[int][]uint16 // 模板数据 map[{rtp}-{类型}-{换表序号}]map[{列号}][]{标签id}  rtp=>类型=>索引=>标签

	AbTestIds []uint
}

func (c *Config) GetTemplate(typ, rtp, which int) (map[int][]uint16, error) {
	key := fmt.Sprintf("%d-%d-%d", rtp, typ, which)
	if tem, ok := c.Templates[key]; ok {
		return tem, nil
	}
	return nil, fmt.Errorf("not found template typ:%d rtp:%d which:%d ", typ, rtp, which)
}

// GetSlotConfig 获取slot配置 从缓存中获取，如果没有则从数据库中获取
func GetSlotConfig(slotId uint, demo bool) (*Config, error) {
	isDemo := helper.If(demo, enum.Yes, enum.No)
	c, exist := GetConfigByCache(slotId, isDemo)
	if exist {
		return c, nil
	}

	rawData, err := NewDbRawDataBySlotId(slotId)
	if err != nil {
		return nil, err
	}

	conf, confDemo := rawData.NewSlotConfig()
	SetConfigCache(slotId, conf, enum.No)
	SetConfigCache(slotId, confDemo, enum.Yes)
	if demo {
		return confDemo, nil
	}

	return conf, nil
}

// GetSlotConfigByIndicateNum 通过AB测编号获取机台配置
func GetSlotConfigByIndicateNum(slotId uint, demo bool, indicateNum uint8) (*Config, error) {
	// 查看一下源机台 是否开启ab test
	c, err := GetSlotConfig(slotId, demo)
	if err != nil {
		return nil, err
	}

	// 没有开启 AB测 或者开启 AB测 但是用户编号没有设置
	if c.IsOpenABTest != enum.OpenABTest || indicateNum == 0 {
		return c, err
	}

	// 获取AB测机台ID
	c, err = GetSlotConfig(business.EncodeSlotId(slotId, indicateNum, len(c.AbTestIds)), demo)
	if err != nil {
		return nil, err
	}
	c.IsOpenABTest = enum.OpenABTest
	return c, err
}

func GetConfigCacheKey(slotId uint, isDemo int) string {
	return fmt.Sprintf("slot_config_%d_%d", slotId, isDemo)
}

func GetConfigByCache(slotId uint, isDemo int) (*Config, bool) {
	res, exist := Cache.Get(GetConfigCacheKey(slotId, isDemo))
	if exist {
		return res.(*Config), true
	}
	return nil, false
}

func SetConfigCache(slotId uint, c *Config, isDemo int) {
	Cache.Set(GetConfigCacheKey(slotId, isDemo), c, local_cache.NoExpire)

}

func DeleteConfigCache(slotId uint) {
	Cache.Delete(GetConfigCacheKey(slotId, enum.No))
	Cache.Delete(GetConfigCacheKey(slotId, enum.Yes))
}

func FlushConfigCache() {
	Cache.Flush()
}

func (c *Config) GetTag(tagName string) *base.Tag {
	tag, ok := c.tagMap[tagName]
	if ok {
		return tag.Copy()
	}
	return base.NewTag(-1, "", 1, 0)
}

func (c *Config) GetReelMap() map[int][]*Reel {
	return c.reel
}

func (c *Config) GetTagById(id int) *base.Tag {
	tag, ok := c.tagIdMap[id]
	if ok {
		return tag.Copy()
	}
	return base.NewTag(-1, "", 1, 0)
}

func (c *Config) GetTags(tagNames ...string) []*base.Tag {
	var tags []*base.Tag
	for _, tagName := range tagNames {
		tag, ok := c.tagMap[tagName]
		if !ok {
			tag = base.NewTag(-1, tagName, 1, 0)
		}
		tags = append(tags, tag.Copy())
	}
	return tags
}

func (c *Config) GetTagsAndExclude(tagNames ...string) []*base.Tag {
	var tags []*base.Tag
	for _, tag := range c.tagMap {
		if helper.InArr(tag.Name, tagNames) {
			continue
		}
		tags = append(tags, tag.Copy())
	}
	return tags
}

func (c *Config) GetAllTag() []base.Tag {
	var tags []base.Tag
	for _, tag := range c.tagMap {
		tags = append(tags, *tag)
	}
	return tags
}

func (c *Config) GetAllTagQuote() []*base.Tag {
	var tags []*base.Tag
	for _, tag := range c.tagMap {
		tags = append(tags, tag.Copy())
	}
	return tags
}

func (c *Config) GetTagIdMap() map[int]base.Tag {
	m := map[int]base.Tag{}
	for _, tag := range c.tagMap {
		m[tag.Id] = *tag
	}
	return m
}

func (c *Config) GetTagsLayout(str string) []uint16 {
	var ids []uint16
	strs := strings.Split(utils.FormatCommandStr(str), ",")
	for _, s := range strs {
		if s == "" {
			continue
		}
		id := uint16(helper.Atoi(s))
		ids = append(ids, id)
	}
	return ids
}

// GetTagsLayoutTest 测试用
func (c *Config) GetTagsLayoutTest(str string, col int) []uint16 {
	var ids []uint16
	//strs := strings.Split(utils.FormatCommandStr(str), ",")
	//str := ""
	if col == 0 {
		str = "66,62,69,67,67,67,67,68,65,69,67,69,63,67,66,65,66,68,68,63,63,64,64,64,65,67,67,67,67,69,68,69,62,63,63,64,68,64,65,62,68,62,62,69,64,69,69,62,66,62,67,62,64"
	}

	if col == 1 {
		str = "66,62,69,67,67,67,67,68,65,69,67,69,69,67,66,64,65,67,64,69,66,64,69,69,62,67,62,67,68,62,65,68,68,68,64,64,69,64,69,69,62,62,62,67,62,64,62,65,62,65,64,62,66"
	}

	if col == 2 {
		str = "67,69,69,67,68,69,63,67,66,66,68,69,63,63,66,64,67,65,66,66,64,64,69,69,66,67,62,67,62,64,62,68,62,62,64,64,68,62,65,62,68,62,65,69,64,69,69,62,66,62,67,62,64"
	}

	if col == 3 {
		str = "69,69,67,64,64,67,64,66,65,68,68,63,63,65,65,65,67,65,63,66,64,69,69,66,66,64,68,62,69,66,69,69,62,67,64,69,64,69,69,68,68,62,67,62,64,62,65,62,65,64,62,67,64"
	}
	if col == 4 {
		str = "66,62,69,64,64,67,64,68,65,69,67,69,63,67,65,64,67,67,69,66,66,69,68,66,63,64,64,64,64,62,69,69,68,64,64,64,68,62,65,62,68,62,65,65,64,65,69,62,66,62,67,62,64"
	}

	if col == 5 {
		str = "67,69,64,64,69,63,64,67,65,64,66,68,68,63,65,64,67,64,65,66,69,68,68,66,64,62,64,62,63,69,64,69,69,62,66,64,68,62,65,62,68,62,65,69,64,69,69,62,66,62,67,62,64"
	}
	strs := strings.Split(str, ",")
	for _, s := range strs {
		if s == "" {
			continue
		}
		id := uint16(helper.Atoi(s))
		ids = append(ids, id)
	}
	return ids
}

// place     []int                   //后台配置指定的排布索引
// freePlace []int                   //后台配置指定的排布索引
func (c *Config) GetPlace() []int {
	return c.place
}

func (c *Config) GetFreePlace() []int {
	return c.freePlace
}
