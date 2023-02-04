package nstd

import (
	"fmt"
)

func Panic(a ...any) bool {
	panic(fmt.Sprint(a...))
	return true
}

func Panicln(a ...any) bool {
	panic(fmt.Sprintln(a...))
	return true
}

func Panicf(format string, a ...any) bool {
	panic(fmt.Sprintf(format, a...))
	return true
}

func Assert(b bool, a ...any) {
	if !b {
		panic(fmt.Sprint(a...))
	}
}

func Assertln(b bool, a ...any) {
	if !b {
		panic(fmt.Sprintln(a...))
	}
}

func Assertf(b bool, format string, a ...any) {
	if !b {
		Panicf(format, a...)
	}
}

// https://github.com/golang/go/issues/58280
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func ZeroOf[T any]() T {
	var x T
	return x
}

func ZeroFor[T any](_ T) T {
	var x T
	return x
}

func Zero[T any](p *T) {
	var x T
	*p = x
}

// https://github.com/golang/go/issues/27804
// https://github.com/golang/go/issues/36449
func Eval[T any](v T) T {
	return v
}
