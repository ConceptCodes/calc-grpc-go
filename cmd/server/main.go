package server

import (
	"context"

	"calc/internal/util"
	"calc/pb"
	"calc/pkg/logger"
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
	util.LogRequest("ADD", req, log)
	res := &pb.CalculationResponse{
		Result: req.A + req.B,
	}
	util.LogResponse("ADD", res, log)
	return res, nil
}
