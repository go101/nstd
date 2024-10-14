package nstd_test

import (
	"math/rand"
	"testing"

	"go101.org/nstd"
)

func TestMakeSlice(t *testing.T) {
	const N = 1024
	var m = make(map[int]blank, N)
	for i := range [N]blank{} {
		m[cap(nstd.MakeSlice[[]byte](i))] = blank{}
	}
	if len(m) == N {
		t.Fatalf("MakeSlice: len(m) == N (%v vs. %v)", len(m), N)
	}
}

func TestMakeSliceWithMinCap(t *testing.T) {
	var caps = []int{0, 1, 5, 8, 11, 12, 33, 35, 99, 1020, 1021, 1023, 1024}
	var n = 0
	for _, c := range caps {
		var s = nstd.MakeSliceWithMinCap[[]bool](c)
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

func TestZeroMap(t *testing.T) {
	var m = map[int]bool{1: true, 0: false}
	{
		var bm = nstd.ZeroMap(m, 32)
		if len(bm) != 0 {
			t.Fatalf("ZeroMap: len(bm) != 0 (%v)", len(bm))
		}
	}
	{
		var bm = nstd.ZeroMap[map[int]bool](nil, 32)
		if len(bm) != 0 {
			t.Fatalf("ZeroMap: len(bm) != 0 (%v)", len(bm))
		}
	}
}

func TestCollectMapKeys(t *testing.T) {
	const N = 1024
	var m = make(map[int]blank, N)
	for range [N]blank{} {
		m[rand.Intn(N)] = blank{}
	}
	var s = nstd.CollectMapKeys(m)
	if len(s) != len(m) {
		t.Fatalf("CollectMapKeys: len(s) != len(m) (%v vs. %v)", len(s), len(m))
	}
}

func TestAppendMapKeys(t *testing.T) {
	const N = 1024
	var m = make(map[int]blank, N)
	for range [N]blank{} {
		m[rand.Intn(N)] = blank{}
	}
	var s = nstd.AppendMapKeys(nil, m)
	if len(s) != len(m) {
		t.Fatalf("AppendMapKeys: len(s) != len(m) (%v vs. %v)", len(s), len(m))
	}
}
