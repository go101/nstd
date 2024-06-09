package nstd_test

import (
	"bytes"
	"testing"

	"go101.org/nstd"
)

func TestByteSeqCommonPrefix(t *testing.T) {
	testByteSeqCommonPrefix_4_cases("abcNNN", "abcMMM", 3, t)
	testByteSeqCommonPrefix_4_cases("abc", "ab", 2, t)
	testByteSeqCommonPrefix_4_cases("abc", "abc", 3, t)
	testByteSeqCommonPrefix_4_cases("xyz", "abc", 0, t)
	testByteSeqCommonPrefix_4_cases("", "abc", 0, t)
	testByteSeqCommonPrefix_4_cases("", "", 0, t)
}

func testByteSeqCommonPrefix_4_cases(x, y string, expectedLen int, t *testing.T) {
	testByteSeqCommonPrefix(string(x), string(y), expectedLen, t)
	testByteSeqCommonPrefix(string(x), []byte(y), expectedLen, t)
	testByteSeqCommonPrefix([]byte(x), string(y), expectedLen, t)
	testByteSeqCommonPrefix([]byte(x), []byte(y), expectedLen, t)
}

func hasPrefix[T nstd.ByteSeq](x, x2 T) bool {
	return bytes.HasPrefix([]byte(x), []byte(x2))
}

func testByteSeqCommonPrefix[X, Y nstd.ByteSeq](x X, y Y, expectedLen int, t *testing.T) {
	x2, y2 := nstd.ByteSeqCommonPrefix(x, y)
	if len(x2) != expectedLen {
		t.Fatalf("ByteSeqCommonPrefix(%T(%s), %T(%s)). Expected length (%v) != len(%s)", x, x, y, y, expectedLen, x2)
	}
	if len(y2) != expectedLen {
		t.Fatalf("ByteSeqCommonPrefix(%T(%s), %T(%s)). Expected length (%v) != len(%s)", x, x, y, y, expectedLen, y2)
	}
	if !hasPrefix(x, x2) {
		t.Fatalf("ByteSeqCommonPrefix(%T(%s), %T(%s)). %s is not prefix of %s", x, x, y, y, x2, x)
	}
	if !hasPrefix(y, y2) {
		t.Fatalf("ByteSeqCommonPrefix(%T(%s), %T(%s)). %s is not prefix of %s", x, x, y, y, y2, y)
	}
	if string(x2) != string(y2) {
		t.Fatalf("ByteSeqCommonPrefix(%T(%s), %T(%s)). %s != %s", x, x, y, y, x2, y2)
	}
}
