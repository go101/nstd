package nstd

// MakeSlice makes a slice with the specified length.
// Different from the built-in make function, the capacity
// of the result slice might be larger than the length.
func MakeSlice[S ~[]E, E any](len int) S {
	return append(S(nil), make(S, len)...)
}

// CollectMapKeys collects all the keys in a map into freshly
// created slice. The length and capacity of the result slice
// are both the length of the map.
//
// See:
//     https://github.com/golang/go/issues/68261
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
