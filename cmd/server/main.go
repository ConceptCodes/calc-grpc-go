package server

import (
	"context"

	"calc/internal/util"
	"calc/pb"
	"calc/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CalcServer struct {
	pb.UnimplementedCalculatorServer
}

func NewCalcServer() *CalcServer {
	return &CalcServer{}
}

var log = logger.New()

func (s *CalcServer) Add(
	ctx context.Context, req *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	util.LogRequest("Add", req, log)
	
	res := &pb.CalculationResponse{
		Result: req.A + req.B,
	}

	util.LogResponse("Add", res, log)
	return res, nil
}

func (s *CalcServer) Subtract(
	ctx context.Context, req *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	util.LogRequest("Subtract", req, log)

	res := &pb.CalculationResponse{
		Result: req.A - req.B,
	}

	util.LogResponse("Subtract", res, log)
	return res, nil
}

func (s *CalcServer) Multiply(
	ctx context.Context, req *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	util.LogRequest("Multiply", req, log)

	res := &pb.CalculationResponse{
		Result: req.A * req.B,
	}

	util.LogResponse("Multiply", res, log)
	return res, nil
}

func (s *CalcServer) Divide(
	ctx context.Context, req *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	util.LogRequest("Divide", req, log)

	if req.B == 0 {
		err := status.Errorf(
			codes.InvalidArgument, "cannot divide by zero",
		)
		util.LogError("Divide", err, log)
		return nil, err
	}

	res := &pb.CalculationResponse{
		Result: req.A / req.B,
	}

	util.LogResponse("Divide", res, log)
	return res, nil
}
