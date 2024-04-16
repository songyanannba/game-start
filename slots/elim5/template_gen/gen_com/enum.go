package gen_com

// Status
const (
	Ok          = 0
	DownTrim    = 1
	UpTrim      = 2
	MaxRange    = 3
	UpScatter   = 4
	DownScatter = 5

	ErrorCode = 100
)

const (
	GainRatioCond  = "gainRatioCond"
	RemoveRateCond = "removeRateCond"
	WinRateCond    = "winRateCond"
	ScaTriggerCond = "scaTriggerCond"
)

const (
	CompareOk = iota //结果可以
	CompareDown
	CompareUp
)
