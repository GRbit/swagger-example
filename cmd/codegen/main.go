package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grbit/swagger-example/internal/codegen"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/xerrors"
)

func main() {
	configureLogging()

	addr := "localhost:8080"

	srv := &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		ReadTimeout:  5 * time.Second,
		WriteTimeout: time.Minute,
		IdleTimeout:  5 * time.Minute,
		Handler:      codegen.NewRouter(log.Logger),
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Logger.Info().Msg("Starting API server on " + addr)

		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Logger.Info().Msg("Server closed with SIGTERM or SIGINT")
		} else if err != nil {
			panic(xerrors.Errorf("server stopped: %w", err))
		}
	}()

	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGINT, syscall.SIGTERM)
	<-termCh
	log.Logger.Info().Msg("kill signal received. Waiting for graceful shutdown ...")

	// Create a deadline
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Err(err).Msg("Server shutdown error")
	}

	log.Logger.Info().Msg("Server gracefully stopped")
}

func configureLogging() {
	log.Logger = log.With().
		Caller().
		Timestamp().
		Logger().Output(
		zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampMicro},
	)

	zerolog.TimeFieldFormat = time.StampMicro
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}
