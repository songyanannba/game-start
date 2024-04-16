package gen_com

import "fmt"

type RunResult struct {
	Gain           int
	ScatterTrigger int
	SpinCount      int
}

type LogicResult struct {
	GainRatio    float64 `json:"返还比"`
	ScatterRatio float64 `json:"Sca触发率"`
	WinRatio     float64 `json:"中奖率"`
	RemoveRate   float64 `json:"去除率"`
	Status       int     `json:"_"`
}

func (l *LogicResult) String() string {
	return fmt.Sprintf("返还比:%.6f,"+
		"Sca触发率:%.6f,"+
		"中奖率:%.6f,"+
		"消除率:%.6f", l.GainRatio, l.ScatterRatio, l.WinRatio, l.RemoveRate)
}
