package template

type Queue[T any] struct {
	first *node[T]
	last  *node[T]
	n     int
}

type node[T any] struct {
	item T
	next *node[T]
}

type Level struct {
	CoreCount int // 核心数量
	EmitCount int // 发射数量
	WildMul   int // 万能数量

}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// IsEmpty 判断是否为空
func (q *Queue[T]) IsEmpty() bool {
	return q.n == 0
}

// Size 获取队列长度
func (q *Queue[T]) Size() int {
	return q.n
}

// EnQueue 入队
func (q *Queue[T]) EnQueue(t T) {
	oldLast := q.last
	q.last = &node[T]{}
	q.last.item = t
	q.last.next = nil
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.n++
}

// DeQueue 出队
func (q *Queue[T]) DeQueue() T {
	if q.IsEmpty() {
		return *new(T)
	}
	item := q.first.item
	q.first = q.first.next
	if q.IsEmpty() {
		q.last = nil
	}
	q.n--
	return item
}
