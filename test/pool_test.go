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
		time.Sleep(time.Second * 2)
		n++
		fmt.Println(n)
		return nil
	})
	p := tnpool.NewPool(8)
	go func() {
		for i := 0; i < 8; i++ {
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
