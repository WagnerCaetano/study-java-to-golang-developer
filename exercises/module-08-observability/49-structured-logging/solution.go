package logging

import (
	"io"
	"log/slog"
)

// NewJSONLogger returns a logger that writes JSON lines to w.
func NewJSONLogger(w io.Writer) *slog.Logger {
	// FIX ME: use slog.NewJSONHandler, not TextHandler
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo}))
}

// LogInfo writes an Info log with structured attributes.
func LogInfo(logger *slog.Logger, msg string, attrs ...slog.Attr) {
	// FIX ME: include attrs in the log record
	logger.Info(msg)
}

// LogError writes an Error log with err and structured attributes.
func LogError(logger *slog.Logger, msg string, err error, attrs ...slog.Attr) {
	// FIX ME: include err (field "error") and attrs
	logger.Error(msg)
}
