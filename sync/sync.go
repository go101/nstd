// Package sync is based upon the standard [sync] package,
// to provide some conveniences for Go programming.
package sync

import (
	"sync"
)

// Methods of *Mutex return a *Mutex result, so that
// these methods may be called in a chain.
// The main purpose of this type is to support
// the following use case:
//
//	func foo() {
//		...
//		defer aMutex.Lock().Unlock()
//		...
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
