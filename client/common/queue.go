package common

import (
	"container/list"
	"sync"
	"sync/atomic"
)

type synQueue struct {
	inLen    int32
	outLen   int32
	inLock   *sync.Mutex
	outLock  *sync.Mutex
	inQueue  *list.List
	outQueue *list.List
}

func NewSynQueue() *synQueue {
	return &synQueue{
		inLock:   new(sync.Mutex),
		outLock:  new(sync.Mutex),
		inQueue:  list.New(),
		outQueue: list.New(),
	}
}

//添加队列
func (q *synQueue) Add(i interface{}) {
	q.inLock.Lock()
	defer q.inLock.Unlock()
	q.inQueue.PushBack(i)
	atomic.AddInt32(&q.inLen, 1)
}

//交换 队列
func (q *synQueue) SwapQue() {
	q.inLock.Lock()
	q.outLock.Lock()

	q.inQueue, q.outQueue = q.outQueue, q.inQueue
	q.inLen, q.outLen = q.outLen, q.inLen

	q.outLock.Unlock()
	q.inLock.Unlock()
}

//删除
func (q *synQueue) RemoveQue() {
	q.outQueue.Remove(q.outQueue.Front())
}
