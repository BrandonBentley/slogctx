package slogctx

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalFactory(t *testing.T) {
	buf := &bytes.Buffer{}

	ctx := context.Background()
	defaultLogger := slog.Default()
	assert.Equal(t, defaultLogger, GetContextLogger(ctx))

	newDefaultLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(newDefaultLogger)
	assert.Equal(t, newDefaultLogger, GetContextLogger(ctx))

	newLogger := slog.New(slog.NewJSONHandler(buf, nil))
	SetRootLoggerFactory(NewFactory(newLogger))
	slog.SetDefault(newLogger)
	assert.Equal(t, newLogger, GetContextLogger(ctx))

	ctx2 := AddAttributesToContextLogger(ctx,
		"key1", "value1",
		slog.Bool("bool1", true),
	)

	assert.Equal(t, newLogger, GetContextLogger(ctx))
	assert.NotEqual(t, newLogger, GetContextLogger(ctx2))

	GetContextLogger(ctx).Info(
		"a message",
	)

	testStruct := struct {
		Msg   string `json:"msg"`
		Key1  string `json:"key1"`
		Bool1 bool   `json:"bool1"`
	}{}

	jd := json.NewDecoder(buf)
	err := jd.Decode(&testStruct)
	assert.NoError(t, err)

	assert.Equal(t, "a message", testStruct.Msg)
	assert.Equal(t, "", testStruct.Key1)
	assert.False(t, testStruct.Bool1)

	GetContextLogger(ctx2).Info(
		"a ctx2 message",
	)

	err = jd.Decode(&testStruct)
	assert.NoError(t, err)

	assert.Equal(t, "a ctx2 message", testStruct.Msg)
	assert.Equal(t, "value1", testStruct.Key1)
	assert.True(t, testStruct.Bool1)
}

func TestImpossibleError(t *testing.T) {
	NewFactoryNoOp()

	buf := &bytes.Buffer{}

	newLogger := slog.New(slog.NewJSONHandler(buf, nil))
	f := NewFactory(newLogger)

	defaultLogger := f.GetContextLogger(context.Background())

	assert.NotNil(t, defaultLogger)

	ctx := context.WithValue(context.Background(), loggerContextKey, true)
	actualLogger := f.GetContextLogger(ctx)

	assert.Equal(t, defaultLogger, actualLogger)

	strings.Contains(buf.String(), "Value at logger context key is not a *slog.Logger. Defaulting to standard logger")
}
