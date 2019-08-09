package test

import (
	"../tnwb"
	"testing"
)

func TestTree(t *testing.T) {
	tre := tnwb.NewTree()
	tre.Insert("k1", "values")
	tre.Insert("k12", "values")
	tre.Insert("k2", "values")
	tre.Insert("k21", "values")
	//tre.Insert("k1","values")
	//tre.Insert("k1","values")
	//tre.Insert("k1","values")
	//tre.Insert("k2","valuess")
	//tre.Insert("k3","asdddd")
	//tre.Insert("k3a","asdddd")
	//tnwb.Insert(tre, []byte("/a/b"),len([]byte("/a/b")))
	//tnwb.Insert(tre, []byte("/a/b/c"),len([]byte("/a/b/c")))
	//
	//time.Sleep(time.Second*1)
	//
	//t.Log(tnwb.Find(tre,[]byte("/a"),len([]byte("/a"))))
	//t.Log(tnwb.Find(tre,[]byte("/a/asd"),len([]byte("/a/asd"))))
	//t.Log(tnwb.Find(tre,[]byte("/a/b/c"),len([]byte("/a/b/c"))))

}
