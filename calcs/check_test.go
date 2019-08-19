package calcs

import "testing"

func Check(t *testing.T, actual float64, expected float64) {
	if expected != actual {
		t.Errorf("Result was incorrect, got: %g, wanted: %g", actual, expected)
	}
}

func CheckWithPrecision(t *testing.T, actual float64, expected float64, precision int) {
	if expected != Round(actual, precision) {
		t.Errorf("Result was incorrect, got: %g, wanted: %g, precision: %d", actual, expected, precision)
	}
}

func CheckSlice(t *testing.T, actual []float64, expected []float64) {
	Check(t, float64(len(expected)), float64(len(actual)))
	for i := 0; i < len(actual); i++ {
		Check(t, actual[i], expected[i])
	}
}

func CheckSliceWithPrecision(t *testing.T, actual []float64, expected []float64, precision int) {
	Check(t, float64(len(expected)), float64(len(actual)))
	for i := 0; i < len(actual); i++ {
		CheckWithPrecision(t, actual[i], expected[i], precision)
	}
}
