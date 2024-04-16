package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type SlotFileUploadAndDownloadSearch struct {
	business.SlotFileUploadAndDownload
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type GameFileSearch struct {
	business.GameFile
	IsEmpty        bool       `json:"isEmpty" form:"isEmpty"`
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
