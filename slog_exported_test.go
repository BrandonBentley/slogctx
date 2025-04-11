package slogctx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgsToAttr_Mismatch(t *testing.T) {
	attr, a := argsToAttr([]any{
		"keyWithoutValue",
	})
	assert.Nil(t, a)
	assert.Equal(t, attr.Key, badKey)
	assert.Equal(t, attr.Value.String(), "keyWithoutValue")
}

func TestArgsToAttr_NonStringKey(t *testing.T) {
	attr, a := argsToAttr([]any{
		1,
	})
	assert.Empty(t, a)
	assert.Equal(t, attr.Key, badKey)
	assert.Equal(t, attr.Value.String(), "1")
}
