package slogctx

import (
	"log/slog"
)

type attrSet struct {
	attrSlice []slog.Attr
	attrMap   map[string]int
}

// higher index values with matching attr.Key will
// override the previous values.
// ex: attrs[0][*].Key == attrs[1][*].Key { value = attrs[1][*] }
func newAttrSet(attrs ...[]slog.Attr) *attrSet {
	if len(attrs) == 0 {
		return &defaultEmptyAttrSet
	} else if len(attrs) > 1 {
		return newAttrSetMultipleSlices(attrs...)
	}

	a := &attrSet{
		attrSlice: make([]slog.Attr, len(attrs[0])),
		attrMap:   make(map[string]int, len(attrs[0])),
	}
	i := 0
	for _, attr := range attrs[0] {
		if index, ok := a.attrMap[attr.Key]; ok {
			a.attrSlice[index] = attr
			continue
		}
		a.attrSlice[i] = attr
		a.attrMap[attr.Key] = i
		i++
	}
	if i < len(a.attrSlice) {
		a.attrSlice = a.attrSlice[:i]
	}
	return a
}

// higher index values with matching attr.Key will
// override the previous values.
// ex: attrs[0][*].Key == attrs[1][*].Key { value = attrs[1][*] }
func newAttrSetMultipleSlices(attrs ...[]slog.Attr) *attrSet {
	minLength := minLen(attrs)
	a := &attrSet{
		attrSlice: make([]slog.Attr, minLength),
		attrMap:   make(map[string]int, minLength),
	}
	i := 0
	for _, attrSlice := range attrs {
		for _, attr := range attrSlice {
			if index, ok := a.attrMap[attr.Key]; ok {
				a.attrSlice[index] = attr
				continue
			}
			a.attrMap[attr.Key] = i
			if i >= minLength {
				a.attrSlice = append(a.attrSlice, attr)
			} else {
				a.attrSlice[i] = attr
			}
			i++
		}
	}
	if i < len(a.attrSlice) {
		a.attrSlice = a.attrSlice[:i]
	}
	return a
}

func (a *attrSet) newRecord(r slog.Record) slog.Record {
	if r.NumAttrs() == 0 {
		r.AddAttrs(a.attrSlice...)
		return r
	}

	attrsCopy := make([]slog.Attr, len(a.attrSlice))

	copy(attrsCopy, a.attrSlice)

	r.Attrs(func(attr slog.Attr) bool {
		if index, ok := a.attrMap[attr.Key]; ok {
			attrsCopy[index] = attr
		}
		attrsCopy = append(attrsCopy, attr)
		return true
	})

	rr := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
	rr.AddAttrs(attrsCopy...)
	return rr
}

func (a *attrSet) with(attrs []slog.Attr) *attrSet {
	if a.attrMap == nil {
		return newAttrSet(attrs)
	}
	newAttrSet := a.copy()
	newAttrSet.setNewAttrs(attrs)
	return newAttrSet
}

func (a *attrSet) copy() *attrSet {
	newSet := &attrSet{
		attrSlice: make([]slog.Attr, len(a.attrSlice)),
		attrMap:   make(map[string]int, len(a.attrMap)),
	}

	// deep copy slice
	copy(newSet.attrSlice, a.attrSlice)

	// deep copy map
	for key, index := range a.attrMap {
		newSet.attrMap[key] = index
	}

	return newSet
}

func (a *attrSet) setNewAttrs(attrs []slog.Attr) {
	for _, attr := range attrs {
		if index, ok := a.attrMap[attr.Key]; ok {
			a.attrSlice[index] = attr
			continue
		}
		a.attrMap[attr.Key] = len(a.attrSlice)
		a.attrSlice = append(a.attrSlice, attr)
	}
}

func minLen(attrs [][]slog.Attr) int {
	length := len(attrs[0])
	for _, attrSlice := range attrs[1:] {
		if len(attrSlice) > length {
			length = len(attrSlice)
		}
	}
	return length
}
