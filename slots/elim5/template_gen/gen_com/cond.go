package gen_com

import "elim5/utils/helper"

type Cond struct {
	Name     string
	MinCond  float64
	MaxCond  float64
	UpFunc   func(t *GenTemplate)
	DownFunc func(t *GenTemplate)
	Count    int
}

func NewCond(str string, min, max float64) *Cond {
	switch str {
	case GainRatioCond:
		return &Cond{
			Name:     str,
			MinCond:  min,
			MaxCond:  max,
			UpFunc:   TemAdjTrimUp,
			DownFunc: TemAdjTrimDown,
			Count:    3,
		}
	case RemoveRateCond:
		return &Cond{
			Name:     str,
			MinCond:  min,
			MaxCond:  max,
			UpFunc:   TemRemoveUp,
			DownFunc: TemRemoveDown,
			Count:    5,
		}
	case WinRateCond:
		return &Cond{
			Name:     str,
			MinCond:  min,
			MaxCond:  max,
			UpFunc:   TemRemoveUp,
			DownFunc: TemRemoveDown,
			Count:    3,
		}
	case ScaTriggerCond:
		return &Cond{
			Name:     str,
			MinCond:  min,
			MaxCond:  max,
			UpFunc:   TemAdjScatterUp,
			DownFunc: TemAdjScatterDown,
			Count:    1,
		}
	default:
		return nil
	}
}

// Adjust  f:比率 ; tem:模版
func (c *Cond) Adjust(f float64, tem *GenTemplate) string {

	//tem.AdjRecords++
	//if tem.AdjRecords/len(tem.CondMap) >= tem.TemGen.Reset && tem.TemGen.Reset > 0 {
	//	err := tem.InitTem()
	//	if err != nil {
	//		return err.Error()
	//	}
	//	tem.AdjRecords = 0
	//	return "Reset"
	//}
	//
	//if c.MaxCond == 0 {
	//	return "ok"
	//}
	//count := helper.RandInt(c.Count) + 1
	//if f < c.MinCond {
	//	for i := 0; i < count; i++ {
	//		c.UpFunc(tem)
	//	}
	//	return "up"
	//} else if f > c.MaxCond {
	//	for i := 0; i < count; i++ {
	//		c.DownFunc(tem)
	//	}
	//	return "down"
	//} else {
	//	return "ok"
	//}

	//if tem.TemGen.Reset == 1 {
	//	err := tem.InitTem()
	//	if err != nil {
	//		return err.Error()
	//	}
	//	return "Reset"
	//}
	if c.MaxCond == 0 {
		return "ok"
	}
	if f < c.MinCond {
		if tem.TemGen.Reset != 1 {
			c.UpFunc(tem)
		}
		err := tem.InitTem()
		if err != nil {
			return err.Error()
		}
		return "reset"
	} else if f > c.MaxCond {
		if tem.TemGen.Reset != 1 {
			c.DownFunc(tem)
		}
		err := tem.InitTem()
		if err != nil {
			return err.Error()
		}
		return "reset"
	} else {
		return "ok"
	}
}

func (c *Cond) GetDisparity(f float64) float64 {
	return helper.Abs(f - (c.MaxCond+c.MinCond)/2)
}

func (c *Cond) Compare(f float64) int {
	if c.MaxCond == 0 {
		return CompareOk
	}
	if f < c.MinCond {
		return CompareUp
	} else if f > c.MaxCond {
		return CompareDown
	} else {
		return CompareOk
	}
}
