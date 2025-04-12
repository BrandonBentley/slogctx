package slogctx

import (
	"log/slog"
)

type attrSet struct {
	attrMap map[string]slog.Attr
}

func newAttrSet(attrs ...slog.Attr) *attrSet {
	a := &attrSet{
		attrMap: make(map[string]slog.Attr, len(attrs)),
	}

	for _, attr := range attrs {
		a.attrMap[attr.Key] = attr
	}
	return a
}

func (a *attrSet) newRecord(r slog.Record) slog.Record {
	attrCopy := a.copyWith()

	r.Attrs(func(attr slog.Attr) bool {
		attrCopy.attrMap[attr.Key] = attr
		return true
	})

	rr := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
	rr.AddAttrs(attrCopy.attrs()...)
	return rr
}

func (a *attrSet) copyWith(attrs ...slog.Attr) *attrSet {
	attrCopy := &attrSet{
		attrMap: make(map[string]slog.Attr, len(a.attrMap)),
	}

	// deep copy map
	for key, attr := range a.attrMap {
		attrCopy.attrMap[key] = attr
	}

	for _, attr := range attrs {
		attrCopy.attrMap[attr.Key] = attr
	}

	return attrCopy
}

func (a *attrSet) attrs() []slog.Attr {
	attrs := make([]slog.Attr, len(a.attrMap))
	i := 0
	for _, attr := range a.attrMap {
		attrs[i] = attr
		i++
	}
	return attrs
}
