package ordered

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

type TestValue struct {
	typ   int
	value int
}

// go test -v -run TestNewMap utils/ordered/*
func TestNewMap(t *testing.T) {
	arr := []*TestValue{
		{typ: 5, value: 1},
		{typ: 6, value: 2},
		{typ: 7, value: 3},
		{typ: 8, value: 3},
	}

	m := NewMap[int, *TestValue]()

	for _, v := range arr {
		m.Set(v.typ, v)
	}

	// 遍历所有值查看是否有序
	for i, value := range m.Values() {
		t.Log("values()", i, value)
	}

	// Range所有键值对
	m.Range(func(i int, v *TestValue) bool {
		t.Log("Range()", i, v)
		return true
	})
}

// go test -v -run TestGroupBy utils/ordered/*
func TestGroupBy(t *testing.T) {
	arr := []*TestValue{
		{typ: 1, value: 1},
		{typ: 1, value: 2},
		{typ: 2, value: 3},
		{typ: 2, value: 4},
	}

	s := GroupBy[int, *TestValue](arr, func(v *TestValue) int {
		return v.typ
	})

	// 测试Get方法
	type1Arr, _ := s.Get(1)

	// 结果需等于前两个type为1的值
	assert.Equal(t, type1Arr, []*TestValue{
		{typ: 1, value: 1},
		{typ: 1, value: 2},
	}, "Get()测试结果有误")

	// 测试Values方法
	typeArrList := s.Values()

	typeArrListWant := [][]*TestValue{
		{
			{typ: 1, value: 1},
			{typ: 1, value: 2},
		},
		{
			{typ: 2, value: 3},
			{typ: 2, value: 4},
		},
	}

	assert.Equal(t, typeArrList, typeArrListWant, "Values()测试结果有误")
}
