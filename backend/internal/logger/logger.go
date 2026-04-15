package logger

import (
	"log/slog"
	"os"
)

func NewLogger() {
	logger := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	})
	slog.SetDefault(slog.New(logger))
}
