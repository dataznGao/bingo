package task

import (
	"sync"
)

type Group struct {
	tasks []func()
	num   int
	wg    *sync.WaitGroup
	index int
}

func NewGroup(num int) *Group {
	wg := &sync.WaitGroup{}
	wg.Add(num)
	return &Group{
		tasks: make([]func(), num),
		num:   num,
		wg:    wg,
		index: 0,
	}
}

func (g *Group) Start() {
	for _, task := range g.tasks {
		task := task
		go func() {
			if task != nil {
				task()
				g.wg.Done()
			} else {
				g.wg.Done()
			}
		}()
	}
}

func (g *Group) Wait() {
	g.wg.Wait()
}

func (g *Group) Add(fu func()) {
	if g.index == g.num {
		g.index = 0
	}
	g.tasks[g.index] = fu
	g.index++
}
