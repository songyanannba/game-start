package template_gen

import (
	"elim5/logicPack/component"
	"elim5/service/template_gen/gen_com"
	"elim5/service/template_gen/unit"
	"elim5/utils/helper"
	"fmt"
)

func GoRun(t *gen_com.GenTemplate) (totalInterval float64, str string, status bool, err error) {
	var unitFace unit.SlotFace
	switch t.TemGen.SlotId {
	case 8, 15:
		unitFace = unit.GetUnit8(t)
	case 17:
		unitFace = unit.GetUnit17(t)
	case 9:
		unitFace = unit.GetUnit9(t)
	case 5:
		unitFace = unit.GetUnit5(t)
	case 6:
		unitFace = unit.GetUnit6(t)
	case 7:
		unitFace = unit.GetUnit7(t)
	case 14:
		unitFace = unit.GetUnit14(t)
	//case 15:
	//	unitFace = unit.GetUnit15(t)
	case 16:
		unitFace = unit.GetUnit16(t)
	case 22:
		unitFace = unit.GetUnit22(t)
	case 25:
		unitFace = unit.GetUnit25(t)
	case 33:
		unitFace = unit.GetUnit33(t)
	case 43:
		unitFace = unit.GetUnit43(t)
	case 44:
		unitFace = unit.GetUnit44(t)
	case 49:
		unitFace = unit.GetUnit49(t)
	case 50:
		unitFace = unit.GetUnit50(t)
	case 52:
		unitFace = unit.GetUnit52(t)
	case 47:
		unitFace = unit.GetUnit47(t)
	case 48:
		unitFace = unit.GetUnit48(t)
	case 53:
		unitFace = unit.GetUnit53(t)
	case 54:
		unitFace = unit.GetUnit54(t)

	default:
		err = fmt.Errorf("未知的机台")
		return
	}
	for _, tags := range t.Template {
		if len(t.Template[0]) != len(tags) {
			err = fmt.Errorf("模版长度不一致")
			return 0, err.Error(), false, err
		}
	}
	ch, errCh, _ := helper.Parallel[[]*component.Spin](t.TemGen.Count, 1000, func() (spins []*component.Spin, err error) {
		spins, err = unitFace.RunTem()
		return
	})

	a := 0
	for {
		a++
		select {
		case err = <-errCh:
			return 0, err.Error(), false, err
		case v, beforeClosed := <-ch:

			if !beforeClosed {
				goto end
			}
			unitFace.Calculate(v)
		}
	}
end:
	totalInterval, str, status = unitFace.GetStatus()
	return
}

func GoRunTest(t *gen_com.GenTemplate) (str string, status bool, err error) {
	var unitFace unit.SlotFace

	switch t.Config.SlotId {
	case 8, 15:
		unitFace = unit.GetUnit8(t)
	case 17:
		unitFace = unit.GetUnit17(t)
	case 9:
		unitFace = unit.GetUnit9(t)
	case 5:
		unitFace = unit.GetUnit5(t)
	case 6:
		unitFace = unit.GetUnit6(t)
	}

	result, _ := unitFace.RunTem() // 9 GoRunTest
	unitFace.Calculate(result)     //9

	result1, _ := unitFace.RunTem()
	unitFace.Calculate(result1)

	_, str, status = unitFace.GetStatus() //9 test

	return str, status, err
}
