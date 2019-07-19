package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiVal()(int ,int )  {
	return rand.Intn(10),rand.Intn(20)
}
func TestFuncOne(t *testing.T) {
	t.Log(returnMultiVal())
}

func calculationTime(inner func (op int) int) func(op int) int  {
	return func(n int) int {
		start :=time.Now()
		ret :=inner(n)
		fmt.Println("Time spend", time.Since(start).Seconds())
		return ret
	}
}
func sleepIt(op int) int{
	time.Sleep(1*time.Second)
	return op
}
func TestFuncTwo(t *testing.T) {
	calculationTime(sleepIt)(10)
}
func addLate(value int) func()int{
	return func() int {
		value++
		return value
	}
}
func TestFuncThree(t *testing.T) {
	add := addLate(1)
	val :=0
	for i:=0;i<10;i++ {
		val +=add()
	}
	t.Log(val)
}

func closeFunc(n int)func()(int,string){
	s:="hello"
	return func()(int ,string){
		return n,s
	}
}
func TestCloseFunc(t *testing.T) {
	n,_ := closeFunc(100)()
	t.Log(n)
	t.Log(closeFunc(200)())
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func TestParms(t *testing.T) {
	MyPrintf(12)
}

func timeCalculation(fun func()){
	start := time.Now()
	fun()
	end := time.Now()
	result := end.Sub(start)
	fmt.Printf("该函数执行完成耗时: %s\n", result)
}
func add(n int)  {
	sum := n
	for i := 0; i < 100000000; i++ {
		sum += i
	}
}
func TestTimeCalculation(t *testing.T)  {
	add := add
	timeCalculation(func() {
		add(10)
	})
}

