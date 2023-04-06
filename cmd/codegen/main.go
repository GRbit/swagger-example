package main

import (
	"os"
	"time"

	"github.com/grbit/swagger-example/internal/codegen"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	configureLogging()

	if err := codegen.StartServer("localhost:8080", log.Logger); err != nil {
		panic(err)
	}
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
