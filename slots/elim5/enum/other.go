package enum

// 用户状态
const (
	UserStatus1Normal = 1 // 正常
	UserStatus2Frozen = 2 // 冻结
)

// 用户类型
const (
	ApiType0Common          = 0 // 通用钱包
	ApiType1SeamlessWallet  = 1 // 简单钱包
	ApiType2BalanceTransfer = 2 // 余额转结
)

// slot 滚轮类型
const (
	SlotReelType1Normal = 1 // slot普通滚轮
	SlotReelType2FS     = 2 // slot免费滚轮
	SlotReelType12Re    = 2 // slot免费滚轮
)

// slot 赢钱组合类型
const (
	SlotPayTableType1Common = 1 // slot普通赢钱组合
	SlotPayTableType2Any    = 2 // slot任意赢钱组合
)

// slot 测试类型
const (
	SlotTestType1Time     = 1 // 指定次数
	SlotTestType2Die      = 2 // 死亡次数
	SlotTestType3Once     = 3 // 单次执行
	SlotTestType4User     = 4 // 指定用户
	SlotTestType5Pressure = 5 // 压力测试
	SlotTestType5RTP      = 6 // RTP波动测试
)

// 后台操作类型
const (
	BackendOperateType1RefreshGameCache     = 1 // 后台操作刷新游戏缓存
	BackendOperateType2RefreshMerchantCache = 2 // 后台操作刷新商户缓存
)

// ClusterOperateType1KickAccount 集群操作类型
const (
	ClusterOperateType1KickAccount  = 1 // 集群操作踢号
	ClusterOperateType2UpdateOnline = 2 // 更新在线玩家
)

// slot配置默认tag
const (
	ConfigNameSlotDefaultTag = "slot_default_tag" // slot配置默认tag
	ConfigNameSlotFreeTag    = "slot_free_tag"    // slot配置默认tag
)

// slot 事件类型
const (
	SlotEvent1ChangeTable = 1 // slot事件换表
)

const (
	BetTypeCommon  = 1
	BetTypeRaise   = 2
	BetTypeBuyFree = 3
	BetTypeBuyRe   = 4
)

// slot spin 玩法类型
const (
	SlotSpinType1Normal = 1 // slot普通转
	SlotSpinType2Fs     = 2 // slot免费转
	SlotSpinType3Respin = 3 // slot重转
	SlotSpinType4FsRs   = 4 // slot免费重转
	SlotSpinUpLevel     = 5 // slot升级
)

// 金币流水操作
const (
	MoneyAction1Play     = 1 // 游玩
	MoneyAction2Cash     = 2 // 现金
	MoneyAction3System   = 3 // 系统
	MoneyAction4Activity = 4 // 活动
)

// 机台汇总类型
const (
	MoneySlotTypeAll = iota + 1
	MoneySlotTypeMonth
	MoneySlotTypeDay
)

// 金币流水类型
const (
	MoneyType1Bet    = 1 // 操作类型下注
	MoneyType2Win    = 2 // 操作类型赢钱
	MoneyType3Refund = 3 // 操作类型退款

	MoneyType200Recharge = 200 // 操作类型充值

	MoneyType300Give = 300 // 操作类型赠送

	MoneyType400Sign = 400 // 操作类型签到

)

// 中奖类型
const (
	JackpotType  = "0"
	WildStrType  = "1"
	SingStrType  = "2"
	PattableType = "3"
)

// txn Status 状态
const (
	TxnStatus0NoBet             = 0 // 未完成下注
	TxnStatus1InProgress        = 1 // 玩家已开始游戏回合但尚未结束
	TxnStatus2CompleteInProcess = 2 // 游戏回合在数据库中被标 记为已完成；但是Result请求没有得到正确回复
	TxnStatus3CancelInProcess   = 3 // 退款处于异步队列中并正被发 送给运营商
	TxnStatus4SpecialProcess    = 4 // 玩家游戏正在特殊回合中

	TxnStatus9Completed = 9  // 玩家已完成游戏回合
	TxnStatus10Canceled = 10 // 退款已完成
)

// merchant api code
const (
	MerchantApiCode0Success             = 0  // 成功
	MerchantApiCode7Fail                = 7  // 失败
	MerchantApiCode10AmountInsufficient = 10 // 余额不足
)

const (
	BumperTypeEmpty = 0
	BumperTypeTop   = 1
	BumperTypeLow1  = 2
	BumperTypeLow2  = 3
)

const (
	CacheKeyMoneySlot = "moneySlotCache"
)
