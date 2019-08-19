package calcs

import (
	"math"
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		param    float64 // input
		expected float64 // expected
	}{
		{0.1, 10.0},
		{1.0, 1.0},
		{10.0, 0.1},
		{0.0, math.Inf(1)},
		{math.Inf(1), 0.0},
	}
	for _, test := range tests {
		Check(t, A(test.param), test.expected)
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a        float64 // input
		b        float64 // input
		expected float64 // expected
	}{
		{0, 0, 0},
		{-1, 1, 0},
		{1, 2, 3},
	}

	for _, test := range tests {
		Check(t, Add(test.a, test.b), test.expected)
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		a        float64 // input
		b        float64 // input
		expected float64 // expected
	}{
		{0, 0, 0},
		{-1, 1, -2},
		{1, -2, 3},
	}

	for _, test := range tests {
		Check(t, Sub(test.a, test.b), test.expected)
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		a        float64 // input
		b        float64 // input
		expected float64 // expected
	}{
		{0, 0, 0},
		{-1, 1, -1},
		{1, -2, -2},
	}

	for _, test := range tests {
		Check(t, Mul(test.a, test.b), test.expected)
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		a        float64 // input
		b        float64 // input
		expected float64 // expected
	}{
		{1, 0, math.Inf(1)},
		{-1, 1, -1},
		{1, -2, -0.5},
	}

	for _, test := range tests {
		Check(t, Div(test.a, test.b), test.expected)
	}
}

func TestSum(t *testing.T) {
	Check(t, Sum([]float64{1.0, 2.0, 3.0}), 6.0)
}

func TestAvg(t *testing.T) {
	Check(t, Avg([]float64{1.0, 2.0, 3.0}), 2.0)
}

func TestY(t *testing.T) {
	tests := []struct {
		arg      float64
		expected float64
	}{
		{0.0, 1.5},
		{10.0, 6.5},
		{20.0, 11.5},
	}
	tested := Y{M: 0.5, C: 1.5}

	for _, test := range tests {
		Check(t, tested.Calculate(test.arg), test.expected)
	}
}
