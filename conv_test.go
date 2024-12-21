package nstd

import (
	"testing"
)

func TestBtoi(t *testing.T) {
	if Btoi(true) != 1 {
		t.Fatal("Btoi(true) should be 1")
	}
	if Btoi(false) != 0 {
		t.Fatal("Btoi(true) should be 0")
	}
}
