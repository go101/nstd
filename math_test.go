package nstd

import (
	"testing"
	"time"
)

func TestScale(t *testing.T) {
	var d = time.Second / 10
	assert(Scale(time.Second, 0.1) == d, "Scale(time.Second, 0.1) != d (%v != %v)", Scale(time.Second, 0.1), d)
	assert(Scale2(time.Second, 1, 10) == d, "Scale(time.Second, 1, 10) != d (%v != %v)", Scale2(time.Second, 1, 10), d)

	var n byte = 200
	var k byte = 200 * 3 / 5
	assert(Scale(n, 0.6) == k, "Scale(n, 0.6) != k (%v != %v)", Scale(n, 0.6), k)
	assert(Scale2(n, 3, 5) == k, "Scale2(n, 3, 5) != k (%v != %v)", Scale2(n, 3, 5), k)
}
