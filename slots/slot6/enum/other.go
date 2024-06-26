package enum

import (
	"slot6/utils"
)

// 用户状态
const (
	UserStatus1Normal = 1 // 正常
	UserStatus2Frozen = 2 // 冻结
)

// logicPack 滚轮类型
const (
	SlotReelType1Normal = 1 // slot普通滚轮
	SlotReelType2FS     = 2 // slot免费滚轮
)

// logicPack 赢钱组合类型
const (
	SlotPayTableType1Common = 1 // slot普通赢钱组合
	SlotPayTableType2Any    = 2 // slot任意赢钱组合
)

// logicPack 测试类型
const (
	SlotTestType1Time   = 1 // 指定次数
	SlotTestType2Die    = 2 // 死亡次数
	SlotTestType3Once   = 3 // 单次执行
	SlotTestType4Result = 4 // 指定结果
	SlotTestType5       = 5 // 指定结果
	SlotTestType6User   = 6 // 指定用户
)

// 后台操作类型
const (
	BackendOperateType1RefreshCache = 1 // 后台操作刷新缓存
)

// 系统默认管理员token
var AdminDefaultToken = utils.MD5V([]byte("sys_admin_default_token_123456"))

// slot配置默认tag
const (
	ConfigNameSlotDefaultTag = "slot_default_tag" // slot配置默认tag
	ConfigNameSlotFreeTag    = "slot_free_tag"    // slot配置默认tag
)

// logicPack 事件类型
const (
	SlotEvent1ChangeTable = 1 // slot事件换表
)

// logicPack spin 玩法类型
const (
	SlotSpinType1Normal = 1 // slot普通转
	SlotSpinType2Fs     = 2 // slot免费转
	SlotSpinType3Respin = 3 // slot重转
	SlotSpinType4FsRs   = 4 // slot免费重转
)

// 金币流水操作
const (
	MoneyAction1Play     = 1 // 游玩
	MoneyAction2Cash     = 2 // 现金
	MoneyAction3System   = 3 // 系统
	MoneyAction4Activity = 4 // 活动
)

// 金币流水类型
const (
	MoneyType1Spin = 1 // 操作类型转动

	MoneyType2Recharge = 200 // 操作类型充值

	MoneyType3Give = 300 // 操作类型赠送

	MoneyType4Sign = 400 // 操作类型签到

	MoneyType5Refund = 500 // 操作类型退款
)

// 中奖类型
const (
	JackpotType  = "0"
	WildStrType  = "1"
	SingStrType  = "2"
	PattableType = "3"
)

// spin ack 响应类型
const (
	SpinAckType1Normal = 1 // 普通响应
	SpinAckType2Demo   = 2 // demo响应
)

// txn Status 状态
const (
	TxnStatus1InProgress        = 1  // 玩家已开始游戏回合但尚未结束
	TxnStatus2CompleteInProcess = 2  // 游戏回合在数据库中被标 记为已完成；但是Result请求没有得到正确回复
	TxnStatus3CancelInProcess   = 3  // 退款处于异步队列中并正被发 送给运营商
	TxnStatus4Completed         = 9  // 玩家已完成游戏回合
	TxnStatus5Canceled          = 10 // 退款已完成
)
