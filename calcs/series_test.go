package calcs

import "testing"

func TestJoin(t *testing.T) {
	expected := []float64{0.0, 1.0, 1.0, 3.0}
	actual := Join(Add, []float64{0.0, 1.0, 0.0, 1.0}, []float64{0.0, 0.0, 1.0, 2.0})
	CheckSlice(t, actual, expected)
}

func TestAccumulate(t *testing.T) {
	Check(t, Accumulate(Add, []float64{0.0, 1.0, 0.0, 1.0}), 2.0)
}

func TestSplit(t *testing.T) {
	// Arrange
	a := []float64{0.0, 1.0, 2.0, 4.0}
	expectedA := []float64{0.0, 0.6, 1.2, 2.4}
	expectedB := []float64{0.0, 0.4, 0.8, 1.6}
	splitter := func(x float64) (float64, float64) { return 0.6 * x, 0.4 * x }

	// Act
	actualA, actualB := Split(splitter, a)

	// Assert
	CheckSlice(t, actualA, expectedA)
	CheckSlice(t, actualB, expectedB)
}

func TestTransform(t *testing.T) {
	data := []float64{0.0, 1.0, 2.0, 4.0}
	actual := Transform(func(x float64) float64 { return 0.5 * x }, data)
	expected := []float64{0.0, 0.5, 1.0, 2.0}
	CheckSliceWithPrecision(t, actual, expected, 5)
}
