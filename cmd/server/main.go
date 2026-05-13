package server

import (
	"context"
	"errors"

	"calc/internal/calculator"
	"calc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CalcServer struct {
	pb.UnimplementedCalculatorServer
	calculator *calculator.Calculator
}

func NewCalcServer(calc *calculator.Calculator) *CalcServer {
	if calc == nil {
		calc = calculator.New()
	}

	return &CalcServer{calculator: calc}
}

func (s *CalcServer) engine() *calculator.Calculator {
	if s.calculator == nil {
		s.calculator = calculator.New()
	}

	return s.calculator
}

func (s *CalcServer) Add(
	ctx context.Context, req *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	res := &pb.CalculationResponse{
		Result: s.engine().Add(req.A, req.B),
	}

	return res, nil
}

func (s *CalcServer) Subtract(
	ctx context.Context, req *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	res := &pb.CalculationResponse{
		Result: s.engine().Subtract(req.A, req.B),
	}

	return res, nil
}

func (s *CalcServer) Multiply(
	ctx context.Context, req *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	res := &pb.CalculationResponse{
		Result: s.engine().Multiply(req.A, req.B),
	}

	return res, nil
}

func (s *CalcServer) Divide(
	ctx context.Context, req *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	result, err := s.engine().Divide(req.A, req.B)
	if err != nil {
		if errors.Is(err, calculator.ErrDivideByZero) {
			return nil, status.Errorf(
				codes.InvalidArgument, err.Error(),
			)
		}
		return nil, err
	}

	res := &pb.CalculationResponse{
		Result: result,
	}

	return res, nil
}
