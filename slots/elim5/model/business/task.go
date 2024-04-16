// 自动生成模板Task
package business

import (
	"elim5/global"
)

// Task 结构体
type Task struct {
	global.GVA_MODEL
	Name    string `json:"name" form:"name" gorm:"column:name;comment:名称;size:255;"`
	Type    uint8  `json:"type" form:"type" gorm:"column:type;default:1;comment:类型;size:8;"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:内容;"`
	Rate    int8   `json:"rate" form:"rate" gorm:"column:rate;default:0;comment:进度;size:8;"`
	Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
	Status  uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
}

// TableName Task 表名
func (Task) TableName() string {
	return "b_task"
}
