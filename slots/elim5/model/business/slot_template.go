// 自动生成模板SlotTemplate
package business

import (
	"elim5/global"
)

// SlotTemplate 结构体
type SlotTemplate struct {
	global.GVA_MODEL
	SlotId    int    `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:游戏编号;size:32;"`
	Type      uint8  `json:"type" form:"type" gorm:"column:type;default:1;comment:类型;size:8;"`
	Column    int    `json:"column" form:"column" gorm:"column:column;default:0;comment:列号;size:32;"`
	Layout    string `json:"layout" form:"layout" gorm:"column:layout;type:mediumtext;comment:排布;"`
	GenId     int    `json:"genId" form:"genId" gorm:"column:gen_id;default:0;comment:生成编号;size:32;"`
	Lock      uint8  `json:"lock" form:"lock" gorm:"column:lock;default:0;comment:锁定;size:8;"`
	MergeInfo string `json:"mergeInfo" form:"mergeInfo" gorm:"column:merge_info;type:text;comment:合并信息;"`
	Rtp       int    `json:"rtp" form:"rtp" gorm:"column:rtp;default:1;comment:RTP;size:16;"`
	Which     int    `json:"which" form:"which" gorm:"column:which;default:0;comment:换表序号;size:32;"`
}

// TableName SlotTemplate 表名
func (SlotTemplate) TableName() string {
	return "b_slot_template"
}

type SlotTemplateExcel struct {
	Title string              `json:"title" form:"title" `
	Datas map[string][]string `json:"datas" form:"datas"`
}
