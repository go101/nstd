package nstd

// MakeSlice makes a slice with the specified length.
// Different from the built-in [make] function, the capacity
// of the result slice might be larger than the length.
func MakeSlice[S ~[]E, E any](len int) S {
	return append(S(nil), make(S, len)...)
}

// MakeSliceWithMinCap makes a slice with capacity not smaller than cap.
// The length of the result slice is zero.
//
// See: https://github.com/golang/go/issues/69872
func MakeSliceWithMinCap[S ~[]E, E any](cap int) S {
	return append(S(nil), make(S, cap)...)[:0]
}

// CloneSlice clones a slice.
// Different from [slices.Clone], the result of CloneSlice
// always has equal length and capacity.
func CloneSlice[S ~[]T, T any](s S) S {
	var r = make(S, len(s))
	copy(r, s)
	return r
}

// UnnamedSlice converts s to unnamed type.
func UnnamedSlice[S ~[]T, T any](s S) []T {
	return s
}

// SliceElemPointers returns an iterator which iterates element pointers of a slice.
func SliceElemPointers[E any](s []E) func(func(*E) bool) {
	return func(yield func(*E) bool) {
		for i := range s {
			if !yield(&s[i]) {
				break
			}
		}
	}
}
