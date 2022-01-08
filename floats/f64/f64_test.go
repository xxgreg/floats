package f64

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	f := func(n, m int) {
		dst := make([]float64, n)
		a := slice(1, n)
		AddTo(dst[:m], dst[:m], a[:m])
		fmt.Println(dst, a)
	}

	for i := 0; i <= 17; i++ {
		f(24, i)
	}

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
	MulTo(dst[:16], dst[:16], a[:16])
	fmt.Println(dst)

	n = 24
	dst = make([]float64, n)
	a = slice(1, n)
	b := slice(2, n)
	MulTo(dst[:16], a[:16], b[:16])
	fmt.Println(dst)
}

func TestAddMany(t *testing.T) {

	{
		n := 16
		dst := make([]float64, n)
		a := slice(2, n)
		b := slice(3, n)
		c := slice(4, n)
		d := slice(4, n)

		AddManyToNoAsm(dst, a, b, c, d)
		fmt.Println("AddManyNoAsm", dst)
	}

	{
		n := 16
		dst := make([]float64, n)
		a := slice(2, n)
		b := slice(3, n)
		c := slice(4, n)
		d := slice(5, n)

		AddManyToSimple(dst, a, b, c, d)
		fmt.Println("AddManySimple", dst)
	}

	{
		n := 16
		dst := make([]float64, n)
		a := slice(2, n)
		b := slice(3, n)
		c := slice(4, n)
		d := slice(5, n)

		AddManyTo(dst, a, b, c, d)
		fmt.Println("AddMany", dst)
	}
}

func slice(x float64, n int) []float64 {
	s := make([]float64, n)
	for i := 0; i < n; i++ {
		s[i] = x * float64(i+1)
	}
	return s
}

func AddManyToNoAsm(dst, a, b, c, d []float64) {
	n := len(dst)
	if len(a) != n || len(b) != n || len(c) != n || len(d) != n {
		panic(BadLen)
	}
	for i := 0; i < n; i++ {
		dst[i] = dst[i] + a[i] + b[i] + c[i] + d[i]
	}
}

func AddManyToSimple(dst, a, b, c, d []float64) {
	AddTo(dst, dst, a)
	AddTo(dst, dst, b)
	AddTo(dst, dst, c)
	AddTo(dst, dst, d)
}

func BenchmarkAddManyNoAsm(b *testing.B) {

	n := 600
	dst := make([]float64, n)
	a := slice(2, n)
	b2 := slice(3, n)
	c := slice(4, n)
	d := slice(5, n)

	for i := 0; i < b.N; i++ {
		AddManyToNoAsm(dst, a, b2, c, d)
	}
}

func BenchmarkAddManySimple(b *testing.B) {

	n := 600
	dst := make([]float64, n)
	a := slice(2, n)
	b2 := slice(3, n)
	c := slice(4, n)
	d := slice(5, n)

	for i := 0; i < b.N; i++ {
		AddManyToSimple(dst, a, b2, c, d)
	}
}

func BenchmarkAddMany(b *testing.B) {

	n := 600
	dst := slice(1, n)
	a := slice(2, n)
	b2 := slice(3, n)
	c := slice(4, n)
	d := slice(5, n)

	for i := 0; i < b.N; i++ {
		AddManyTo(dst, a, b2, c, d)
	}
}
