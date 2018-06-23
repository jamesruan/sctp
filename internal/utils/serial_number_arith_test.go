package utils

import (
	"testing"
)

func TestSerial32_Add(t *testing.T) {
	a := MakeSerial32(1<<32 - 1)
	b := a.Add(1)
	if !b.Eq(MakeSerial32(0)) {
		t.Fail()
	}
	if a.GreaterThan(b) {
		t.Fail()
	} else if !a.LessThan(b) {
		t.Fail()
	}
}
