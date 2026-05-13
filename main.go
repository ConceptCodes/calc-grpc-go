package main

import (
	"calc/cmd/server"
	"calc/internal/bootstrap"
	"calc/internal/calculator"
	"calc/pkg/logger"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
)

const port = ":8080"

func main() {
	log := logger.New()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := bootstrap.Run(ctx, bootstrap.Config{
		Address: port,
		Logger:  log,
		NewServer: func(logger zerolog.Logger) bootstrap.Server {
			return server.NewGRPCServer(logger, calculator.New())
		},
	}); err != nil {
		log.Err(err).Msg("server stopped")
		os.Exit(1)
	}
}
