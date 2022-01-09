package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/aws/aws-lambda-go/lambda"
	"gonum.org/v1/gonum/floats"

	"github.com/xxgreg/floats/f64"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "local" {
			_, _ = Handler(nil, Request{})
			return
		}
	}
	lambda.Start(Handler)
}

type Request struct{}

type Response struct{}

func Handler(ctx context.Context, req Request) (Response, error) {

	fmt.Println(
		"using_asm", UsingAsm,
		"num_cpu", runtime.NumCPU(),
		"lambda_mem", os.Getenv("AWS_LAMBDA_FUNCTION_MEMORY_SIZE"))

	bs, err := os.ReadFile("/proc/cpuinfo")
	if err == nil {
		fmt.Println(string(bs))
	}

	Test()
	BenchmarkAll()

	return Response{}, nil
}

func Test() {
	{
		n := 16
		dst := make([]float64, n)
		f64.AddConst(2, dst)
		fmt.Println(dst)
	}
	{
		n := 16
		dst := slice(3, n)
		f64.ScaleTo(dst, 2, dst)
		fmt.Println(dst)
	}
	fmt.Println("add scaled to")
	{
		n := 16
		dst := make([]float64, n)
		x := slice(2, n)
		y := slice(3, n)
		f64.AddScaledTo(dst, x, 2, y)
		fmt.Println(dst)
	}
	{
		n := 7
		dst := make([]float64, n)
		x := slice(2, n)
		y := slice(3, n)
		f64.AddScaledTo(dst, x, 2, y)
		fmt.Println(dst)
	}
}

func slice(x float64, n int) []float64 {
	s := make([]float64, n)
	for i := 0; i < n; i++ {
		s[i] = x * float64(i+1)
	}
	return s
}

func BenchmarkAll() {

	size := 0

	bms := []struct {
		name string
		fn   func(b *testing.B)
	}{
		{"AddTo", func(b *testing.B) { benchToOp(b, f64.AddTo, size) }},
		{"AddConst", func(b *testing.B) { benchAddConst(b, f64.AddConst, size) }},
		{"AddConst Gonum", func(b *testing.B) { benchAddConst(b, floats.AddConst, size) }},
		{"ScaleTo", func(b *testing.B) { benchScaleTo(b, f64.ScaleTo, size) }},
		{"ScaleTo Gonum", func(b *testing.B) { benchScaleTo(b, floats.ScaleTo, size) }},
		{"AddScaledTo", func(b *testing.B) { benchAddScaledTo(b, f64.AddScaledTo, size) }},
		{"AddScaledTo Gonum", func(b *testing.B) { benchAddScaledTo(b, floats.AddScaledTo, size) }},
		//{"MulTo", func(b *testing.B) { benchToOp(b, f64.MulTo, size) }},
		//{"AddTo Gonum", func(b *testing.B) { benchToOp(b, floats.AddTo, size) }},
		//{"MulTo Gonum", func(b *testing.B) { benchToOp(b, floats.MulTo, size) }},
		//{"AddManyTo", func(b *testing.B) { BenchmarkAddManyTo(b, size) }},
	}

	fmt.Println("function,asm,elements,ns/op")

	for _, s := range []int{15, 600} {
		size = s
		for _, bm := range bms {
			r := testing.Benchmark(bm.fn)

			asm := "asm"
			if !UsingAsm {
				asm = "noasm"
			}

			fmt.Printf("%v,%v,%v,%v\n", bm.name, asm, size, r.NsPerOp())
		}
	}
}

func BenchmarkAddManyTo(b *testing.B, size int) {
	dst := make([]float64, size)

	a := make([]float64, size)
	for i := range a {
		a[i] = float64(i) * 0.078
	}

	b2 := make([]float64, size)
	for i := range a {
		b2[i] = float64(i) * 0.045
	}

	c := make([]float64, size)
	for i := range a {
		c[i] = float64(i) * 0.067
	}

	d := make([]float64, size)
	for i := range a {
		d[i] = float64(i) * 0.023
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f64.AddManyTo(dst, a, b2, c, d)
	}
}

func benchOp(b *testing.B, f func(dst, a []float64), size int) {
	dst := make([]float64, size)
	for i := range dst {
		dst[i] = float64(i) * 0.036
	}

	y := make([]float64, size)
	for i := range y {
		y[i] = float64(i) * 0.078
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(dst, y)
	}
}

func benchToOp(b *testing.B, f func(dst, a, b []float64) []float64, size int) {
	dst := make([]float64, size)

	x := make([]float64, size)
	for i := range x {
		x[i] = float64(i) * 0.036
	}

	y := make([]float64, size)
	for i := range y {
		y[i] = float64(i) * 0.078
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(dst, x, y)
	}
}

func benchAddConst(b *testing.B, f func(c float64, dst []float64), size int) {
	dst := make([]float64, size)

	x := make([]float64, size)
	for i := range x {
		x[i] = float64(i) * 0.036
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(2, dst)
	}
}

func benchScaleTo(b *testing.B, f func(dst []float64, c float64, a []float64) []float64, size int) {
	dst := make([]float64, size)

	x := make([]float64, size)
	for i := range x {
		x[i] = float64(i) * 0.036
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(dst, 2, x)
	}
}

func benchAddScaledTo(b *testing.B, f func(dst, x []float64, a float64, y []float64) []float64, size int) {
	dst := make([]float64, size)

	x := make([]float64, size)
	for i := range x {
		x[i] = float64(i) * 0.036
	}

	y := make([]float64, size)
	for i := range y {
		y[i] = float64(i) * 0.078
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(dst, x, 2, y)
	}
}
