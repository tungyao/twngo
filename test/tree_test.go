package test

import (
	"../tnwb"
	"testing"
)

func TestTree(t *testing.T) {
	tre := tnwb.NewTrie()
	tre.Insert("ab", "value->ab")
	//tre.Insert("abc", "value->ab")
	tre.Find("abc")
	//tre.Insert("aab", "value->aab")
	//tre.Insert("aabb", "value->aabb")

}
