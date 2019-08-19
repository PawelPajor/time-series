package calcs

func Exercise(m float64, c float64, channel []float64) float64 {
	y := Transform(Y{M: m, C: c}.Calculate, channel)
	a := Transform(func(x float64) float64 { return 1 / x }, channel)
	b := Join(Add, y, a)
	result := Avg(b)
	return result
}
