package nstd

import (
	"strings"
	"unicode/utf8"
	"unsafe"
)

// A ByteSeq is either a string or a byte slice.
type ByteSeq interface{ ~string | ~[]byte }

// ByteSeqCommonPrefixes returns the common prefixes of two [ByteSeq] values.
func ByteSeqCommonPrefix[X, Y ByteSeq](x X, y Y) (x2 X, y2 Y) {
	// Tried several coding mamners and found
	// the current one is inline-able, with cost 80.
	min := minOfTwo(len(x), len(y))
	x2, y2 = x[:min], y[:min]
	if len(x2) == len(y2) { // BCE hint
		for i := 0; i < len(x2); i++ {
			if x2[i] != y2[i] {
				return x2[:i], y2[:i]
			}
		}
	}
	return
}

var (
	_, _ = ByteSeqCommonPrefix("", []byte{})
	_, _ = ByteSeqCommonPrefix("", "")
	_, _ = ByteSeqCommonPrefix([]byte{}, "")
	_, _ = ByteSeqCommonPrefix([]byte{}, []byte{})
)

// ReverseBytes inverts the bytes in a byte slice.
// The argument is returned so that calls to ReverseBytes can be used as expressions.
func ReverseBytes[Bytes ~[]byte](s Bytes) Bytes {
	if len(s) == 0 {
		return s[:0]
	}
	var i, j = len(s) - 1, 0
	for i > j {
		s[i], s[j] = s[j], s[i]
		j++
		i--
	}
	return s
}

// ReverseByteSeq returnes a copy (in the form of byte slice) of
// a byte sequence (in the form of either string or byte slice) but reversed.
func ReverseByteSeq[Seq ByteSeq](s Seq) []byte {
	if len(s) == 0 {
		return []byte(s[:0])
	}
	var into = make([]byte, len(s))
	var j = 0
	for i := len(s) - 1; i >= 0; i-- {
		into[j] = s[i]
		j++
	}
	return into
}

// ReverseRuneSeq returnes a copy (in the form of byte slice) of
// a rune sequence (in the form of either string or byte slice) but reversed.
//
// See:
//
// * https://github.com/golang/go/issues/14777
// * https://github.com/golang/go/issues/68348
func ReverseRuneSeq[Seq ByteSeq](s Seq) []byte {
	if len(s) == 0 {
		return []byte(s[:0])
	}

	var into = make([]byte, len(s))
	var bytes = []byte(s) // doesn't allocate since Go toolchain 1.22
	var i = len(bytes)
	for {
		_, size := utf8.DecodeRune(bytes)
		if size == 0 {
			if i != 0 {
				Panicf("i (%v) != 0", i)
			}
			break
		}
		if i < size {
			Panicf("i (%v) < size (%v)", i, size)
		}
		i -= size
		var j, k = i, 0
		for k < size {
			into[j] = bytes[k]
			j++
			k++
		}
		bytes = bytes[size:]
	}

	return into
}

// MakeDirtyBytes makes a dirty byte slice with the specified length.
// Here, "dirty" means byte values will not be reset to 0.
func MakeDirtyBytes(len int) []byte {
	var b strings.Builder
	b.Grow(len)
	var p = unsafe.StringData(b.String())
	return unsafe.Slice(p, b.Cap())[:len]
}
