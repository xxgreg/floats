//go:build noasm || (!arm64 && !amd64)

package f64

const BadLen = "floats: mismatched slice lengths"

func AddTo(dst, a, b []float64) []float64 {
	if len(dst) != len(a) || len(dst) != len(b) {
		panic(BadLen)
	}
	for i := range dst {
		dst[i] = a[i] + b[i]
	}
	return dst
}

func AddConst(c float64, dst []float64) {
	for i := range dst {
		dst[i] += c
	}
}
func AddManyTo(dst, a, b, c, d []float64) []float64 {
	n := len(dst)
	if len(a) != n || len(b) != n || len(c) != n || len(d) != n {
		panic(BadLen)
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] + b[i] + c[i] + d[i]
	}
	return dst
}

func MulTo(dst, a, b []float64) []float64 {
	if len(dst) != len(a) || len(dst) != len(b) {
		panic(BadLen)
	}
	for i := range dst {
		dst[i] = a[i] * b[i]
	}
	return dst
}

func ScaleTo(dst []float64, c float64, s []float64) []float64 {
	if len(dst) != len(s) || len(dst) != len(s) {
		panic(BadLen)
	}
	for i := range dst {
		dst[i] = s[i] * c
	}
	return dst
}

func AddScaledTo(dst, y []float64, a float64, x []float64) []float64 {
	n := len(dst)
	if len(dst) != n || len(y) != n || len(x) != n {
		panic(BadLen)
	}
	for i := range x {
		dst[i] = a*x[i] + y[i]
	}
	return dst
}

func SubTo(dst, a, b []float64) []float64 {
	if len(dst) != len(a) || len(dst) != len(b) {
		panic(BadLen)
	}
	for i := range dst {
		dst[i] = a[i] - b[i]
	}
	return dst
}

func DivTo(dst, a, b []float64) []float64 {
	if len(dst) != len(a) || len(dst) != len(b) {
		panic(BadLen)
	}
	for i := range dst {
		dst[i] = a[i] / b[i]
	}
	return dst
}
