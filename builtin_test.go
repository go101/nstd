package nstd

import (
	"reflect"
	"testing"
)

func TestZeroOf(t *testing.T) {
	testZeroOf(1, 0, t)
	testZeroOf("go", "", t)
	testZeroOf(true, false, t)
	testZeroOf(struct{}{}, struct{}{}, t)
	testZeroOf([2]byte{1, 2}, [2]byte{}, t)
	testZeroOf([]byte{1, 2}, nil, t)
}

func testZeroOf[T any](v, zero T, t *testing.T) {
	if z := ZeroOf(v); !reflect.DeepEqual(z, zero) {
		t.Fatalf("Zero(%v) != %v (but %v)", v, z, zero)
	}
}

func TestZero(t *testing.T) {
	testZero(1, 0, t)
	testZero("go", "", t)
	testZero(true, false, t)
	testZero(struct{}{}, struct{}{}, t)
	testZero([2]byte{1, 2}, [2]byte{}, t)
	testZero([]byte{1, 2}, nil, t)
}

func testZero[T any](v, zero T, t *testing.T) {
	var old = v
	var p = &v
	if Zero(p); !reflect.DeepEqual(v, zero) {
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
	if p := New(v); p == pv {
		t.Fatalf("New(&%v) == &%v", v, v)
	} else if !reflect.DeepEqual(*p, *pv) {
		t.Fatalf("*New(&%v) != %v (but %v)", v, v, *p)
	}
}

func TestTypeAssert(t *testing.T) {
	testTypeAssert(1, (*int)(nil), true, t)
	testTypeAssert(1, New(123), true, t)
	testTypeAssert(1, New(true), false, t)
	testTypeAssert(true, New(true), true, t)
	testTypeAssert(1, New(true), false, t)
	testTypeAssert(true, New(any(0)), true, t)
}

func testTypeAssert[T any](v any, p *T, shouldOkay bool, t *testing.T) {
	if TypeAssert(v, p) != shouldOkay {
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
	if HasEntry(m, key) != shouldOkay {
		t.Fatalf("*HasEntry(%v, %v) != %v", m, key, shouldOkay)
	}
}
