//go:build noasm || (!arm64 && !amd64)

package f64

const BadLen = "floats: mismatched slice lengths"

func Add(dst, a []float64) {
	if len(dst) != len(a) {
		panic(BadLen)
	}
	for i := range dst {
		dst[i] = dst[i] + a[i]
	}
}

func AddTo(dst, a, b []float64) []float64 {
	if len(dst) != len(a) || len(dst) != len(b) {
		panic(BadLen)
	}
	for i := range dst {
		dst[i] = a[i] + b[i]
	}
	return dst
}

func Mul(dst, a []float64) {
	if len(dst) != len(a) {
		panic(BadLen)
	}
	for i := range dst {
		dst[i] = dst[i] * a[i]
	}
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
