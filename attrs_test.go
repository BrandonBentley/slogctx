package slogctx

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAttrSet_Empty(t *testing.T) {
	assert.Equal(t, &defaultEmptyAttrSet, newAttrSet())
}

func TestNewAttrSet_Single(t *testing.T) {
	set := []slog.Attr{
		slog.String("k1", "v1"),
		slog.String("k2", "v2"),
		slog.String("k1", "v2"),
	}
	as := newAttrSet(set)
	assert.Len(t, as.attrSlice, 2)
	assert.Len(t, as.attrMap, 2)
	assertNoEmptiesAndUniquenessInAttrSet(t, as)
}

func TestNewAttrSet_Multiple(t *testing.T) {
	attrsSlices := [][]slog.Attr{
		{
			slog.String("k1", "v1"),
			slog.String("k2", "v2"),
			slog.String("k1", "v1.5"),
			slog.String("k3", "v3"),
		},
		{
			slog.String("k1", "v1"),
			slog.String("k2", "v2"),
		},
	}
	as := newAttrSet(attrsSlices...)
	assert.Len(t, as.attrSlice, 3)
	assert.Len(t, as.attrMap, 3)
	assertNoEmptiesAndUniquenessInAttrSet(t, as)
}

func TestNewAttrSet_Multiple_AppendingRequired(t *testing.T) {
	attrsSlices := [][]slog.Attr{
		{
			slog.String("k1", "v1"),
		},
		{
			slog.String("k2", "v2"),
			slog.String("k3", "v3"),
		},
		{
			slog.String("k4", "v4"),
		},
	}
	as := newAttrSet(attrsSlices...)
	assert.Len(t, as.attrSlice, 4)
	assert.Len(t, as.attrMap, 4)
	assertNoEmptiesAndUniquenessInAttrSet(t, as)
}

func TestNewRecord(t *testing.T) {
	set := []slog.Attr{
		slog.String("k1", "v1"),
		slog.String("k2", "v2"),
		slog.String("k1", "v2"),
	}
	as := newAttrSet(set)
	assertNoEmptiesAndUniquenessInAttrSet(t, as)
}

func assertNoEmptiesAndUniquenessInAttrSet(t *testing.T, as *attrSet) {
	duplicateKeysMap := map[string]bool{}
	for _, a := range as.attrSlice {
		assert.NotEmpty(t, a.Value)
		assert.NotEmpty(t, a.Key)
		found := duplicateKeysMap[a.Key]
		if found {
			t.Fatal("duplicate keys in slice", as.attrSlice, as.attrMap)
		}
		duplicateKeysMap[a.Key] = true
	}
	duplicateValuesMap := map[int]bool{}
	for k, index := range as.attrMap {
		assert.NotEmpty(t, k)
		found := duplicateValuesMap[index]
		if found {
			t.Fatal("duplicate indexes in map", as.attrSlice, as.attrMap)
		}
		duplicateValuesMap[index] = true
	}
}
