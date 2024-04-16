// 自动生成模板Slot_spin
package business

import (
	"elim5/global"
)

// SlotSpin 结构体
type SlotSpin struct {
	global.GVA_MODEL
	UserId         int    `json:"user_id" form:"user_id" gorm:"column:user_id;default:0;comment:用户编号;size:32;"`
	SlotId         int    `json:"slot_id" form:"slot_id" gorm:"column:slot_id;default:0;comment:机器编号;size:32;"`
	Dir            int    `json:"dir" form:"dir" gorm:"column:dir;default:0;comment:方向;size:32;"`
	Step           int    `json:"step" form:"step" gorm:"column:step;default:0;comment:移动几次;size:32;"`
	Bet            int    `json:"bet" form:"bet" gorm:"column:bet;default:0;comment:基础押注;size:32;"`
	Type           int    `json:"type" form:"type" gorm:"column:type;default:0;comment:类型 1:普通砖;2:free_spin;size:32;"`
	Trigger        int    `json:"trigger" form:"trigger" gorm:"column:trigger;default:0;comment:是否是触发free_spin 1 是;size:32;"`
	ProgressStatus int8   `json:"progress_status" form:"progress_status" gorm:"column:progress_status;default:0;comment:状态 1 未完成 2 完成 （长标签的进度）;size:32;"`
	Layout         string `json:"layout" form:"layout" gorm:"column:layout;comment:窗口排布;"`
}

// TableName SlotSpin 表名
func (SlotSpin) TableName() string {
	return "b_slot_spin"
}
