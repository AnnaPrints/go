package main

import (
	"testing"
)

func TestDivideBasic(t *testing.T) {

	for i := 10000; i < 100000; i++ {
		result := i / i
		if result != 0 {
			t.Error("Unexpected result", result)
		}
	}
}

func Test_Goland_WithVeryLongName(t *testing.T) {
}
