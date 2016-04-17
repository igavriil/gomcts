package main

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}
