package slogctx

import (
	"context"
	"fmt"
	"io"
	"log/slog"
)

type uniqueLoggerContextKey string

const loggerContextKey = uniqueLoggerContextKey("context-logger")

type Factory struct {
	defaultLogger *slog.Logger
}

func NewFactory(logger *slog.Logger) *Factory {
	return &Factory{
		defaultLogger: logger,
	}
}

func NewFactoryNoOp() *Factory {
	return &Factory{
		defaultLogger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

func (f *Factory) GetContextLogger(ctx context.Context) *slog.Logger {
	return f.getLoggerFromContext(ctx)
}

func (f *Factory) AddAttributesToContextLogger(ctx context.Context, attrs ...any) context.Context {
	return context.WithValue(ctx, loggerContextKey, f.getLoggerFromContext(ctx).With(attrs...))
}

func (f *Factory) getLoggerFromContext(ctx context.Context) *slog.Logger {
	val := ctx.Value(loggerContextKey)

	logger := f.getDefaultLogger()

	if val != nil {
		if contextLogger, ok := val.(*slog.Logger); ok {
			logger = contextLogger
		} else {
			logger.Error(
				"Value at logger context key is not a *slog.Logger. Defaulting to standard logger",
				"type", fmt.Sprintf("%T", val),
			)
		}
	}

	return logger
}

func (f *Factory) getDefaultLogger() *slog.Logger {
	if f.defaultLogger != nil {
		return f.defaultLogger
	}
	return slog.Default()
}
