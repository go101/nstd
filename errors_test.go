package nstd

import (
	"errors"
	"strings"
	"testing"
)

func TestNewError(t *testing.T) {
	var base = strings.Repeat("abc", 3)
	var e1 = Error(base[0:3])
	var e2 = Error(base[3:6])
	var e3 = Error(base[6:9])
	if e1 != e2 {
		t.Fatal("NewError: e1 != e2")
	}
	if e1 != e3 {
		t.Fatal("NewError: e1 != e3")
	}
	if e3 != e2 {
		t.Fatal("NewError: e3 != e2")
	}
}

type IncomparableError struct {
	x []int
	error
}

func (ie *IncomparableError) Unwrap() error {
	return ie.error
}

type ComparableError struct {
	x [8]int
	error
}

func TestTrackError(t *testing.T) {
	var ie = IncomparableError{
		x:     []int{1, 2},
		error: errors.New("test"),
	}
	func() {
		defer func() {
			if v := recover(); v == nil {
				t.Fatal("TrackError(ie, ie) should panic.")
			}
		}()

		TrackError(ie, ie)
		panic("unreachable")
	}()

	func() {
		defer func() {
			if v := recover(); v == nil {
				t.Fatal("TrackError(ce, ce) should panic.")
			}
		}()

		var ce = ComparableError{error: ie}
		TrackError(ce, ce)
		panic("unreachable")
	}()

	var pie = &ie
	if !TrackError(pie, pie) {
		t.Fatal("TrackError(pie, pie) should return true.")
	}
	var ce = ComparableError{error: pie}
	if !TrackError(ce, ce) {
		t.Fatal("TrackError(ce, ce) should return true.")
	}
}

func TestTrackErrorOf(t *testing.T) {
	var ce = ComparableError{}

	var ie = IncomparableError{
		error: ce,
	}

	if TrackErrorOf[ComparableError](ie) != nil {
		t.Fatal("TrackErrorOf[ComparableError](ie) != nil")
	}
	if TrackErrorOf(ie, ce) != nil {
		t.Fatal("TrackErrorOf(ie, ce) != nil")
	}

	if TrackErrorOf[ComparableError](&ie) == nil {
		t.Fatal("TrackErrorOf[ComparableError](&ie) == nil")
	}
	if TrackErrorOf(&ie, ce) == nil {
		t.Fatal("TrackErrorOf(&ie, ce) == nil")
	}

	if TrackErrorOf[IncomparableError](ce) != nil {
		t.Fatal("TrackErrorOf[IncomparableError](ce) != nil")
	}
	if TrackErrorOf(ce, ie) != nil {
		t.Fatal("TrackErrorOf(ce, ie) != nil")
	}

	if TrackErrorOf[IncomparableError](&ce) != nil {
		t.Fatal("TrackErrorOf[IncomparableError](ce) != nil")
	}
	if TrackErrorOf(&ce, ie) != nil {
		t.Fatal("TrackErrorOf(&ce, ie) != nil")
	}
}
