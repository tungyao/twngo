package test

import "testing"

func TestArray(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	t.Log(arr[1:5])
}
func TestFor(t *testing.T) {
	n := 0
	for n > 100 {
		n++
	}
	t.Log(n)
	for {
		n++
		if n > 200 {
			break
		}
	}
	t.Log(n)
}
