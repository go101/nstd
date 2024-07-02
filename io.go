package nstd

import (
	"fmt"
	"io"
)

// CheckWriteResult checks whether or not a Write method is badly implemented.
//
// See:
//   - https://github.com/golang/go/issues/67921
//   - https://github.com/golang/go/issues/9096
func WriteWithCheck(w io.Writer, p []byte) (n int, err error) {
	n, err = w.Write(p)
	if n < 0 {
		n = 0
		if err != nil {
			err = fmt.Errorf("errBadWrite: n (%d) < 0 with additional error: %w", n, err)
		} else {
			err = fmt.Errorf("errBadWrite: n (%d) < 0", n)
		}
	} else if n > len(p) {
		n = len(p)
		if err != nil {
			err = fmt.Errorf("errBadWrite: n (%d) > len (%d) with additional error: %w", n, len(p), err)
		} else {
			err = fmt.Errorf("errBadWrite: n (%d) > len (%d)", n, len(p))
		}
	} else if n < len(p) && err == nil {
		err = fmt.Errorf("errBadWrite: n (%d) < len (%d) but no errors", n, len(p))
	}
	return
}

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

		_, err := WriteWithCheck(w, x)
		n += len(x)
		if err != nil {
			return n, err
		}

		s = s[len(x):]
	}

	return n, nil
}
