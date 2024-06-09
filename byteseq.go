package nstd

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
