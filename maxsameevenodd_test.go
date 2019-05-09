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

func TestFourItemsEvenEqual(t *testing.T) {
	assert([]int{1, 2, 1, 4}, 3, t)
}

func TestFourItemsOddEqual(t *testing.T) {
	assert([]int{1, 2, 3, 2}, 3, t)
}

func TestFourItemsEvenOddEqual(t *testing.T) {
	assert([]int{1, 2, 1, 2}, 4, t)
}

func TestFiveItemsFirst3EvenOddEqual(t *testing.T) {
	assert([]int{1, 2, 1, 3, 4}, 3, t)
}

func TestFiveItemsLast3EvenOddEqual(t *testing.T) {
	assert([]int{1, 2, 3, 4, 3}, 3, t)
}

func TestFiveItemsNoneEqual(t *testing.T) {
	assert([]int{1, 2, 3, 4, 5}, 2, t)
}

func TestFiveItems4EvenOddEqual(t *testing.T) {
	assert([]int{3, 2, 1, 2, 1}, 4, t)
}

func TestFiveItems5EvenOddEqual(t *testing.T) {
	assert([]int{1, 2, 1, 2, 1}, 5, t)
}
