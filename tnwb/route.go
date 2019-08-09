package tnwb

import "fmt"

const sonNum = 52
const base = 'a'

type Trie struct {
	num  int64
	node *Son
}
type Son struct {
	key      string
	value    string
	deep     int
	child    map[rune]*Son
	terminal bool
}

func NewSon(key string, value string, deep int) *Son {
	return &Son{
		key:   key,
		value: value,
		deep:  deep,
		child: make(map[rune]*Son),
	}
}
func NewTree() *Trie {
	return &Trie{
		num:  0,
		node: NewSon("k", "kkk", 1),
	}
}
func (t *Trie) Insert(key string, value string) {
	var currentNode = t.node
	if key != currentNode.key {
		for _, v := range key {
			son, ok := currentNode.child[v]
			fmt.Println(v, ok)
			if !ok {
				son = NewSon(key, value, currentNode.deep+1)
				currentNode.child[v] = son
			}
			currentNode = son

		}
	}
	currentNode.value = value
	currentNode.terminal = true
	t.node = currentNode
	fmt.Println(t.node)
}
