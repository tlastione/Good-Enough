package internal

import (
	"log/slog"
	"os"
)

func Greet() string {
	stdoutHandler := slog.NewJSONHandler(os.Stdout, nil)
	log := slog.New(stdoutHandler)

	log.Info("ehhe")
	return "hehe"
}
