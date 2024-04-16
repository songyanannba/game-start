// 自动生成模板BlackList
package business

import (
	"elim5/global"
)

// BlackList 结构体
type BlackList struct {
	global.GVA_MODEL
	Country string `json:"country" form:"country" gorm:"column:country;comment:国家;size:50;"`
	Status  uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
	Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
}

// TableName BlackList 表名
func (BlackList) TableName() string {
	return "b_black_list"
}
