package f64

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	n := 10
	dst := slice(1, n)
	a := slice(2, n)
	Add(dst, a)
	fmt.Println(dst)

	//n = 24
	//dst = make([]float64, n)
	//a = slice(1, n)
	//b := slice(2, n)
	//AddTo(dst[:10], a[:10], b[:10])
	//fmt.Println(dst)
}

func TestMul(t *testing.T) {

	n := 24
	dst := slice(1, n)
	a := slice(2, n)
	Mul(dst[:16], a[:16])
	fmt.Println(dst)

	n = 24
	dst = make([]float64, n)
	a = slice(1, n)
	b := slice(2, n)
	MulTo(dst[:16], a[:16], b[:16])
	fmt.Println(dst)
}

func slice(x float64, n int) []float64 {
	s := make([]float64, n)
	for i := 0; i < n; i++ {
		s[i] = x * float64(i)
	}
	return s
}
