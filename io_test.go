package nstd_test

import (
	"bytes"
	"testing"

	"go101.org/nstd"
)

func TestWriteStringWithBuffer(t *testing.T) {
	testWriteStringWithBuffer("", t)
	testWriteStringWithBuffer("a", t)
	testWriteStringWithBuffer("  abcdefghijklmnopqrstuvwxyz    -", t)
	testWriteStringWithBuffer(string([]byte{100000: 'x'}), t)
}

func testWriteStringWithBuffer(s string, t *testing.T) {
	for _, size := range []int{1, 64, 100, 789, 1024} {
		buf := make([]byte, size)
		var w bytes.Buffer
		n, err := nstd.WriteStringWithBuffer(&w, s, buf)
		if err != nil {
			continue
		}
		if n != len(s) {
			t.Fatalf("WriteStringWithBuffer: n (%d) != len(s) (bufer size: %d, s: %s)", n, buf, s)
		}
		if r := w.String(); r != s {
			t.Fatalf("WriteStringWithBuffer: bad write: %s (bufer size: %d, s: %s)", r, buf, s)
		}
	}
}
