package slogctx

import (
	"context"
)

type uniqueContextAttrsKey string

const attrSetContextKey = uniqueContextAttrsKey("context-logger-attr-set")

func getAttrSetFromContext(ctx context.Context) (as *attrSet) {
	val := ctx.Value(attrSetContextKey)
	if val == nil {
		return newAttrSet()
	}

	attrs, ok := val.(*attrSet)
	if !ok {
		return newAttrSet()
	}

	return attrs
}
