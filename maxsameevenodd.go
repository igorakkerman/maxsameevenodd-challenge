package main

func MaxSame(a []int) int {
	if len(a) <= 2 {
		return len(a)
	}

	even := a[0]
	odd := a[1]
	last := 2
	max := 2
	for i, value := range a[2:] {
		if i%2 == 0 && value == even || i%2 == 1 && value == odd {
			last++
			if last > max {
				max = last
			}
		} else {
			last = 2
		}
		if i%2 == 0 {
			even = value
		} else {
			odd = value
		}
	}

	return max
}
