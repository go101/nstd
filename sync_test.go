package nstd_test

import (
	"testing"

	"go101.org/nstd"
)

func TestMutexAndWaitGroup(t *testing.T) {
	const N = 1000
	var n = 0
	defer func() {
		if n != N*5 {
			t.Fatalf("n != %d", N*5)
		}
	}()

	var wg nstd.WaitGroup
	defer wg.Wait()

	var m nstd.Mutex
	for range [1000]struct{}{} {
		wg.Go(func() {
			defer m.Lock().Unlock()
			n += 1
		}, func() {
			defer m.Lock().Unlock()
			n += 1
		})
		wg.GoN(3, func() {
			defer m.Lock().Unlock()
			n += 1
		})
	}
}
