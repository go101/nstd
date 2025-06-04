package nstd

import (
	"unique"
)

// Error returns an error that formats as the given text.
// Different from errors.New in the std library, two calls
// to Error return a distinct error value if the texts are identical.
func Error(text string) error {
	return errorString{unique.Make(text)}
}

type errorString struct {
	s unique.Handle[string]
}

func (e errorString) Error() string {
	return e.s.Value()
}
