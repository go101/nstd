package nstd

// MakeSlice makes a slice with the specified length.
// Different from the built-in make function, the capacity
// of the result slice might be larger than the length.
func MakeSlice[S ~[]E, E any](len int) S {
	return append(S(nil), make(S, len)...)
}

// BlankMap returns a blank map which has the same type as the input map.
// * Usage 1: `ZeroMap[MapType](nil, 32)
// * Usage 2: `ZeroMap(aMap, 8)
func ZeroMap[M ~map[K]E, K comparable, E any](m M, capHint int) M {
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
