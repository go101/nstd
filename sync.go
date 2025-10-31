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
// See: https://github.com/golang/go/issues/63941
func (m *Mutex) Do(f func()) {
	defer m.Lock().Unlock()
	f()
}

// WaitGroup extends sync.WaitGroup.
// Each WaitGroup maintains an internal count which initial value is zero.
type WaitGroup struct {
	wg sync.WaitGroup
}

// GoN starts several concurrent tasks and increases the internal count by len(fs).
// The internal count will be descreased by one when each of the task is done.
// Note: if a call to any function in fs panics, then the whole program crashes.
//
// See: https://github.com/golang/go/issues/18022 and https://github.com/golang/go/issues/76126
func (wg *WaitGroup) Go(fs ...func()) {
	for i, f := range fs {
		if f == nil {
			Panicf("fs[%d] is nil", i)
		}
	}
	wg.wg.Add(len(fs))
	for _, f := range fs {
		f := f
		go func() {
			defer func() {
				if x := recover(); x != nil {
					panic(x)
				}

				wg.wg.Done()
			}()

			f()
		}()
	}
}

// GoN starts a task n times concurrently and increases the internal count by n.
// The internal count will be descreased by one when each of the task instances is done.
// Note: if the f call panics, then the whole program crashes.
//
// See: https://github.com/golang/go/issues/18022 and https://github.com/golang/go/issues/76126
func (wg *WaitGroup) GoN(n int, f func()) {
	if n < 0 {
		panic("the count must not be negative")
	}
	if f == nil {
		panic("f is nil")
	}
	wg.wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer func() {
				if x := recover(); x != nil {
					panic(x)
				}

				wg.wg.Done()
			}()
			f()
		}()
	}
}

// Wait blocks until the internal counter is zero.
func (wg *WaitGroup) Wait() {
	wg.wg.Wait()
}

// WaitChannel returns a channel which reads will block until the internal counter is zero.
func (wg *WaitGroup) WaitChannel() <-chan struct{} {
	var c = make(chan struct{})

	go func() {
		wg.wg.Wait()
		close(c)
	}()

	return c
}
