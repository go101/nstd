package nstd

import (
	"testing"
)

func TestMakeSlice(t *testing.T) {
	const N = 1024
	var m = make(map[int]blank, N)
	for i := range [N]blank{} {
		m[cap(MakeSlice[[]byte](i))] = blank{}
	}
	if len(m) == N {
		t.Fatalf("MakeSlice: len(m) == N (%v vs. %v)", len(m), N)
	}
}

func TestMakeSliceWithMinCap(t *testing.T) {
	var caps = []int{0, 1, 5, 8, 11, 12, 33, 35, 99, 1020, 1021, 1023, 1024}
	var n = 0
	for _, c := range caps {
		var s = MakeSliceWithMinCap[[]bool](c)
		if len(s) != 0 {
			t.Fatalf("MakeSliceWithMinCap: len(s) == 0 (len=%d, c=%d)", len(s), c)
		}
		if cap(s) < c {
			t.Fatalf("MakeSliceWithMinCap: cap(s) is too small (cap=%d, c=%d)", cap(s), c)
		}
		if cap(s) > c {
			n++
		}
	}
	if n == 0 {
		t.Fatal("MakeSliceWithMinCap: no capacity larger than min")
	}
}

func TestSlice(t *testing.T) {
	type S []int
	var x = S{1, 2, 3}
	var y = Slice(x).Clone()
	if TypeOf(x) != TypeOf(y) {
		t.Fatalf("SliceClone: TypeOf(x) != TypeOf(y)\n\t%s\n\t%s", TypeOf(x), TypeOf(y))
	}
	for v := range Slice(x).RefIter {
		*v *= 2
	}
	for i := range x {
		if x[i] != y[i]*2 {
			t.Fatalf("SliceClone: x[%d] != y[%d] * 2 (%d : %d)", i, i, x[i], y[i])
		}
	}

	var z = Slice(x).Unnamed()
	if TypeOf(x) == TypeOf(z) {
		t.Fatalf("SliceClone: TypeOf(x) == TypeOf(z)\n\t%s\n\t%s", TypeOf(x), TypeOf(z))
	}
	z[0] = 99
	if x[0] != z[0] {
		t.Fatalf("SliceClone: x[0] != z[0] (%d : %d)", x[0], z[0])
	}
	if y[0] == z[0] {
		t.Fatalf("SliceClone: y[0] == z[0] (%d : %d)", y[0], z[0])
	}
}
