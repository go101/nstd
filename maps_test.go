package nstd

import (
	"math/rand"
	"testing"
)

func TestZeroMap(t *testing.T) {
	var m = map[int]bool{1: true, 0: false}
	{
		var bm = ZeroMap(m, 32)
		if len(bm) != 0 {
			t.Fatalf("ZeroMap: len(bm) != 0 (%v)", len(bm))
		}
	}
	{
		var bm = ZeroMap[map[int]bool](nil, 32)
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
	var s = CollectMapKeys(m)
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
	var s = AppendMapKeys(nil, m)
	if len(s) != len(m) {
		t.Fatalf("AppendMapKeys: len(s) != len(m) (%v vs. %v)", len(s), len(m))
	}
}

func TestBoolKeyMap(t *testing.T) {
	var m BoolKeyMap[bool, int]
	m.Put(true, 123)
	m.Put(false, 789)
	var x = m.Get(true)
	assert(x == 123, "BoolKeyMap: Get(true) != 123 (%d)", x)
	var y = m.Get(false)
	assert(y == 789, "BoolKeyMap: Get(false) != 789 (%d)", y)
	m.Put(true, 321)
	m.Put(false, 987)
	x = m.Get(true)
	assert(x == 321, "BoolKeyMap: Get(true) != 321 (%d)", x)
	y = m.Get(false)
	assert(y == 987, "BoolKeyMap: Get(false) != 987 (%d)", y)
}

func TestBoolElementMap(t *testing.T) {
	var m BoolElementMap[int, bool]
	m.Put(1, false)
	m.Put(2, false)
	m.Put(3, false)
	assert(m.m == nil, "BoolElementMap: m.m == nil")
	m.Put(1, true)
	assert(len(m.m) == 1, "BoolElementMap: len(m.m) == 1 (%d)", len(m.m))
	m.Put(1, true)
	assert(len(m.m) == 1, "BoolElementMap: len(m.m) == 1 (%d)", len(m.m))
	m.Put(2, true)
	assert(len(m.m) == 2, "BoolElementMap: len(m.m) == 2 (%d)", len(m.m))
	var x = m.Get(1)
	var y = m.Get(2)
	var z = m.Get(3)
	assert(x, "BoolElementMap: x is false")
	assert(y, "BoolElementMap: y is false")
	assert(!z, "BoolElementMap: z is true")
	m.Put(1, false)
	assert(len(m.m) == 1, "BoolElementMap: len(m.m) == 1 (%d)", len(m.m))
	x = m.Get(1)
	y = m.Get(2)
	z = m.Get(3)
	assert(!x, "BoolElementMap: x is true")
	assert(y, "BoolElementMap: y is false")
	assert(!z, "BoolElementMap: z is true")
}
