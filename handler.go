package slogctx

import (
	"context"
	"io"
	"log/slog"
)

// ContextHandler enables this package to handle passing
// values from context to log messages when using
// the slog context functions
// ex: slog.InfoContext(ctx, msg, ...args)
type ContextHandler struct {
	handler slog.Handler
}

// NewJSONContextHandler creates a new ContextHandler
// that wraps a slog.JSONHandler
func NewJSONContextHandler(w io.Writer, opts *slog.HandlerOptions) *ContextHandler {
	return NewContextHandler(slog.NewJSONHandler(w, opts))
}

// NewContextHandler creates a new ContextHandler
// that wraps a slog.Handler passed in
func NewContextHandler(handler slog.Handler) *ContextHandler {
	return &ContextHandler{handler}
}

// Enabled is to required implement slog.Handler
func (h *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

// Handle is to required implement slog.Handler
func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.handler.Handle(ctx, getAttrSetFromContext(ctx).newRecord(r))
}

// WithAttrs is required to implement slog.Handler
func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{h.handler.WithAttrs(attrs)}
}

// WithGroup is required to implement slog.Handler
func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{h.handler.WithGroup(name)}
}
