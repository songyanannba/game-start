package enum

// event
const (
	Slot5NormalInterval = iota      //第五台普通转区间
	Slot5NormalSca      = iota + 15 //第五台普通转出sca个数
	Slot5NormalTNum                 //scatter出现个数@（3-7个触发的Freespin次数）（普通转）
	Slot5FreeInterval               //第五台免费转区间
	Slot5FreeX2                     // 第五台免费转x2权重
	Slot5FreeSca                    //第五台免费转转出sca个数
	Slot5FreeTNum                   //scatter出现个数@（3-7个触发的Freespin次数）（免费转）
)

const (
	Slot5TriggerFreeNum = 3 //第五台普通转区间
	Slot5FreeLineNum    = 5 //第五台普通转出sca个数
)

const (
	Slot5EventFreeX2 = iota //第五台免费转区间
)
