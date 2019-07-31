package tnpool

import (
	"fmt"
)

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	t := Task{f: f}
	return &t
}
func (t *Task) Execute() {
	t.f()
}

type Pool struct {
	EntryChannel  chan *Task
	InnerChanel   chan *Task
	MaxWorkNumber int
}

func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel:  make(chan *Task),
		InnerChanel:   make(chan *Task),
		MaxWorkNumber: cap,
	}
	return &p
}
func (p *Pool) Worker(worked int) {
	for task := range p.InnerChanel {
		task.Execute()
		fmt.Println("运行：", task, "完毕")

	}
}
func (p *Pool) Close() {
	close(p.EntryChannel)
	close(p.InnerChanel)
}
func (p *Pool) Run() {
	for i := 0; i < p.MaxWorkNumber; i++ {
		go p.Worker(i)
		fmt.Println("开启线程:", i)
	}
	for task := range p.EntryChannel {
		p.InnerChanel <- task
	}

}
