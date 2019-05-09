package main

func MaxSame(a []int) int {
	if len(a) == 0 {
		return 0
	}

	if len(a) == 1 {
		return 1
	}

	if len(a) == 2 {
		return 2
	}

	if a[2] == a[0] {
		return 3
	} else {
		return 2
	}
}
