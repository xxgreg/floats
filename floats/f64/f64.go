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

func Div(dst, x []float64) {
	DivTo(dst, x, dst)
}

func Sub(dst, x []float64) {
	SubTo(dst, x, dst)
}
