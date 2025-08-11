package nstd

import (
	"unique"
)

// Error returns an error that formats as the given text.
// Different from errors.New in the std library, two calls
// to Error return an identical error value if the texts are identical.
func Error(text string) error {
	return errorString{unique.Make(text)}
}

// Another approach: type errorString string.
// But that results fat interface values.
type errorString struct {
	s unique.Handle[string]
}

func (e errorString) Error() string {
	return e.s.Value()
}
