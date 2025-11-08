package logger

import (
	"log/slog"
	"os"
)

func NewLoggerText() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}
