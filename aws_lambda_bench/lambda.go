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
	n := 16
	{
		a := make([]float64, 24)
		b := make([]float64, 24)
		for i := 0; i < 16; i++ {
			a[i] = float64(i + 1)
			b[i] = float64((i + 1) * 2)
		}
		f64.Add(a[:n], b[:n])
		fmt.Println("Add", a, b)
	}
}

func BenchmarkAll() {

	size := 0

	bms := []struct {
		name string
		fn   func(b *testing.B)
	}{
		{"Add", func(b *testing.B) { benchOp(b, f64.Add, size) }},
		{"AddTo", func(b *testing.B) { benchToOp(b, f64.AddTo, size) }},
		{"Mul", func(b *testing.B) { benchOp(b, f64.Mul, size) }},
		{"MulTo", func(b *testing.B) { benchToOp(b, f64.MulTo, size) }},
		{"Add Gonum", func(b *testing.B) { benchOp(b, floats.Add, size) }},
		{"AddTo Gonum", func(b *testing.B) { benchToOp(b, floats.AddTo, size) }},
		{"Mul Gonum", func(b *testing.B) { benchOp(b, floats.Mul, size) }},
		{"MulTo Gonum", func(b *testing.B) { benchToOp(b, floats.MulTo, size) }},
	}

	for _, s := range []int{8, 64, 600, 1000, 10000} {
		size = s
		for _, bm := range bms {
			r := testing.Benchmark(bm.fn)
			nsPerElement := fmt.Sprintf("%v ns/el", float64(r.NsPerOp())/float64(size))
			fmt.Println(bm.name, size, r, nsPerElement)
		}
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
