package component

import (
	"elim5/enum"
	"elim5/global"
	"elim5/logicPack/base"
	"elim5/logicPack/eliminate"
	"elim5/logicPack/template"
	"elim5/model/business"
	"elim5/pbs/common"

	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type Line struct {
	Name string
	Tags []*base.Tag
	Win  int64
}

func (l Line) String() string {
	return fmt.Sprintf("%s %v %d\n", l.Name, l.Tags, l.Win)
}

type Spin struct {
	*Options

	Config           *Config       `json:"-"`
	InitDataList     [][]*base.Tag `json:"-"` // 初始数据
	InitDataListCopy [][]*base.Tag `json:"-"` // 初始数据
	ResDataList      [][]*base.Tag // 结果数据
	WildList         []*base.Tag   // 百搭数据
	SingleList       []*base.Tag   // 单出数据
	WinLines         []*Line       // 赢钱线

	Jackpot        *JackpotData     // 最终奖池
	PayTables      []*base.PayTable // 最终payTable
	Bet            int              // 压注
	Gain           int              // 赢钱
	OtherMul       int              // 其他倍数
	FreeSpinParams FreeSpinParams   // 免费游戏参数

	Table    *eliminate.Table
	Id       int   // spinId
	ParentId int   // 父级ID
	PSpin    *Spin // 父级spin
	GroupId  int   // 组ID

	Which    int                // 选择哪个轮子
	SpinInfo *template.SpinInfo // spinInfo

	IsWildReSpin    bool //wild-reSpin
	MoveDataList    []*MoveDataList
	MoveResDataList [][]*base.Tag
	Redundancy      string //冗余数据
	Record          string
	ExtraGain       int
	ExtraMul        int
	TriggerMode     int //触发模式

	IsProcess bool // 是否有中间处理

	MateDataSlot21
	MateDataSlot23
	MateDataSlot26
	ExpandTagData
	MateDataSlot51
	MateDataSlot57
	MateDataSlot58

	AddList []*base.Tag
}

type Rtp struct {
}

// GetNormalTags 获取普通tag集合s
func (s *Spin) GetNormalTags() []*base.Tag {
	return lo.Filter(s.GetAllTags(), func(tag *base.Tag, i int) bool {
		return tag.Name != "scatter" && tag.Name != "multiplier" && tag.IsWild == false && tag.Name != enum.MegastackName && tag.Name != enum.PlusName && tag.Name != enum.SambaName
	})
}

func (s *Spin) GetRedundancyInt() int {
	if s.Redundancy == "" {
		return 0
	}
	strs := strings.Split(s.Redundancy, "-")
	if len(strs) != 2 {
		return 0
	}
	atoi, err := strconv.Atoi(strs[1])
	if err != nil {
		return 0
	}
	return atoi
}

func (s *Spin) String() string {

	initCop := s.PrintInitCopy("initCopy")
	res := ""
	for _, v := range lo.Slice(s.ResDataList, 0, 100) {
		res += fmt.Sprintf("%v", v)
		res += "\n"
	}
	if len(s.ResDataList) > 100 {
		res += "..."
	}

	text := fmt.Sprintf("%s\nresData: \n%v \n"+
		"winLineList: %v \n"+
		"wildList: %v \n"+
		"singleList: %v \n"+
		"jackpot: %v \n"+
		"bet: %v \n"+
		"gain: %v \n"+
		"TriggerMode: %v \n"+
		"freeSpin: %+v \n"+
		"options: %+v \n"+
		"locationInfo: %+v \n",
		initCop, res, s.WinLines, s.WildList, s.SingleList, s.Jackpot, s.Bet, s.Gain, s.TriggerMode, s.FreeSpinParams, s.Options, s.LocationInfo())

	text += s.LinkInfo()
	text += fmt.Sprintf("ExtraGain: %d \n", s.ExtraGain)
	return text
}

func (s *Spin) LocationInfo() string {
	str := ""
	if s.GetSlotId() != enum.SlotId13 {
		return str
	}

	for _, moveData := range s.MoveDataList {
		str += fmt.Sprintf("Start: %v ; Start tag %v \n",
			moveData.Start, moveData.Tags[moveData.Start])
	}

	return str
}

func (s *Spin) LinkInfo() string {
	str := ""
	if s.GetSlotId() != enum.SlotId23 {
		return str
	}

	if s.MateDataSlot23.AllMult <= 0 {
		return str
	}

	//中间列的最后一个值
	//str += fmt.Sprintf(" 中间列的最后一个标签值 tag: %v ", s.InitDataList[1][3])
	str += "slot23LinkTags:赢钱  " + helper.Itoa(s.MateDataSlot23.AllMult)
	for _, tag := range s.MateDataSlot23.LinkTags {
		str += fmt.Sprintf(" 标签 link tag: %v ", tag)
	}

	return str
}

func (s *Spin) Get51MutArr() string {
	return fmt.Sprintf("倍率集合: %+v \n", s.MateDataSlot51.MutArr)
}

func (s *Spin) Type() int {
	if s.IsReSpin {
		if s.IsReSpinLink {
			return enum.SpinAckType4ReSpinLink
		} else {
			return enum.SpinAckType3ReSpin
		}
	} else if s.IsFree {
		return enum.SpinAckType2FreeSpin
	}
	return enum.SpinAckType1NormalSpin
}

// GateDetail 获取详情
func (s *Spin) GateDetail() string {
	str := s.String()
	//替换[]符号变成]/n[
	str = strings.Replace(str, "] [", "]\n[", -1)
	return str
}

func NewSpin(slotId uint, amount int, options ...Option) (*Spin, error) {
	s := &Spin{
		Options:      &Options{},
		Bet:          amount,
		Gain:         0,
		MoveDataList: make([]*MoveDataList, 0),
	}
	c, err := GetSlotConfig(slotId, s.Demo)
	if err != nil {
		return s, err
	}
	s.Config = c
	s.setOption(options...)
	if s.Options.RatioConfirm == 0 {
		s.Options.RatioConfirm = 1
	}
	return s, nil
}

func (s *Spin) setOption(options ...Option) {
	s.Options.Spin = s
	for _, option := range options {
		option(s.Options)
	}
}

func (s *Spin) AddInitData(x int, arr []string) {
	var tags []*base.Tag
	for y, v := range arr {
		tag := *s.Config.GetTag(v)
		tags = append(tags, &tag)
		tag.X = x
		tag.Y = y
	}
	s.InitDataList = append(s.InitDataList, tags)
}

// JackpotMatch 请保证Multiple从高到低排序 以优先匹配最高金额
func (s *Spin) JackpotMatch() *JackpotData {
	for _, tags := range s.ResDataList {
		for _, jackpot := range s.Config.JackpotList {
			if line, ok := jackpot.Match(tags); ok {
				s.Jackpot = jackpot
				// 计算赢钱
				line.Win = helper.MulToInt(s.Bet, s.jackpotSum(jackpot))
				s.WinLines = append(s.WinLines, line)
				return jackpot
			}
		}
	}
	return nil
}

// PayTableOnceMatch 请保证PayTableList按Multiple从高到低排序 以优先匹配最高金额
func (s *Spin) PayTableOnceMatch() *base.PayTable {
	for _, tags := range s.ResDataList {
		for _, table := range s.Config.PayTableList {
			if ok, newTable := table.Match(tags); ok {
				line := &Line{
					Tags: newTable.Tags,
					Win:  helper.MulToInt(s.Bet, newTable.Multiple),
				}
				s.WinLines = append(s.WinLines, line)

				s.PayTables = append(s.PayTables, newTable)
				return newTable
			}
		}
	}
	return nil
}

// PayTableMatch 匹配所有payTable
func (s *Spin) PayTableMatch() {
	for _, tags := range s.ResDataList {
		for _, table := range s.Config.PayTableList {
			if ok, newTable := table.Match(tags); ok {
				s.WinLines = append(s.WinLines, &Line{
					Tags: newTable.Tags,
					Win:  helper.MulToInt(s.Bet, newTable.Multiple),
				})

				s.PayTables = append(s.PayTables, newTable)
				break
			}
		}
	}
}

func (s *Spin) GetPayTableALlTags() []*base.Tag {
	var tags []*base.Tag
	for _, table := range s.PayTables {
		tags = append(tags, table.Tags...)
	}
	return tags
}

func ParseDefaultTag(v string, index int) []int {
	res := make([]int, index)
	if v == "" {
		return res
	}
	arr := helper.SplitInt[int](v, ",")
	for k, vv := range arr {
		res[k] = vv
		if k >= index-1 {
			break
		}
	}
	return res
}

// SetDebugInitData 设置调试初始数据
func (s *Spin) SetDebugInitData(userId uint) {
	debugType := uint8(1)
	playType := uint8(1)
	if s.IsTest {
		debugType = 2
	}
	if s.IsReSpin {
		playType = 3
	}
	if s.IsFree {
		playType = 2
	}
	debugs := lo.Filter(s.Config.Debugs, func(item *business.DebugConfig, index int) bool {
		return item.DebugType == debugType && item.PalyType == playType && (item.UserId == int(userId) || item.UserId == 0)
	})
	if len(debugs) == 0 {
		return
	}
	userDebugs := lo.Filter(debugs, func(item *business.DebugConfig, index int) bool {
		return item.UserId == int(userId)
	})
	if len(userDebugs) > 0 {
		s.JsonToTags(userDebugs[0].ResultData)
	} else {
		s.JsonToTags(debugs[0].ResultData)
	}
	s.IsSetDebug = true
}

func (s *Spin) JsonToTags(jsonStr string) {
	s.InitDataList = [][]*base.Tag{}
	var arr [][]string
	if err := global.Json.Unmarshal([]byte(jsonStr), &arr); err != nil {
		panic(err)
	}
	arr = helper.ArrVertical(arr)
	for k, v := range arr {
		s.AddInitData(k, v)
	}
	for a, tags := range s.InitDataList {
		for b, tag := range tags {
			if tag.Id == -1 {
				newTag := s.Config.GetTag(enum.NullName)
				newTag.X = tag.X
				newTag.Y = tag.Y
				s.InitDataList[a][b] = newTag
			}
		}
	}
}

func (s *Spin) GetDebugData() string {
	debugType := uint8(1)
	playType := uint8(1)
	if s.IsTest {
		debugType = 2
	}
	if s.IsReSpin {
		playType = 2
	}
	if s.IsFree {
		playType = 3
	}
	debugs := lo.Filter(s.Config.Debugs, func(item *business.DebugConfig, index int) bool {
		return item.DebugType == debugType && item.PalyType == playType
	})
	if len(debugs) == 0 {
		return ""
	}
	return debugs[0].ResultData
}

func (s *Spin) FinalCardList() []*common.Cards {
	finalCardList := []*common.Cards{}
	for _, tags := range s.InitDataList {
		cards := common.Cards{}
		for _, tag := range tags {
			if s.MoveDataList[tag.X].Step == 0 {
				tagCard := tag.ToCard()
				cards.Cards = append(cards.Cards, tagCard)
			} else {
				dataList := s.MoveDataList[tag.X]
				initSite := dataList.InitSite
				step := dataList.Step
				dir := dataList.Dir
				n := 0
				if tag.Y == 0 {
					n = 1
				} else if tag.Y == 1 {
					n = 0
				} else {
					n = -1
				}
				if dir == 1 {
					nTag := dataList.Tags[initSite-n+step]
					nTag.X = tag.X
					nTag.Y = tag.Y
					tagCard := nTag.ToCard()
					cards.Cards = append(cards.Cards, tagCard)
				} else if dir == 2 {
					nTag := dataList.Tags[initSite-n-step]
					nTag.X = tag.X
					nTag.Y = tag.Y
					tagCard := nTag.ToCard()
					cards.Cards = append(cards.Cards, tagCard)
				} else {
					continue
				}
			}
		}
		finalCardList = append(finalCardList, &cards)
	}
	return finalCardList
}

// FindPayTableByTagsName 查找指定名称的payTable
func (s *Spin) FindPayTableByTagsName(Include ...string) []*base.PayTable {
	var pTs []*base.PayTable
	for _, table := range s.PayTables {
		tagsName := base.GetTagsName(table.Tags)
		if helper.InArr(tagsName, Include) {
			pTs = append(pTs, table)
		}
	}
	return pTs
}

func (s *Spin) PrintDisplaySlot21(str string) string {
	str += ":\n"
	for _, col := range s.Display {
		str += fmt.Sprintf("%s\t", strconv.Itoa(col.X)+":"+strconv.Itoa(col.Y)+" "+"Mult:"+strconv.Itoa(int(col.Multiple))+" "+helper.If(col.Name == "", "🀆", col.Name))
		str += "\r\n"
	}
	//fmt.Println(str)
	return str + "\r\n"
}

func (s *Spin) PrintInitCopy(str string) string {
	if len(s.InitDataListCopy) == 0 {
		return ""
	}
	str += ":\n"
	for _, tags := range helper.ArrVertical(s.InitDataListCopy) {
		for _, tag := range tags {
			s := fmt.Sprintf("%d:%d%s", tag.X, tag.Y, tag.Name)
			str += s + sep(10-len(s)) + "\t"
		}
		str += "\r\n"
	}
	return str
}

func sep(int2 int) string {
	str := ""
	for i := 0; i < int2; i++ {
		str += " "
	}
	return str
}

func (s *Spin) PrintParamSlot21(str string, tags []*base.Tag) string {
	//str := ""
	str += ":\n"
	for _, col := range tags {

		str += fmt.Sprintf("%s\t", strconv.Itoa(col.X)+":"+strconv.Itoa(col.Y)+" "+strconv.Itoa(int(col.Multiple))+" "+helper.If(col.Name == "", "🀆", col.Name))
		str += "\r\n"
	}
	//fmt.Println(str)
	return str + "\r\n"
}

func (s *Spin) IsNextByName(name string) bool {
	next := false
	for _, tag := range s.MateDataSlot21.Display {
		if tag.Name == name {
			next = true
		}
	}
	return next
}
