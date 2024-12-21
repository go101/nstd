package nstd

import (
	"testing"
)

func TestSign(t *testing.T) {
	testSign[int](100, 1, t)
	testSign[int](-23, -1, t)
	testSign[int](0, 0, t)
	testSign[int32](55, 1, t)
	testSign[int32](-1, -1, t)
	testSign[int32](0, 0, t)
}

func testSign[S Signed](v S, sign int, t *testing.T) {
	if sn := Sign(v); sn != sign {
		t.Fatalf("Sign(%v) != %v (but %v)", v, sign, sn)
	}
}
