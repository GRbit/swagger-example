package codegen

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/xerrors"
)

func StartServer(addr string, lg zerolog.Logger, handler http.Handler) error {
	srv := &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		ReadTimeout:  5 * time.Second,
		WriteTimeout: time.Minute,
		IdleTimeout:  5 * time.Minute,
		Handler:      handler,
	}

	// error channel to catch errors from the server
	errCh := make(chan error, 1)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		lg.Info().Msg("Starting API server on " + addr)

		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			lg.Info().Msg("Server closed with SIGTERM or SIGINT")
		} else if err != nil {
			errCh <- xerrors.Errorf("server stopped: %w", err)
		}
	}()

	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errCh:
		return err
	case <-termCh:
		lg.Info().Msg("kill signal received. Waiting for graceful shutdown ...")
	}

	// Create a deadline
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		lg.Err(err).Msg("Server shutdown error")
		return err
	}

	lg.Info().Msg("Server gracefully stopped")

	return nil
}
