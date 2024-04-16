package enum

const (
	Slot22FreeLineNum = 5
)

const (
	Slot22RemoveReasonType1 = 0 // 普通消除
	Slot22RemoveReasonType2 = 1 // 转换寿命
)

var Slot22RemoveReasonMap = map[int]string{
	Slot22RemoveReasonType1: "普通消除",
	Slot22RemoveReasonType2: "转换寿命",
}
