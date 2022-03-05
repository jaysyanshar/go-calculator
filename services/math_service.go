package services

type MathService interface {
	Add(a, b float64) float64
	Sub(a, b float64) float64
	Mul(a, b float64) float64
	Div(a, b float64) float64
}

type MathServiceImpl struct{}

func (s *MathServiceImpl) Add(a, b float64) float64 {
	return a + b
}

func (s *MathServiceImpl) Sub(a, b float64) float64 {
	return a - b
}

func (s *MathServiceImpl) Mul(a, b float64) float64 {
	return a * b
}

func (s *MathServiceImpl) Div(a, b float64) float64 {
	return a / b
}
