package slogctx

import (
	"context"
	"log/slog"
	"testing"
)

func TestGetLogger(t *testing.T) {
	ctx := With(
		context.Background(),
		"someKey", "someVal",
	)
	GetLogger(ctx)
}

func TestGetLogger_DefaultNil(t *testing.T) {
	ctx := With(
		context.Background(),
		"someKey", "someVal",
		slog.Bool("itsTheTruth", false),
	)
	defaultLogger = nil
	GetLogger(ctx)
}
