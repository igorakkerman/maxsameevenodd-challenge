package main

import "testing"

func TestEmpty_0(t *testing.T) {
	a := []int{}

	maxSame := MaxSame(a)

	if maxSame != 0 {
		t.Errorf("MaxSame wrong. Expected: %d. Got %d", 0, maxSame)
	}
}

func TestSingleton_1(t *testing.T) {
	a := []int{1}

	maxSame := MaxSame(a)

	if maxSame != 1 {
		t.Errorf("MaxSame wrong. Expected: %d. Got %d", 1, maxSame)
	}
}
