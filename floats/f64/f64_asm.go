//go:build !noasm && (arm64 || amd64)

package f64

const BadLen = "floats: mismatched slice lengths"

func Add(dst, a []float64) {
	n := len(dst)
	if len(a) != n {
		panic(BadLen)
	}
	tail := n % 8
	if n >= 8 {
		add8(&dst[0], &dst[0], &a[0], n-tail)
	}
	for i := n - tail; i < n; i++ {
		dst[i] = dst[i] + a[i]
	}
}

func AddTo(dst, a, b []float64) []float64 {
	n := len(dst)
	if len(a) != n || len(b) != n {
		panic(BadLen)
	}
	tail := n % 8
	if n >= 8 {
		add8(&dst[0], &a[0], &b[0], n-tail)
	}
	for i := n - tail; i < n; i++ {
		dst[i] = a[i] + b[i]
	}
	return dst
}

func AddManyTo(dst, a, b, c, d []float64) []float64 {
	n := len(dst)
	if len(a) != n || len(b) != n || len(c) != n || len(d) != n {
		panic(BadLen)
	}
	tail := n % 8
	if n >= 8 {
		add8_4(&dst[0], &a[0], &b[0], &c[0], &d[0], n-tail)
	}
	for i := n - tail; i < n; i++ {
		dst[i] = a[i] + b[i] + c[i] + d[i]
	}
	return dst
}

func Mul(dst, a []float64) {
	n := len(dst)
	if len(a) != n {
		panic(BadLen)
	}
	tail := n % 8
	if n >= 8 {
		mul8(&dst[0], &dst[0], &a[0], n-tail)
	}
	for i := n - tail; i < n; i++ {
		dst[i] = dst[i] + a[i]
	}
}

func MulTo(dst, a, b []float64) []float64 {
	n := len(dst)
	if len(a) != n || len(b) != n {
		panic(BadLen)
	}
	tail := n % 8
	if n >= 8 {
		mul8(&dst[0], &a[0], &b[0], n-tail)
	}
	for i := n - tail; i < n; i++ {
		dst[i] = a[i] + b[i]
	}
	return dst
}

func add8(dst, a, b *float64, n int)

func mul8(dst, a, b *float64, n int)

func add8_4(dst, a, b, c, d *float64, n int)
