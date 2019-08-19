package calcs

import "log"

func Join(joiner Binary, a []float64, b []float64) []float64 {
	l := len(a)
	if l != len(b) {
		log.Panic("Mismatch in lengths")
	}

	var result = make([]float64, l)
	for i := 0; i < l; i++ {
		result[i] = joiner(a[i], b[i])
	}
	return result
}

func Accumulate(accumulator Binary, numbers []float64) float64 {
	length := len(numbers)
	if length == 0 {
		log.Panic("Empty slice")
	}

	var result = numbers[0]

	for i := 1; i < length; i++ {
		result = accumulator(result, numbers[i])
	}

	return result
}

func Split(splitter Splitter, arg []float64) ([]float64, []float64) {
	l := len(arg)
	var resultA = make([]float64, l)
	var resultB = make([]float64, l)
	for i := 0; i < l; i++ {
		a, b := splitter(arg[i])
		resultA[i] = a
		resultB[i] = b
	}
	return resultA, resultB
}

func Transform(unary Unary, numbers []float64) []float64 {
	length := len(numbers)
	var result = make([]float64, length)
	for i := 0; i < length; i++ {
		result[i] = unary(numbers[i])
	}
	return result
}
