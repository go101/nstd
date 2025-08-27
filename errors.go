package nstd

import (
	"reflect"
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

var errorType = reflect.TypeOf((*error)(nil)).Elem()

// TrackError reports whether any error in err's tree matches target.
// It behaves almost the same as [errors.Is], except that
//
// * TrackError(anIncomparbleErr, anIncomparbleErr) panics.
// * TrackError(&aComparableErr, aComparableErr) returns true.
// * It panics if the type of target is a pointer which base type's size is 0.
//
// See: https://github.com/golang/go/issues/74488
func TrackError(err, target error) bool {
	if err == nil || target == nil {
		return err == target
	}

	return trackError(err, target)
}

func trackError(err, target error) bool {
	for {
		if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
			return true
		}
		if err == target {
			tt := reflect.TypeOf(target)
			if tt.Kind() == reflect.Pointer && tt.Elem().Size() == 0 {
				panic("target should not be a pointer pointing to a zero-size value")
			}
			return true
		}
		errValue := reflect.ValueOf(err)
		errType := errValue.Type()
		if errType.Kind() == reflect.Pointer && errType.Elem().Implements(errorType) && !errValue.IsNil() {
			if trackError(errValue.Elem().Interface().(error), target) {
				return true
			}
		}

		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				return false
			}
		case interface{ Unwrap() []error }:
			for _, err := range x.Unwrap() {
				if trackError(err, target) {
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
//
// Different from [errors.As],
//
//   - If *ErrorType is also an error type, then TrackErrorOf
//     doesn't distinguish ErrorType and *ErrorType.
//   - If ErrorType is pointer type and its base type is also an error type,
//     then TrackErrorOf doesn't distinguish ErrorType and its base type.
func TrackErrorOf[ErrorType error](err error, _ ...ErrorType) *ErrorType {
	if err == nil {
		return nil
	}

	var target = new(ErrorType)
	var targetValue = reflect.ValueOf(target)
	if n, value := trackErrorOf(err, target, targetValue); n != 0 {
		if n != 1 {
			targetValue.Elem().Set(value)
		}
		return target
	}

	return nil
}

func trackErrorOf(err error, target any, targetValue reflect.Value) (int, reflect.Value) {
	var Type = targetValue.Type().Elem()
	var info = targetInfo{
		value:        target,
		reflectValue: targetValue,
		reflectType:  Type,
	}
	if trackOf(err, &info) {
		return 1, reflect.Value{}
	}

	if Type.Kind() == reflect.Pointer {
		Type = Type.Elem()
		if Type.Implements(errorType) {
			Value := reflect.New(Type)
			info = targetInfo{
				value:        Value.Interface(),
				reflectValue: Value,
				reflectType:  Type,
			}
			if trackOf(err, &info) {
				return 2, Value
			}
		}
	} else if Type.Kind() != reflect.Interface {
		Type = reflect.PointerTo(Type)
		Value := reflect.New(Type)
		info = targetInfo{
			value:        Value.Interface(),
			reflectValue: Value,
			reflectType:  Type,
		}
		if trackOf(err, &info) {
			return 3, Value.Elem().Elem()
		}
	}

	return 0, reflect.Value{}
}

type targetInfo struct {
	value        any
	reflectValue reflect.Value
	reflectType  reflect.Type
}

func trackOf(err error, info *targetInfo) bool {
	for {
		if x, ok := err.(interface{ As(any) bool }); ok && x.As(info.value) {
			return true
		}
		if reflect.TypeOf(err).AssignableTo(info.reflectType) {
			info.reflectValue.Elem().Set(reflect.ValueOf(err))
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
				if err == nil {
					continue
				}
				if trackOf(err, info) {
					return true
				}
			}
			return false
		default:
			return false
		}
	}
}
