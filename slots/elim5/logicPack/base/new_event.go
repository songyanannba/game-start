package base

import (
	"elim5/enum"
	"elim5/global"
	"elim5/model/business"
	"elim5/utils"
	"elim5/utils/helper"
	"strconv"
	"strings"
)

// Unit5NewEvent 第五台特殊事件解析
func (e *Event) Unit5NewEvent(events []*business.SlotEvent) {
	for i, event := range events {
		if i == 0 || i == 18 {
			e.m[i] = GetIntervalRatioEvent(event.Event1)
		} else {
			e.Add(i, event.Event1)
		}
	}
}

// Unit6NewEvent 第六台特殊事件解析
func (e *Event) Unit6NewEvent(events []*business.SlotEvent) {
	for i, event := range events {
		e.m[i] = Unit6ParseWeightDataStr(event.Event1)
	}

	//for i, event := range events {
	//	if i == 0 {
	//		e.m[i] = GetIntervalRatioEvent(event.Event1)
	//	} else {
	//		switch i {
	//		case 16, 17, 18, 19, 20:
	//			e.m[i] = Unit6ParseWeightDataStr(event.Event1)
	//		default:
	//			e.Add(i, event.Event1)
	//		}
	//
	//	}
	//}
}

func Unit6ParseWeightDataStr(str string) *Unit6LevelEvent {
	result := &Unit6LevelEvent{}
	// 按@符号拆分成两个字符串
	split := strings.Split(str, "@")
	if len(split) != 6 {
		global.GVA_LOG.Info("输入字符串格式不正确")
		return nil
	}
	result.Collect = helper.Atoi(split[0])
	result.CoreCount = helper.Atoi(split[1])
	result.EmitEvent = ParseWeightData(split[2] + "@" + split[3])
	result.WildEvent = ParseWeightData(split[4] + "@" + split[5])
	return result
}

func GetIntervalRatioEvent(str string) *IntervalRatioEvent {
	// 按@符号拆分成两个字符串
	split := strings.Split(str, "@")
	if len(split) != 2 {
		global.GVA_LOG.Info("输入字符串格式不正确")
		return nil
	}

	arrStr := strings.Split(split[0], "&")
	var arr [][2]int

	for _, str := range arrStr {
		// 去除首尾的方括号
		str = strings.Trim(str, "[]")
		// 按逗号拆分
		values := strings.Split(str, ",")
		if len(values) != 2 {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		// 将字符串转换为整数
		start, err := strconv.Atoi(values[0])
		if err != nil {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		end, err := strconv.Atoi(values[1])
		if err != nil {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		// 添加到二维数组
		arr = append(arr, [2]int{start, end})
	}
	// 拆分后面的字符串按&符号
	numStr := strings.Split(split[1], "&")
	var nums []int

	// 遍历拆分的字符串进行转换
	for _, str := range numStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		// 添加到数组
		nums = append(nums, num)
	}
	return &IntervalRatioEvent{
		Data:   arr,
		weight: nums,
	}
}

func GetIntervalRatioEventTem(str string) *IntervalRatioEvent {
	// 按@符号拆分成两个字符串
	split := strings.Split(str, "@")
	if len(split) != 2 {
		global.GVA_LOG.Info("输入字符串格式不正确")
		return nil
	}

	arrStr := strings.Split(split[0], "&")
	var arr [][2]int

	for _, str := range arrStr {
		// 去除首尾的方括号
		str = strings.Trim(str, "[]")
		// 按逗号拆分
		values := strings.Split(str, ",")
		if len(values) != 2 {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		// 将字符串转换为整数
		start, err := strconv.Atoi(values[0])
		if err != nil {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		end, err := strconv.Atoi(values[1])
		if err != nil {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		// 添加到二维数组
		arr = append(arr, [2]int{start, end})
	}
	// 拆分后面的字符串按&符号
	numStr := strings.Split(split[1], "&")
	var nums []int

	// 遍历拆分的字符串进行转换
	for _, str := range numStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		// 添加到数组
		nums = append(nums, num)
	}
	return &IntervalRatioEvent{
		Data:   arr,
		weight: nums,
	}
}

// Unit7NewEvent 第七台特殊事件解析
func (e *Event) Unit7NewEvent(events []*business.SlotEvent) {
	for i, event := range events {
		if i == enum.Slot7EventLevel {
			e.m[i] = ParseEventLevel(event.Event1)
		} else {
			e.Add(i, event.Event1)
		}
	}
}

// Unit8NewEvent 第八台特殊事件解析
func (e *Event) Unit8NewEvent(events []*business.SlotEvent) {
	for i, event := range events {
		e.m[i] = ParseWeightData(event.Event1)
	}
}

func (e *Event) AddEventJoinLayout(typ int, params string) {
	if params == "" {
		return
	}

	var (
		tagStr                 = ""
		weightStr              = ""
		configurationStr       = ""
		weightAndConfiguration = ""
		weight                 = []int{}
		configuration          = []int{}
	)

	params = utils.FormatCommandStr(params)

	tagStr, weightAndConfiguration, _ = strings.Cut(params, "@")
	tags := helper.SplitInt[int](tagStr, "&")

	if strings.Contains(weightAndConfiguration, "@") {
		weightStr, configurationStr, _ = strings.Cut(weightAndConfiguration, "@")
		weight = helper.SplitInt[int](weightStr, "&")
		configuration = helper.SplitInt[int](configurationStr, "&")
	} else {
		weight = helper.SplitInt[int](weightAndConfiguration, "&")
	}

	if len(weight)-1 != len(tags) {
		return
	}

	result := &Unit51LevelEvent{
		Data:   tags,
		Weight: weight,
		MutArr: configuration,
	}

	e.m[typ] = result
}
