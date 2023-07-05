package eliminate

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"slot6/global"
	"slot6/src/base"
	"slot6/utils/helper"

	"strconv"
)

type Table struct {
	Row          int               // 行数
	Col          int               // 列数
	Tags         []base.Tag        // 所有tag
	WildTags     []*base.Tag       // wild标签
	Scatter      *base.Tag         // scatter
	TagList      [][]*base.Tag     // 二维列表
	InitTable    [][]*base.Tag     // 初始表
	Target       *Target           // 目标
	Mul          float64           // 倍数
	SkillMul     float64           // 技能倍数
	PayTableList []*base.PayTable  // 所有的paytable
	AlterFlows   []*base.AlterFlow // 改变的流程

	RmCount int //移除的个数

	//TagListMapNum map[string]int
	PayTableListMaps map[string]map[int]*base.PayTable //paytable 划线的map

	SlotId uint // 机器id
	RankId int

	TableWildPoint map[[2]int]*base.Tag //全局
}

func (t *Table) PayTableListMap() map[string]map[int]*base.PayTable {
	if t.PayTableListMaps != nil {
		return t.PayTableListMaps
	}
	var ptm = make(map[string]map[int]*base.PayTable)
	for _, v := range t.PayTableList {
		nameKay := v.Tags[0].Name
		if ptm[nameKay] == nil {
			ptm[nameKay] = make(map[int]*base.PayTable)
		}
		ptm[nameKay][len(v.Tags)] = v
	}
	t.PayTableListMaps = ptm
	return ptm
}

// NeedFill 需要填充的空白Tag
func (t *Table) NeedFill() []*base.Tag {
	return lo.Filter(helper.ListToArr(t.TagList), func(tag *base.Tag, i int) bool {
		return tag.Name == ""
	})
}

// NeedFillEdge 获取空白边缘的标签
func (t *Table) NeedFillEdge() []*base.Tag {
	v := make(map[[2]int]bool, 0)
	needFillTags := t.NeedFill()
	for _, fillTag := range needFillTags {
		for _, tag := range t.GetAdjacent(fillTag.X, fillTag.Y) {
			if tag.Name != "" && tag.Name != "scatter" {
				v[[2]int{tag.X, tag.Y}] = true
			}
		}
	}
	tags := make([]*base.Tag, 0)
	for k := range v {
		tags = append(tags, t.TagList[k[0]][k[1]])
	}
	return tags
}

// QueryTags  查询指定名称的标签
func (t *Table) QueryTags(name string) []*base.Tag {
	return lo.Filter(t.ToArr(), func(tag *base.Tag, i int) bool {
		return tag.Name == name
	})
}

// GetTag 获取tag
func (t *Table) GetTag(x, y int) *base.Tag {
	return t.TagList[x][y]
}

func (t *Table) NameGetTag(name string) *base.Tag {
	if name == "scatter" {
		return t.Scatter.Copy()
	}

	for _, tag := range t.Tags {
		if tag.Name == name {
			return tag.Copy()
		}
	}

	for _, tag := range t.WildTags {
		if tag.Name == name {
			return tag.Copy()
		}
	}

	return &base.Tag{}
}

// SetTag 设置tag
func (t *Table) SetTag(x, y int, tag base.Tag) {
	tag.X = x
	tag.Y = y
	t.TagList[x][y] = &tag
}

// randTag 随机一个tag
func (t *Table) randTag() base.Tag {
	return t.Tags[helper.RandInt(len(t.Tags))]
}

// Count 统计某个name的个数
func (t *Table) Count(name string) (count int) {
	for _, tag := range t.ToArr() {
		if tag.Name == name {
			count++
		}
	}
	return
}

// ToArr 转换为一维数组
func (t *Table) ToArr() []*base.Tag {
	return helper.ListToArr(t.TagList)
}

// GetColumn 获取某一列
func (t *Table) GetColumn(Y int) []*base.Tag {
	var cols []*base.Tag
	for _, tags := range t.TagList {
		cols = append(cols, tags[Y])
	}
	return cols
}

// GetAdjacent 获取附近四个方向的tag
func (t *Table) GetAdjacent(x, y int) []*base.Tag {
	var adjacent []*base.Tag
	if x > 0 {
		adjacent = append(adjacent, t.TagList[x-1][y].Copy())
	}
	if x < t.Row-1 {
		adjacent = append(adjacent, t.TagList[x+1][y].Copy())
	}
	if y > 0 {
		adjacent = append(adjacent, t.TagList[x][y-1].Copy())
	}
	if y < t.Col-1 {
		adjacent = append(adjacent, t.TagList[x][y+1].Copy())
	}
	return adjacent
}

// GetBiasAdjacent 获取斜角方向
func (t *Table) GetBiasAdjacent(x, y int) []*base.Tag {
	var adjacent []*base.Tag
	if x > 0 && y > 0 {
		adjacent = append(adjacent, t.TagList[x-1][y-1].Copy())
	}
	if x > 0 && y < t.Col-1 {
		adjacent = append(adjacent, t.TagList[x-1][y+1].Copy())
	}

	if x < t.Row-1 && y > 0 {
		adjacent = append(adjacent, t.TagList[x+1][y-1].Copy())
	}
	if x < t.Row-1 && y < t.Col-1 {
		adjacent = append(adjacent, t.TagList[x+1][y+1].Copy())
	}

	return adjacent
}

// GetAllAdjacent 获取附近8个方向的tag
func (t *Table) GetAllAdjacent(x, y int) []*base.Tag {
	var adjacent []*base.Tag
	adjacent = append(adjacent, t.GetAdjacent(x, y)...)
	adjacent = append(adjacent, t.GetBiasAdjacent(x, y)...)
	return adjacent
}

// GetAdjacentOne  获取附近四个方向的tag随机取一个
func (t *Table) GetAdjacentOne(x, y int) *base.Tag {
	var adjacent []*base.Tag
	if x > 0 {
		adjacent = append(adjacent, t.TagList[x-1][y].Copy())
	}
	if x < t.Row-1 {
		adjacent = append(adjacent, t.TagList[x+1][y].Copy())
	}
	if y > 0 {
		adjacent = append(adjacent, t.TagList[x][y-1].Copy())
	}
	if y < t.Col-1 {
		adjacent = append(adjacent, t.TagList[x][y+1].Copy())
	}
	return adjacent[helper.RandInt(len(adjacent))]
}

// RandPosition 返回随机位置
func (t *Table) RandPosition(verify *Verify) [2]int {
	var x, y int
	count := 0
	for {
		if count > 100 {
			global.GVA_LOG.Error("随机位置超过100次" + t.PrintTable(""))
		}
		count++
		x = helper.RandInt(t.Row)
		y = helper.RandInt(t.Col)
		if !verify.GetVerify(x, y) {
			break
		}
	}
	return [2]int{
		x, y,
	}
}

func (t *Table) PrintTable(str string) string {
	str += ":\n"
	for _, row := range t.TagList {
		for _, col := range row {
			str += fmt.Sprintf("%s\t", strconv.Itoa(col.X)+":"+strconv.Itoa(col.Y)+" "+helper.If(col.Name == "", "🀆", col.Name))
		}
		str += "\r\n"
	}
	//fmt.Println(str)
	return str + "\r\n"
}

func (t *Table) PrintList(tags []*base.Tag) {
	str := ":\n"
	for r, row := range t.TagList {
		for c, col := range row {

			tagNeeds := lo.Filter(tags, func(tag *base.Tag, i int) bool {
				return tag.X == r && tag.Y == c
			})
			if len(tagNeeds) > 0 {
				str += fmt.Sprintf("%s\t", strconv.Itoa(col.X)+":"+strconv.Itoa(col.Y)+helper.If(tagNeeds[0].Name == "", "🀆", tagNeeds[0].Name))
			} else {
				str += fmt.Sprintf("%s\t", strconv.Itoa(col.X)+":"+strconv.Itoa(col.Y)+"🀆")
			}

		}
		str += "\r\n"
	}
	fmt.Println(str)
}

func (t *Table) AddMul(mul float64) {
	t.Mul, _ = decimal.NewFromFloat(t.Mul).Add(decimal.NewFromFloat(mul)).Float64()
}
func (t *Table) AddSkillMul(mul float64) {
	t.SkillMul, _ = decimal.NewFromFloat(t.SkillMul).Add(decimal.NewFromFloat(mul)).Float64()
}

// SetCoordinates 设置坐标
func (t *Table) SetCoordinates() {
	for i, tags := range t.TagList {
		for i2, tag := range tags {
			tag.X = i
			tag.Y = i2
		}
	}
}

// TableReset 初始化列表
func (t *Table) TableReset() {
	t.TagList = helper.NewTable[*base.Tag](t.Row, t.Col, func(x, y int) *base.Tag {
		return &base.Tag{X: x, Y: y, Name: ""}
	})
}

func (t *Table) Copy() *Table {
	newT := *t
	return &newT
}

// GetGraph 获取布局
func (t *Table) GetGraph() [][]*base.Tag {
	tagList := make([][]*base.Tag, t.Row)
	for i, tags := range t.TagList {
		tagList[i] = make([]*base.Tag, t.Col)
		for i2, tag := range tags {
			tagList[i][i2] = tag.Copy()
		}
	}
	return tagList
}

func (t *Table) GetInformation() string {
	//MinMul      float64 // 最小倍数
	//MaxMul      float64 // 最大倍数
	//InitNum     int     // 初始个数
	//ScatterNum  int     // scatter次数
	mul := t.Mul
	return fmt.Sprintf(
		"最大倍率:%g,最小倍率:%g,初始个数:%d,scatter次数:%d,倍率:%g",
		t.Target.MaxMul,
		t.Target.MinMul,
		t.Target.InitNum,
		t.Target.ScatterNum,
		mul,
	)

}

//func (t *Table) GetTagListKeyNum() map[string]int {
//	// 现在只能在用到的时候重新计算
//	//todo 优化方向：每次重新布局的时候 填充数据
//	var m = make(map[string]int)
//	for _, tags := range t.TagList {
//		for _, tag := range tags {
//			m[tag.Name]++
//		}
//	}
//	return m
//}

func (t *Table) GetUpDownLeftRight(x, y int) []*base.Tag {
	var adjacent []*base.Tag
	//上
	if x > 0 {
		if t.TagList[x-1][y].Copy().Name == "" {
			adjacent = append(adjacent, t.TagList[x-1][y].Copy())
		}
	}

	//下
	if x < t.Row-1 {
		if t.TagList[x+1][y].Copy().Name == "" {
			adjacent = append(adjacent, t.TagList[x+1][y].Copy())
		}

	}

	//左
	if y > 0 {
		if t.TagList[x][y-1].Copy().Name == "" {
			adjacent = append(adjacent, t.TagList[x][y-1].Copy())
		}
	}
	//右
	if y < t.Col-1 {
		if t.TagList[x][y+1].Copy().Name == "" {
			adjacent = append(adjacent, t.TagList[x][y+1].Copy())
		}
	}

	return adjacent
}

func (t *Table) GetRoundTagsByCenterPoint(x, y int) []*base.Tag {
	var adjacent []*base.Tag
	//上
	if x > 0 {
		if t.TagList[x-1][y].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x-1][y].Copy())
		}
	}
	//左上
	if x > 0 && y > 0 {
		if t.TagList[x-1][y-1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x-1][y-1].Copy())
		}
	}

	//右上
	if x > 0 && y < t.Col-1 {
		if t.TagList[x-1][y+1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x-1][y+1].Copy())
		}
	}

	//下
	if x < t.Row-1 {
		if t.TagList[x+1][y].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x+1][y].Copy())
		}
	}

	//左下
	if x < t.Row-1 && y > 0 {
		if t.TagList[x+1][y-1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x+1][y-1].Copy())
		}
	}

	//右下
	if x < t.Row-1 && y < t.Col-1 {
		if t.TagList[x+1][y+1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x+1][y+1].Copy())
		}
	}

	//左
	if y > 0 {
		if t.TagList[x][y-1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x][y-1].Copy())
		}
	}

	//右
	if y < t.Col-1 {
		if t.TagList[x][y+1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x][y+1].Copy())
		}
	}
	return adjacent
}
