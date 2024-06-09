package nstd

import (
	"reflect"
)

// Type returns a reflect.Type which represents type T,
// which may be either an non-interface type or interface type.
func Type[T any]() reflect.Type {
	var v T
	return ValueOf(v).Type()
}

// TypeOf returns a reflect.Type which represents the type of v,
// which may be either an non-interface value or interface value.
func TypeOf[T any](v T) reflect.Type {
	return ValueOf(v).Type()
}

// Value returns a reflect.Value which represents the value v,
// which may be either an non-interface value or interface value.
func ValueOf[T any](v T) reflect.Value {
	// make sure r.CanAddr() and r.CanSet() both always return false,
	// even if the passed argument is an interface.
	return reflect.ValueOf([1]T{v}).Index(0)
}
