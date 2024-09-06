package nstd_test

import (
	"testing"

	"go101.org/nstd"
)

func TestMutexAndWaitGroup(t *testing.T) {
	const N = 1000
	var n = 0
	defer func() {
		if expected := N * 12; n != expected {
			t.Fatalf("n != %d", expected)
		}
	}()

	var wg nstd.WaitGroup

	var f = func() {
		var m nstd.Mutex
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
