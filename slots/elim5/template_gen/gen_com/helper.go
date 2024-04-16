package gen_com

import (
	"elim5/logicPack/base"
	"elim5/utils/helper"
	"errors"
	"fmt"
)

// IntervalPlacement 按照间隔放置
func IntervalPlacement(tagIds []uint16, fillTag *base.Tag, num, minInterval, maxInterval int) error {
	if len(tagIds) < num*minInterval+num {
		return errors.New(fmt.Sprintf("tags length is %d, but num is %d, interval is %d", len(tagIds), num, minInterval))
	}
	if fillTag.Id < 1 {
		return errors.New("fillTag id is error")
	}
	mayIndexMap := helper.ArrConversionMap(tagIds, func(item uint16, i int) (int, bool) {
		return i, item > 0
	})

	for i := 0; i < num; i++ {
		mayIndexs := helper.MapFilter(mayIndexMap, func(item int, i bool) (int, bool) {
			return item, !i
		})
		if len(mayIndexs) == 0 {
			continue
		}
		rand := helper.RandInt(len(mayIndexs))
		randIndex := mayIndexs[rand]
		tagIds[randIndex] = uint16(fillTag.Id)
		mayIndexMap[randIndex] = true
		delete(mayIndexMap, randIndex)

		for a := 0 - minInterval; a < minInterval+1; a++ {
			useIndex := randIndex + a
			if useIndex < 0 {
				useIndex = len(tagIds) + useIndex
			}
			if useIndex >= len(tagIds) {
				useIndex = useIndex - len(tagIds)
			}
			mayIndexMap[useIndex] = true
			delete(mayIndexMap, useIndex)
		}
	}

	return nil
}

func FilterIndex[T any](list []T, f func(item T, i int) bool) []int {
	indexs := make([]int, 0)
	for i, t := range list {
		if f(t, i) {
			indexs = append(indexs, i)
		}
	}
	return indexs
}

func TagsFill(tags []uint16, fillTags []uint16) {
	fillIndex := 0
	for i, tag := range tags {
		if tag < 1 {
			tags[i] = fillTags[fillIndex]
			fillIndex++
		}
	}
}

// GetRandAdjacency 随机取某个标签位置,多个相邻算一个
func GetRandAdjacency(allTags []uint16, name *base.Tag) int {
	adjacency := make([]int, 0)
	isLx := true
	for i, tag := range allTags {
		if tag == uint16(name.Id) && isLx {
			adjacency = append(adjacency, i)
			isLx = false
		}
		if tag != uint16(name.Id) {
			isLx = true
		}
	}
	if len(adjacency) == 0 {
		return -1
	}
	return adjacency[helper.RandInt(len(adjacency))]
}

// GetRandAdjacencyMul 随机取某个标签,只取重叠的
func GetRandAdjacencyMul(allTags []*base.Tag, name *base.Tag) int {
	adjacency := make([]int, 0)
	for i, tag := range allTags {
		lIndex := i - 1
		if lIndex < 0 {
			lIndex = len(allTags) - 1
		}
		rIndex := i + 1
		if rIndex >= len(allTags) {
			rIndex = 0
		}
		if tag.Name == name.Name && (allTags[lIndex].Name == name.Name || allTags[rIndex].Name == name.Name) {
			adjacency = append(adjacency, i)
		}
	}
	if len(adjacency) == 0 {
		return -1
	}
	return adjacency[helper.RandInt(len(adjacency))]
}

func GetMinMax(str string) (min, max float64) {
	min, max = 0.0, 0.0
	_, err := fmt.Sscanf(str, "%f-%f", &min, &max)
	if err != nil {
		return 0, 0
	}
	return min, max
}

// ArrayInsert 在数组指定位置插入值
func ArrayInsert[T any](slice []T, index int, value T) []T {
	// 如果索引是数组长度减一，则插入到开头
	if index == len(slice)-1 {
		return append([]T{value}, slice...)
	}

	// 如果索引是0或负数，则插入到末尾
	if index <= 0 {
		return append(slice, value)
	}
	// 在指定索引位置插入值
	return append(slice[:index], append([]T{value}, slice[index:]...)...)
}

// ArrayRemove 在数组指定位置删除值
func ArrayRemove[T any](slice []T, index int) []T {
	if index >= 0 && index < len(slice) {
		return append(slice[:index], slice[index+1:]...)

	} else {
		return slice
	}
}

func MapMinBy[T any, N helper.Number](collection map[N]T, comparison func(a T, b T) bool) (T, N) {
	var minAny T
	var minKey N

	if len(collection) == 0 {
		return minAny, minKey
	}
	minAny = *new(T)
	minKey = *new(N)
	for n, t := range collection {
		if comparison(t, minAny) {
			minAny = t
			minKey = n
		}
	}
	return minAny, minKey
}
