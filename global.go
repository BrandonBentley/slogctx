package slogctx

import (
	"context"
	"log/slog"
)

var rootLoggerFactory = NewFactory(nil)

func SetRootFactory(factory *Factory) {
	rootLoggerFactory = factory
}

func GetLogger(ctx context.Context) *slog.Logger {
	return rootLoggerFactory.GetLogger(ctx)
}

func AddAttributesToLogger(ctx context.Context, fields ...any) context.Context {
	return rootLoggerFactory.AddAttributesToLogger(ctx, fields...)
}
