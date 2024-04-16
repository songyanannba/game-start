package gen_com

import (
	"elim5/enum"
	"elim5/global"
	"elim5/logicPack/base"
	"elim5/logicPack/component"
	"elim5/model/business"
	"elim5/utils"
	"elim5/utils/helper"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"strconv"
	"strings"
	"sync"
)

func NewGenTemplate(tem *business.SlotTemplateGen) (gen *GenTemplate, err error) {
	var (
		config *component.Config
		newMap map[int][]*Scale
	)
	config, err = component.GetSlotConfig(uint(tem.SlotId), false)
	if err != nil {
		return nil, err
	}
	gen = &GenTemplate{
		Config:        config,
		TemGen:        tem,
		InitialWeight: map[int]map[string]int{},
		LargeScale:    map[int][]*WeightInterval{},
		Interval:      []*WeightInterval{},
		TrimDown:      map[int][]*Scale{},
		TrimUp:        map[int][]*Scale{},
		Template:      map[int][]uint16{},
		ExtraTemplate: map[int][]uint16{},
		SpecialWeight: map[int]any{},
		Adjacent:      map[int]map[string]int{},
		CondMap:       map[string]*Cond{},
		AdjRecords:    0,
		Raise:         helper.If(tem.Type == enum.SpinAckType1NormalSpin, true, false),
		mu:            sync.Mutex{},
		RatioConfirm:  tem.Rtp,
		Which:         tem.Which,
	}

	err = gen.GetInitialWeight(tem.InitialWeight)
	if err != nil {
		return nil, err
	}
	err = gen.GetLargeScale(tem.LargeScale)
	if err != nil {
		return nil, err
	}
	newMap, err = gen.GetTrim(tem.TrimDown) // 向下微调
	if err != nil {
		return nil, err
	}
	gen.TrimDown = newMap
	newMap, err = gen.GetTrim(tem.TrimUp) // 向上微调
	if err != nil {
		return nil, err
	}
	gen.TrimUp = newMap
	err = gen.SetSpecialWeight(tem.SpecialConfig)
	if err != nil {
		return nil, err
	}
	err = gen.SetConds(tem.OtherCond)
	if err != nil {
		return nil, err
	}
	err = gen.InitTem()
	if err != nil {
		return nil, err
	}
	return gen, nil
}

func NewTestGenTemplate(tem *business.SlotTemplateGen) (gen *GenTemplate, err error) {
	var (
		config *component.Config
		//newMap map[int][]*Scale
	)
	config, err = component.GetSlotConfig(uint(tem.SlotId), false)
	if err != nil {
		return nil, err
	}
	gen = &GenTemplate{
		Config:        config,
		TemGen:        tem,
		InitialWeight: map[int]map[string]int{},
		LargeScale:    map[int][]*WeightInterval{},
		Interval:      []*WeightInterval{},
		TrimDown:      map[int][]*Scale{},
		TrimUp:        map[int][]*Scale{},
		Template:      map[int][]uint16{},
		SpecialWeight: map[int]any{},
		CondMap:       map[string]*Cond{},
	}
	var tems []*business.SlotTemplate
	global.NOLOG_DB.Where("gen_id = ?", tem.ID).Find(&tems)

	for _, template := range tems {
		gen.Template[template.Column] = helper.SplitInt[uint16](template.Layout, ",")
	}

	//err = gen.InitTestTem(tem.Template)
	if err != nil {
		return nil, err
	}
	return gen, nil
}

// GetColInfo 获取列号拆分
func GetColInfo(str string) (int, string, error) {

	colW := strings.Split(str, ":")
	if len(colW) != 2 {
		return 0, "", fmt.Errorf("initial weight format error")
	}
	colNum, err := strconv.Atoi(colW[0])
	if err != nil {
		return 0, "", err
	}
	return colNum, colW[1], nil
}

// GetColMap 获取列号map
func GetColMap(str string) (map[int]string, error) {
	if str == "" {
		return map[int]string{}, nil
	}
	strs := utils.FormatCommand(str)
	strMap := map[int]string{}
	for _, str := range strs {
		if str == "" {
			continue
		}
		info, s, err := GetColInfo(str)
		if err != nil {
			return nil, err
		}
		strMap[info] = s
	}
	return strMap, nil
}

// GetInterval 间隔配置
func GetInterval(str string) (map[string][2]int, error) {
	if str == "" {
		return map[string][2]int{}, nil
	}
	strMap := map[string][2]int{}
	info := strings.Split(str, ",")
	for _, s := range info {
		if s == "" {
			continue
		}
		nameInt := strings.Split(s, "=")
		if len(nameInt) != 2 {
			return nil, fmt.Errorf("interval format error")
		}
		name := nameInt[0]
		interval := strings.Split(nameInt[1], "-")
		if len(interval) != 2 {
			return nil, fmt.Errorf("interval format error")
		}
		min, err := strconv.Atoi(interval[0])
		if err != nil {
			min = 0
		}
		max, err := strconv.Atoi(interval[1])
		if err != nil {
			max = 0
		}
		if min >= max {
			return nil, fmt.Errorf("interval format error")
		}
		strMap[name] = [2]int{min, max}
	}
	return strMap, nil
}

// GetInitialWeight 初始权重
func (t *GenTemplate) GetInitialWeight(inWei string) error {
	if inWei == "" {
		return nil
	}
	t.InitialWeight = map[int]map[string]int{}
	colMap, err := GetColMap(inWei)
	if err != nil {
		return err
	}

	for col := 0; col < t.Config.Index; col++ {
		t.InitialWeight[col] = map[string]int{}
		t.Adjacent[col] = map[string]int{}
		s := ""
		if colMap[col] != "" {
			s = colMap[col]
		} else {
			s = colMap[0]
		}
		strs, ints := GetInitialWeightFunc(s)
		for i, str := range strs {
			t.InitialWeight[col][str] = ints[i+1][0] - ints[i][0]
			t.Adjacent[col][str] = ints[i+1][1]
		}
	}
	return nil
}

func GetInitialWeightFunc(inWei string) (strs []string, ints [][2]int) {
	if inWei == "" {
		return []string{}, [][2]int{}
	}
	splitInt := helper.SplitStr(inWei, "@")
	if len(splitInt) != 2 {
		return []string{}, [][2]int{}
	}
	strs = helper.SplitStr(splitInt[0], "&")
	ints = [][2]int{}
	for _, s := range helper.SplitStr(splitInt[1], "&") {
		if s == "" {
			continue
		}
		var min, max int
		var err error
		interval := helper.SplitStr(s, ",")
		if len(interval) == 1 {
			min, err = strconv.Atoi(interval[0])
			if err != nil {
				min = 0
			}
			max = 0
		} else {
			min, err = strconv.Atoi(interval[0])
			if err != nil {
				min = 0
			}
			max, err = strconv.Atoi(interval[1])
			if err != nil {
				max = 0
			}
		}
		ints = append(ints, [2]int{min, max})
	}
	if len(strs) != len(ints)-1 {
		return []string{}, [][2]int{}
	}
	return

}

// GetLargeScale 大比例调整
func (t *GenTemplate) GetLargeScale(inWei string) error {
	if inWei == "" {
		return nil
	}
	colMap, err := GetColMap(inWei)
	if err != nil {
		return err
	}

	for colNum, s := range colMap {
		var (
			tagCs []*WeightInterval
		)
		interval, err := GetInterval(s)
		if err != nil {
			return err
		}
		for name, interval := range interval {
			tagCs = append(tagCs, &WeightInterval{
				Tag:      t.Config.GetTag(name),
				MinCount: interval[0],
				MaxCount: interval[1],
			})
		}
		t.LargeScale[colNum] = tagCs
	}

	return nil
}

// GetTrim 小比例调整
func (t *GenTemplate) GetTrim(inWei string) (map[int][]*Scale, error) {
	if inWei == "" {
		return map[int][]*Scale{}, nil
	}
	newMap := map[int][]*Scale{}
	colMap, err := GetColMap(inWei) //inWei == 解析结构 列:被调整的列值->期望调整为的列值; (2:low_1=>high_1,low_2=>high_2)
	if err != nil {
		return nil, err
	}

	for i := 0; i < t.Config.Index; i++ {
		s := ""
		if colMap[i] != "" {
			s = colMap[i]
		} else {
			s = colMap[0]
		}
		var (
			tagCs []*Scale
		)
		trims := strings.Split(s, ",")
		for _, trim := range trims {
			if trim == "" {
				continue
			}
			twoTag := strings.Split(trim, "=>")
			if len(twoTag) != 2 {
				return nil, fmt.Errorf("trim format error")
			}
			fillTag := t.Config.GetTag(twoTag[0])
			if fillTag.Name == "" || fillTag.Id == -1 {
				return nil, fmt.Errorf("trim format error col:%d name:%v", i, twoTag[0])
			}
			fillReplaceTag := t.Config.GetTag(twoTag[1])
			if fillReplaceTag.Name == "" || fillReplaceTag.Id == -1 {
				return nil, fmt.Errorf("trim format error col:%d name:%v", i, twoTag[1])
			}

			tagCs = append(tagCs, &Scale{
				Tag:        t.Config.GetTag(twoTag[0]),
				ReplaceTag: t.Config.GetTag(twoTag[1]),
			})
		}
		newMap[i] = tagCs
	}
	return newMap, nil
}

// InitTem 初始化模板
func (t *GenTemplate) InitTem() error {

	group := sync.WaitGroup{}
	for col, counts := range t.InitialWeight {
		go func(col int, counts map[string]int) {
			err := t.ColHandle(col, counts)
			if err != nil {
				global.GVA_LOG.Error("colHandle error:" + err.Error())
				//panic(err)
				return
			}
			group.Done()
		}(col, counts)
		group.Add(1)
	}
	group.Wait()
	return nil
}

func (t *GenTemplate) ColHandle(col int, counts map[string]int) error {

	sumCount := 0
	for _, count := range counts {
		sumCount += count //对应列的所有的标签的总个数
	}
	allTags := make([]uint16, sumCount) //每列的标签总和
	// 填充需要间隔的标签
	for _, interval := range t.Interval {
		err := IntervalPlacement(allTags, interval.Tag, counts[interval.Tag.Name], interval.MinCount, interval.MaxCount)
		if err != nil {
			return err
		}
	}

	//填充其他标签,并打乱
	var tags []uint16
	for name, count := range counts {
		fillTag := t.Config.GetTag(name)
		if fillTag.Name == "" || fillTag.Id < 1 {
			global.GVA_LOG.Error("tag name is empty " + name)
			return errors.New("tag name is empty " + name)
		}
		_, b := lo.Find(t.Interval, func(item *WeightInterval) bool {
			return item.Tag.Name == name
		})
		if b {
			continue
		}
		for i := 0; i < count; i++ {
			tags = append(tags, uint16(fillTag.Id))
		}
	}
	helper.SliceShuffle(tags)

	emptyTags := lo.Filter(allTags, func(item uint16, i int) bool {
		return item == 0
	})

	if len(emptyTags) != len(tags) {
		return fmt.Errorf("init template_gen error" + fmt.Sprintf("all:%d , empty:%d , tags:%d", len(allTags), len(emptyTags), len(tags)))
	}

	//填充普通的标签
	TagsFill(allTags, tags)

	t.mu.Lock()
	t.Template[col] = t.InsertAdjacency(allTags, col)
	t.mu.Unlock()

	return nil
}

// SetSpecialWeight 特殊权重
func (t *GenTemplate) SetSpecialWeight(weight string) error {
	if weight == "" {
		return nil
	}
	strs := utils.FormatCommand(weight)
	for _, str := range strs {
		if str == "" {
			continue
		}
		infos := strings.Split(str, ":")
		if len(infos) != 2 {
			return fmt.Errorf("special weight format error")
		}
		weightName := infos[0]
		weightStr := infos[1]
		switch weightName {
		case enum.MulTagWeight:
			t.SpecialWeight[enum.MulTagWeightId] = base.ParseWeightData(weightStr)
		case enum.FillWeight:
			t.SpecialWeight[enum.FillWeightId] = base.ParseWeightDataStr(weightStr)
		case enum.Interval:
			interval, err := GetInterval(weightStr)
			if err != nil {
				return err
			}
			for name, interval := range interval {
				t.Interval = append(t.Interval, &WeightInterval{
					Tag:      t.Config.GetTag(name),
					MinCount: interval[0],
					MaxCount: interval[1],
				})
			}
		case enum.Adjacent:
			adjs := strings.Split(weightStr, ",")
			for i := 0; i < t.Config.Index; i++ {
				t.Adjacent[i] = map[string]int{}
				for _, adj := range adjs {
					if adj == "" {
						continue
					}
					inserts := strings.Split(adj, "=")
					if len(inserts) != 2 {
						return fmt.Errorf("adjacent format error")
					}
					num, err := strconv.Atoi(inserts[1])
					if err != nil {
						num = 0
					}
					t.Adjacent[i][inserts[0]] = num
				}
			}
		}
	}
	return nil
}

func (t *GenTemplate) SetConds(weight string) error {

	t.CondMap[GainRatioCond] = NewCond(GainRatioCond, t.TemGen.MinRatio, t.TemGen.MaxRatio)
	t.CondMap[ScaTriggerCond] = NewCond(ScaTriggerCond, t.TemGen.MinScatter, t.TemGen.MaxScatter)
	if weight == "" {
		return nil
	}
	strs := utils.FormatCommand(weight)
	for _, str := range strs {
		if str == "" {
			continue
		}
		infos := strings.Split(str, ":")
		if len(infos) != 2 {
			return fmt.Errorf("setConds weight format error")
		}
		weightName := infos[0]
		weightStr := infos[1]
		min, max := GetMinMax(weightStr)
		t.CondMap[weightName] = NewCond(weightName, min, max)
	}
	return nil
}

// InitTestTem 初始化测试模板
func (t *GenTemplate) InitTestTem(tem string) error {
	temMap := map[int][]uint16{}
	colMap, err := GetColMap(tem)
	if err != nil {
		return err
	}
	for i, s := range colMap {
		tagNames := strings.Split(s, ",")
		for _, name := range tagNames {
			if name != "" {
				tag := t.Config.GetTag(name)
				temMap[i] = append(temMap[i], uint16(tag.Id))
			}
		}
	}
	t.Template = temMap
	return nil
}

// InsertAdjacency 插入相邻标签
func (t *GenTemplate) InsertAdjacency(allTags []uint16, col int) (newAllTags []uint16) {
	for s, num := range t.Adjacent[col] {
		if num > 0 {
			for i := 0; i < num; i++ {
				tag := t.Config.GetTag(s)
				adjacency := GetRandAdjacency(allTags, tag)
				if adjacency == -1 {
					continue
				}
				allTags = ArrayInsert(allTags, adjacency, uint16(tag.Id))
			}
		}
	}
	return allTags
}
