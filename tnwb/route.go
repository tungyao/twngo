package tnwb

import (
	"fmt"
	"strings"
)

const sonNum = 52
const base = 'a'

type Trie struct {
	num  int64
	root *Son
}
type Son struct {
	key      string
	path     string
	value    string
	deep     int
	child    map[string]*Son
	terminal bool
}

func NewSon(path string, value string, deep int) *Son {
	return &Son{
		key:   path,
		value: value,
		path:  path,
		deep:  deep,
		child: make(map[string]*Son),
	}
}
func NewTrie() *Trie {
	return &Trie{
		num:  1,
		root: NewSon("a", "value->a", 1),
	}
}
func (t *Trie) Insert(path string, value string) {
	son := t.root
	res := strings.Split(path, "")
	if son.key != path { //匹配不成功才加入链表
		for _, key := range res {
			node, ok := son.child[key] //看字son中存不存在该 key
			if !ok {                   //不存在就加进去
				node = NewSon(key, value, son.deep+1) //这个其实是son的son
			}
			son.child[key] = node

		}
	}
	son.terminal = true
	son.value = value
	son.path = path
	t.root = son
}
func (t *Trie) Find(key string) {
	son := t.root
	res := strings.Split(key, "")
	if son.key != key { //匹配不成功才加入链表
		for _, key := range res {
			node, ok := son.child[key]
			fmt.Println(node)

			if ok {
				son = node
				fmt.Println("找到：", node.key, key)
			} else {
				fmt.Println("没找到", key)
				return
			}
		}
	}
}
