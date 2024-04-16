package enum

import "errors"

const (
	Yes = 1 // 是
	No  = 2 // 否
)

const (
	FUN          = "FUN" // 测试商户货币
	DemoCurrency = "USD" // 测试默认货币

	TESTCurrency = "TEST"                                        // 测试用户货币
	TESTAmount   = 100000                                        // 测试固定金额
	TESTTxID     = "test_tx"                                     // 测试TxID
	TESTPassword = "123testBot123"                               // 测试用户密码
	TESTToken    = "222434-751ee37d-c77f-4888-9107-9c68806aa2fa" // 测试用户token

	TestAgent = "test" // 测试商户标识
	TestUID   = 1
)

const (
	CommonStatusUnknown    = 0  // 通用状态未知
	CommonStatusBegin      = 1  // 通用状态开始
	CommonStatusProcessing = 2  // 通用状态处理中
	CommonStatusClose      = 8  // 通用状态关闭
	CommonStatusError      = 9  // 通用状态异常
	CommonStatusFinish     = 10 // 通用状态完成
	CommonStatusAwait      = 11 // 通用状态等待
	CommonStatusPause      = 12 // 通用状态暂停
)

// 客户端通用错误消息
var (
	ErrUnknown        = errors.New("unknown")
	ErrBusy           = errors.New("Service_Busy")
	ErrNoMoney        = errors.New("No_Money")
	ErrNoNet          = errors.New("No_NET")
	ErrNoServer       = errors.New("No_Server")
	ErrSysError       = errors.New("Sys_Error")
	ErrTokenInvalid   = errors.New("Token_Invalid")
	ErrRecordNotExist = errors.New("Record_Not_Exist")
	ErrSlotError      = errors.New("Slot_Error")
	ErrParamsError    = errors.New("Params_Error")

	ErrBet        = errors.New("bet gear not match")
	ErrGameIsNo   = errors.New("the current game does not exist")
	ErrUnrealized = errors.New("unrealized")
	ErrBlackList  = errors.New("current country is not supported")
	ErrNotOpen    = errors.New("the current game is not yet open")
	ErrCommon     = errors.New("an error has occurred, please contact the administrator")
	ErrReConn     = errors.New("an error was encountered while getting an unfinished process")
	ErrStatus     = errors.New("status has changed, please login again")
)

const DemoAmount = 10000000

const (
	SpinAckType1NormalSpin    = 1  // 普通转
	SpinAckType2FreeSpin      = 2  // 免费转
	SpinAckType3ReSpin        = 3  // 重转
	SpinAckType4ReSpinLink    = 4  // 重转带次数
	SpinAckType1BranchRaise   = 11 //加注转
	SpinAckType1BranchBuyFree = 12 //购买免费转
)

var SpinAckTypeMap = map[int]string{
	SpinAckType1NormalSpin: "普通转",
	SpinAckType2FreeSpin:   "免费转",
	SpinAckType3ReSpin:     "重转",
	SpinAckType4ReSpinLink: "重转带次数",
}

const (
	NoBuy   = 0 // 未购买
	BuyFree = 1 // 购买免费转
	BuyRe   = 2 // 购买重转
)

const (
	VueTypeDefault = "default"
	VueTypePrimary = "primary"
	VueTypeSuccess = "success"
	VueTypeInfo    = "info"
	VueTypeWarning = "warning"
	VueTypeDanger  = "danger"
)

const ClusterChannel = "cluster"

const (
	SurvivalTypeMerchant = 1
	SurvivalTypeSlot     = 2
)
