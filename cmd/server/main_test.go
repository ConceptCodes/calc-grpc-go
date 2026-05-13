package server

import (
	"context"
	"testing"

	"calc/internal/calculator"
	"calc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCalcServerOperations(t *testing.T) {
	t.Parallel()

	svc := NewCalcServer(calculator.New())

	tests := []struct {
		name string
		call func() int64
		want int64
	}{
		{
			name: "add",
			call: func() int64 {
				resp, err := svc.Add(context.Background(), &pb.CalculationRequest{A: 3, B: 4})
				return mustValue(t, resp, err)
			},
			want: 7,
		},
		{
			name: "subtract",
			call: func() int64 {
				resp, err := svc.Subtract(context.Background(), &pb.CalculationRequest{A: 10, B: 2})
				return mustValue(t, resp, err)
			},
			want: 8,
		},
		{
			name: "multiply",
			call: func() int64 {
				resp, err := svc.Multiply(context.Background(), &pb.CalculationRequest{A: 6, B: 7})
				return mustValue(t, resp, err)
			},
			want: 42,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.call()
			if got != tt.want {
				t.Fatalf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCalcServerDivide(t *testing.T) {
	t.Parallel()

	svc := NewCalcServer(calculator.New())

	got, err := svc.Divide(context.Background(), &pb.CalculationRequest{A: 12, B: 3})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Result != 4 {
		t.Fatalf("got %d, want %d", got.Result, 4)
	}

	_, err = svc.Divide(context.Background(), &pb.CalculationRequest{A: 12, B: 0})
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument, got %v", err)
	}
}

func mustValue(t *testing.T, resp *pb.CalculationResponse, err error) int64 {
	t.Helper()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	return resp.Result
}
