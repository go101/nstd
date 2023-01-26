/*
Package nstd provides some common used functions and types
missing in the Go standard library.
*/
package nstd

import (
	"fmt"
)

func Panicf(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}

func Assert(b bool, msg any) {
	if !b {
		panic(msg)
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
