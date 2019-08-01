package test

import (
	"../tnjson"
	"github.com/tungyao/twngo/tnpool"
	"sync"
	"testing"
)

func TestJSON(t *testing.T) {
	data := map[string]interface{}{"a": "a", "c": &map[string]interface{}{"d": 123}, "b": 1}
	json := new(tnjson.JSON)

	var p tnpool.OPool
	n := 0
	p.Init(8, 90)
	for i := 0; i < 8; i++ {
		p.AddTask(func() error {
			for j := 0; j < 20000; j++ {
				n++
				json.Encode(data)
			}
			return nil
		})
	}
	p.SetFinishCallback(func() {
	})
	p.Start()
	p.Stop()
}
func TestJSON2(t *testing.T) {
	data := map[string]interface{}{"a": "a", "c": &map[string]interface{}{"d": 123}, "b": 1}
	json := new(tnjson.JSON)
	var wg sync.WaitGroup
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < 20000; j++ {

				json.Encode(data)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
