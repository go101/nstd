/*
Package slice provides several common slice operations,
with more options than the ones provided in the [builtin]
and [slices] standard packages.
*/
package slices

// A MakeOp is used to make a new slice.
// By default, a MakeOp tries to use up hosting memory blocks of slice elements.
//
// ToDo: move to the top "nstd" package?
type MakeOp[S ~[]T, T any] struct {
	Clipped bool
}

func NewMakeOp[S ~[]T, T any](_ S) *MakeOp[S, T] {
	return &MakeOp[S, T]{}
}

// The Clipped property os a MakeOp is used to control try to
// use up hosting memory blocks of slice elements or not.
func (maker *MakeOp[S, T]) SetClipped(b bool) *MakeOp[S, T] {
	maker.Clipped = b
	return maker
}

func (maker *MakeOp[S, T]) Do(len int) S {
	if maker.Clipped {
		return make(S, len)
	}
	return append(S(nil), make(S, len)...)
}

// A MergeOp is used to merge multipl slices (with the same element type).
// By default, a MergeOp tries to use up hosting memory blocks of slice elements.
//
// ToDo: move to the top "nstd" package?
type MergeOp[S ~[]T, T any] struct {
	Clipped bool
}

func NewMergeOp[S ~[]T, T any](_ S) *MergeOp[S, T] {
	return &MergeOp[S, T]{}
}

// The Clipped property os a MergeOp is used to control try to
// use up hosting memory blocks of slice elements or not.
func (merger *MergeOp[S, T]) SetClipped(b bool) *MergeOp[S, T] {
	merger.Clipped = b
	return merger
}

func (merger *MergeOp[S, T]) Do(ss ...S) S {
	var n, allNils, k = 0, true, -1
	for i := range ss {
		if m := len(ss[i]); n != 0 {
			n += m
			if n < 0 {
				panic("sum of lengths is too large")
			}
		} else if m > 0 {
			n = m
			allNils = false
			k = i
		} else {
			allNils = allNils && ss[i] == nil
		}
	}
	if allNils {
		return nil
	}
	if n == 0 {
		return S{}
	}

	return merger.do(n, ss[k], ss[k+1:]...)
}

func (merger *MergeOp[S, T]) do(n int, first S, ss ...S) S {
	var r []T
	if merger.Clipped {
		// Make use of this optimization:
		// https://github.com/golang/go/commit/6ed4661807b219781d1aa452b7f210e21ad1974b
		r = make(S, n)
		copy(r, first)
		r = r[:len(first)]
	} else {
		if m := len(first); m == n {
			return NewCloneOp(S{}).Do(first) // guarantee result and first don't share backing array
		} else {
			r = append(first[:len(first):len(first)], make(S, n-m)...)
		}
	}

	for _, s := range ss {
		r = append(r, s...)
	}
	return r
}

// A CloneOp is used to clone a slice.
// By default, a CloneOp tries to use up hosting memory blocks of slice elements.
//
// ToDo: move to the top "nstd" package?
type CloneOp[S ~[]T, T any] struct {
	Clipped bool
}

func NewCloneOp[S ~[]T, T any](_ S) *CloneOp[S, T] {
	return &CloneOp[S, T]{}
}

// The Clipped property os a CloneOp is used to control try to
// use up hosting memory blocks of slice elements or not.
func (cloner *CloneOp[S, T]) SetClipped(b bool) *CloneOp[S, T] {
	cloner.Clipped = b
	return cloner
}

func (cloner *CloneOp[S, T]) Do(s S) S {
	// Preserve nil in case it matters.
	if s == nil {
		return nil
	}

	if cloner.Clipped {
		// Make use of this optimization:
		// https://github.com/golang/go/commit/6ed4661807b219781d1aa452b7f210e21ad1974b
		r := make(S, len(s))
		copy(r, s)
		return s
	}

	return append(S(nil), s...)
}

// MultiCopy copies multiple (source) slices into the target slice (into).
// It returns the number of copied elements, the number of fully copied slices,
// and the number of copied elements of the partically copied slices.
//
// Warning: this function doesn't detect element overlapping. So it is best
// to make sure the target slice doesn't overlap any source slice.
//
// ToDo: move to the top "nstd" package?
func MultiCopy[S ~[]T, T any](into S, ss ...S) (nElements, nCompletes, nElementsOfIncomplete int) {
	var n = 0
	for i, s := range ss {
		m := copy(into[n:], s)
		n += m
		if m < len(s) {
			return n, i, m
		}
	}
	return n, len(ss), 0
}

// MultiAppend appends (elements of) multiple (source) slices
// into the target slice (into) and returns a slice composed of the elements
// if the target slice and all source slices.
// A new backing array will be allocated if the free slots of the target slice
// are not enough to hold all of the appended elements, in which case,
// the append operation actually merges the target slice and the source slices.
//
// Warning: this function doesn't detect element overlapping. So it is best
// to make sure the target slice doesn't overlap any source slice.
//
// ToDo: move to the top "nstd" package?
func MultiAppend[S ~[]T, T any](into S, ss ...S) S {
	var n = 0
	for i := range ss {
		n += len(ss[i])
		if n < 0 {
			panic("sum of lengths is too large")
		}
	}

	if n+len(into) < 0 {
		panic("sum of lengths is too large")
	}

	if n+len(into) <= cap(into) {
		MultiCopy(into[len(into):len(into)+n], ss...)
		return into[:len(into)+n]
	}

	return NewMergeOp(S{}).do(n+len(into), into, ss...)
}

// ElemPtrs builds a new slice which elements are the pointers
// to the elements in the input slice.
func ElemPtrs[S ~[]T, T any](s S) []*T {
	if s == nil {
		return nil
	}
	r := make([]*T, len(s), cap(s))
	for i := range s {
		r[i] = &s[i]
	}
	return r
}
