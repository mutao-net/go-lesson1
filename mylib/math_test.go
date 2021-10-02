package mylib

import "testing"

var Debug bool = false

func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("Debug skipping")
	}
	v := Avereage([]int{1, 2, 3, 4, 5, 6})
	if v != 3 {
		t.Error("expected 3, got ", v)
	}
}
