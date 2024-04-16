// 自动生成模板IpWhite
package business

import (
	"elim5/global"
)

// IpWhite 结构体
type IpWhite struct {
	global.GVA_MODEL
	MerchantId int    `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;default:0;comment:商户ID;size:32;"`
	Ip         string `json:"ip" form:"ip" gorm:"column:ip;comment:IP;size:20;"`
	Remark     string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
	Status     uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
}

// TableName IpWhite 表名
func (IpWhite) TableName() string {
	return "b_ip_white"
}
