package test

import (
	"../tnpool"
	"fmt"
	"testing"
)

func TestPool2(t *testing.T) {
	n := 0
	e := tnpool.NewTask(func() error {
		n++
		fmt.Println(n)
		return nil
	})
	p := tnpool.NewPool(10)
	go func() {

		for {
			p.EntryChannel <- e
		}
	}()

	p.Run()
}

//func TestPool(t *testing.T) {
//	var p tnpool.OPool
//	p.Init(9, 90)
//	for i := 0; i < 9; i++ {
//		p.AddTask(func() error {
//			time.Sleep(time.Second * 2)
//			fmt.Println(i)
//			return nil
//		})
//	}
//	p.SetFinishCallback(func() {
//		fmt.Println("输出完成")
//	})
//	p.Start()
//	p.Stop()
//}

//func TestChan(t *testing.T) {
//	ch := make(chan string)
//	go func() {
//		ch <- "send"
//	}()
//
//}

//func TestChan1(t *testing.T) {
//	g := make(chan func())
//	quit := make(chan bool)
//
//	go func() {
//		for {
//			select {
//			case v := <-g:
//				v()
//			case <-quit:
//				fmt.Println("B退出")
//				return
//			}
//		}
//	}()
//
//	for i := 0; i < 3; i++ {
//		g <- func() {
//			time.Sleep(time.Second*1)
//		}
//	}
//	quit <- true
//	fmt.Println("testAB退出")
//}
