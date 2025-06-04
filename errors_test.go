package nstd

import (
	"strings"
	"testing"
)

func TestNewError(t *testing.T) {
	var base = strings.Repeat("abc", 3)
	var e1 = NewError(base[0:3])
	var e2 = NewError(base[3:6])
	var e3 = NewError(base[6:9])
	if e1 != e2 {
		t.Fatal("NewError: e1 != e2")
	}
	if e1 != e3 {
		t.Fatal("NewError: e1 != e3")
	}
	if e3 != e2 {
		t.Fatal("NewError: e3 != e2")
	}
}
