package nstd

import (
	"unique"
)

// New returns an error that formats as the given text.
// Different from errors.New in the std library, two calls
// to NewError return a distinct error value if the texts are identical.
func NewError(text string) error {
	return errorString{unique.Make(text)}
}

type errorString struct {
	s unique.Handle[string]
}

func (e errorString) Error() string {
	return e.s.Value()
}
