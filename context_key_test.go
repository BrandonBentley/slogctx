package slogctx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAttrSetFromContext_Panic(t *testing.T) {
	ctx := context.WithValue(
		context.Background(),
		attrSetContextKey,
		"InvalidValue",
	)

	assert.Equal(t, &defaultEmptyAttrSet, getAttrSetFromContext(ctx))
}
