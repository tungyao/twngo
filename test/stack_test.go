package test

import "testing"

func calc(a, b int) int {
	var c int
	c = a * b
	var x int
	x = c * 10
	return x
}
func dummy(a int) int {
	var b int
	b = a
	return b

}
func void() {

}
func TestStack(t *testing.T) {
	t.Log(calc(10, 20))

}
func TestAuto(t *testing.T) {
	var a int
	void()
	t.Log(a)
	t.Log(dummy(10))
}

type Data struct {
}

func dummys() *Data {
	var c Data
	return &c
}
func TestStruct(t *testing.T) {
	t.Log(dummys())

}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	tuesday
	wednesday
	Thursday
	Friday
	Saturday
)

func TestWeek(t *testing.T) {
	t.Log(Weekday(Monday))
}

type ChipType int

const (
	None ChipType = iota
	CPU           // 中央处理器
	GPU           // 图形处理器
)

func (b ChipType) String() string {
	switch b {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}
func TestChipType(t *testing.T) {
	t.Log(CPU)
}
func TestMapWithFunValue(t *testing.T) {
	mp := map[int]int{}
	t.Logf("%p", &mp)
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](3))
}
func TestMapForSet(t *testing.T) {
	myset := map[int]bool{}
	myset[1] = true
	n := 1
	if myset[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	t.Log(len(myset))
}
func TestString(t *testing.T) {
	s := "救"
	t.Log(len(s))
	c := []rune(s)
	t.Log(len(c))
}
func TestArrOne(t *testing.T) {
	var a [3]int
	for _, v := range a {
		t.Log(v)
	}
}
