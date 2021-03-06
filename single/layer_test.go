package single

import (
	"github.com/gonum/matrix/mat64"
	"testing"
)

func TestAffine(t *testing.T) {
	for _, c := range []struct {
		title string
		w     *mat64.Dense
		b     *mat64.Vector
		x     *mat64.Vector
		y     *mat64.Vector
		dout  *mat64.Vector
		dx    *mat64.Vector
	}{
		{
			title: "TestAffine",
			w: mat64.NewDense(2, 3, []float64{
				1, 2, 3,
				4, 5, 6,
			}),
			b:    mat64.NewVector(3, []float64{7, 8, 9}),
			x:    mat64.NewVector(2, []float64{1, 2}),
			y:    mat64.NewVector(3, []float64{16, 20, 24}),
			dout: mat64.NewVector(3, []float64{1, 2, 3}),
			dx:   mat64.NewVector(2, []float64{14, 32}),
		},
	} {
		layer := NewAffineLayer(c.w, c.b, nil)
		y := layer.Forward(c.x)
		if !mat64.Equal(y, c.y) {
			t.Fatalf("%s expect %v but got %v", c.title, c.y, y)
		}
		dx := layer.Backward(c.dout)
		if !mat64.Equal(dx, c.dx) {
			t.Fatalf("%s expect %v but got %v", c.title, c.dx, dx)
		}
	}
}

func TestReLULayer(t *testing.T) {
	for _, c := range []struct {
		title  string
		x      *mat64.Vector
		y      *mat64.Vector
		dout   *mat64.Vector
		expect *mat64.Vector
	}{
		{
			title:  "TestReLU",
			x:      mat64.NewVector(3, []float64{2, -1, -2}),
			y:      mat64.NewVector(3, []float64{2, 0, 0}),
			dout:   mat64.NewVector(3, []float64{-1, 1, -1}),
			expect: mat64.NewVector(3, []float64{-1, 0, 0}),
		},
	} {
		l := &ReLULayer{}
		y := l.Forward(c.x)
		if !mat64.Equal(y, c.y) {
			t.Fatalf("%s expect %v but got %v", c.title, c.y, y)
		}
		expect := l.Backward(c.dout)
		if !mat64.Equal(expect, c.expect) {
			t.Fatalf("%s expect %v but got %v", c.title, c.expect, expect)
		}
	}
}
