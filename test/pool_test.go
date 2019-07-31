package test

import (
	"../tnpool"
	"fmt"
	"testing"
	"time"
	//"time"
)

func TestPool(t *testing.T) {
	n := 0
	e := tnpool.NewTask(func() error {
		time.Sleep(time.Second * 1)
		n++
		fmt.Println(n)
		return nil
	})
	p := tnpool.NewPool(10)

	//defer close(p.EntryChannel)

	go func() {
		for i := 0; i < 10; i++ {
			p.EntryChannel <- e
		}

	}()

	p.Run()

}
func TestChan(t *testing.T) {
	ch := make(chan string)
	go func() {
		ch <- "send"
	}()

}
