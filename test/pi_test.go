package test

import (
	"math"
	"testing"
)

func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

//func TestPI(t *testing.T) {
//	t.Log(pi(500000))
//}
func TestFBNAQIE(t *testing.T) {
	for n:=10;n<100000;n++{
		ans := 0
		temp := n
		lens := 0             //记录n的位数
		for temp/10!=0 || temp%10!=0 {
			temp/=10
			lens ++
		}
		ts := n
		for ts!=0 {
			te := ts%10
			ans += int(math.Pow(float64(te),float64(lens)))
			ts = ts/10
		}
		if ans == n {
			t.Log(ans)
		}
	}
}
