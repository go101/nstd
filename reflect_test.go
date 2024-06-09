package nstd_test

import (
	"reflect"
	"testing"

	"go101.org/nstd"
)

func TestReflect(t *testing.T) {
	var v = 123
	var x any = v

	if nstd.TypeOf(x).Kind() != reflect.Interface {
		t.Fatal("type of x should be Inferface")
	}
	if nstd.ValueOf(x).Kind() != reflect.Interface {
		t.Fatal("type of x should be Inferface")
	}
	if nstd.Type[any]().Kind() != reflect.Interface {
		t.Fatal("any should be Inferface")
	}

	if nstd.TypeOf(v).Kind() == reflect.Interface {
		t.Fatal("type of v should not be Inferface")
	}
	if nstd.ValueOf(v).Kind() == reflect.Interface {
		t.Fatal("type of v should not be Inferface")
	}
	if nstd.Type[int]().Kind() == reflect.Interface {
		t.Fatal("int should not be Inferface")
	}
}
