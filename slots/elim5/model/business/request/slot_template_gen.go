package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type SlotTemplateGenSearch struct {
	business.SlotTemplateGen
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type MergerIncrease struct {
	Ids            []int      `json:"ids" form:"ids"`
	Extra          float64    `json:"extra" form:"extra"`
	Specify        string     `json:"specify" form:"specify"`
	EliminateType  int        `json:"eliminateType" form:"eliminateType"`
	TriggerLineNum int        `json:"triggerLineNum" form:"triggerLineNum"` // 触发的线数
	TriggerSca     int        `json:"triggerSca" form:"triggerSca"`         // 触发的scatter个数
	ExtraTags      []ExtraTag `json:"extraTags" form:"extraTags"`
	TagAdditional  map[string]int
}

type ExtraTag struct {
	Name  string `json:"name" form:"name"`
	Count int    `json:"count" form:"count"`
	Max   int    `json:"max" form:"max"`
}

type SpecifyTag struct {
	Tag   string `json:"tag" form:"tag"`
	Count int    `json:"count" form:"count"`
}
