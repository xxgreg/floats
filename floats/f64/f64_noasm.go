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
