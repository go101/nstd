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
//    var aMutext nstd.Mutex
//
//	func foo() {
//		defer aMutex.Lock().Unlock()
//		... // do something
//	}
type Mutex struct {
	mu sync.Mutex
}

func (m *Mutex) Lock() *Mutex {
	m.mu.Lock()
	return m
}

func (m *Mutex) Unlock() *Mutex {
	m.mu.Unlock()
	return m
}
