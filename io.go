package nstd

import (
	"io"
)

// WriteStringWithBuffer writes a string into an [io.Writer] with a provided buffer.
// The buffer is used to avoid a string-> []byte conversion (which might allocate).
// This function is like [io.CopyBuffer] but much simpler.
func WriteStringWithBuffer(w io.Writer, s string, buffer []byte) (int, error) {
	if len(buffer) == 0 {
		panic("the buffer is")
	}

	var n = 0
	for len(s) > 0 {
		x := buffer[:copy(buffer, s)]

		k, err := w.Write(x)
		n += k
		if err != nil {
			return n, err
		}
		if k != len(x) {
			return n, io.ErrShortWrite
		}

		s = s[len(x):]
	}

	return n, nil
}
