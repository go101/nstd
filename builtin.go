package nstd

import (
	"fmt"
)

func Panic(a ...any) bool {
	panic(fmt.Sprint(a...))
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

func Assertf(b bool, format string, a ...any) {
	if !b {
		Panicf(format, a...)
	}
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
