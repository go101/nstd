package nstd

// MakeSlice makes a slice with the specified length.
// Different from the built-in make function, the capacity
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

type _slice[S ~[]T, T any] []T

// Slice returns an internal representation of s to do some operations.
// Call methods of the internal representation to perform the operations.
func Slice[S ~[]T, T any](s S) _slice[S, T] {
	return []T(s)
}

// Unnamed converts s to unnamed type.
func (s _slice[S, T]) Unnamed() []T {
	return s
}

// Clone clones s and converts the clone result to S.
// Different from [slices.Clone], the result of Clone always has equal length and capacity.
func (s _slice[S, T]) Clone() S {
	var r = make([]T, len(s))
	copy(r, s)
	return r
}

// RefIter is an iterator which iterates references of elements of s.
func (s _slice[S, T]) RefIter(yield func(*T, int) bool) {
	for i := range s {
		if !yield(&s[i], i) {
			return
		}
	}
}
