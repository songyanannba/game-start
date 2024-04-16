package template

import (
	"elim5/global"
	"elim5/logicPack/base"
	"elim5/logicPack/template/flow"
	"elim5/model/business"
	"elim5/pbs/common"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
	"sort"
	"strconv"
)

type SpinInfo struct {
	LoadFunc
	infoType   int             // 信息类型  1普通转	2免费转
	Config     TemConfig       // 游戏配置
	Display    [][]*base.Tag   // 游戏展示
	initWindow [][]*base.Tag   // 初始化展示
	temInit    []int32         // 模版初始化行数
	SpinFlow   []flow.SpinFlow // 游戏流程

	templateRowMap map[int]int // 模版行数映射 map[列]标签id

	//第八台游戏需要数据
	Scatter    *flow.WinLine // scatter数量
	Multiplier *flow.WinLine // multiplier数量

	payTable map[string][]*NumMul // 划线对应赢钱倍率  标签名字=>[]{标签个数 ,标签赔率}
	Gain     int                  // 收益    // 赢钱
	Counter
	RmCount        int //截止到当前 消除的个数（历史游戏流程 每次消除个数的总记录）
	SingleTags     []SingleTag
	Level          int               //等级
	HideProperties map[[2]int]string //初始技能位置
	Which          int               //模板类型
	Slot43Data
}

type LoadFunc struct {
	CustomFill      map[int]func(s *SpinInfo) error                       //自定义填充函数
	CustomTag       map[string]func(tag *base.Tag, s *SpinInfo) *base.Tag //自定义标签修改
	CustomEffective func(x, y int) bool                                   //自定义有效性
	InitConvert     func(s *SpinInfo) [][]*base.Tag                       //初始化转换

}

func (s *SpinInfo) GetInfo() string {
	str := fmt.Sprintf("等级:%d\n初始取值:%v\nScatter:%d个%g倍\n翻倍标签:", s.Level, s.temInit, len(s.Scatter.Tags), s.Scatter.Mul)
	if s.Multiplier != nil {
		for _, i2 := range s.Multiplier.Tags {
			str += fmt.Sprintf("%g,", i2.Multiple)
		}
	}
	str += "\n"
	return str
}

func NewGameInfo(config TemConfig, spinType int) (*SpinInfo, error) {
	game := &SpinInfo{
		infoType:       spinType,
		Config:         config,
		templateRowMap: make(map[int]int),
		Scatter: &flow.WinLine{
			Tags: make([]*base.Tag, 0),
			Mul:  0,
		},
		initWindow: make([][]*base.Tag, 0),
		temInit:    make([]int32, 0),
		Counter: Counter{
			LeftCount: 0,
		},
		SingleTags:     make([]SingleTag, 0),
		HideProperties: make(map[[2]int]string),
		Which:          config.GetWhich(),
	}

	game.CustomFill = map[int]func(s *SpinInfo) error{}
	game.CustomTag = map[string]func(tag *base.Tag, s *SpinInfo) *base.Tag{}
	game.InitConvert = func(s *SpinInfo) [][]*base.Tag {
		return helper.CopyListArr(s.Display)
	}
	game.CustomEffective = func(x, y int) bool {
		return true
	}
	//templateRowMap 用于每一列的X坐标从那里取值（列上随机一个点,既是x的初始取值,每列上的x值也不一样）
	err := game.SetIndexMap()
	if err != nil {
		return nil, err
	}
	//坐标初始化 表里面的数据（标签）为空
	game.Display = helper.NewTable(config.GetCol(), config.GetRow(), func(x, y int) *base.Tag {
		return &base.Tag{X: x, Y: y, Name: ""}
	})
	game.SetPayTable() //结构化 赢钱标签对应的赔率和个数

	return game, nil
}

func NewGameInfoDeformity(config TemConfig, spinType int, display [][]*base.Tag) (*SpinInfo, error) {
	game := &SpinInfo{
		infoType:       spinType,
		Config:         config,
		templateRowMap: make(map[int]int),
		Scatter: &flow.WinLine{
			Tags: make([]*base.Tag, 0),
			Mul:  0,
		},
		initWindow: make([][]*base.Tag, 0),
		temInit:    make([]int32, 0),
		Counter: Counter{
			LeftCount: 0,
		},
		SingleTags:     make([]SingleTag, 0),
		HideProperties: make(map[[2]int]string),
		Which:          config.GetWhich(),
		Display:        display,
	}

	game.CustomFill = map[int]func(s *SpinInfo) error{}
	game.CustomTag = map[string]func(tag *base.Tag, s *SpinInfo) *base.Tag{}
	game.InitConvert = func(s *SpinInfo) [][]*base.Tag {
		return helper.CopyListArr(s.Display)
	}
	game.CustomEffective = func(x, y int) bool {
		return true
	}

	//templateRowMap 用于每一列的X坐标从那里取值（列上随机一个点,既是x的初始取值,每列上的x值也不一样）
	err := game.SetIndexMap()
	if err != nil {
		return nil, err
	}
	//坐标初始化 表里面的数据（标签）为空
	game.SetPayTable() //结构化 赢钱标签对应的赔率和个数

	return game, nil
}

type NumMul struct {
	Num int
	Mul float64
}

func (s *SpinInfo) Copy() *SpinInfo {
	newSpinInfo := *s
	return &newSpinInfo
}

func (s *SpinInfo) SetDebugConfig() error {
	config := s.Config
	debugConfigs := config.GetDebugConfigs()
	var (
		debugs []*business.DebugConfig
		ints   []int
	)
	debugs = lo.Filter(debugConfigs, func(item *business.DebugConfig, index int) bool {
		if int(item.PalyType) != s.infoType {
			return false
		}
		if config.GetIsTest() && item.DebugType != 2 {
			return false
		} else if !config.GetIsTest() && item.DebugType != 1 {
			return false
		}
		userId := int(config.GetUserId())
		if userId != 0 && item.UserId != 0 && userId != item.UserId {
			return false
		}

		return true
	})
	if len(debugs) == 0 {
		return nil
	}
	ints = helper.SplitInt[int](debugs[0].TemIndex, ",")
	if len(ints) <= 1 {
		randIndex, err := s.Config.GetInitTemIndex(s.infoType, s.Which)
		if err != nil {
			return err
		}
		ints = append(ints, randIndex)
	}
	if ints[0] != -1 {
		s.temInit = make([]int32, 0)
		for y := 0; y < config.GetCol(); y++ {
			s.templateRowMap[y] = int(ints[0])
			s.temInit = append(s.temInit, int32(ints[0]))
		}
	}

	var arr [][]string
	if err := global.Json.Unmarshal([]byte(debugs[0].ResultData), &arr); err != nil {
		panic(err)
	}

	for i, strings := range arr {
		for i2, s2 := range strings {
			if s.Display[i][i2].IsFixed() {
				continue
			}
			nameMul := helper.SplitStr(s2, "*")
			tagName := ""
			if len(nameMul) == 2 {
				tagName = nameMul[0]
				mul, _ := strconv.Atoi(nameMul[1])
				if i < len(s.Display) && i2 < len(s.Display[i]) {
					s.Display[i][i2] = s.Config.GetTagByName(tagName).Copy()
					s.Display[i][i2].Multiple = float64(mul)
				}
			} else {
				tagName = s2
				if i < len(s.Display) && i2 < len(s.Display[i]) {
					s.Display[i][i2] = s.Config.GetTagByName(s2).Copy()
				}
			}

			if f, ok := s.CustomTag[tagName]; ok {
				s.Display[i][i2] = f(s.Display[i][i2], s)
			}
		}
	}
	return nil
}

func (s *SpinInfo) SetPayTable() {
	s.payTable = make(map[string][]*NumMul)

	payTableMap := lo.GroupBy(s.Config.GetPayTables(), func(item *base.PayTable) string {
		return item.Tags[0].Name
	})
	for name, tables := range payTableMap {
		numMuls := lo.FilterMap(tables, func(item *base.PayTable, i int) (*NumMul, bool) {
			return &NumMul{
				Num: len(item.Tags),
				Mul: item.Multiple,
			}, true
		})
		sort.Slice(numMuls, func(i, j int) bool {
			return numMuls[i].Num <= numMuls[j].Num
		})
		s.payTable[name] = numMuls
	}
}

// GetGameType 获取游戏类型
func (s *SpinInfo) GetGameType() int {
	return s.infoType
}

// GetInitAck  转换ack 初始化展示
func (s *SpinInfo) GetInitAck() (initList []*common.Tags) {
	initList = make([]*common.Tags, 0)
	for _, tags := range s.initWindow {
		pbTags := &common.Tags{
			Tags: make([]*common.Tag, 0),
		}
		for _, tag := range tags {
			pbTags.Tags = append(pbTags.Tags, tag.ToTagAck())
		}
		initList = append(initList, pbTags)
	}
	for ints, m := range s.HideProperties {
		initList[ints[0]].Tags[ints[1]].Redundancy = m
	}
	return
}

func (s *SpinInfo) GetTemInitRows() []int32 {
	return s.temInit
}

func (s *SpinInfo) GetSingleTags() []*common.SingleTag {

	var singTags []*common.SingleTag
	for _, singleTags := range s.SingleTags {
		sTag := &common.Tag{
			TagId:    int32(singleTags.Tag.Id),
			X:        int32(singleTags.Tag.X),
			Y:        int32(singleTags.Tag.Y),
			Multiple: 0,
			IsWild:   true,
		}

		PosOder := make([]*common.Pos, 0)
		for _, val := range singleTags.PosOder {
			pos := &common.Pos{
				X: int32(val.X),
				Y: int32(val.Y),
			}
			PosOder = append(PosOder, pos)
		}

		singleTag := &common.SingleTag{
			Single:    sTag,
			FlowIndex: int32(singleTags.FlowIndex),
			PosOder:   PosOder,
			SubId:     int32(singleTags.SubId), //标签上面的倍数
		}

		singTags = append(singTags, singleTag)
	}

	return singTags
}

func (s *SpinInfo) GetInfoType() int {
	return s.infoType
}

func (s *SpinInfo) AddRmCount(winLines []*flow.WinLine) {
	for _, e := range winLines {
		s.RmCount += e.Integral
	}
}
