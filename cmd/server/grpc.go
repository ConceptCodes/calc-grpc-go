package server

import (
	"context"
	"path"
	"strings"

	"calc/internal/calculator"
	"calc/pb"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(log zerolog.Logger, calc *calculator.Calculator) *grpc.Server {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(unaryLoggingInterceptor(log)))
	reflection.Register(grpcServer)
	pb.RegisterCalculatorServer(grpcServer, NewCalcServer(calc))

	return grpcServer
}

func unaryLoggingInterceptor(log zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		operation := operationName(info.FullMethod)
		logRequest(operation, req, log)

		resp, err := handler(ctx, req)
		if err != nil {
			logError(operation, err, log)
			return resp, err
		}

		logResponse(operation, resp, log)
		return resp, nil
	}
}

func operationName(fullMethod string) string {
	name := path.Base(fullMethod)
	if strings.TrimSpace(name) == "" || name == "." {
		return fullMethod
	}

	return name
}

func logRequest(target string, req interface{}, log zerolog.Logger) {
	r, ok := req.(*pb.CalculationRequest)
	if !ok {
		return
	}

	log.Debug().
		Int64("A", r.A).
		Int64("B", r.B).
		Msgf("Incoming %s request", target)
}

func logResponse(target string, resp interface{}, log zerolog.Logger) {
	r, ok := resp.(*pb.CalculationResponse)
	if !ok {
		return
	}

	log.Debug().
		Int64("result", r.Result).
		Msgf("Result for %s is %d", target, r.Result)
}

func logError(target string, err error, log zerolog.Logger) {
	log.Err(err).Msgf("Error w/ %s", target)
}
