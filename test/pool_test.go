package test

import (
	"../tnpool"
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var p tnpool.OPool
	p.Init(9, 90)
	for i := 0; i < 9; i++ {
		p.AddTask(func() error {
			time.Sleep(time.Second * 2)
			return nil
		})
	}
	p.SetFinishCallback(func() {
		fmt.Println("输出完成")
	})
	p.Start()
	p.Stop()
}

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
