package main

import (
	"calc/cmd/server"
	"calc/pb"
	"calc/pkg/logger"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":8080"

var log = logger.New()

func main() {
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Err(err).Msg("failed to create listener")
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterCalculatorServer(s, &server.CalcServer{})

	if err := s.Serve(listener); err != nil {
		log.Err(err).Msg("failed to start server")
	}

	log.Info().Msg("Server is listening on port 8000")
}
