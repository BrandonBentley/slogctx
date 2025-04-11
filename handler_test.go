package slogctx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextHandler(t *testing.T) {
	buf := &bytes.Buffer{}

	logger := slog.New(NewContextHandler(slog.NewJSONHandler(buf, nil)))

	logger.Handler()
	slog.SetDefault(slog.New(NewJSONContextHandler(buf, nil)))
	ctx := AddAttributesToLogger(
		context.Background(),
		"key1", "value1",
	)
	slog.ErrorContext(ctx, "test")
	ctx = AddAttributesToLogger(
		ctx,
		"key1", "value2",
		"dif", "noop",
	)
	slog.ErrorContext(ctx, "test")
	logs := strings.Split(buf.String(), "\n")

	var log1 map[string]any
	json.Unmarshal([]byte(logs[0]), &log1)
	fmt.Println(log1)

	assert.Equal(t, "value1", log1["key1"])

	var log2 map[string]any
	json.Unmarshal([]byte(logs[1]), &log2)
	fmt.Println(log2)

	assert.Equal(t, "value2", log2["key1"])
	assert.Equal(t, "noop", log2["dif"])
}
