package nstd

import (
	"errors"
	"unique"
)

// Error returns an error that formats as the given text.
// Different from [errors.New] in the std library, two calls
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

// TrackError reports whether any error in err's tree matches target.
// It behaves almost the same as [errors.Is], except that
// TrackError(anIncomparbleErr, anIncomparbleErr) panics.
//
// See: https://github.com/golang/go/issues/74488
func TrackError(err, target error) bool {
	if err == nil || target == nil {
		return err == target
	}

	return is(err, target)
}

func is(err, target error) bool {
	for {
		if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
			return true
		}
		if err == target {
			return true
		}

		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				return false
			}
		case interface{ Unwrap() []error }:
			for _, err := range x.Unwrap() {
				if is(err, target) {
					return true
				}
			}
			return false
		default:
			return false
		}
	}
}

// TrackErrorOf finds the first error in err's tree that matches ErrorType,
// and if one is found, returns a pointer to a copy of that error.
// Otherwise, it returns nil.
func TrackErrorOf[ErrorType error](err error, _ ...ErrorType) *ErrorType {
	var e ErrorType
	if errors.As(err, &e) {
		return &e
	}

	return nil
}
