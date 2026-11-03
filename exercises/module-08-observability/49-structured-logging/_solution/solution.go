package logging

import (
	"context"
	"io"
	"log/slog"
)

func NewJSONLogger(w io.Writer) *slog.Logger {
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo}))
}

func LogInfo(logger *slog.Logger, msg string, attrs ...slog.Attr) {
	logger.LogAttrs(context.Background(), slog.LevelInfo, msg, attrs...)
}

func LogError(logger *slog.Logger, msg string, err error, attrs ...slog.Attr) {
	all := append([]slog.Attr{slog.String("error", err.Error())}, attrs...)
	logger.LogAttrs(context.Background(), slog.LevelError, msg, all...)
}
