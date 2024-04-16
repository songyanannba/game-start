package template

import "errors"

type OQ interface {
	GetId() int
}

type OrderQueue[T OQ] struct {
	Skills map[int]T
}

func NewOrderQueue[T OQ]() OrderQueue[T] {
	return OrderQueue[T]{
		Skills: map[int]T{},
	}
}

func (oq *OrderQueue[T]) Clear() {
	for i, _ := range oq.Skills {
		delete(oq.Skills, i)
	}
}

// IsEmpty 判断是否为空
func (oq *OrderQueue[T]) IsEmpty() bool {
	return oq.Size() == 0
}

// Size 获取队列长度
func (oq *OrderQueue[T]) Size() int {
	return len(oq.Skills)
}

// GetMinIndex 获取最小索引
func (oq *OrderQueue[T]) GetMinIndex() int {
	var minIndex int
	first := true
	for key := range oq.Skills {
		if first || key < minIndex {
			minIndex = key
			first = false
		}
	}
	return minIndex
}

// EnQueue 入队
func (oq *OrderQueue[OQ]) EnQueue(T OQ) error {
	if T.GetId() == 0 {
		return errors.New("id is 0")
	}
	oq.Skills[T.GetId()] = T
	return nil
}

// DeQueueMin 出队最小
func (oq *OrderQueue[T]) DeQueueMin() T {
	minIndex := oq.GetMinIndex()
	defer delete(oq.Skills, minIndex)
	if minIndex == 0 {
		return *new(T)
	}
	skill := oq.Skills[minIndex]
	return skill
}

// DeQueueById 出队指定id
func (oq *OrderQueue[T]) DeQueueById(index int) T {
	defer delete(oq.Skills, index)
	if index == 0 {
		return *new(T)
	}
	skill := oq.Skills[index]
	return skill
}

// GetList 获取队列
func (oq *OrderQueue[T]) GetList() []T {
	var list []T
	for _, v := range oq.Skills {
		list = append(list, v)
	}
	return list
}
