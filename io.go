package nstd

import (
	"io"
)

// WriteStringWithBuffer writes a string into an [io.Writer].
// The buffer is used to avoid a string-> []byte conversion (which might allocate).
func WriteStringWithBuffer(w io.Writer, s string, buffer []byte) (int, error) {
	if len(buffer) == 0 {
		panic("the buffer is")
	}

	var n = 0
	for len(s) > 0 {
		x := copy(buffer, s)
		y, err := w.Write(buffer[:x])
		n += y
		if err != nil {
			return n, err
		}
		s = s[y:]
	}

	return n, nil
}
