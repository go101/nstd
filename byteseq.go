package nstd

// A ByteSeq is either a string or a byte slice.
type ByteSeq interface{ ~string | ~[]byte }

// CommonPrefixes returns the common prefixes of two [ByteSeq] values.
func ByteSeqCommonPrefixes[X, Y ByteSeq](x X, y Y) (x2 X, y2 Y) {
	// Tried several coding mamners and found
	// the current one is inline-able, with cost 80.
	min := Min(len(x), len(y))
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
	_, _ = ByteSeqCommonPrefixes("", []byte{})
	_, _ = ByteSeqCommonPrefixes("", "")
	_, _ = ByteSeqCommonPrefixes([]byte{}, "")
	_, _ = ByteSeqCommonPrefixes([]byte{}, []byte{})
)

// Compare compares two [ByteSeq] values.
// It returns a negative value for string(x) < string(y);
// It returns 0 for string(x) == string(y);
// It returns a positive value for string(x) > string(y);
func ByteSeqCompare[X, Y ByteSeq](x X, y Y) int {
	min := Min(len(x), len(y))
	x2, y2 := x[:min], y[:min]
	if len(x2) == len(y2) { // BCE hint
		// Here, the possible string->[]byte conversion
		// will not allocate. And the inline cost of a
		// for-range loop is lower than a for;; loop,
		// so that the instanitations of the function
		// are inline-able.
		for i := range []byte(x2) {
			if v := int(x2[i]) - int(y2[i]); v != 0 {
				return v
			}
		}
	}

	return len(x) - len(y)
}

var (
	_ = ByteSeqCompare("", []byte{})
	_ = ByteSeqCompare("", "")
	_ = ByteSeqCompare([]byte{}, "")
	_ = ByteSeqCompare([]byte{}, []byte{})
)
