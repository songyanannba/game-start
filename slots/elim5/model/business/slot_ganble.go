// 自动生成模板Ganble
package business

import (
	"elim5/global"
)

// Gamble 结构体
type Gamble struct {
	global.GVA_MODEL
	SlotId   int  `json:"slot_id" form:"slot_id" gorm:"column:slot_id;default:0;comment:机器编号;size:32;"`
	RecordId int  `json:"record_id" form:"record_id" gorm:"column:record_id;default:0;comment:spin编号;size:32;"`
	Gamble   int  `json:"gamble" form:"gamble" gorm:"column:gamble;default:0;comment:gamble倍数;size:32;"`
	Status   uint `json:"status" form:"status" gorm:"column:status;default:0;comment:状态;size:32;"`
}

// TableName Gamble 表名
func (Gamble) TableName() string {
	return "b_slot_gamble"
}
