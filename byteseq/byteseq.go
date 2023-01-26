/*
Package byteseq provides some operations capable of
handling both strings and byte slices.
*/
package byteseq

// A ByteSeq is either a string or a byte slice.
type ByteSeq interface{ ~string | ~[]byte }

// CommonPrefixLen returns the length of the common prefix
// of two [ByteSeq] values.
func CommonPrefixLen[X, Y ByteSeq](x X, y Y) int {
	if len(x) < len(y) {
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] {
				return i
			}
		}
		return len(x)
	}

	for i := 0; i < len(y); i++ {
		if x[i] != y[i] {
			return i
		}
	}
	return len(y)
}
