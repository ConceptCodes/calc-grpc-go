package bootstrap

import (
	"context"
	"errors"
	"io"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func TestRunReturnsListenError(t *testing.T) {
	t.Parallel()

	want := errors.New("boom")
	err := Run(context.Background(), Config{
		Address: ":0",
		Logger:  zerolog.Nop(),
		Listen: func(network, address string) (net.Listener, error) {
			return nil, want
		},
		NewServer: func(logger zerolog.Logger) Server {
			t.Fatal("new server should not be called when listen fails")
			return nil
		},
	})

	if !errors.Is(err, want) {
		t.Fatalf("expected wrapped listen error, got %v", err)
	}
}

func TestRunStopsOnContextCancel(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error, 1)
	srv := newTestServer()

	go func() {
		done <- Run(ctx, Config{
			Address: ":0",
			Logger:  zerolog.Nop(),
			Listen: func(network, address string) (net.Listener, error) {
				return stubListener{}, nil
			},
			NewServer: func(logger zerolog.Logger) Server {
				return srv
			},
		})
	}()

	srv.waitForServe(t)
	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("expected graceful shutdown, got %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("run did not exit after context cancellation")
	}
}

type stubListener struct{}

func (stubListener) Accept() (net.Conn, error) { return nil, io.EOF }
func (stubListener) Close() error              { return nil }
func (stubListener) Addr() net.Addr            { return stubAddr("stub") }

type stubAddr string

func (a stubAddr) Network() string { return string(a) }
func (a stubAddr) String() string  { return string(a) }

type testServer struct {
	started chan struct{}
	stop    chan struct{}
	once    sync.Once
}

func newTestServer() *testServer {
	return &testServer{
		started: make(chan struct{}),
		stop:    make(chan struct{}),
	}
}

func (s *testServer) Serve(net.Listener) error {
	close(s.started)
	<-s.stop
	return nil
}

func (s *testServer) GracefulStop() {
	s.once.Do(func() {
		close(s.stop)
	})
}

func (s *testServer) waitForServe(t *testing.T) {
	t.Helper()

	select {
	case <-s.started:
	case <-time.After(2 * time.Second):
		t.Fatal("serve did not start")
	}
}
