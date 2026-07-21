package nstd

import (
	"math"
	"runtime"
	"runtime/debug"
	"testing"
)

func TestMutexAndWaitGroup(t *testing.T) {
	const N = 1000
	n := 0
	defer func() {
		if expected := N * 12; n != expected {
			t.Fatalf("n != %d", expected)
		}
	}()

	var wg WaitGroup

	f := func() {
		var m Mutex
		for range [1000]struct{}{} {
			wg.Go(func() {
				defer m.Lock().Unlock()
				n += 1
			}, func() {
				defer m.Lock().Unlock()
				n += 2
			})
			wg.GoN(3, func() {
				defer m.Lock().Unlock()
				n += 1
			})
		}
	}

	f()
	<-wg.WaitChannel()

	defer wg.Wait()
	f()
}

func TestPool(t *testing.T) {
	// disable GC so we can control when it happens.
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	p := Pool[*rune]{}
	if p.Get() != nil {
		t.Fatal("expected empty")
	}

	runes := [...]rune{'a', 'b', 'c'}

	pinner := new(runtime.Pinner)
	pinner.Pin(&p)
	p.Put(&runes[0])
	p.Put(&runes[1])
	if g := p.Get(); g == nil || *g != runes[0] {
		t.Fatalf("got %#v; want %c", g, runes[0])
	}
	if g := p.Get(); g == nil || *g != runes[1] {
		t.Fatalf("got %#v; want %c", g, runes[1])
	}
	if g := p.Get(); g != nil {
		t.Fatalf("got %#v; want nil", g)
	}
	pinner.Unpin()

	// Put in a large number of objects so they spill into
	// stealable space.
	for range 100 {
		p.Put(&runes[2])
	}
	// After one GC, the victim cache should keep them alive.
	runtime.GC()
	if g := p.Get(); g == nil || *g != runes[2] {
		t.Fatalf("got %#v; want %c after GC", g, runes[2])
	}
	// A second GC should drop the victim cache.
	runtime.GC()
	if g := p.Get(); g != nil {
		t.Fatalf("got %#v; want nil after second GC", g)
	}
}

func TestPoolNew(t *testing.T) {
	// disable GC so we can control when it happens.
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	i := 0
	p := Pool[int]{
		New: func() int {
			i++
			return i
		},
	}
	if v := p.Get(); v != 1 {
		t.Fatalf("got %v; want 1", v)
	}
	if v := p.Get(); v != 2 {
		t.Fatalf("got %v; want 2", v)
	}

	pinner := new(runtime.Pinner)
	pinner.Pin(&p)
	p.Put(42)
	if v := p.Get(); v != 42 {
		t.Fatalf("got %v; want 42", v)
	}
	pinner.Unpin()

	if v := p.Get(); v != 3 {
		t.Fatalf("got %v; want 3", v)
	}
}

func TestPoolReset(t *testing.T) {
	// disable GC so we can control when it happens.
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	to := math.MinInt
	p := Pool[*int]{
		Reset: func(v *int) *int {
			*v = to
			return v
		},
	}

	pinner := new(runtime.Pinner)
	pinner.Pin(&p)
	x := 42
	p.Put(&x)
	if v := p.Get(); v == nil || *v != to {
		t.Fatalf("got %v; want %d", v, to)
	}
	pinner.Unpin()
}
