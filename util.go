package main

func Min(a ...int) int {
	min := int(^uint(0) >> 1) // largest possible int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func Max(a ...int) int {
	max := -(int(^uint(0)>>1) - 1) // smallest int
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func Percent(percent, of int) int {
	return of * percent / 100
}

// Used to debug area gen
func RuneFromInt(i int) rune {
	if i == 0 {
		return '0'
	}
	if i == 1 {
		return '1'
	}
	if i == 2 {
		return '2'
	}
	if i == 3 {
		return '3'
	}
	if i == 4 {
		return '4'
	}
	if i == 5 {
		return '5'
	}
	if i == 6 {
		return '6'
	}
	if i == 7 {
		return '7'
	}
	return 'X'
}
