package nstd_test

import (
	"reflect"
	"testing"

	"go101.org/nstd"
)

type blank = struct{} // used internally

func TestZero(t *testing.T) {
	testZero(1, 0, t)
	testZero("go", "", t)
	testZero(true, false, t)
	testZero(struct{}{}, struct{}{}, t)
	testZero([2]byte{1, 2}, [2]byte{}, t)
	testZero([]byte{1, 2}, nil, t)
}

func testZero[T any](v, zero T, t *testing.T) {
	if z := nstd.Zero(v); !reflect.DeepEqual(z, zero) {
		t.Fatalf("Zero(%v) != %v (but %v)", v, z, zero)
	}
}

func TestZeroIt(t *testing.T) {
	testZeroIt(1, 0, t)
	testZeroIt("go", "", t)
	testZeroIt(true, false, t)
	testZeroIt(struct{}{}, struct{}{}, t)
	testZeroIt([2]byte{1, 2}, [2]byte{}, t)
	testZeroIt([]byte{1, 2}, nil, t)
}

func testZeroIt[T any](v, zero T, t *testing.T) {
	var old = v
	var p = &v
	if nstd.ZeroIt(p); !reflect.DeepEqual(v, zero) {
		t.Fatalf("ZeroIt(&%v) != %v (but %v)", old, v, zero)
	}
}

func TestNew(t *testing.T) {
	testNew(1, t)
	testNew("go", t)
	testNew(true, t)
	testNew(struct{}{}, t)
	testNew([2]byte{1, 2}, t)
	testNew([]byte{1, 2}, t)
}

func testNew[T any](v T, t *testing.T) {
	var pv = &v
	if p := nstd.New(v); p == pv {
		t.Fatalf("New(&%v) == &%v", v, v)
	} else if !reflect.DeepEqual(*p, *pv) {
		t.Fatalf("*New(&%v) != %v (but %v)", v, v, *p)
	}
}

func TestTypeAssert(t *testing.T) {
	testTypeAssert(1, (*int)(nil), true, t)
	testTypeAssert(1, nstd.New(123), true, t)
	testTypeAssert(1, nstd.New(true), false, t)
	testTypeAssert(true, nstd.New(true), true, t)
	testTypeAssert(1, nstd.New(true), false, t)
	testTypeAssert(true, nstd.New(any(0)), true, t)
}

func testTypeAssert[T any](v any, p *T, shouldOkay bool, t *testing.T) {
	if nstd.TypeAssert(v, p) != shouldOkay {
		t.Fatalf("*TypeAssert[%T](%v, %T) != %v", *p, v, p, shouldOkay)
	} else if shouldOkay && p != nil && !reflect.DeepEqual(v, *p) {
		t.Fatalf("*TypeAssert[%T](%v, %T) fails. (got %v)", *p, v, p, *p)
	}
}

func TestHasEntry(t *testing.T) {
	testHasEntry(map[int]int{1: 2}, 1, true, t)
	testHasEntry(map[int]int{1: 2}, 2, false, t)
	testHasEntry(map[bool]int{true: 2}, true, true, t)
	testHasEntry(map[bool]int{false: 2}, true, false, t)
}

func testHasEntry[K comparable, E any](m map[K]E, key K, shouldOkay bool, t *testing.T) {
	if nstd.HasEntry(m, key) != shouldOkay {
		t.Fatalf("*HasEntry(%v, %v) != %v", m, key, shouldOkay)
	}
}
