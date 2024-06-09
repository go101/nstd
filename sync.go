package nstd

import (
	"sync"
)

// Methods of *Mutex return a *Mutex result, so that
// these methods may be called in a chain.
// It is just a simpler wrapper of the [sync.Mutex].
// The main purpose of this type is to support
// the following use case:
//
//	var aMutex nstd.Mutex
//
//	func foo() {
//		defer aMutex.Lock().Unlock()
//		... // do something
//	}
type Mutex struct {
	mu sync.Mutex
}

// Lock return m, so that the methods of m can be called in chain.
func (m *Mutex) Lock() *Mutex {
	m.mu.Lock()
	return m
}

// Unlock return m, so that the methods of m can be called in chain.
func (m *Mutex) Unlock() *Mutex {
	m.mu.Unlock()
	return m
}

// Do guards the execution of a function in Lock() and Unlock()
//
// See:
// * https://github.com/golang/go/issues/63941
func (m *Mutex) Do(f func()) {
	defer m.Lock().Unlock()
	f()
}

// WaitGroup extends sync.WaitGroup.
type WaitGroup struct {
	sync.WaitGroup
}

// GoN starts several concurrent tasks waited by wg.
//
// See:
//     https://github.com/golang/go/issues/18022
func (wg *WaitGroup) Go(fs ...func()) {
	for i, f := range fs {
		if f == nil {
			Panicf("fs[%d] is nil", i)
		}
	}
	wg.Add(len(fs))
	for _, f := range fs {
		f := f
		go func() {
			defer wg.Done()
			f()
		}()
	}
}

// GoN starts a concurrent task n times. These task gorutines are waited by wg,
func (wg *WaitGroup) GoN(n int, f func()) {
	if f == nil {
		panic("f is nil")
	}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			f()
		}()
	}
}
