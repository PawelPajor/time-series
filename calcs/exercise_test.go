package calcs

import "testing"

func TestExercise(t *testing.T) {
	CheckWithPrecision(t, Exercise(2.0, 0.5, []float64{1.0, 2.0, 3.0}), 5.11, 3)
}
