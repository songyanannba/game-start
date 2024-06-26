package base

import (
	"elim5/enum"
	"elim5/pbs/common"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
	"sort"
	"strconv"
	"strings"
)

type Tag struct {
	Id       int
	Name     string
	Include  []string `json:"-"`
	Multiple float64

	X int
	Y int

	IsLine      bool `json:"-"`
	IsPayTable  bool `json:"-"`
	IsWild      bool
	IsSingle    bool `json:"-"` // 是否单出
	IsJackpot   bool `json:"-"`
	IsLock      bool `json:"-"` // 是否锁定
	IsHead      bool
	IsMegastack bool   //特殊标签
	Redundancy  string //冗余字段
	TagExt
	Life    int //生命值
	MergeId int //合并id
	Rank    int //等级
	Valid   bool
	Gold    bool
	Fixed   bool
}

type TagExt struct {
	IsSpecial       bool
	InitSite        int
	Dir             int // 1上 2 下
	Step            int
	RelativeDir     int //相对点
	MidWildId       int
	Which           int
	MoveAfterTag    *Tag
	LinkCollectMult float64
	IsLinkCollect   bool //是否已经使用
	IsLinkMult      bool //是否已经使用
	IsMystery       bool
	IsConverter     bool //转换标签 是否使用过
	IsLinkNudgeUp   bool //转换标签 是否使用过
}

func (t *Tag) String() string {

	return fmt.Sprintf("%d :%s %d-%d", t.Id, t.Name, t.X, t.Y)
}

type TagList []*Tag

func GetFixedTag() *Tag {
	return &Tag{
		Name:  "fixed",
		Fixed: true,
	}
}

func (t *Tag) IsEmpty() bool {
	if t == nil {
		return true
	}
	return t.Name == "" || t.Id == -1
}

func (t *Tag) IsFixed() bool {
	return t.Fixed
}

func (t *Tag) Match() []string {
	if t.IsWild {
		return t.Include
	} else {
		return []string{t.Name}
	}
}

func (ts *TagList) Copy() TagList {
	var list TagList
	for _, v := range *ts {
		list = append(list, v.Copy())
	}
	return list
}

type Label struct {
	Name   string `json:"name"`
	Colour string `json:"colour"` //Default Primary Success Into Warning Danger
	Hide   string `json:"hide"`
}

func (t *Tag) Dump(typ ...int) *Label {
	label := &Label{
		Name:   t.Name,
		Colour: enum.VueTypeInfo,
	}
	if t.IsPayTable {
		label.Colour = enum.VueTypePrimary
	}
	if t.IsSingle {
		label.Colour = enum.VueTypeSuccess
	}
	if t.IsJackpot {
		label.Colour = enum.VueTypeWarning
	}
	if t.IsWild {
		label.Colour = enum.VueTypeDanger
	}

	if t.Multiple > 1 {
		label.Name = fmt.Sprintf("%s*%d", label.Name, int(t.Multiple))
	}
	if t.Valid || t.IsMegastack {
		label.Hide = "*"
	}
	label.Hide += t.Redundancy
	return label
}

func (t *Tag) parseAttr(s string) {
	id, attr, _ := strings.Cut(s, "|")
	bin := helper.DecToBin(attr)
	attrs := helper.SplitInt[int](bin, "")
	arr := make([]int, 5)
	copy(arr, attrs)
	t.Id = helper.Atoi(id)
	t.IsLine = arr[0] == 1
	t.IsPayTable = arr[1] == 1
	t.IsWild = arr[2] == 1
	t.IsSingle = arr[3] == 1
	t.IsJackpot = arr[4] == 1
}

func (t Tag) ToCard() *common.Card {
	return &common.Card{
		CardId:     int32(t.Id),
		IsWild:     t.IsWild,
		IsPayTable: t.IsPayTable,
		X:          int32(t.X),
		Y:          int32(t.Y),
		IsLock:     t.IsLock,
		Redundancy: t.Redundancy,
	}
}

func (t Tag) ToTagAck() *common.Tag {
	return &common.Tag{
		TagId:       int32(t.Id),
		X:           int32(t.X),
		Y:           int32(t.Y),
		Multiple:    int32(t.Multiple),
		IsWild:      t.IsWild,
		Redundancy:  t.Redundancy,
		IsMegastack: t.IsMegastack,
	}
}

func (t *Tag) Copy() *Tag {
	c := *t
	return &c
}

func NewTag(id int, v string, multiple float64, single uint8, include ...string) *Tag {
	var isWild bool
	if len(include) > 0 {
		if include[0] != "" {
			isWild = true
		}
	}
	return &Tag{
		Id:       id,
		Name:     v,
		Include:  include,
		IsWild:   isWild,
		Multiple: multiple,
		IsSingle: helper.If(single == 1, true, false),
	}
}

func (t *Tag) SetInclude(include ...string) {
	t.Include = append(t.Include, include...)
	if len(t.Include) > 0 {
		t.IsWild = true
	}
}

// IsMatchName 匹配名称 自身考虑wild 其他只考虑名称
func (t *Tag) IsMatchName(v string) bool {
	if t.IsWild {
		for _, i := range t.Include {
			if i == v {
				return true
			}
		}
		return false
	}
	return t.Name == v
}

// IsMatchTag 匹配名称 自身考虑wild 其他只考虑wild
func (t *Tag) IsMatchTag(tag *Tag) bool {
	if t.IsWild {
		for _, i := range t.Include {
			if i == tag.Name {
				return true
			}
		}
		return false
	}

	if tag.IsWild {
		for _, s := range tag.Include {
			if s == t.Name {
				return true
			}
		}
		return false
	}

	return t.Name == tag.Name
}

// MatchTag 匹配标签 自身和其他都考虑wild
func (t *Tag) MatchTag(tag *Tag) bool {
	if t.IsWild && tag.IsWild {
		for _, v := range t.Include {
			if tag.IsMatchName(v) {
				return true
			}
		}
		return false
	}
	if t.IsWild {
		return t.IsMatchName(tag.Name)
	}
	if tag.IsWild {
		return tag.IsMatchName(t.Name)
	}
	return t.Id == tag.Id
}

// InTags 判断当前tag是否在标签中 被匹配的标签不考虑wild
func (t *Tag) InTags(tags []*Tag) bool {
	for _, tag := range tags {
		if t.IsMatchName(tag.Name) {
			return true
		}
	}
	return false
}

// GetTagsName 获取该组标签并排除百搭标签的名称
func GetTagsName(tags []*Tag, exclude ...string) string {
	if len(tags) == 0 {
		return ""
	}
	for _, tag := range tags {
		if helper.InArr(tag.Name, exclude) {
			continue
		}
		return tag.Name
	}
	return ""
}

func GetTagsNameByFunc(tags []*Tag, fn func(*Tag) bool) string {
	if len(tags) == 0 {
		return ""
	}
	for _, tag := range tags {
		if !fn(tag) {
			continue
		}
		return tag.Name
	}
	return ""
}

// SameTags 使用多叉树 递归匹配相同的标签
type SameTags struct {
	Tag  *Tag
	Name string
	Key  int
	Last []*SameTags
	Ok   bool
}

func newSameTag(t *Tag, k int, name string) *SameTags {
	if name == "" {
		if !t.IsWild {
			name = t.Name
		}
	}
	return &SameTags{
		Tag:  t,
		Key:  k,
		Name: name,
	}
}

func (s *SameTags) Add(t *Tag, k int) {
	s.Last = append(s.Last, newSameTag(t, k, s.Name))
}

func (s *SameTags) Match(t *Tag, k int) bool {
	// 当前k是否对应
	if s.Key+1 == k {
		if s.Name != "" {
			if t.IsMatchName(s.Name) {
				s.Add(t, s.Key+1)
				s.Ok = true
				return true
			}
		} else {
			// name为空 则为百搭
			if s.Tag.MatchTag(t) {
				s.Add(t, s.Key+1)
				s.Ok = true
				return true
			}
		}
	} else {
		if len(s.Last) > 0 {
			res := false
			for _, tags := range s.Last {
				if tags.Match(t, k) {
					res = true
					s.Ok = true
				}
			}
			return res
		}
	}
	return false
}

// SumRes 递归结果
func (s *SameTags) SumRes(tags []*Tag) [][]*Tag {
	var ts = append(tags, s.Tag)
	if len(s.Last) > 0 {
		var res [][]*Tag
		for _, son := range s.Last {
			if ts != nil {
				var tsCopy = make([]*Tag, len(ts))
				copy(tsCopy, ts)
				res = append(res, son.SumRes(tsCopy)...)
			}
		}
		return res
	}
	return [][]*Tag{ts}
}

// MatchSameTagList 从初始数据中匹配相同的标签列表 此时数据为纵向排列
func MatchSameTagList(tagList [][]*Tag, length int, selectTags ...*Tag) [][]*Tag {
	if len(tagList) == 0 {
		return nil
	}
	var sameTags []*SameTags //第一列数据
	for _, tag := range tagList[0] {
		if len(selectTags) != 0 {
			if tag.InTags(selectTags) {
				sameTags = append(sameTags, newSameTag(tag, 0, enum.EmptyTagName))
			}
		} else {
			sameTags = append(sameTags, newSameTag(tag, 0, enum.EmptyTagName))
		}
	}
	if len(sameTags) == 0 {
		return nil
	}
	for i := 1; i < len(tagList); i++ {
		for _, tag := range tagList[i] {
			for _, sameTag := range sameTags {
				sameTag.Match(tag, i)
			}
		}
		if i < length {
			sameTags = lo.Filter(sameTags, func(item *SameTags, index int) bool {
				return item.Ok
			})
			for _, same := range sameTags {
				same.Ok = false
			}
		}
	}
	var res [][]*Tag
	for _, same := range sameTags {
		res = append(res, same.SumRes(nil)...)
	}
	for _, tags := range res {
		for _, tag := range tags {
			tag.IsLine = true
		}
	}
	return res
}

// MatchSameTags 横向匹配相同的标签
func MatchSameTags(tags []*Tag, length int, selectName ...string) bool {
	if len(tags) < 2 {
		return true
	}
	sameTag := newSameTag(tags[0], 0, "")
	for i := 1; i < len(tags); i++ {
		sameTag.Add(tags[i], i)
		if !sameTag.Match(tags[i], i) {
			return false
		}
		if i+1 >= length {
			break
		}
	}
	return true
}

func FilterSameTagList(tagList [][]*Tag, length int) [][]*Tag {
	var arr [][]*Tag
	for _, v := range tagList {
		if MatchSameTags(v, length) {
			arr = append(arr, v)
		}
	}
	return arr
}

// MatchSameTagName 一组标签中标签名称是否相同
func MatchSameTagName(tags []*Tag, length int, name string) bool {
	if len(tags) == 0 {
		return false
	}
	for i, tag := range tags {
		if tag.Name != name {
			return false
		}
		if i+1 >= length {
			break
		}
	}
	return true
}

func GetSpecialTags(tagList [][]*Tag, names ...string) []*Tag {
	Result := make([]*Tag, 0)
	for _, tags := range tagList {
		for _, tag := range tags {
			if lo.Contains(names, tag.Name) {
				var tagNew = *tag
				Result = append(Result, &tagNew)
			}
		}
	}
	return Result
}

func CountTagName(tags []*Tag, name string) (count int) {
	for _, tag := range tags {
		if tag.Name == name {
			count++
		}
	}
	return
}

// TagRecord 标签记录
type TagRecord struct {
	Tag
	WinType int
}

// GetRecordStr 根据标签生成记录
func (t Tag) GetRecordStr(id int, typ string) string {
	if id == 0 {
		id = t.Id
	}
	return fmt.Sprintf("%s,%d,%d,%d,%d,%f",
		typ, id, t.X, t.Y, helper.If(t.IsWild, 1, 0), t.Multiple,
	)
}

// ParseTagRecord 解析标签记录
func ParseTagRecord(s string) TagRecord {
	t := TagRecord{}
	splitTag := strings.Split(s, ",")
	if len(splitTag) != 6 {
		return t
	}
	t.WinType, _ = strconv.Atoi(splitTag[0])
	t.Id, _ = strconv.Atoi(splitTag[1])
	t.X, _ = strconv.Atoi(splitTag[2])
	t.Y, _ = strconv.Atoi(splitTag[3])
	t.IsWild, _ = strconv.ParseBool(splitTag[4])
	t.Multiple, _ = strconv.ParseFloat(splitTag[5], 64)
	return t
}

func (t Tag) TagSplit(Split string) (name string, serial int) {

	splitTag := strings.Split(t.Name, Split)
	if len(splitTag) != 2 {
		return "", 0
	}
	name = splitTag[0]
	serial, _ = strconv.Atoi(splitTag[1])

	return name, serial
}

func (t *Tag) TagExtAssign(site, step int) {
	t.TagExt.IsSpecial = true
	t.TagExt.InitSite = site
	t.TagExt.Step = step
}

func ExpMatchSameTagList(tagList [][]*Tag, cols []int, length int) [][]*Tag {
	var newCol []int
	for _, v := range cols {
		newCol = append(newCol, v)
	}
	for _, tags := range tagList {
		for _, tag := range tags {
			isIn := false
			if tag.IsWild {
				for _, c := range cols {
					if tag.X == c {
						isIn = true
					}
				}
				if !isIn {
					newCol = append(newCol, tag.X)
				}
			}
		}
	}

	sort.Ints(newCol) //生序

	var sameTags []*SameTags
	for _, tag := range tagList[newCol[0]] {
		sameTags = append(sameTags, newSameTag(tag, 0, enum.EmptyTagName))
	}
	if len(sameTags) == 0 {
		return nil
	}
	for i := newCol[0] + 1; i < len(tagList); i++ {
		for _, tag := range tagList[i] {
			for _, sameTag := range sameTags {
				sameTag.ExpMatch(tag, i, newCol)
			}
		}

		if i < length {
			//sameTags = lo.Filter(sameTags, func(item *SameTags, index int) bool {
			//	return item.Ok
			//})
			//for _, same := range sameTags {
			//	same.Ok = false
			//}
		}
	}

	var res [][]*Tag
	for _, same := range sameTags {
		res = append(res, same.SumRes(nil)...)
	}

	var newRes [][]*Tag
	for _, tags := range res {
		if len(tags) >= length {
			newRes = append(newRes, tags)
		}
	}
	return newRes
}

func (s *SameTags) ExpMatch(t *Tag, k int, cols []int) bool {
	// 当前k是否对应
	// 当前k是否对应
	pre := 0
	next := 0
	isIn := false
	for i, v := range cols {
		if s.Key == v {
			pre = i
		}
		if k == v {
			next = i
		}
		if k == v {
			isIn = true
		}
	}
	if !isIn {
		return false
	}

	if pre+1 == next {
		//
		if s.Name != "" {
			if t.IsMatchName(s.Name) {
				s.Add(t, k)
				s.Ok = true
				return true
			}
		} else {
			// name为空 则为百搭
			if s.Tag.MatchTag(t) {
				s.Add(t, k)
				s.Ok = true
				return true
			}
		}

	} else {
		//
		if len(s.Last) > 0 {
			res := false
			for _, tags := range s.Last {
				if tags.ExpMatch(t, k, cols) {
					res = true
					s.Ok = true
				}
			}
			return res
		}
	}
	return false
}
