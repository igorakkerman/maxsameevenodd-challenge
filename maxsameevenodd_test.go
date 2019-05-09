package main

import "testing"

func assert(a []int, expected int, t *testing.T) {
	actual := MaxSame(a)
	if actual != expected {
		t.Errorf("MaxSame wrong. Expected: %d. Got %d", expected, actual)
	}
}

func TestEmpty(t *testing.T) {
	assert([]int{}, 0, t)
}

func TestSingleton(t *testing.T) {
	assert([]int{1}, 1, t)
}

func TestTwoItems(t *testing.T) {
	assert([]int{1, 2}, 2, t)
}

func TestThreeItemsNoneEqual(t *testing.T) {
	assert([]int{1, 2, 3}, 2, t)
}

func TestThreeItemsEvenEqual(t *testing.T) {
	assert([]int{1, 2, 1}, 3, t)
}

func TestFourItemsNoneEqual(t *testing.T) {
	assert([]int{1, 2, 3, 4}, 2, t)
}

func TestFourItemsOddEqual(t *testing.T) {
	assert([]int{1, 2, 3, 4}, 2, t)
}
