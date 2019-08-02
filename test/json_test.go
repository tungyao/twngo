package test

import (
	"../tnjson"
	"testing"
)

func TestJSON(t *testing.T) {
	data := map[string]interface{}{"a": "a"}
	json := new(tnjson.JSON)
	t.Log(json.Encode(data))
	data = map[string]interface{}{"a": "a", "b": "b"}
	t.Log(json.Encode(data))
	data = map[string]interface{}{"a": "a", "c": &map[string]interface{}{"d": 123}, "b": 1}
	t.Log(json.Encode(data))

}

//
//func TestJSON(t *testing.T) {
//	data := map[string]interface{}{"a": "a", "c": &map[string]interface{}{"d": 123}, "b": 1}
//	json := new(tnjson.JSON)
//
//	var p tnpool.OPool
//	n := 0
//	p.Init(4, 4)
//	for i := 0; i < 4; i++ {
//		p.AddTask(func() error {
//			for j := 0; j < 25000; j++ {
//				n++
//				json.Encode(data)
//			}
//			return nil
//		})
//	}
//	p.SetFinishCallback(func() {
//	})
//	p.Start()
//	p.Stop()
//}
//func TestJSON2(t *testing.T) {
//	data := map[string]interface{}{"a": "a", "c": &map[string]interface{}{"d": 123}, "b": 1}
//	json := new(tnjson.JSON)
//	var wg sync.WaitGroup
//	wg.Add(4)
//	for i := 0; i < 4; i++ {
//		go func() {
//			for j := 0; j < 25000; j++ {
//
//				json.Encode(data)
//			}
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//}
//
//func TestJSON3(t *testing.T) {
//	data := map[string]interface{}{"a": "a", "c": &map[string]interface{}{"d": 123}, "b": 1}
//	json := new(tnjson.JSON)
//	for i := 0; i < 100000; i++ {
//		json.Encode(data)
//	}
//}
