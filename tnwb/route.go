package tnwb

import (
	"fmt"
	"strings"
)

type Trie struct {
	num  int64
	root *Son
}
type Son struct {
	key      string
	path     string
	deep     int
	child    map[string]*Son
	terminal bool
	handler  func()
}

func NewSon(path string, handler func(), deep int) *Son {
	return &Son{
		key:     path,
		path:    path,
		deep:    deep,
		handler: handler,
		child:   make(map[string]*Son),
	}
}
func NewTrie() *Trie {
	return &Trie{
		num: 1,
		root: NewSon("a", func() {
			fmt.Println("/a")
		}, 1),
	}
}
func (t *Trie) Insert(path string, handler func()) {
	son := t.root
	res := strings.Split(path, "")
	if son.key != path { //匹配不成功才加入链表
		for _, key := range res {
			if son.child[key] == nil {
				node := NewSon(key, handler, son.deep+1)
				node.child = make(map[string]*Son)
				node.terminal = false
				son.child[key] = node
			}
			son = son.child[key]
		}

	}
	son.terminal = true
	son.handler = handler
	son.path = path
}
func (t *Trie) Find(key string) func() {
	son := t.root
	res := strings.Split(key, "")
	path := ""
	var han func()
	if son.key != key { //匹配不成功才加入链表
		for _, key := range res {
			if son.child[key] == nil {
				return nil
			} else {
				path += son.child[key].key
				han = son.child[key].handler
			}
			son = son.child[key]
		}
	}
	fmt.Println(path)
	return han
}
