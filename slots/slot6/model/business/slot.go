// 自动生成模板Slot
package business

import "slot6/global"

// Slot 结构体
type Slot struct {
	global.GVA_MODEL
	Name                          string  `json:"name" form:"name" gorm:"column:name;comment:名称;size:100;"`
	Icon                          string  `json:"icon" form:"icon" gorm:"column:icon;comment:使用的标签图集;size:255;"`
	IconShadow                    string  `json:"iconShadow" form:"iconShadow" gorm:"column:icon_shadow;comment:使用的模糊标签图集;size:255;"`
	PaylineNo                     int     `json:"paylineNo" form:"paylineNo" gorm:"column:payline_no;comment:线数编号;default:0;size:32;"`
	BigWin                        string  `json:"bigWin" form:"bigWin" gorm:"column:big_win;comment:赢钱最大区间;size:255;"`
	BetNum                        string  `json:"betNum" form:"betNum" gorm:"type:text;column:bet_num;comment:机器押注;"`
	Raise                         float64 `json:"raise" form:"raise" gorm:"column:raise;comment:加注;type:decimal(14,2);default:0;"`
	BuyFreeSpin                   float64 `json:"buyFreeSpin" form:"buyFreeSpin" gorm:"column:buy_free_spin;comment:购买免费旋转;type:decimal(14,2);default:0;"`
	BuyReSpin                     float64 `json:"buyReSpin" form:"buyReSpin" gorm:"column:buy_re_spin;comment:购买重旋转;type:decimal(14,2);default:0;"`
	IsDrawAllPayline              uint8   `json:"isDrawAllPayline" form:"isDrawAllPayline" gorm:"column:is_draw_all_payline;type:tinyint;comment:是否划所有线;size:8;"`
	DrawAllPaylineTime            int     `json:"drawAllPaylineTime" form:"drawAllPaylineTime" gorm:"column:draw_all_payline_time;comment:划所有线时间;default:0;size:32;"`
	DrawAllPaylineTimeInterval    int     `json:"drawAllPaylineTimeInterval" form:"drawAllPaylineTimeInterval" gorm:"column:draw_all_payline_time_interval;comment:划所有线间隔;default:0;size:32;"`
	DrawAllPaylineCircle          int     `json:"drawAllPaylineCircle" form:"drawAllPaylineCircle" gorm:"column:draw_all_payline_circle;comment:划所有线循环几次;default:0;size:32;"`
	IsDrawSinglePayline           uint8   `json:"isDrawSinglePayline" form:"isDrawSinglePayline" gorm:"column:is_draw_single_payline;type:tinyint;comment:是否划单条线;size:8;"`
	DrawSinglePaylineTime         int     `json:"drawSinglePaylineTime" form:"drawSinglePaylineTime" gorm:"column:draw_single_payline_time;comment:每条线划线时间;default:0;size:32;"`
	DrawSinglePaylineTimeInterval int     `json:"drawSinglePaylineTimeInterval" form:"drawSinglePaylineTimeInterval" gorm:"column:draw_single_payline_time_interval;comment:每条线划线间隔;default:0;size:32;"`
	JackpotRule                   string  `json:"jackpotRule" form:"jackpotRule" gorm:"column:jackpot_rule;comment:奖池规则;size:255;"`
	Status                        uint8   `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
	Url                           string  `json:"url" form:"url" gorm:"column:url;comment:游戏地址;size:255;"`
	TopMul                        int     `json:"topMul" form:"topMul" gorm:"column:top_mul;comment:最高倍数;default:0;size:32;"`
}

// TableName Slot 表名
func (Slot) TableName() string {
	return "b_slot"
}
