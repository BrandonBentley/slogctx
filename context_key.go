package slogctx

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type uniqueContextAttrsKey string

const attrSetContextKey = uniqueContextAttrsKey("context-logger-attr-set")

var defaultEmptyAttrSet = attrSet{}

func getAttrSetFromContext(ctx context.Context) (as *attrSet) {
	val := ctx.Value(attrSetContextKey)
	if val == nil {
		return &defaultEmptyAttrSet
	}

	defer func() {
		if recover() != nil {
			as = &defaultEmptyAttrSet
			freshStdoutLogger().Error(
				"slogctx context key has been overwritten with incorrect type",
				slog.String("type", fmt.Sprintf("%T", val)),
			)
		}
	}()

	return val.(*attrSet)
}

func freshStdoutLogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)
}
