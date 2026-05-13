package bootstrap

import (
	"context"
	"fmt"
	"net"

	"github.com/rs/zerolog"
)

type Server interface {
	Serve(net.Listener) error
	GracefulStop()
}

type Config struct {
	Address   string
	Logger    zerolog.Logger
	Listen    func(network, address string) (net.Listener, error)
	NewServer func(logger zerolog.Logger) Server
}

func Run(ctx context.Context, cfg Config) error {
	if ctx == nil {
		ctx = context.Background()
	}

	listen := cfg.Listen
	if listen == nil {
		listen = net.Listen
	}

	listener, err := listen("tcp", cfg.Address)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", cfg.Address, err)
	}

	if cfg.NewServer == nil {
		return fmt.Errorf("new server factory is required")
	}

	grpcServer := cfg.NewServer(cfg.Logger)
	cfg.Logger.Info().Msgf("Server is listening on port %s", cfg.Address)

	serveErr := make(chan error, 1)
	go func() {
		serveErr <- grpcServer.Serve(listener)
	}()

	select {
	case err := <-serveErr:
		if err != nil {
			return fmt.Errorf("serve gRPC on %s: %w", cfg.Address, err)
		}
		return nil
	case <-ctx.Done():
		cfg.Logger.Info().Msg("Shutting down server")
		grpcServer.GracefulStop()

		if err := <-serveErr; err != nil {
			return fmt.Errorf("serve gRPC on %s: %w", cfg.Address, err)
		}
		return nil
	}
}
