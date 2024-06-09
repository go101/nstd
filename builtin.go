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

// Zero[T]() returns the zero value of type T.
func Zero[T any](...T) T {
	var x T
	return x
}

// ZeroIt(p) zeros the value referenced by the pointer p.
// ZeroIt is useful for resetting values of some unexported types,
// or resetting values of some other packages but without importing those packages.
func ZeroIt[T any](p *T) {
	var x T
	*p = x
}

// New allocates a T value and initialize it with the specified one.
func New[T any](v T) *T {
	return &v
}

// SliceFrom is used to create a slice from some values of the same type.
// Some use scenarios:
// 1. Convert multiple results of a function call to a []any slice,
//    then use the slice in fmt.Printf alike functions.
// 2. Construct a []T slice from some T values without using the []T{...} form.
//
// NOTE: SliceFrom(aSlice...) returns aSlice,
//
// See:
// * https://github.com/golang/go/issues/61213
func SliceFrom[T any](vs ...T) []T {
	return vs
}

// IsOfType checks whether or not the concrete type of x
// is of type T.
//
// See:
//     https://github.com/golang/go/issues/65846
func IsOfType[T any](x any) bool {
	_, ok := x.(T)
	return ok
}

// AssertInto asserts an interface value x into the value referenced by t.
// If the assertion succeeds, true is returned, othewise, false is returned.
func AssertInto[T any](x any, t *T) (ok bool) {
	*t, ok = x.(T)
	return
}

// HasMapEntry checks whether or not a map contains an entry
// with the specified key.
func HasEntry[K comparable, E any](m map[K]E, key K) bool {
	_, ok := m[key]
	return ok
}
