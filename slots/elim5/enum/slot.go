package enum

import (
	"elim5/utils/helper"
	"errors"
)

const (
	SlotMaxSpinNum     = 10000 // slot最大转动
	SlotMaxSpinStr     = "the number of times exceeded 10000. Procedure"
	SlotMaxFreeSpinErr = "more than 1000 free plays"
	SlotMaxReSpinNum   = 1000
)

const (
	GameType0               = 0        // 未知
	GameType1Slot           = 1        // slot
	GameType2Match          = 2        // 消消乐
	MatchSlotType1Link      = iota + 1 // 消消乐+划线
	MatchSlotType2Count                // 消消乐+数个数
	MatchSlotType2SpeLink              // 消消乐+左右连线
	MatchSlotType3Underline            // 划线消除
	MatchSlotType4Mahjong              // 划线消除
)

const (
	ComZero  = 0
	ComOne   = 1
	ComEmpty = ""
)

// Slot 第六台等级
const (
	Rand1 = iota + 1
	Rand2
	Rand3
	Rand4
	Rand5
)
const (
	ScatterName    = "scatter"
	MultiplierName = "multiplier"
	WildName       = "wild"
	NullName       = "null"
	SambaName      = "samba"
	LinkTagName    = "link"
	MegastackName  = "megastack"
	High1Name      = "high_1"
	PlusName       = "plus"
	BonusName      = "bonus"
	CloverName     = "clover"
	JackpotName    = "jackpot"
	GoldName       = "gold"
)

const (
	EmptyTagName   = "" //空标签
	SlotWild       = "wild"
	SlotWild1      = "wild_1"
	SlotWildReSpin = "wild_respin"
	SlotWild2      = "wild_2"
	SlotWild5      = "wild_5"
	NullSp         = "null_sp"
	Wild10         = "wild_10"
	Wild20         = "wild_20"
	Wild100        = "wild_100"
)

var SlotNameMap = map[string]string{
	Wild10:  Wild10,
	Wild20:  Wild20,
	Wild100: Wild100,
}

const (
	CoreTagMinNum = 4 //中心点周围 最少填充的tag个数
	CoreTagMaxNum = 8 //中心点周围 最多填充的tag个数

	GetLine = 5 //默认匹配连续标签的最少个数
	NumLine = 8 //默认匹配随机标签的最少个数

	QuantityTagNum = 8 //第八台逻辑 最少8个标签才能连消

	IsMustFreeScatterNum = 4 // 购买免费转填充4个 Scatter
	ScatterNum4Datum     = 15

	MinFreeScatterNum = 3 // 免费转 3个 Scatter
	ScatterNum3Datum  = 5

	Slot2FreeNum = 8 //第二台获取免费转的次数

	SameTagLen = 3 //在第二台 和 第三台用

	InitFillTagNum = 5 //初始化 填充tag 的数量

	IncFreeIntVal       = 10
	IncFreeIntValByFree = 5
)

// 机台ID
const (
	SlotId1 = iota + 1
	SlotId2
	SlotId3
	SlotId4
	SlotId5 //机台5
	SlotId6 //机台6
	SlotId7
	SlotId8
	SlotId9
	SlotId10
	SlotId11
	SlotId12
	SlotId13
	SlotId14
	SlotId15
	SlotId16
	SlotId17
	SlotId18
	SlotId19
	SlotId20
	SlotId21
	SlotId22
	SlotId23
	SlotId24 = 24
	SlotId25 = 25
	SlotId26 = 26
	SlotId27 = 27
	SlotId28 = 28
	SlotId29 = 29
	SlotId30 = 30
	SlotId31 = 31
	SlotId32 = 32
	SlotId33 = 33
	SlotId34 = 34
	SlotId35 = 35
	SlotId36 = 36
	SlotId37 = 37
	SlotId38 = 38
	SlotId39 = 39
	SlotId40 = 40
	SlotId41 = 41
	SlotId42 = 42
	SlotId43 = 43
	SlotId44 = 44
	SlotId45 = 45
	SlotId46 = 46
	SlotId47 = 47
	SlotId48 = 48
	SlotId49 = 49
	SlotId50 = 50
	SlotId51 = 51
	SlotId52 = 52
	SlotId53 = 53
	SlotId54 = 54
	SlotId57 = 57
	SlotId58 = 58
	SlotId59 = 59
	SlotId60 = 60
	SlotId61 = 61
	SlotId62 = 62
)

// SlotTypeMap 机台类型
var SlotTypeMap = map[int32]int{
	SlotId1: GameType1Slot, // 钻石
	SlotId2: GameType1Slot, // 宙斯
	SlotId3: GameType1Slot, // 金筹码
	SlotId4: GameType1Slot, // 挖矿

	SlotId5: GameType2Match, // 水果消消1
	SlotId6: GameType2Match, // 甜点消消
	SlotId7: GameType2Match, // 魔法消消
	SlotId8: GameType2Match, // 龙消消
	SlotId9: GameType2Match, // 忍者消消+划线

	SlotId10: GameType1Slot, // book 还不知道
	SlotId11: GameType0,     // buffalo 还不知道

	SlotId12: GameType1Slot, // 经典wild-respin
	SlotId13: GameType1Slot, // 经典长wild

	SlotId14: GameType2Match, // 水果消消2
	SlotId15: GameType2Match, // 宙斯换皮
	SlotId16: GameType2Match, // Sughar Rush 消消
	SlotId17: GameType2Match, // 金筹码换皮
	SlotId18: GameType1Slot,  // 挖矿换皮
	SlotId19: GameType1Slot,  // 经典wild-respin换皮
	SlotId21: GameType1Slot,
	SlotId23: GameType1Slot,
	SlotId22: GameType2Match, // 忍者消消+左右连线
	SlotId24: GameType1Slot,  // 经典长wild换皮
	SlotId25: GameType2Match, // 水果消消3
	SlotId26: GameType1Slot,
	SlotId31: GameType1Slot,  // 金矿
	SlotId30: GameType1Slot,  // 金矿
	SlotId29: GameType1Slot,  // 金矿
	SlotId28: GameType1Slot,  // 金矿
	SlotId27: GameType1Slot,  // 金矿
	SlotId32: GameType1Slot,  // 金矿
	SlotId33: GameType2Match, // 金矿
	SlotId43: GameType2Match, // 金矿
	SlotId46: GameType1Slot,  //
	SlotId44: GameType1Slot,  //
	SlotId48: GameType2Match, // 龙1 （土 水 火 巨龙四种 模式）
	SlotId49: GameType2Match, // 龙2 （土 水 火 巨龙四种 模式）
	SlotId50: GameType2Match, //溏心风暴
	SlotId51: GameType1Slot,  //
	SlotId52: GameType2Match, //

	SlotId53: GameType2Match,
	SlotId54: GameType2Match,
	//SlotId55: ,
	//SlotId56: ,
	SlotId57: GameType1Slot,
	SlotId58: GameType1Slot,
	//SlotId59: ,
	//SlotId60: ,
	//SlotId61: ,
}

// SlotPlayMethod 机台玩法
var SlotPlayMethod = map[uint8][]uint{
	// 购买加注
	BetTypeRaise: {SlotId3, SlotId8, SlotId15, SlotId17, SlotId30, SlotId33, SlotId51},
	// 购买免费转
	BetTypeBuyFree: {SlotId5, SlotId8, SlotId9, SlotId14, SlotId16, SlotId22, SlotId15, SlotId21, SlotId17, SlotId30, SlotId33, SlotId54},
	// 购买重转
	BetTypeBuyRe: {SlotId6, SlotId21},
}

// ReconnectSlot 断线重连机台
var ReconnectSlot = []int{SlotId21, SlotId29, SlotId43}

// IsReconnectSlot 是否为断线重连机台
func IsReconnectSlot[T int | uint | int32](slotId T) bool {
	return helper.InArr(int(slotId), ReconnectSlot)
}

// GetMatchSlotInfo 消消乐机台类型细分
func GetMatchSlotInfo(slotId int) int {
	switch slotId {
	case SlotId5:
		return MatchSlotType1Link
	case SlotId6:
		return MatchSlotType1Link
	case SlotId7:
		return MatchSlotType1Link
	case SlotId8, SlotId15:
		return MatchSlotType2Count
	case SlotId9:
		return MatchSlotType2SpeLink
	case SlotId14:
		return MatchSlotType1Link
	case SlotId16:
		return MatchSlotType1Link
	case SlotId22:
		return MatchSlotType1Link
	case SlotId33:
		return MatchSlotType2Count
	case SlotId43:
		return MatchSlotType3Underline
	case SlotId52:
		return MatchSlotType3Underline
	case SlotId53:
		return MatchSlotType3Underline
	case SlotId54:
		return MatchSlotType4Mahjong
	default:
		return GameType0
	}
}

func IsAloneAck(slotId uint) bool {
	if slotId > SlotId17 {
		return true
	}
	return false
}

// reel_data which
const (
	Which1 = 1 //
	Which2 = 2 //
)

const (
	AddReSpin1 = 1 // 检测中间是否是 wild_reSpin
	AddReSpin2 = 2 // 检测是否赢钱
)

const (
	SiteUP = iota + 1
	SiteDown
	SiteLeft
	SiteRight
)

const (
	Initial   = "初始"  //
	Eliminate = "消除"  //
	Drop      = "掉落"  //
	DropAfter = "填充后" //
)

const (
	HidePropertiesSkill = iota + 1 // 隐藏属性技能
	HidePropertiesMul              // 隐藏属性倍数
)

const (
	OpenABTest  = 1 //开启ab test
	CloseABTest = 0 //关闭ab test

	UserMachineAccordNum = 99    //生成用户路由到机器 所依据的随机基本数值 （AB test）
	IndicateNumPrefix    = "100" //ab测机台的前缀
	AbTestMinSlotId      = 100000
	IndicateNumLen       = 2    //长度 随机数长度 (逻辑长度 已经固定 为 2)
	ABTestLen            = 6    //ab test的 id 长度 （长度固定）
	IndicateNumStart     = "01" //随机数的开始（字符串标示）
	IndicateNumEnd       = "99" //随机数的结束（字符串标示）
	SlotFillPrefix       = "0"  //
)

var (
	SlotIdNotFound   = errors.New("slotId not found")
	ABSlotIdNotFound = errors.New("ab test slotId not found")
)

const (
	MachineVersion1 = "1.0.0" // 机台版本1.0.0
	MachineVersion2 = "2.0.0" // 机台版本2.0.0
)

var MachineVersionMap = map[int]string{
	SlotId1:  MachineVersion1,
	SlotId2:  MachineVersion1,
	SlotId3:  MachineVersion1,
	SlotId4:  MachineVersion1,
	SlotId5:  MachineVersion1,
	SlotId6:  MachineVersion1,
	SlotId7:  MachineVersion1,
	SlotId8:  MachineVersion1,
	SlotId9:  MachineVersion1,
	SlotId10: MachineVersion1,
	SlotId11: MachineVersion1,
	SlotId12: MachineVersion1,
	SlotId13: MachineVersion1,
	SlotId14: MachineVersion1,
	SlotId15: MachineVersion1,
	SlotId16: MachineVersion1,
	SlotId17: MachineVersion1,
	SlotId18: MachineVersion2,
	SlotId19: MachineVersion2,
	SlotId20: MachineVersion1,
	SlotId21: MachineVersion2,
	SlotId22: MachineVersion2,
	SlotId23: MachineVersion2,
	SlotId24: MachineVersion2,
	SlotId25: MachineVersion2,
	SlotId26: MachineVersion2,
	SlotId27: MachineVersion2,
	SlotId28: MachineVersion2,
	SlotId29: MachineVersion2,
	SlotId30: MachineVersion2,
	SlotId31: MachineVersion2,
	SlotId32: MachineVersion2,
	SlotId33: MachineVersion1,
	SlotId34: MachineVersion1,
	SlotId35: MachineVersion1,
	SlotId36: MachineVersion2,
	SlotId37: MachineVersion2,
	SlotId38: MachineVersion2,
	SlotId39: MachineVersion2,
	SlotId40: MachineVersion2,
	SlotId41: MachineVersion2,
	SlotId42: MachineVersion2,
	SlotId43: MachineVersion2,

	SlotId46: MachineVersion2,

	SlotId48: MachineVersion2,
	SlotId49: MachineVersion2,
	SlotId50: MachineVersion2,
	SlotId51: MachineVersion2,
	SlotId52: MachineVersion2,
	SlotId53: MachineVersion2,
	SlotId54: MachineVersion2,

	SlotId57: MachineVersion2,
	SlotId58: MachineVersion2,
}
