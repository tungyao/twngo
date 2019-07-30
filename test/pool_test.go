package test

import (
	"../tnpool"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestPool(t *testing.T) {
	var n int = 0
	e := tnpool.NewTask(func() error {
		n++
		fmt.Println(n)
		if n >= 10000 {
			fmt.Println("超过10000")
			return errors.New("STOP")
		}
		return nil
	})
	p := tnpool.NewPool(4)
	go func() {
		for {
			p.EntryChannel <- e
		}
	}()
	p.Run()
}
