package component

import (
	"elim5/global"
	"elim5/logicPack/base"
)

type Options struct {
	// 从父级继承时无需变动的参数
	IsTest           bool  // 是否测试
	Raise            int64 // 额外加注
	JackpotStartTime int64 // 奖池开始时间

	// 从父级继承时需要变动的参数
	IsFree       bool // 是否免费转
	IsReSpin     bool // 是否respin
	IsReSpinLink bool // 是否respinLink

	BuyFreeCoin int64 // 购买免费次数
	BuyReCoin   int64 // 购买免费次数

	Rank     int // 进度
	NextRank int // 下一次进度

	TagsLock   []*base.Tag // 锁定的tag
	IsMustFree bool        // 购买免费Free
	IsMustRes  bool        // 购买Res

	Demo bool // 是否是demo

	FreeNum int // 剩余免费次数
	ResNum  int //  剩余Respin次数

	DebugConfig string

	Spin *Spin `json:"-"`

	IsTemGen bool // 是否是临时生成的

	UserId uint // 用户id

	OptionIndex int // 1 三个scatter 2 四个scatter 3 五个scatter 4 scatter随机 5 全部神秘标签
	RecordId    int

	IsSetDebug bool // 是否设置debug数据

	RatioConfirm int

	IndicateNum uint8 //用户 ab test 编号

	SpecialConfig string //特殊配置
}

func (s Options) IsRaise() bool {
	return s.Raise > 0
}

func WithOptionIndex(optionIndex int) Option {
	return func(o *Options) {
		o.OptionIndex = optionIndex
	}
}

func WithRecordId(recordId int) Option {
	return func(o *Options) {
		o.RecordId = recordId
	}
}

func (s Options) String() string {
	toString, err := global.Json.MarshalToString(s)
	if err != nil {
		return ""
	}
	return toString
}

type Option func(*Options)

func GetOptions(opts ...Option) *Options {
	o := &Options{}
	for _, option := range opts {
		option(o)
	}
	return o
}

func WithTest() Option {
	return func(o *Options) {
		o.IsTest = true
	}
}

func WithRaise(n int64) Option {
	return func(o *Options) {
		o.Raise = n
	}
}

func WithReSpin() Option {
	return func(o *Options) {
		o.IsReSpin = true
	}
}

func WithFreeSpin() Option {
	return func(o *Options) {
		o.IsFree = true
	}
}

func WithTagsLock(tags []*base.Tag) Option {
	return func(o *Options) {
		o.TagsLock = make([]*base.Tag, 0)
		for _, tag := range tags {
			o.TagsLock = append(o.TagsLock, &base.Tag{
				Id:       tag.Id,
				Name:     tag.Name,
				X:        tag.X,
				Y:        tag.Y,
				Multiple: tag.Multiple,
				IsLock:   tag.IsLock,
				//IsPayTable: true,
			})
		}
	}
}

func WithIsMustFree() Option {
	return func(o *Options) {
		o.IsMustFree = true
	}
}

func WithIsMustRes() Option {
	return func(o *Options) {
		o.IsMustRes = true
	}
}

func WithDemo() Option {
	return func(o *Options) {
		o.Demo = true
	}
}

func SetFreeNum(num int) Option {
	return func(o *Options) {
		o.FreeNum = num
	}
}

func SetResNum(num int) Option {
	return func(o *Options) {
		o.ResNum = num
	}
}

func SetDebugConfig(userId uint) Option {
	return func(o *Options) {
		s := o.Spin
		s.SetDebugInitData(userId)
		o.UserId = userId
	}
}

// ConfirmRatio 设置返还比选择
func ConfirmRatio(ratio int) Option {
	return func(o *Options) {
		o.RatioConfirm = ratio
	}
}

// SetIndicateNum 设置用户 ab test 编号
func SetIndicateNum(indicateNum uint8) Option {
	return func(o *Options) {
		o.IndicateNum = indicateNum
	}
}

// SetSpecialConfig 设置特殊配置
func SetSpecialConfig(specialConfig string) Option {
	return func(o *Options) {
		o.SpecialConfig = specialConfig
	}
}
