package slogctx

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAttrSet_Empty(t *testing.T) {
	as := newAttrSet()
	assert.Len(t, as.attrMap, 0)
	assert.NotNil(t, as.attrMap)
}

func TestNewAttrSet(t *testing.T) {
	set := []slog.Attr{
		slog.String("k1", "v1"),
		slog.String("k2", "v2"),
		slog.Bool("k1", true),
	}
	as := newAttrSet(set...)
	assert.Len(t, as.attrMap, 2)
	assert.True(t, as.attrMap["k1"].Value.Bool())
	assert.Equal(t, "v2", as.attrMap["k2"].Value.String())
}

func TestGetAttrSetFromContext_Panic(t *testing.T) {
	ctx := context.WithValue(
		context.Background(),
		attrSetContextKey,
		"InvalidValue",
	)

	assert.Len(t, getAttrSetFromContext(ctx).attrMap, 0)
}
