package slogctx

import (
	"context"
	"log/slog"
)

// WithAttrs is used to add slog.Attrs to the context for logging purposes
// these slog.Attrs will appear in logs of any child contexts
// where slogctx.GetLogger(ctx) is called or any slog context functions
// for example slog.InfoContext(ctx, msg, ...args)
func WithAttrs(ctx context.Context, attrs ...slog.Attr) context.Context {
	as := getAttrSetFromContext(ctx)
	return context.WithValue(ctx, attrSetContextKey, as.with(attrs))
}

// With is used to add values to the context for logging purposes
// these args will appear in logs of any child contexts
// where slogctx.GetLogger(ctx) is called or any slog context functions
// for example slog.InfoContext(ctx, msg, ...args)
func With(ctx context.Context, args ...any) context.Context {
	return WithAttrs(ctx, argsToAttrSlice(args)...)
}
