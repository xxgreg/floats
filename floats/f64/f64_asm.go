//go:build !noasm && (arm64 || amd64)

package f64

const BadLen = "floats: mismatched slice lengths"

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
		dst[i] = a[i] * b[i]
	}
	return dst
}

func AddConst(c float64, dst []float64) {
	n := len(dst)
	tail := n % 8
	if n >= 8 {
		addConst8(&dst[0], &dst[0], c, n-tail)
	}
	for i := n - tail; i < n; i++ {
		dst[i] = dst[i] + c
	}
}

func ScaleTo(dst []float64, a float64, x []float64) []float64 {
	n := len(dst)
	if len(x) != n {
		panic(BadLen)
	}
	tail := n % 8
	if n >= 8 {
		scale8(&dst[0], &x[0], a, n-tail)
	}
	for i := n - tail; i < n; i++ {
		dst[i] = x[i] * a
	}
	return dst
}

func AddScaled(dst []float64, a float64, x []float64) {
	n := len(dst)
	if len(x) != n {
		panic(BadLen)
	}
	tail := n % 8
	if n >= 8 {
		//FIXME FMA isn't much faster than add scaled, so I'll just delete this.
		fma8(&dst[0], &x[0], a, n-tail)
	}
	for i := n - tail; i < n; i++ {
		dst[i] = a*x[i] + dst[i]
	}
}

func AddScaledTo(dst, x []float64, a float64, y []float64) []float64 {
	n := len(dst)
	if len(x) != n || len(y) != n {
		panic(BadLen)
	}
	tail := n % 8
	if n >= 8 {
		//FIXME I've swapped the order of x and y somewhere. in x86?
		addScaled8(&dst[0], &y[0], &x[0], a, n-tail)
	}
	for i := n - tail; i < n; i++ {
		dst[i] = a*x[i] + y[i]
	}
	return dst
}

func DivTo(dst, x, y []float64) []float64 {
	panic("TODO")
	//TODO what about div by zero = NaN?
}

func SubTo(dst, x, y []float64) []float64 {
	n := len(dst)
	if len(x) != n || len(y) != n {
		panic(BadLen)
	}
	tail := n % 8
	if n >= 8 {
		sub8(&dst[0], &x[0], &y[0], n-tail)
	}
	for i := n - tail; i < n; i++ {
		dst[i] = x[i] - y[i]
	}
	return dst
}

//TODO rename to x,y
func add8(dst, a, b *float64, n int)

func sub8(dst, a, b *float64, n int)

func mul8(dst, a, b *float64, n int)

func add8_4(dst, a, b, c, d *float64, n int)

func addConst8(dst, x *float64, c float64, n int)

func scale8(dst, x *float64, c float64, n int)

func fma8(dst, x *float64, a float64, n int)

func addScaled8(dst, x, y *float64, a float64, n int)
