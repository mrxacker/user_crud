package app

import (
	"context"
	"time"

	"github.com/mrxacker/user_service/configs"
	"github.com/mrxacker/user_service/internal/server"
)

const (
	ReadTimeout      = 5 * time.Second
	ErrorChannelSize = 2
)

func Run(ctx context.Context) error {

	config, err := configs.Load()
	if err != nil {
		return err
	}

	err = initServer(ctx, config)
	if err != nil {
		return err
	}

	return nil
}

func initServer(ctx context.Context, config *configs.Config) error {

	srv, err := setupServer(*config)
	if err != nil {
		return err
	}

	err = runServer(ctx, srv)
	if err != nil {
		return err
	}

	return nil
}

func setupServer(config configs.Config) (*server.Server, error) {
	srv, err := server.NewServer(&config)
	if err != nil {
		return nil, err
	}

	return srv, nil

}

func runServer(ctx context.Context, s *server.Server) error {
	errCh := make(chan error, ErrorChannelSize)

	go func() {
		if err := s.Run(); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
		defer cancel()

		err := s.Shutdown(shutdownCtx)
		if err != nil {
			return err
		}

	case err := <-errCh:
		return err
	}

	return nil
}
