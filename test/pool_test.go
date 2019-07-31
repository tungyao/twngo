package test

import (
	"../tnpool"
	"testing"
	//"time"
)

func TestPool(t *testing.T) {
	e := tnpool.NewTask(func() error {
		return nil
	})
	p := tnpool.NewPool(8)
	go func() {
		defer close(p.EntryChannel)
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
