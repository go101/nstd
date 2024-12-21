package nstd

import (
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	var v = 123
	var x any = v

	if TypeOf(x).Kind() != reflect.Interface {
		t.Fatal("type of x should be Inferface")
	}
	if ValueOf(x).Kind() != reflect.Interface {
		t.Fatal("type of x should be Inferface")
	}
	if Type[any]().Kind() != reflect.Interface {
		t.Fatal("any should be Inferface")
	}

	if TypeOf(v).Kind() == reflect.Interface {
		t.Fatal("type of v should not be Inferface")
	}
	if ValueOf(v).Kind() == reflect.Interface {
		t.Fatal("type of v should not be Inferface")
	}
	if Type[int]().Kind() == reflect.Interface {
		t.Fatal("int should not be Inferface")
	}
}
