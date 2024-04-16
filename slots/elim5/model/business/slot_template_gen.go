// 自动生成模板SlotTemplateGen
package business

import (
	"elim5/enum"
	"elim5/global"
	"elim5/utils"
	"go.uber.org/zap"
	"strings"
)

// SlotTemplateGen 结构体
type SlotTemplateGen struct {
	global.GVA_MODEL
	SlotId        int            `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:游戏编号;size:32;"`
	Type          uint8          `json:"type" form:"type" gorm:"column:type;default:1;comment:类型;size:8;"`
	MinRatio      float64        `json:"minRatio" form:"minRatio" gorm:"column:min_ratio;default:0;type:decimal(14,6);comment:最小返奖率;size:32;"`
	MaxRatio      float64        `json:"maxRatio" form:"maxRatio" gorm:"column:max_ratio;default:0;type:decimal(14,6);comment:最大返奖率;size:32;"`
	MinScatter    float64        `json:"minScatter" form:"minScatter" gorm:"column:min_scatter;type:decimal(14,6);default:0;comment:最小消散符号个数;size:32;"`
	MaxScatter    float64        `json:"maxScatter" form:"maxScatter" gorm:"column:max_scatter;type:decimal(14,6);default:0;comment:最大消散符号个数;size:32;"`
	OtherCond     string         `json:"otherCond" form:"otherCond" gorm:"column:other_cond;type:text;comment:其他条件;"`
	InitialWeight string         `json:"initialWeight" form:"initialWeight" gorm:"column:initial_weight;type:text;comment:初始权重;"`
	SpecialConfig string         `json:"specialConfig" form:"specialConfig" gorm:"column:special_config;type:text;comment:特殊配置;"`
	LargeScale    string         `json:"largeScale" form:"largeScale" gorm:"column:large_scale;type:text;comment:大范围调整;"`
	TrimDown      string         `json:"trimDown" form:"trimDown" gorm:"column:trim_down;type:text;comment:向下微调;"`
	TrimUp        string         `json:"trimUp" form:"trimUp" gorm:"column:trim_up;type:text;comment:向上微调;"`
	FinalWeight   string         `json:"finalWeight" form:"finalWeight" gorm:"column:final_weight;type:text;comment:最终权重;"`
	Template      string         `json:"template" form:"template" gorm:"column:template;comment:模版;type:mediumtext;"`
	Schedule      string         `json:"schedule" form:"schedule" gorm:"column:schedule;type:mediumtext;comment:进度;"`
	Remarks       string         `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注;size:500;"`
	State         uint8          `json:"state" form:"state" gorm:"column:state;default:0;comment:状态;size:8;"`
	Lock          uint8          `json:"lock" form:"lock" gorm:"column:lock;default:0;comment:锁定;size:8;"`
	Count         int            `json:"count" form:"count" gorm:"column:count;default:0;comment:转动次数;size:32;"`
	Reset         int            `json:"reset" form:"reset" gorm:"column:reset;default:0;comment:重置次数;size:32;"`
	Rtp           int            `json:"rtp" form:"rtp" gorm:"column:rtp;default:1;comment:RTP;size:16;"`
	SpecialWeight map[int]string `json:"specialWeight" form:"specialWeight" gorm:"-"`
	Which         int            `json:"which" form:"which" gorm:"column:which;default:0;comment:换表序号;size:32;"`
}

// TableName SlotTemplateGen 表名
func (*SlotTemplateGen) TableName() string {
	return "b_slot_template_gen"
}

// Start 开始
func (t *SlotTemplateGen) Start() {
	t.Remarks = ""
	t.Schedule = ""
	t.Template = ""
	t.State = enum.CommonStatusBegin
	if t.Rtp == 0 {
		t.Rtp = 1
	}
	global.NOLOG_DB.Save(t)
}

// Stop 停止
func (t *SlotTemplateGen) Stop() {
	t.State = enum.CommonStatusClose
	global.NOLOG_DB.Save(t)
}

// Finish 结束
func (t *SlotTemplateGen) Finish() {
	t.State = enum.CommonStatusFinish
	err := global.NOLOG_DB.Save(t).Error
	if err != nil {
		global.GVA_LOG.Info("err", zap.Error(err))
	}
}

// WriteProgress 写入进度
func (t *SlotTemplateGen) WriteProgress(str string) {
	t.Schedule = str + t.Schedule
	t.State = enum.CommonStatusProcessing
	global.NOLOG_DB.Save(t)
}

func (t *SlotTemplateGen) SetSpecialWeight() {
	strs := utils.FormatCommand(t.SpecialConfig)
	t.SpecialWeight = map[int]string{}
	for _, str := range strs {
		if str == "" {
			continue
		}
		infos := strings.Split(str, ":")
		if len(infos) != 2 {
			global.GVA_LOG.Error("special weight format error")
			t.SpecialWeight = nil
			return
		}
		weightName := infos[0]
		weightStr := infos[1]
		switch weightName {
		case enum.MulTagWeight:
			t.SpecialWeight[enum.MulTagWeightId] = weightStr
		case enum.FillWeight:
			t.SpecialWeight[enum.FillWeightId] = weightStr
		}
	}
}

func (t *SlotTemplateGen) GetWeight(id int) string {
	if t.SpecialWeight == nil {
		return ""
	}
	if v, ok := t.SpecialWeight[id]; ok {
		return v
	}
	return ""
}
