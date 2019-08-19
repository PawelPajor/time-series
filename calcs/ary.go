package calcs

type Unary func(float64) float64
type Binary func(float64, float64) float64
type Splitter func(float64) (float64, float64)

func Add(a float64, b float64) float64 { return a + b }
func Sub(a float64, b float64) float64 { return a - b }
func Mul(a float64, b float64) float64 { return a * b }
func Div(a float64, b float64) float64 { return a / b }

type Y struct {
	M float64
	C float64
}

func (y Y) Calculate(x float64) float64 {
	return y.M*x + y.C
}

func A(x float64) float64 {
	return 1.0 / x
}

func Sum(numbers []float64) float64 {
	return Accumulate(Add, numbers)
}

func Avg(numbers []float64) float64 {
	return Sum(numbers) / float64(len(numbers))
}
