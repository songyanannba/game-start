package helper

import (
	"github.com/samber/lo"
	"sync"
)

type PoolRunner interface {
	Run()
}

type Pool struct {
	sync.Mutex
	m     map[int]PoolRunner
	idArr []int
	ch    chan struct{}
}

func NewPool() *Pool {
	p := &Pool{
		m:  make(map[int]PoolRunner),
		ch: make(chan struct{}, 1000),
	}
	go p.start()
	return p
}

func (p *Pool) Put(id int, t PoolRunner) {
	p.Lock()
	defer p.Unlock()
	p.m[id] = t
	p.idArr = append(p.idArr, id)
	go p.runOnce()
}

func (p *Pool) start() {
	// 接受到信号时取出任务并执行
	for range p.ch {
		t := p.Pop()
		if t != nil {
			t.Run()
		}
	}
}

// Get 从池中获取指定任务
func (p *Pool) Get(id int) PoolRunner {
	p.Lock()
	defer p.Unlock()

	return p.m[id]
}

// PopByID 从池中取出指定任务 可作为删除
func (p *Pool) PopByID(id int) PoolRunner {
	p.Lock()
	defer p.Unlock()

	return p.popByID(id)
}

// Pop 从池中取出第一个任务
func (p *Pool) Pop() PoolRunner {
	p.Lock()
	defer p.Unlock()
	if len(p.idArr) == 0 {
		return nil
	}

	return p.popByID(p.idArr[0])
}

// popByID 从池中取出指定任务 内部使用 无锁
func (p *Pool) popByID(id int) PoolRunner {
	r, ok := p.m[id]
	if !ok {
		return nil
	}
	delete(p.m, id)
	p.idArr = lo.Filter(p.idArr, func(i, k int) bool {
		return i != id
	})
	return r
}

// GetIDArr 获取任务ID组
func (p *Pool) GetIDArr() []int {
	p.Lock()
	defer p.Unlock()
	return CopySlice(p.idArr)
}

// runOnce 发送一个运行信号
func (p *Pool) runOnce() {
	p.ch <- struct{}{}
}
