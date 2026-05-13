package calculator

import "errors"

var ErrDivideByZero = errors.New("cannot divide by zero")

type Calculator struct{}

func New() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Add(a, b int64) int64 {
	return a + b
}

func (c *Calculator) Subtract(a, b int64) int64 {
	return a - b
}

func (c *Calculator) Multiply(a, b int64) int64 {
	return a * b
}

func (c *Calculator) Divide(a, b int64) (int64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	return a / b, nil
}
