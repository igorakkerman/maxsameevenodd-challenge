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

func Test1(t *testing.T) {
	assert([]int{1}, 1, t)
}

func Test2(t *testing.T) {
	assert([]int{1, 2}, 2, t)
}

func Test3NoneEqual(t *testing.T) {
	assert([]int{1, 2, 3}, 2, t)
}

func Test3EvenEqual(t *testing.T) {
	assert([]int{1, 2, 1}, 3, t)
}

func Test4NoneEqual(t *testing.T) {
	assert([]int{1, 2, 3, 4}, 2, t)
}

func Test4EvenEqual(t *testing.T) {
	assert([]int{1, 2, 1, 4}, 3, t)
}

func Test4OddEqual(t *testing.T) {
	assert([]int{1, 2, 3, 2}, 3, t)
}

func Test4EvenOddEqual(t *testing.T) {
	assert([]int{1, 2, 1, 2}, 4, t)
}

func Test5First3EvenOddEqual(t *testing.T) {
	assert([]int{1, 2, 1, 3, 4}, 3, t)
}

func Test5Last3EvenOddEqual(t *testing.T) {
	assert([]int{1, 2, 3, 4, 3}, 3, t)
}

func Test5NoneEqual(t *testing.T) {
	assert([]int{1, 2, 3, 4, 5}, 2, t)
}

func Test54EvenOddEqual(t *testing.T) {
	assert([]int{3, 2, 1, 2, 1}, 4, t)
}

func Test55EvenOddEqual(t *testing.T) {
	assert([]int{1, 2, 1, 2, 1}, 5, t)
}

func Test7WithNegatives4EvenOddEqual(t *testing.T) {
	assert([]int{3, 2, -1, 2, -1, 3, -1, 3}, 4, t)
}
