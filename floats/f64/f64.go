package f64

func Add(dst, x []float64) {
	AddTo(dst, dst, x)
}

func Mul(dst, x []float64) {
	MulTo(dst, dst, x)
}

func Scale(a float64, dst []float64) {
	ScaleTo(dst, a, dst)
}

func Sub(dst, x []float64) {
	SubTo(dst, dst, x)
}

func Div(dst, x []float64) {
	DivTo(dst, dst, x)
}

// TODO implement simd asm for div
// Is it ok to use CPU instruction semantics instead of go semantics?
// What are the differences?
func DivTo(dst, a, b []float64) []float64 {
	if len(dst) != len(a) || len(dst) != len(b) {
		panic(BadLen)
	}
	for i := range dst {
		dst[i] = a[i] / b[i]
	}
	return dst
}

func AddScaled(dst []float64, a float64, x []float64) {
	AddScaledTo(dst, x, a, dst)
}
