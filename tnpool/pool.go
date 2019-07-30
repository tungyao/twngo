package tnpool

import "fmt"

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
	InnerChanerl  chan *Task
	MaxWorkNumber int
}

func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel:  make(chan *Task),
		InnerChanerl:  make(chan *Task),
		MaxWorkNumber: cap,
	}
	return &p
}
func (p *Pool) Worker(workid int) {
	for task := range p.InnerChanerl {
		task.Execute()
		fmt.Println("workid", workid, "执行完毕")
	}
}
func (p *Pool) Run() {
	for i := 0; i < p.MaxWorkNumber; i++ {
		go p.Worker(i)
	}

	for task := range p.EntryChannel {
		p.InnerChanerl <- task
	}
}
