package pkg

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	level := slog.LevelInfo

	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		level = slog.LevelDebug
	case "erroe":
		level = slog.LevelError
	case "warn":
		level = slog.LevelWarn
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}
