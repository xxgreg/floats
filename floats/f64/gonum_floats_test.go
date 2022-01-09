// This source is a modified version of https://github.com/gonum/gonum/blob/master/floats/floats_test.go

// Copyright ©2013 The Gonum Authors. All rights reserved.
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://github.com/gonum/gonum/blob/master/LICENSE

// Copyright ©2013 The Gonum Authors. All rights reserved.
//
//Redistribution and use in source and binary forms, with or without
//modification, are permitted provided that the following conditions are met:
//    * Redistributions of source code must retain the above copyright
//      notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above copyright
//      notice, this list of conditions and the following disclaimer in the
//      documentation and/or other materials provided with the distribution.
//    * Neither the name of the Gonum project nor the names of its authors and
//      contributors may be used to endorse or promote products derived from this
//      software without specific prior written permission.
//
//THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
//ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
//WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
//DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
//FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
//DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
//SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
//CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
//OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
//OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package f64

import (
	"fmt"
	"testing"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/floats/scalar"
)

const (
	EqTolerance = 1e-14
	Small       = 10
	Medium      = 1000
	Large       = 100000
	Huge        = 10000000
)

func areSlicesEqual(t *testing.T, truth, comp []float64, str string) {
	if !floats.EqualApprox(comp, truth, EqTolerance) {
		t.Errorf(str+". Expected %v, returned %v", truth, comp)
	}
}

func areSlicesSame(t *testing.T, truth, comp []float64, str string) {
	ok := len(truth) == len(comp)
	if ok {
		for i, a := range truth {
			if !scalar.EqualWithinAbsOrRel(a, comp[i], EqTolerance, EqTolerance) && !scalar.Same(a, comp[i]) {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf(str+". Expected %v, returned %v", truth, comp)
	}
}

func Panics(fun func()) (b bool) {
	defer func() {
		err := recover()
		if err != nil {
			b = true
		}
	}()
	fun()
	return
}

func TestAdd2(t *testing.T) {
	t.Parallel()
	a := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := []float64{4, 5, 6, 7, 8, 9, 10, 11, 12}
	c := []float64{7, 8, 9, 11, 13, 15, 17, 19, 21}
	truth := []float64{12, 15, 18, 22, 26, 30, 34, 38, 42}
	n := make([]float64, len(a))

	Add(n, a)
	Add(n, b)
	Add(n, c)
	areSlicesEqual(t, truth, n, "Wrong addition of slices new receiver")
	Add(a, b)
	Add(a, c)
	areSlicesEqual(t, truth, n, "Wrong addition of slices for no new receiver")

	// Test that it panics
	if !Panics(func() { Add(make([]float64, 2), make([]float64, 3)) }) {
		t.Errorf("Did not panic with length mismatch")
	}
}

func TestAddTo(t *testing.T) {
	t.Parallel()
	a := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := []float64{4, 5, 6, 7, 8, 9, 10, 11, 12}
	truth := []float64{5, 7, 9, 11, 13, 15, 17, 19, 21}
	n1 := make([]float64, len(a))

	n2 := AddTo(n1, a, b)
	areSlicesEqual(t, truth, n1, "Bad addition from mutator")
	areSlicesEqual(t, truth, n2, "Bad addition from returned slice")

	// Test that it panics
	if !Panics(func() { AddTo(make([]float64, 2), make([]float64, 3), make([]float64, 3)) }) {
		t.Errorf("Did not panic with length mismatch")
	}
	if !Panics(func() { AddTo(make([]float64, 3), make([]float64, 3), make([]float64, 2)) }) {
		t.Errorf("Did not panic with length mismatch")
	}
}

func TestAddConst2(t *testing.T) {
	t.Parallel()
	s := []float64{3, 4, 1, 7, 5, 3, 4, 1, 7}
	c := 6.0
	truth := []float64{9, 10, 7, 13, 11, 9, 10, 7, 13}
	AddConst(c, s)
	areSlicesEqual(t, truth, s, "Wrong addition of constant")
}

func TestAddScaled(t *testing.T) {
	t.Parallel()
	s := []float64{3, 4, 1, 7, 5, 3, 4, 1, 7}
	alpha := 6.0
	dst := []float64{1, 2, 3, 4, 5, 1, 2, 3, 4}
	ans := []float64{19, 26, 9, 46, 35, 19, 26, 9, 46}
	AddScaled(dst, alpha, s)
	if !floats.EqualApprox(dst, ans, EqTolerance) {
		t.Errorf("Adding scaled did not match")
	}
	short := []float64{1}
	if !Panics(func() { AddScaled(dst, alpha, short) }) {
		t.Errorf("Doesn't panic if s is smaller than dst")
	}
	if !Panics(func() { AddScaled(short, alpha, s) }) {
		t.Errorf("Doesn't panic if dst is smaller than s")
	}
}

func TestAddScaledTo(t *testing.T) {
	t.Parallel()
	s := []float64{3, 4, 1, 7, 5, 3, 4, 1, 7}
	alpha := 6.0
	y := []float64{1, 2, 3, 4, 5, 1, 2, 3, 4}
	dst1 := make([]float64, 9)
	ans := []float64{19, 26, 9, 46, 35, 19, 26, 9, 46}
	dst2 := AddScaledTo(dst1, y, alpha, s)
	if !floats.EqualApprox(dst1, ans, EqTolerance) {
		t.Errorf("AddScaledTo did not match for mutator")
	}
	if !floats.EqualApprox(dst2, ans, EqTolerance) {
		t.Errorf("AddScaledTo did not match for returned slice")
	}
	AddScaledTo(dst1, y, alpha, s)
	if !floats.EqualApprox(dst1, ans, EqTolerance) {
		t.Errorf("Reusing dst did not match")
	}
	short := []float64{1}
	if !Panics(func() { AddScaledTo(dst1, y, alpha, short) }) {
		t.Errorf("Doesn't panic if s is smaller than dst")
	}
	if !Panics(func() { AddScaledTo(short, y, alpha, s) }) {
		t.Errorf("Doesn't panic if dst is smaller than s")
	}
	if !Panics(func() { AddScaledTo(dst1, short, alpha, s) }) {
		t.Errorf("Doesn't panic if y is smaller than dst")
	}
}

//TODO
//func TestDiv(t *testing.T) {
//	t.Parallel()
//	s1 := []float64{5, 12, 27}
//	s2 := []float64{1, 2, 3}
//	ans := []float64{5, 6, 9}
//	Div(s1, s2)
//	if !EqualApprox(s1, ans, EqTolerance) {
//		t.Errorf("Div doesn't give correct answer")
//	}
//	s1short := []float64{1}
//	if !Panics(func() { Div(s1short, s2) }) {
//		t.Errorf("Did not panic with unequal lengths")
//	}
//	s2short := []float64{1}
//	if !Panics(func() { Div(s1, s2short) }) {
//		t.Errorf("Did not panic with unequal lengths")
//	}
//}

//TODO
//func TestDivTo(t *testing.T) {
//	t.Parallel()
//	s1 := []float64{5, 12, 27}
//	s1orig := []float64{5, 12, 27}
//	s2 := []float64{1, 2, 3}
//	s2orig := []float64{1, 2, 3}
//	dst1 := make([]float64, 3)
//	ans := []float64{5, 6, 9}
//	dst2 := DivTo(dst1, s1, s2)
//	if !EqualApprox(dst1, ans, EqTolerance) {
//		t.Errorf("DivTo doesn't give correct answer in mutated slice")
//	}
//	if !EqualApprox(dst2, ans, EqTolerance) {
//		t.Errorf("DivTo doesn't give correct answer in returned slice")
//	}
//	if !EqualApprox(s1, s1orig, EqTolerance) {
//		t.Errorf("S1 changes during multo")
//	}
//	if !EqualApprox(s2, s2orig, EqTolerance) {
//		t.Errorf("s2 changes during multo")
//	}
//	DivTo(dst1, s1, s2)
//	if !EqualApprox(dst1, ans, EqTolerance) {
//		t.Errorf("DivTo doesn't give correct answer reusing dst")
//	}
//	dstShort := []float64{1}
//	if !Panics(func() { DivTo(dstShort, s1, s2) }) {
//		t.Errorf("Did not panic with s1 wrong length")
//	}
//	s1short := []float64{1}
//	if !Panics(func() { DivTo(dst1, s1short, s2) }) {
//		t.Errorf("Did not panic with s1 wrong length")
//	}
//	s2short := []float64{1}
//	if !Panics(func() { DivTo(dst1, s1, s2short) }) {
//		t.Errorf("Did not panic with s2 wrong length")
//	}
//}

func TestMul2(t *testing.T) {
	t.Parallel()
	s1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans := []float64{1, 4, 9, 16, 25, 36, 49, 64, 81}
	Mul(s1, s2)
	if !floats.EqualApprox(s1, ans, EqTolerance) {
		fmt.Println(s1)
		t.Errorf("Mul doesn't give correct answer")
	}
	s1short := []float64{1}
	if !Panics(func() { Mul(s1short, s2) }) {
		t.Errorf("Did not panic with unequal lengths")
	}
	s2short := []float64{1}
	if !Panics(func() { Mul(s1, s2short) }) {
		t.Errorf("Did not panic with unequal lengths")
	}
}

func TestMulTo(t *testing.T) {
	t.Parallel()
	s1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1orig := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2orig := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	dst1 := make([]float64, 9)
	ans := []float64{1, 4, 9, 16, 25, 36, 49, 64, 81}
	dst2 := MulTo(dst1, s1, s2)
	if !floats.EqualApprox(dst1, ans, EqTolerance) {
		t.Errorf("MulTo doesn't give correct answer in mutated slice")
	}
	if !floats.EqualApprox(dst2, ans, EqTolerance) {
		t.Errorf("MulTo doesn't give correct answer in returned slice")
	}
	if !floats.EqualApprox(s1, s1orig, EqTolerance) {
		t.Errorf("S1 changes during multo")
	}
	if !floats.EqualApprox(s2, s2orig, EqTolerance) {
		t.Errorf("s2 changes during multo")
	}
	MulTo(dst1, s1, s2)
	if !floats.EqualApprox(dst1, ans, EqTolerance) {
		t.Errorf("MulTo doesn't give correct answer reusing dst")
	}
	dstShort := []float64{1}
	if !Panics(func() { MulTo(dstShort, s1, s2) }) {
		t.Errorf("Did not panic with s1 wrong length")
	}
	s1short := []float64{1}
	if !Panics(func() { MulTo(dst1, s1short, s2) }) {
		t.Errorf("Did not panic with s1 wrong length")
	}
	s2short := []float64{1}
	if !Panics(func() { MulTo(dst1, s1, s2short) }) {
		t.Errorf("Did not panic with s2 wrong length")
	}
}

func TestScale(t *testing.T) {
	t.Parallel()
	s := []float64{3, 4, 1, 7, 5, 3, 4, 1, 7}
	c := 5.0
	truth := []float64{15, 20, 5, 35, 25, 15, 20, 5, 35}
	Scale(c, s)
	areSlicesEqual(t, truth, s, "Bad scaling")
}

func TestScaleTo2(t *testing.T) {
	t.Parallel()
	s := []float64{3, 4, 1, 7, 5, 3, 4, 1, 7}
	sCopy := make([]float64, len(s))
	copy(sCopy, s)
	c := 5.0
	truth := []float64{15, 20, 5, 35, 25, 15, 20, 5, 35}
	dst := make([]float64, len(s))
	ScaleTo(dst, c, s)
	if !floats.Same(dst, truth) {
		t.Errorf("Scale to does not match. Got %v, want %v", dst, truth)
	}
	if !floats.Same(s, sCopy) {
		t.Errorf("Source modified during call. Got %v, want %v", s, sCopy)
	}
	if !Panics(func() { ScaleTo(dst, 0, []float64{1}) }) {
		t.Errorf("Expected panic with different slice lengths")
	}
}

//TODO
//func TestSub(t *testing.T) {
//	t.Parallel()
//	s := []float64{3, 4, 1, 7, 5}
//	v := []float64{1, 2, 3, 4, 5}
//	truth := []float64{2, 2, -2, 3, 0}
//	Sub(s, v)
//	areSlicesEqual(t, truth, s, "Bad subtract")
//	// Test that it panics
//	if !Panics(func() { Sub(make([]float64, 2), make([]float64, 3)) }) {
//		t.Errorf("Did not panic with length mismatch")
//	}
//}

//TODO
//func TestSubTo(t *testing.T) {
//	t.Parallel()
//	s := []float64{3, 4, 1, 7, 5}
//	v := []float64{1, 2, 3, 4, 5}
//	truth := []float64{2, 2, -2, 3, 0}
//	dst1 := make([]float64, len(s))
//	dst2 := SubTo(dst1, s, v)
//	areSlicesEqual(t, truth, dst1, "Bad subtract from mutator")
//	areSlicesEqual(t, truth, dst2, "Bad subtract from returned slice")
//	// Test that all mismatch combinations panic
//	if !Panics(func() { SubTo(make([]float64, 2), make([]float64, 3), make([]float64, 3)) }) {
//		t.Errorf("Did not panic with dst different length")
//	}
//	if !Panics(func() { SubTo(make([]float64, 3), make([]float64, 2), make([]float64, 3)) }) {
//		t.Errorf("Did not panic with subtractor different length")
//	}
//	if !Panics(func() { SubTo(make([]float64, 3), make([]float64, 3), make([]float64, 2)) }) {
//		t.Errorf("Did not panic with subtractee different length")
//	}
//}

func randomSlice(l int, src rand.Source) []float64 {
	rnd := rand.New(src)
	s := make([]float64, l)
	for i := range s {
		s[i] = rnd.Float64()
	}
	return s
}

func benchmarkAdd(b *testing.B, size int) {
	src := rand.NewSource(1)
	s1 := randomSlice(size, src)
	s2 := randomSlice(size, src)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Add(s1, s2)
	}
}
func BenchmarkAddSmall(b *testing.B) { benchmarkAdd(b, Small) }
func BenchmarkAddMed(b *testing.B)   { benchmarkAdd(b, Medium) }
func BenchmarkAddLarge(b *testing.B) { benchmarkAdd(b, Large) }
func BenchmarkAddHuge(b *testing.B)  { benchmarkAdd(b, Huge) }

func benchmarkAddTo(b *testing.B, size int) {
	src := rand.NewSource(1)
	s1 := randomSlice(size, src)
	s2 := randomSlice(size, src)
	dst := randomSlice(size, src)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddTo(dst, s1, s2)
	}
}
func BenchmarkAddToSmall(b *testing.B) { benchmarkAddTo(b, Small) }
func BenchmarkAddToMed(b *testing.B)   { benchmarkAddTo(b, Medium) }
func BenchmarkAddToLarge(b *testing.B) { benchmarkAddTo(b, Large) }
func BenchmarkAddToHuge(b *testing.B)  { benchmarkAddTo(b, Huge) }

//func benchmarkDiv(b *testing.B, size int) {
//	src := rand.NewSource(1)
//	s := randomSlice(size, src)
//	dst := randomSlice(size, src)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		Div(dst, s)
//	}
//}
//func BenchmarkDivSmall(b *testing.B) { benchmarkDiv(b, Small) }
//func BenchmarkDivMed(b *testing.B)   { benchmarkDiv(b, Medium) }
//func BenchmarkDivLarge(b *testing.B) { benchmarkDiv(b, Large) }
//func BenchmarkDivHuge(b *testing.B)  { benchmarkDiv(b, Huge) }
//
//func benchmarkDivTo(b *testing.B, size int) {
//	src := rand.NewSource(1)
//	s1 := randomSlice(size, src)
//	s2 := randomSlice(size, src)
//	dst := randomSlice(size, src)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		DivTo(dst, s1, s2)
//	}
//}
//func BenchmarkDivToSmall(b *testing.B) { benchmarkDivTo(b, Small) }
//func BenchmarkDivToMed(b *testing.B)   { benchmarkDivTo(b, Medium) }
//func BenchmarkDivToLarge(b *testing.B) { benchmarkDivTo(b, Large) }
//func BenchmarkDivToHuge(b *testing.B)  { benchmarkDivTo(b, Huge) }

//func benchmarkSub(b *testing.B, size int) {
//	src := rand.NewSource(1)
//	s1 := randomSlice(size, src)
//	s2 := randomSlice(size, src)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		Sub(s1, s2)
//	}
//}
//func BenchmarkSubSmall(b *testing.B) { benchmarkSub(b, Small) }
//func BenchmarkSubMed(b *testing.B)   { benchmarkSub(b, Medium) }
//func BenchmarkSubLarge(b *testing.B) { benchmarkSub(b, Large) }
//func BenchmarkSubHuge(b *testing.B)  { benchmarkSub(b, Huge) }
//
//func benchmarkSubTo(b *testing.B, size int) {
//	src := rand.NewSource(1)
//	s1 := randomSlice(size, src)
//	s2 := randomSlice(size, src)
//	dst := randomSlice(size, src)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		SubTo(dst, s1, s2)
//	}
//}
//func BenchmarkSubToSmall(b *testing.B) { benchmarkSubTo(b, Small) }
//func BenchmarkSubToMed(b *testing.B)   { benchmarkSubTo(b, Medium) }
//func BenchmarkSubToLarge(b *testing.B) { benchmarkSubTo(b, Large) }
//func BenchmarkSubToHuge(b *testing.B)  { benchmarkSubTo(b, Huge) }

func benchmarkAddScaledTo(b *testing.B, size int) {
	src := rand.NewSource(1)
	dst := randomSlice(size, src)
	y := randomSlice(size, src)
	s := randomSlice(size, src)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddScaledTo(dst, y, 2.3, s)
	}
}
func BenchmarkAddScaledToSmall(b *testing.B)  { benchmarkAddScaledTo(b, Small) }
func BenchmarkAddScaledToMedium(b *testing.B) { benchmarkAddScaledTo(b, Medium) }
func BenchmarkAddScaledToLarge(b *testing.B)  { benchmarkAddScaledTo(b, Large) }
func BenchmarkAddScaledToHuge(b *testing.B)   { benchmarkAddScaledTo(b, Huge) }

func benchmarkScale(b *testing.B, size int) {
	src := rand.NewSource(1)
	dst := randomSlice(size, src)
	b.ResetTimer()
	for i := 0; i < b.N; i += 2 {
		Scale(2.0, dst)
		Scale(0.5, dst)
	}
}
func BenchmarkScaleSmall(b *testing.B)  { benchmarkScale(b, Small) }
func BenchmarkScaleMedium(b *testing.B) { benchmarkScale(b, Medium) }
func BenchmarkScaleLarge(b *testing.B)  { benchmarkScale(b, Large) }
func BenchmarkScaleHuge(b *testing.B)   { benchmarkScale(b, Huge) }
