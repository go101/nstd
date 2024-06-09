package nstd_test

import (
	"testing"

	"go101.org/nstd"
)

func TestBtoi(t *testing.T) {
	if nstd.Btoi(true) != 1 {
		t.Fatal("Btoi(true) should be 1")
	}
	if nstd.Btoi(false) != 0 {
		t.Fatal("Btoi(true) should be 0")
	}
}
