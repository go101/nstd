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

func TestReverseBytes(t *testing.T) {
	testReverseBytes("", "", t)
	testReverseBytes("a", "a", t)
	testReverseBytes("你好！", []byte{0x81, 0xbc, 0xef, 0xbd, 0xa5, 0xe5, 0xa0, 0xbd, 0xe4}, t)
	testReverseBytes("abc", "cba", t)
	testReverseBytes("abc你好！", []byte{0x81, 0xbc, 0xef, 0xbd, 0xa5, 0xe5, 0xa0, 0xbd, 0xe4, 'c', 'b', 'a'}, t)
	testReverseBytes("你好！abc", []byte{'c', 'b', 'a', 0x81, 0xbc, 0xef, 0xbd, 0xa5, 0xe5, 0xa0, 0xbd, 0xe4}, t)
}

func testReverseBytes[SeqX, SeqY nstd.ByteSeq](sx SeqX, sy SeqY, t *testing.T) {
	x, y := []byte(sx), []byte(sy)
	if string(nstd.ReverseBytes(x)) != string(sy) {
		t.Fatalf("ReverseBytes(%s) != %s", x, sy)
	}
	if string(nstd.ReverseBytes(y)) != string(sx) {
		t.Fatalf("ReverseBytes(%s) != %s", y, sx)
	}
	if str := string(y); string(nstd.ReverseBytes(nstd.ReverseBytes(y))) != str {
		t.Fatalf("ReverseBytes(nstd.ReverseBytes((%s)) != %s", y, str)
	}
	if str := string(x); string(nstd.ReverseBytes(nstd.ReverseBytes(x))) != str {
		t.Fatalf("ReverseBytes(nstd.ReverseBytes((%s)) != %s", x, str)
	}
}

func TestReverseByteSeq(t *testing.T) {
	testReverseByteSeq("", "", t)
	testReverseByteSeq("a", "a", t)
	testReverseByteSeq("你好！", []byte{0x81, 0xbc, 0xef, 0xbd, 0xa5, 0xe5, 0xa0, 0xbd, 0xe4}, t)
	testReverseByteSeq("abc", "cba", t)
	testReverseByteSeq("abc你好！", []byte{0x81, 0xbc, 0xef, 0xbd, 0xa5, 0xe5, 0xa0, 0xbd, 0xe4, 'c', 'b', 'a'}, t)
	testReverseByteSeq("你好！abc", []byte{'c', 'b', 'a', 0x81, 0xbc, 0xef, 0xbd, 0xa5, 0xe5, 0xa0, 0xbd, 0xe4}, t)
}

func testReverseByteSeq[SeqX, SeqY nstd.ByteSeq](sx SeqX, sy SeqY, t *testing.T) {
	x, y := []byte(sx), []byte(sy)
	if string(nstd.ReverseByteSeq(x)) != string(sy) {
		t.Fatalf("ReverseByteSeq(%s) != %s", x, sy)
	}
	if string(nstd.ReverseByteSeq(y)) != string(sx) {
		t.Fatalf("ReverseByteSeq(%s) != %s", y, sx)
	}
	if str := string(y); string(nstd.ReverseByteSeq(nstd.ReverseByteSeq(y))) != str {
		t.Fatalf("ReverseByteSeq(nstd.ReverseByteSeq((%s)) != %s", y, str)
	}
	if str := string(x); string(nstd.ReverseByteSeq(nstd.ReverseByteSeq(x))) != str {
		t.Fatalf("ReverseByteSeq(nstd.ReverseByteSeq((%s)) != %s", x, str)
	}
}

func TestReverseRuneSeq(t *testing.T) {
	testReverseRuneSeq("", "", t)
	testReverseRuneSeq("a", "a", t)
	testReverseRuneSeq("你好！", "！好你", t)
	testReverseRuneSeq("abc", "cba", t)
	testReverseRuneSeq("abc你好！", "！好你cba", t)
	testReverseRuneSeq("你好！abc", "cba！好你", t)
}

func testReverseRuneSeq[SeqX, SeqY nstd.ByteSeq](sx SeqX, sy SeqY, t *testing.T) {
	x, y := []byte(sx), []byte(sy)
	if string(nstd.ReverseRuneSeq(x)) != string(sy) {
		t.Fatalf("ReverseRuneSeq(%s) != %s", x, sy)
	}
	if string(nstd.ReverseRuneSeq(y)) != string(sx) {
		t.Fatalf("ReverseRuneSeq(%s) != %s", y, sx)
	}
	if str := string(y); string(nstd.ReverseRuneSeq(nstd.ReverseRuneSeq(y))) != str {
		t.Fatalf("ReverseRuneSeq(nstd.ReverseRuneSeq((%s)) != %s", y, str)
	}
	if str := string(x); string(nstd.ReverseRuneSeq(nstd.ReverseRuneSeq(x))) != str {
		t.Fatalf("ReverseRuneSeq(nstd.ReverseRuneSeq((%s)) != %s", x, str)
	}
}
