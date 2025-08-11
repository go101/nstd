package nstd

// BlankMap returns a blank map which has the same type as the input map.
// It is mainly used to avoid writing some verbose map type literals.
//
// * Usage 1: BlankMap[MapType](nil, 32)
// * Usage 2: BlankMap(aMap, 8)
func BlankMap[M ~map[K]E, K comparable, E any](_ M, capHint int) M {
	return make(M, capHint)
}

// CollectMapKeys collects all the keys in a map into a freshly
// created result slice. The length and capacity of the result slice
// are both equal to the length of the map.
//
// See: https://github.com/golang/go/issues/68261
func CollectMapKeys[K comparable, E any](m map[K]E) []K {
	if len(m) == 0 {
		return nil
	}

	var s = make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// AppendMapKeys appends all the keys in a map into the specified slice.
func AppendMapKeys[K comparable, E any](s []K, m map[K]E) []K {
	for k := range m {
		s = append(s, k)
	}
	return s
}

// BoolKeyMap is an optimized version of map[K]E, where K is a bool type.
type BoolKeyMap[K ~bool, E any] struct {
	trueE  E
	falseE E
}

// Put puts an entry {k, e} into m.
func (m *BoolKeyMap[K, E]) Put(k K, e E) {
	if k {
		m.trueE = e
	} else {
		m.falseE = e
	}
}

// Get returns the element indexed by key k.
func (m *BoolKeyMap[K, E]) Get(k K) E {
	if k {
		return m.trueE
	} else {
		return m.falseE
	}
}

// BoolElementMap is optimized version of map[K]E, where E is a bool type.
// Entries with false element value will not be put in BoolElementMap maps.
type BoolElementMap[K comparable, E ~bool] struct {
	m map[K]blank
}

// Put puts an entry {k, e} into m.
// Note, if e is false and the corresponding entry exists, the entry is deleted.
func (m *BoolElementMap[K, E]) Put(k K, e E) {
	if e {
		if m.m == nil {
			m.m = make(map[K]blank)
		}
		m.m[k] = blank{}
	} else if m.m != nil {
		delete(m.m, k)
	}
}

// Get returns the element indexed by key k.
func (m *BoolElementMap[K, E]) Get(k K) E {
	_, has := m.m[k]
	return E(has)
}
