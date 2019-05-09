package main

import "testing"

func TestEmpty(t *testing.T) {
	a := []int{}

	maxSame := MaxSame(a)

	if maxSame != 0 {
		t.Errorf("MaxSame wrong. Expected: %d. Got %d", 0, maxSame)
	}
}

func TestSingleton(t *testing.T) {
	a := []int{1}

	maxSame := MaxSame(a)

	if maxSame != 1 {
		t.Errorf("MaxSame wrong. Expected: %d. Got %d", 1, maxSame)
	}
}

func TestTwoItems(t *testing.T) {
	a := []int{1, 2}

	maxSame := MaxSame(a)

	if maxSame != 2 {
		t.Errorf("MaxSame wrong. Expected: %d. Got %d", 2, maxSame)
	}
}

func TestThreeItemsEvenEqual(t *testing.T) {
	a := []int{1, 2, 1}

	maxSame := MaxSame(a)

	if maxSame != 3 {
		t.Errorf("MaxSame wrong. Expected: %d. Got %d", 3, maxSame)
	}
}

func TestThreeItemsNoneEqual(t *testing.T) {
	a := []int{1, 2, 3}

	maxSame := MaxSame(a)

	if maxSame != 2 {
		t.Errorf("MaxSame wrong. Expected: %d. Got %d", 2, maxSame)
	}
}
