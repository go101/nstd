package nstd

import (
	"fmt"
)

// Generally, Panicf(format, v...) is a short form of panic(fmt.Sprintf(format, v...)).
// When format is blank, then it is a short form of panic(fmt.Sprint(v...)).
func Panicf(format string, a ...any) bool {
	if format == "" {
		panic(fmt.Sprint(a...))
	} else {
		panic(fmt.Sprintf(format, a...))
	}
	return true
}

// Must panics if err is not nil; otherwise, the T value is returned.
//
// See https://github.com/golang/go/issues/58280
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Must2 panics if err is not nil; otherwise, the T1 and T2 values are returned.
func Must2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return v1, v2
}

// Eval is used to ensure the evaluation order of some expressions
// in a statement.
//
// See:
// * https://go101.org/article/evaluation-orders.html
// * https://github.com/golang/go/issues/27804
// * https://github.com/golang/go/issues/36449
func Eval[T any](v T) T {
	return v
}

// ZeroOf[T]() and ZeroOf(valueOfT) both return the zero value of type T.
func ZeroOf[T any](...T) T {
	var x T
	return x
}

// Zero(p) zeros the value referenced by the pointer p.
// Zero is useful for resetting values of some unexported types,
// or resetting values of some other packages but without importing those packages.
func Zero[T any](p *T) {
	var x T
	*p = x
}

// New allocates a T value and initialize it with the specified one.
func New[T any](v T) *T {
	return &v
}

// SliceFrom is used to create a slice from some values of the same type.
// Some use scenarios:
//  1. Convert multiple results of a function call to a []any slice,
//     then use the slice in fmt.Printf alike functions.
//  2. Construct a []T slice from some T values without using the []T{...} form.
//
// NOTE: SliceFrom(aSlice...) returns aSlice,
//
// See:
// * https://github.com/golang/go/issues/61213
func SliceFrom[T any](vs ...T) []T {
	return vs
}

// TypeAssert asserts an interface value x to type T.
// If the assertion succeeds, true is returned, othewise, false is returned.
// If into is not nil, then the concrete value of x will be assigned to
// the value referenced by into.
//
// See:
//
//	https://github.com/golang/go/issues/65846
func TypeAssert[T any](x any, into *T) (ok bool) {
	if into != nil {
		*into, ok = x.(T)
	} else {
		_, ok = x.(T)
	}
	return
}

// HasMapEntry checks whether or not a map contains an entry
// with the specified key.
func HasEntry[K comparable, E any](m map[K]E, key K) bool {
	_, ok := m[key]
	return ok
}
