package main

func MaxSame(a []int) int {
	if len(a) <= 2 {
		return len(a)
	}

	even := a[0]
	odd := a[1]
	last := 2
	max := 2
	for i := 2; i < len(a); i++ {
		if i%2 == 0 && a[i] == even || i%2 == 1 && a[i] == odd {
			last++
			if last > max {
				max = last
			}
		} else {
			last = 2
		}
		if i%2 == 0 {
			even = a[i]
		} else {
			odd = a[i]
		}
	}

	return max
}
