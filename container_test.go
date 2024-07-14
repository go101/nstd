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
