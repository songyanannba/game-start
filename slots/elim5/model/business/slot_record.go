package business

import (
	"elim5/global"
	"elim5/utils"
	"elim5/utils/helper"
	"time"
)

// SlotRecord 结构体
type SlotRecord struct {
	global.GVA_MODEL
	Date       uint   `json:"date" form:"date" gorm:"index;column:date;default:0;comment:日期;size:32;"`
	MerchantId uint   `json:"merchantId" form:"merchantId" gorm:"index;column:merchant_id;default:0;comment:商户编号;size:32;"`
	SetId      uint   `json:"setId" form:"setId" gorm:"index;column:set_id;default:0;comment:集合编号;size:32;"`
	UserId     uint   `json:"userId" form:"userId" gorm:"index;column:user_id;default:0;comment:用户编号;size:64;"`
	SlotId     uint   `json:"slotId" form:"slotId" gorm:"index;column:slot_id;default:0;comment:机器编号;size:64;"`
	TxnId      uint   `json:"txnId" form:"txnId" gorm:"index;column:txn_id;default:0;comment:交易编号;size:64;"`
	BetType    uint8  `json:"betType" form:"betType" gorm:"index;column:bet_type;default:1;comment:押注类型;size:8;"`
	Currency   string `json:"currency" form:"currency" gorm:"index;column:currency;comment:货币;size:10;"` //货币

	Bet      int   `json:"bet" form:"bet" gorm:"column:bet;default:0;comment:基础押注;size:64;"`
	Raise    int64 `json:"raise" form:"raise" gorm:"column:raise;default:0;comment:加注;size:64;"`
	TotalBet int64 `json:"totalBet" form:"totalBet" gorm:"column:total_bet;default:0;comment:总押注;size:64;"`
	Gain     int   `json:"gain" form:"gain" gorm:"column:gain;default:0;comment:赢钱金额;size:64;"`

	IsBk       int    `json:"isBk" form:"isBk" gorm:"column:is_bk;default:2;comment:是否破产;size:16;"`
	Status     uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
	PayTableId string `json:"payTableId" form:"payTableId" gorm:"column:pay_table_id;type:text;comment:赢钱组合编号;"`
	BumperType uint8  `json:"bumperType" form:"bumperType" gorm:"column:bumper_type;default:0;comment:保险杠类型;size:8;"`
	Remark     string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"` //备注
	Ack        []byte `json:"ack" form:"ack" gorm:"type:mediumblob;column:ack;comment:回复;"`   //回复

	ChangeBal int64 `json:"changeBal" form:"changeBal" gorm:"column:change_bal;comment:Change Bal;size:64;"`
	BeforeBal int64 `json:"beforeBal" form:"beforeBal" gorm:"column:before_bal;comment:Before Bal;size:64;"`
	AfterBal  int64 `json:"afterBal" form:"afterBal" gorm:"column:after_bal;comment:After Bal;size:64;"`

	JackpotId  uint    `json:"jackpotId" form:"jackpotId" gorm:"column:jackpot_id;default:0;comment:奖池编号;size:32;"`
	JackpotMul float64 `json:"jackpotMul" form:"jackpotMul" gorm:"column:jackpot_mul;type:decimal(14,2);default:0;comment:奖池倍数;"`

	AckStr string `json:"ackStr" form:"ackStr" gorm:"-"` //回复字符串
	No     string `json:"no" form:"no" gorm:"-"`         //编号

	CreatedAt time.Time `gorm:"index;type:timestamp;size:0"` // 创建时间
	Rtp       int       `json:"rtp" form:"rtp" gorm:"column:rtp;default:1;comment:RTP;size:16;"`
}

// TableName SlotRecord 表名
func (SlotRecord) TableName() string {
	return "b_slot_record"
}

type SlotRecordResult struct {
	SlotRecord
	User         UserBaseInfo `json:"user" gorm:"foreignKey:UserId"`
	PlatformType int          `json:"platformType" form:"platformType" gorm:"-"`
}

func FmtSlotRecordNo(id uint) string {

	front := utils.Base62Encode(helper.IntReverse(int(id)))

	back := utils.Base62Encode(len(front)*3 + 1)

	return front + utils.Base62Encode(int(id)+1000000000) + back
}

func ParseSlotRecordId(no string) uint {
	if len(no) < 5 {
		return 0
	}
	back := (utils.Base62Decode(no[len(no)-1:]) - 1) / 3
	if back < 0 || back > len(no)-2 {
		return 0
	}

	front := no[:back]
	no = no[back : len(no)-1]

	n := utils.Base62Decode(no)
	if n < 1000000000 {
		return 0
	}
	res := n - 1000000000
	if utils.Base62Encode(helper.IntReverse(res)) != front {
		return 0
	}

	return uint(res)
}
