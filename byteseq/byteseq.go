/*
Package byteseq provides some operations capable of
handling both strings and byte slices.
*/
package byteseq

import (
	"go101.org/nstd/ordered"
)

// A ByteSeq is either a string or a byte slice.
type ByteSeq interface{ ~string | ~[]byte }

// CommonPrefixes returns the common prefixes of two [ByteSeq] values.
func CommonPrefixes[X, Y ByteSeq](x X, y Y) (X, Y) {
	min := ordered.Min(len(x), len(y))
	x2, y2 := x[:min], y[:min]
	if len(x2) != len(y2) { // BCE hint
		panic("len(x2) != len(y2)")
	}
	for i := 0; i < len(x2); i++ {
		if x2[i] != y2[i] {
			return x2[:i], y2[:i]
		}
	}
	return x2, y2
}

var (
	_, _ = CommonPrefixes("", []byte{})
	_, _ = CommonPrefixes("", "")
	_, _ = CommonPrefixes([]byte{}, "")
	_, _ = CommonPrefixes([]byte{}, []byte{})
)

// Compare compares two [ByteSeq] values.
// It returns a negative value for string(x) < string(y);
// It returns 0 for string(x) == string(y);
// It returns a positive value for string(x) > string(y);
func Compare[X, Y ByteSeq](x X, y Y) int {
	min := ordered.Min(len(x), len(y))
	x2, y2 := x[:min], y[:min]
	if len(x2) != len(y2) { // BCE hint
		panic("len(x2) != len(y2)")
	}
	for i := 0; i < len(x2); i++ {
		if v := int(x2[i]) - int(y2[i]); v != 0 {
			return v
		}
	}
	
	return len(x) - len(y)
}

var (
	_ = Compare("", []byte{})
	_ = Compare("", "")
	_ = Compare([]byte{}, "")
	_ = Compare([]byte{}, []byte{})
)