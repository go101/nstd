package nstd

type SliceOpConfig struct {
	_    [0]func()
	Clip bool

	// Potential more:
	// Arera 
}

func mergeSlices[S ~[]E, E any](n int, clip bool, first S, ss ...S) S {
	var r []E
	if clip {
		// Make use of this optimization:
		// https://github.com/golang/go/commit/6ed4661807b219781d1aa452b7f210e21ad1974b
		r = make(S, n)
		copy(r, first)
		r = r[:len(first)]
	} else {
		if m := len(first); m == n {
			return append(S(nil), first...) // guarantee result and first don't share backing array
		} else {
			r = append(first[:len(first):len(first)], make(S, n-m)...)
		}
	}

	for _, s := range ss {
		r = append(r, s...)
	}
	return r
}

// CloneSlice merge a list of slices. If clip is true,
// then there are no unused element slots in the result.
// The function always creates a new slice and returns it.
// If there is only one slice in ss, then that slice is cloned.
func MergeSlices[S ~[]E, E any](cfg SliceOpConfig, ss ...S) S {
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

	return mergeSlices(n, cfg.Clip, ss[k], ss[k+1:]...)
}

// CloneSlice clones a slice. If clip is true, then
// there are no unused element slots in the result.
func CloneSlice[S ~[]E, E any](cfg SliceOpConfig, s S) S {
	// Preserve nil in case it matters.
	if s == nil {
		return nil
	}

	if cfg.Clip {
		// Make use of this optimization:
		// https://github.com/golang/go/commit/6ed4661807b219781d1aa452b7f210e21ad1974b
		r := make(S, len(s))
		copy(r, s)
		return s
	}

	return append(S(nil), s...)
}

// MakeSlice makes a slice with the specified minimum capacity.
// The length and capacity of the result slice are equal
// and are both larger than minLength for sure
func MakeSlice[S ~[]E, E any, LenType Integer](minLength LenType) S {
	return append(S(nil), make(S, minLength)...)
}

// MultiCopy copies multiple (source) slices into the target slice (into).
// It returns the number of copied elements, the number of fully copied slices,
// and the number of copied elements of the partically copied slices.
//
// Warning: this function doesn't detect element overlapping. So it is best
// to make sure the target slice doesn't overlap any source slice.
//
// ToDo: move to the top "nstd" package?
func MultiCopy[S ~[]E, E any](into S, ss ...S) (nElements, nCompletes, nElementsOfIncomplete int) {
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
func MultiAppend[S ~[]E, E any](into S, ss ...S) S {
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

	return mergeSlices(n+len(into), false, into, ss...)
}

// ElemPtrs builds a new slice which elements are the pointers
// to the elements in the input slice.
func ElemPtrs[S ~[]E, E any](s S) []*E {
	if s == nil {
		return nil
	}
	r := make([]*E, len(s), cap(s))
	for i := range s {
		r[i] = &s[i]
	}
	return r
}

func Clipped[S ~[]E, E any](s S) []E {
	return s[:len(s):len(s)]
}

func Unused[S ~[]E, E any](s S) []E {
	return s[len(s):]
}