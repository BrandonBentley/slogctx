package slogctx

import (
	"context"
	"log/slog"
)

var rootLoggerFactory = NewFactory(nil)

func SetRootLoggerFactory(factory *Factory) {
	rootLoggerFactory = factory
}

func GetContextLogger(ctx context.Context) *slog.Logger {
	return rootLoggerFactory.GetContextLogger(ctx)
}

func AddAttributesToContextLogger(ctx context.Context, fields ...any) context.Context {
	return rootLoggerFactory.AddAttributesToContextLogger(ctx, fields...)
}
