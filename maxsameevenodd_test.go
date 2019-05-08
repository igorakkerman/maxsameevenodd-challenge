package main

import "testing"

func TestEmpty_0(t *testing.T) {
	a := []int{}

	maxSame := MaxSame(a)

	if maxSame != 0 {
		t.Errorf("MaxSame wrong. Expected: %d. Got %d", 0, maxSame)
	}
}
