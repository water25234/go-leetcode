package main

// Mid Level Math Weekly Contest 481
func mirrorDistance(n int) int {
	return abs(n - reverse(n))
}

func reverse(n int) int {
	var res int = 0

	for n > 0 {
		res *= 10
		res += n % 10
		n = n / 10
	}

	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
