package logger

import (
	"log/slog"
	"os"
)

// New creates and returns a new JSON logger instance
func New() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	return slog.New(handler)
}
