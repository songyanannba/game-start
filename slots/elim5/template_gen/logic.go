package template_gen

import (
	"elim5/model/business"
	"elim5/service/cache"
	"elim5/service/template_gen/gen_com"
	"fmt"
)

// CreateTemplate 创建模版逻辑
func CreateTemplate(tem *business.SlotTemplateGen) (err error) {
	var (
		str           string
		count         int
		status        bool
		t             *gen_com.GenTemplate
		totalInterval float64
	)
	defer func() {
		err := cache.DeleteTemplateGenCache(int(tem.ID))
		if err != nil {
			return
		}
	}()
	if err = cache.SetTemplateGenCache(int(tem.ID)); err != nil {
		return fmt.Errorf("模版正在生成中")
	}
	//写入状态
	tem.Start()
	t, err = gen_com.NewGenTemplate(tem)
	if err != nil {
		return err
	}

	for {
		if !cache.GetTemplateGenCache(int(tem.ID)) {
			tem.Stop()
			return nil
		}
		count++
		totalInterval, str, status, err = GoRun(t)
		if err != nil {
			return err
		}
		str = fmt.Sprintf("%d:%v", count, str)
		if status {
			goto end
		}
		if count >= 1000 {
			str += "超过1000次,停止"
		}
		if totalInterval <= t.Closest || t.Closest == 0 {
			t.Closest = totalInterval
			t.TemGen.FinalWeight = fmt.Sprintf("%s\n%v", str, t.GetFinalWeight(tem.InitialWeight))
		}
		tem.WriteProgress(str)
	}
end:
	tem.Schedule = str + "\n" + tem.Schedule
	tem.Template = t.GetFinalTemplate()
	tem.FinalWeight = fmt.Sprintf("%s\n%v", str, t.GetFinalWeight(tem.InitialWeight))
	tem.Lock = 1
	tem.Finish()
	t.CreateTem(tem)
	return nil
}

func TestTemplateGen(tem *business.SlotTemplateGen) error {
	t, err := gen_com.NewTestGenTemplate(tem)
	if err != nil {
		return err
	}
	totalInterval, str, status, _ := GoRun(t)
	fmt.Printf("%v %v %d", str, status, totalInterval)
	return nil
}

func TestTemplateGen9(tem *business.SlotTemplateGen) error {
	t, err := gen_com.NewGenTemplate(tem) //TestTemplateGen9
	if err != nil {
		return err
	}
	str, status, _ := GoRunTest(t)

	tem.Schedule = str + "\n" + tem.Schedule
	tem.Template = t.GetFinalTemplate()
	tem.FinalWeight = t.GetFinalWeight(tem.InitialWeight)
	tem.Lock = 1
	//tem.Finish()
	//t.CreateTem(tem)
	fmt.Printf("%v %v", str, status)
	return nil
}
