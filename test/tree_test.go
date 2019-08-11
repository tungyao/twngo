package test

import (
	"../tnwb"
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tre := tnwb.NewTrie()
	tre.Insert("ab", func() {
		fmt.Println("/ab")
	})

	tre.Insert("ba", func() {
		fmt.Println("/ba")
	})

	if fun := tre.Find("ba"); fun != nil {
		fun()
	}

}
