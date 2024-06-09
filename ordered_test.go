package nstd_test

import (
	"math"
	"testing"

	"go101.org/nstd"
)

func TestClamp(t *testing.T) {
	testClampOfType[int](t)
	testClampOfType[uint32](t)
	testClampOfType[int64](t)
	testClampOfType[float32](t)
	testClampOfType[float64](t)

	testClamp(math.NaN(), 0, 1, math.NaN(), t)
	testClamp("a", "c", "f", "c", t)
	testClamp("x", "c", "f", "f", t)

	defer func() {
		if recover() == nil {
			t.Fatalf("should panic when min > max")
		}
	}()
	testClamp(math.NaN(), 1, 0, math.NaN(), t)
}

func testClamp[R nstd.Ordered](v, min, max, expected R, t *testing.T) {
	if clamped := nstd.Clamp(v, min, max); clamped != expected {
		if v == v || expected == expected {
			t.Fatalf("Clamp(%v, %v, %v) != %v (but %v)", v, min, max, expected, clamped)
		}
	}
}

func testClampOfType[R nstd.Real](t *testing.T) {
	var min R = 2
	var max R = 8

	var testCases = []struct {
		v, clamped R
	}{
		{1, 2},
		{2, 2},
		{5, 5},
		{8, 8},
		{9, 8},
	}

	for _, tc := range testCases {
		testClamp(tc.v, min, max, tc.clamped, t)
	}

}
