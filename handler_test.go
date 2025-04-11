package slogctx

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextHandler(t *testing.T) {
	buf := &bytes.Buffer{}

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

	slog.ErrorContext(
		ctx,
		"test",
		"someKey", "someVal",
		"key1", "value3",
	)

	logs := strings.Split(buf.String(), "\n")

	var log1 map[string]any
	json.Unmarshal([]byte(logs[0]), &log1)

	assert.Equal(t, "value1", log1["key1"])

	var log2 map[string]any
	json.Unmarshal([]byte(logs[1]), &log2)

	assert.Equal(t, "value3", log2["key1"])
	assert.Equal(t, "noop", log2["dif"])

	buf2 := &bytes.Buffer{}
	logger := slog.New(
		NewContextHandler(slog.NewJSONHandler(buf2, nil)).WithGroup("group1"),
	)

	ctx2 := WithAttrs(
		context.Background(),
		slog.String("someKey", "someVal"),
		slog.String("dontChange", "this"),
	)

	logger.ErrorContext(
		ctx2,
		"Test Error",
		"key", "value",
		slog.String("someKey", "difVal"),
	)

	var groupLog map[string]any
	json.Unmarshal(buf2.Bytes(), &groupLog)

	assert.Equal(t, "ERROR", groupLog["level"])
	assert.Equal(t, "Test Error", groupLog["msg"])

	group1, ok := groupLog["group1"].(map[string]any)
	if !ok {
		t.Fatal("group1 not a map[string]any")
	}

	assert.Equal(t, "difVal", group1["someKey"])
	assert.Equal(t, "this", group1["dontChange"])
	assert.Equal(t, "value", group1["key"])
}
