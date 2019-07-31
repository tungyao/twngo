package tnpool

import (
	"fmt"
)

type Task struct {
	f chan func() error
}

func NewTask(f chan func() error) *Task {
	t := Task{f: f}
	go t.Execute()
	return &t
}
func (t *Task) Execute() {
	for f := range t.f {
		f()
	}
}

type fPool struct {
	EntryChannel  chan *Task
	InnerChanel   chan *Task
	MaxWorkNumber int
	Stop          chan bool
}

func NewPool(cap int) *fPool {
	p := fPool{
		EntryChannel:  make(chan *Task),
		InnerChanel:   make(chan *Task),
		Stop:          make(chan bool),
		MaxWorkNumber: cap,
	}
	return &p
}
func (p *fPool) Worker(worked int) {
	for task := range p.InnerChanel {
		task.Execute()
		fmt.Println("运行：", worked, "完毕")
	}
}
func (p *fPool) Close() {
	close(p.EntryChannel)
	close(p.InnerChanel)
}

func (p *fPool) Run() {
	for i := 0; i < p.MaxWorkNumber; i++ {
		go p.Worker(i)
	}
	for {
		select {
		case ret := <-p.EntryChannel:
			p.InnerChanel <- ret
			break
		case <-p.Stop:
			fmt.Println("线程关闭关闭")
			p.Close()
			return
		}

	}

}
