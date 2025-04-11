package slogctx

import (
	"io"
	"log/slog"
	"os"
)

var defaultLogger *slog.Logger

func getDefultLogger() *slog.Logger {
	if defaultLogger != nil {
		return defaultLogger
	}
	return slog.Default()
}

// SetDefaultLogger sets the default logger
// for this package (slogctx)
func SetDefaultLogger(logger *slog.Logger) {
	defaultLogger = logger
}

// Calls slog.SetDefault with a JSONHandler
// that logs to Stdout wrapped in a ContextHandler
func SetSlogPackageDefault() {
	SetSlogPackageWithOptions(os.Stdout, nil)
}

// Calls slog.SetDefault with options passed to a JSONHandler
// that logs to Stdout wrapped in a ContextHandler
func SetSlogPackageWithOptions(writer io.Writer, opts *slog.HandlerOptions) {
	logger := slog.New(NewContextHandler(slog.NewJSONHandler(writer, opts)))
	slog.SetDefault(logger)
	SetDefaultLogger(logger)
}
