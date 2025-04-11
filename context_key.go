package slogctx

import (
	"context"
)

type uniqueContextAttrsKey string

const attrSetContextKey = uniqueContextAttrsKey("context-logger-attr-set")

func getAttrSetFromContext(ctx context.Context) attrSet {
	val := ctx.Value(attrSetContextKey)
	if val == nil {
		return newAttrSet()
	}
	as, ok := val.(attrSet)
	if !ok {
		return newAttrSet()
	}
	return as
}
