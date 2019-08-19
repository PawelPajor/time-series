package calcs

import "math"

func Round(x float64, decimalPoint int) float64 {
	unit := 5.0 / math.Pow10(decimalPoint)
	return math.Round(x/unit) * unit
}
