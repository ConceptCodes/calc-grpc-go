package calculator

import (
	"errors"
	"testing"
)

func TestCalculatorOperations(t *testing.T) {
	t.Parallel()

	calc := New()

	tests := []struct {
		name string
		got  int64
		want int64
	}{
		{name: "add", got: calc.Add(3, 4), want: 7},
		{name: "subtract", got: calc.Subtract(10, 2), want: 8},
		{name: "multiply", got: calc.Multiply(6, 7), want: 42},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.got != tt.want {
				t.Fatalf("got %d, want %d", tt.got, tt.want)
			}
		})
	}
}

func TestCalculatorDivide(t *testing.T) {
	t.Parallel()

	calc := New()

	got, err := calc.Divide(12, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 4 {
		t.Fatalf("got %d, want %d", got, 4)
	}

	_, err = calc.Divide(12, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Fatalf("expected ErrDivideByZero, got %v", err)
	}
}
