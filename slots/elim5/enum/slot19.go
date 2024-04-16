package enum

const (
	Unit19TriggerReSpin     = iota + 1 //成功触发特殊模式
	Unit19TriggerReSpinFail            //触发特殊模式失败
	Unit19TriggerReSpinNo              //未触发特殊模式
)
const (
	Unit19SpecialPlayWeight = iota //特殊玩法触发权重
	Unit19ChoiceTagWeight          //选择标签权重
	Unit19ChoiceTagCountH1         //选择标签数量
	Unit19ChoiceTagCountH2         //选择标签数量
	Unit19ChoiceTagCountH3         //选择标签数量
	Unit19ChoiceTagCountL1         //选择标签数量
	Unit19ChoiceTagCountL2         //选择标签数量
	Unit19ChoiceTagCountL3         //选择标签数量
	Unit19ChangeTheMeter           //特殊玩法触发权重
)

var Unit19Which = map[string]int{
	"":       1,
	"high_1": 2,
	"high_2": 3,
	"high_3": 4,
	"low_1":  5,
	"low_2":  6,
	"low_3":  7,
}

var Unit19WhichMap = map[int]string{
	1: "",
	2: "high_1",
	3: "high_2",
	4: "high_3",
	5: "low_1",
	6: "low_2",
	7: "low_3",
}

var Unit19TagCountWeight = map[string]int{
	"high_1": Unit19ChoiceTagCountH1,
	"high_2": Unit19ChoiceTagCountH2,
	"high_3": Unit19ChoiceTagCountH3,
	"low_1":  Unit19ChoiceTagCountL1,
	"low_2":  Unit19ChoiceTagCountL2,
	"low_3":  Unit19ChoiceTagCountL3,
}
