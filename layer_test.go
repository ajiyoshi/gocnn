package gocnn

import (
	"github.com/gonum/matrix/mat64"
	"testing"
)

func zeros(n int) *mat64.Vector {
	return mat64.NewVector(n, make([]float64, n))
}

func TestConvLayer(t *testing.T) {

	cases := []struct {
		msg      string
		generate func() (*Convolution, ImageStrage, ImageStrage)
	}{
		{
			msg: "実際に実行してみた結果",
			generate: func() (*Convolution, ImageStrage, ImageStrage) {
				filterNum := 1
				filterSize := 3
				dataNum, chNum, xCol, xRow := 1, 1, 2, 2
				x := NewImages(ImageShape{n: dataNum, ch: chNum, col: xCol, row: xRow}, []float64{
					1, 2,
					3, 4,
				})

				w := NewImages(ImageShape{n: filterNum, ch: chNum, col: filterSize, row: filterSize}, []float64{
					0.01061144, 0.00930966, 0.00157138,
					0.01366734, 0.00596517, -0.00052856,
					-0.0022351, -0.00402149, -0.00019544,
				})
				bias := zeros(filterNum)
				conv := &Convolution{
					Weight: w,
					Bias:   bias,
					stride: 1,
					pad:    1,
				}
				expect := NewImages(ImageShape{dataNum, filterNum, 2, 2}, []float64{
					-0.00793818, 0.00280642,
					0.02823369, 0.09409346,
				})
				return conv, x, expect
			},
		},
		{
			msg: "実際に実行してみた結果",
			generate: func() (*Convolution, ImageStrage, ImageStrage) {
				filterNum := 2
				filterSize := 3
				dataNum, chNum, xCol, xRow := 1, 1, 2, 2
				x := NewImages(ImageShape{n: dataNum, ch: chNum, col: xCol, row: xRow}, []float64{
					1, 2,
					3, 4,
				})

				w := NewImages(ImageShape{n: filterNum, ch: chNum, col: filterSize, row: filterSize}, []float64{
					-0.00205192, -0.01427015, 0.01118195,
					-0.00115402, 0.00920227, -0.00072591,
					0.00013398, 0.00050144, -0.01106872,

					-0.00434496, -0.00634229, 0.01542008,
					-0.01091955, 0.00087254, 0.00059587,
					-0.00545791, 0.00746039, 0.00338262,
				})
				bias := zeros(filterNum)
				conv := &Convolution{
					Weight: w,
					Bias:   bias,
					stride: 1,
					pad:    1,
				}
				expect := NewImages(ImageShape{dataNum, filterNum, 2, 2}, []float64{
					-0.03502013, 0.01965824,
					0.0327969, 0.00275479,

					0.03797594, 0.00429336,
					0.02949898, -0.04629804,
				})
				return conv, x, expect
			},
		},
	}

	for _, c := range cases {
		layer, x, expect := c.generate()
		y := layer.Forward(x)
		if !y.Equal(expect) {
			t.Fatalf("expect \n%v got \n%v", expect, y)
		}
	}
}

func TestHoge(t *testing.T) {

	cases := []struct {
		msg      string
		generate func() (*Convolution, ImageStrage, ImageStrage)
	}{
		{
			msg: "実際に実行してみた結果",
			generate: func() (*Convolution, ImageStrage, ImageStrage) {
				filterNum := 1
				filterSize := 3
				dataNum, chNum, xCol, xRow := 1, 1, 2, 2
				x := NewImages(ImageShape{n: dataNum, ch: chNum, col: xCol, row: xRow}, []float64{
					1, 2,
					3, 4,
				})

				w := NewImages(ImageShape{n: filterNum, ch: chNum, col: filterSize, row: filterSize}, []float64{
					0.01061144, 0.00930966, 0.00157138,
					0.01366734, 0.00596517, -0.00052856,
					-0.0022351, -0.00402149, -0.00019544,
				})
				bias := zeros(filterNum)
				conv := &Convolution{
					Weight: w,
					Bias:   bias,
					stride: 1,
					pad:    1,
				}
				expect := NewImages(ImageShape{dataNum, filterNum, 2, 2}, []float64{
					-0.00793818, 0.00280642,
					0.02823369, 0.09409346,
				})
				return conv, x, expect
			},
		},
		{
			msg: "実際に実行してみた結果",
			generate: func() (*Convolution, ImageStrage, ImageStrage) {
				filterNum := 2
				filterSize := 3
				dataNum, chNum, xCol, xRow := 1, 1, 2, 2
				x := NewImages(ImageShape{n: dataNum, ch: chNum, col: xCol, row: xRow}, []float64{
					1, 2,
					3, 4,
				})

				w := NewImages(ImageShape{n: filterNum, ch: chNum, col: filterSize, row: filterSize}, []float64{
					-0.00205192, -0.01427015, 0.01118195,
					-0.00115402, 0.00920227, -0.00072591,
					0.00013398, 0.00050144, -0.01106872,

					-0.00434496, -0.00634229, 0.01542008,
					-0.01091955, 0.00087254, 0.00059587,
					-0.00545791, 0.00746039, 0.00338262,
				})
				bias := zeros(filterNum)
				conv := &Convolution{
					Weight: w,
					Bias:   bias,
					stride: 1,
					pad:    1,
				}
				expect := NewImages(ImageShape{dataNum, filterNum, 2, 2}, []float64{
					-0.03502013, 0.01965824,
					0.0327969, 0.00275479,

					0.03797594, 0.00429336,
					0.02949898, -0.04629804,
				})
				return conv, x, expect
			},
		},
	}

	for _, c := range cases {
		layer, x, expect := c.generate()
		y := layer.Forward(x)
		if !y.Equal(expect) {
			t.Fatalf("expect \n%v got \n%v", expect, y)
		}
	}
}

func TestRun(t *testing.T) {
	filterNum := 2
	filterSize := 3
	dataNum, chNum, xCol, xRow := 1, 1, 2, 2
	x := NewImages(ImageShape{n: dataNum, ch: chNum, col: xCol, row: xRow}, []float64{
		1, 2,
		3, 4,
	})

	w := NewImages(ImageShape{n: filterNum, ch: chNum, col: filterSize, row: filterSize}, []float64{
		-0.00431678, 0.00845254, 0.00188367,
		0.0036537, -0.00328172, 0.0070477,
		0.0162438, -0.00225665, -0.00838634,

		0.01078964, 0.01474864, 0.00210126,
		-0.01044749, -0.00207768, -0.00652951,
		-0.00074343, -0.00369562, -0.00204102,
	})
	bias := zeros(filterNum)
	conv := &Convolution{
		Weight: w,
		Bias:   bias,
		stride: 1,
		pad:    1,
	}
	y := conv.Forward(x)
	z := conv.Backword(y)
	t.Logf("%s", z)
}