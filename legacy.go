package slogctx

import (
	"context"
	"log/slog"
)

// GetLogger is used to get the context logger
// It is recommended to use slog context functions instead
// ex slog.InfoContext(ctx, msg, ...args)
func GetLogger(ctx context.Context) *slog.Logger {
	return slog.New(getDefultLogger().Handler().WithAttrs(getAttrSetFromContext(ctx).attrSlice))
}

// AddAttributesToLogger is now an alias for With
func AddAttributesToLogger(ctx context.Context, args ...any) context.Context {
	return With(ctx, args...)
}
