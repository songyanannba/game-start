package business

import (
	"elim5/enum"
	"elim5/global"
	"elim5/pbs/common"
	"elim5/utils/helper"
)

// Slot 结构体
type Slot struct {
	global.GVA_MODEL
	Name         string         `json:"name" form:"name" gorm:"column:name;comment:名称;size:100;"`
	NamePkg      string         `json:"namePkg" form:"namePkg" gorm:"column:name_pkg;comment:分包名称;size:100;"`
	Icon         string         `json:"icon" form:"icon" gorm:"column:icon;comment:使用的标签图集;size:255;"`
	IconShadow   string         `json:"iconShadow" form:"iconShadow" gorm:"column:icon_shadow;comment:使用的模糊标签图集;size:255;"`
	PaylineNo    int            `json:"paylineNo" form:"paylineNo" gorm:"column:payline_no;comment:线数编号;default:0;size:32;"`
	BigWin       string         `json:"bigWin" form:"bigWin" gorm:"column:big_win;comment:赢钱最大区间;size:255;"`
	BetNum       string         `json:"betNum" form:"betNum" gorm:"type:text;column:bet_num;comment:机器押注;"`
	Raise        float64        `json:"raise" form:"raise" gorm:"column:raise;comment:加注;type:decimal(14,2);default:0;"`
	BuyFreeSpin  float64        `json:"buyFreeSpin" form:"buyFreeSpin" gorm:"column:buy_free_spin;comment:购买免费旋转;type:decimal(14,2);default:0;"`
	BuyReSpin    float64        `json:"buyReSpin" form:"buyReSpin" gorm:"column:buy_re_spin;comment:购买重旋转;type:decimal(14,2);default:0;"`
	JackpotRule  string         `json:"jackpotRule" form:"jackpotRule" gorm:"column:jackpot_rule;comment:奖池规则;size:255;"`
	Status       uint8          `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
	Url          string         `json:"url" form:"url" gorm:"column:url;comment:游戏地址;size:255;"`
	TopMul       int            `json:"topMul" form:"topMul" gorm:"column:top_mul;comment:最高倍数;default:0;size:32;"`
	ClientConf   string         `json:"clientConf" form:"clientConf" gorm:"column:client_conf;comment:客户端配置;type:text;"`
	SortLevel    int            `json:"sortLevel" form:"sortLevel" gorm:"column:sort_level;comment:排序等级;default:0;size:32;"`
	Label        string         `json:"label" form:"label" gorm:"column:label;comment:标签;size:50;"`
	PlatformType int            `json:"platformType" form:"platformType" gorm:"column:platform_type;comment:平台类型;default:1;size:32;"`
	FileList     []*common.File `json:"fileList" form:"fileList" gorm:"-"` // 文件列表
	OpenABTest   uint8          `json:"openABTest" form:"openABTest" gorm:"column:open_ab_test;comment:ab测试 1 开启;"`
}

// TableName Slot 表名
func (Slot) TableName() string {
	return "b_slot"
}

func GetGameDomainByPlatformType(t int) string {
	switch t {
	case 1:
		return global.GVA_CONFIG.System.GameDomain
	case 2:
		return global.GVA_CONFIG.System.GamePPDomain
	case 3:
		return global.GVA_CONFIG.System.GamePGDomain
	}
	return global.GVA_CONFIG.System.GameDomain
}

// GetAbTestIds 获取ab test的id
func GetAbTestIds(slotId uint) []uint {
	var ids []uint
	if slotId <= enum.AbTestMinSlotId {
		minID := enum.IndicateNumPrefix + helper.Itoa(slotId) + enum.IndicateNumStart
		maxID := enum.IndicateNumPrefix + helper.Itoa(slotId) + enum.IndicateNumEnd
		err := global.GVA_DB.Model(&Slot{}).Where("id >= ? and id <= ?", minID, maxID).Pluck("id", &ids).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
		}
	}
	return ids
}

// EncodeSlotId 根据原机台ID 和 AB测编号 和 机台总数 获取AB测的机台ID （100 + 1 + 01）
func EncodeSlotId(id uint, indicateNum uint8, slotNum int) uint {
	// 编号与机台数量取余
	indicateNum %= uint8(slotNum)
	if indicateNum == 0 {
		indicateNum = uint8(slotNum)
	}
	indicateNumStr := helper.Itoa(indicateNum)
	// 随机编号如果小于10 则补0
	if len(indicateNumStr) < enum.IndicateNumLen {
		indicateNumStr = enum.SlotFillPrefix + indicateNumStr
	}
	// 前缀 100 + 机台ID + 指定编号
	return uint(helper.Atoi(enum.IndicateNumPrefix + helper.Itoa(id) + indicateNumStr))
}

// DecodeSlotId 根据AB测机台ID获取真实的原机台ID
func DecodeSlotId(AbId uint) (slotId uint, err error) {
	AbIdStr := helper.Itoa(AbId)
	if len(AbIdStr) < enum.ABTestLen {
		return 0, enum.ABSlotIdNotFound
	}
	return uint(helper.Atoi(AbIdStr[len(enum.IndicateNumPrefix) : len(AbIdStr)-enum.IndicateNumLen])), nil
}

// GetIndicateNum 通过AB测试的ID 获取AB测试的随机编号
func GetIndicateNum(AbId uint) uint8 {
	newAbId := helper.Itoa(AbId)
	if len(newAbId) < enum.ABTestLen {
		return 0
	}
	// 截取最后两位
	return uint8(helper.Atoi(newAbId[len(newAbId)-enum.IndicateNumLen:]))
}
