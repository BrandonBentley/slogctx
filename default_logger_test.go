package slogctx

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultLogger(t *testing.T) {
	SetSlogPackageDefault()
	buf := &bytes.Buffer{}
	SetSlogPackageWithOptions(buf, nil)
	ctx := WithAttrs(context.Background(), slog.String("keyTest", "testValue"))
	slog.ErrorContext(
		ctx,
		"This is the message",
		"k", "v",
	)

	var log1 map[string]any
	json.Unmarshal(buf.Bytes(), &log1)

	assert.Equal(t, "v", log1["k"])
	assert.Equal(t, "testValue", log1["keyTest"])

}
