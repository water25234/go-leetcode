package main

// You are given an integer n.

// Define its mirror distance as: abs(n - reverse(n))​​​​​​​ where reverse(n) is the integer formed by reversing the digits of n.

// Return an integer denoting the mirror distance of n​​​​​​​.

// abs(x) denotes the absolute value of x.

// Example 1:

// Input: n = 25

// Output: 27

// Explanation:

// reverse(25) = 52.
// Thus, the answer is abs(25 - 52) = 27.
// Example 2:

// Input: n = 10

// Output: 9

// Explanation:

// reverse(10) = 01 which is 1.
// Thus, the answer is abs(10 - 1) = 9.
// Example 3:

// Input: n = 7

// Output: 0

// Explanation:

// reverse(7) = 7.
// Thus, the answer is abs(7 - 7) = 0.

// Constraints:

// 1 <= n <= 109

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

func solution(R []string) int {
	n := len(R)
	m := len(R[0])

	// 0:right, 1:down, 2:left, 3:up
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	// seenState[r][c][dir] = 這個狀態是否出現過
	seenState := make([][][]bool, n)
	for i := 0; i < n; i++ {
		seenState[i] = make([][]bool, m)
		for j := 0; j < m; j++ {
			seenState[i][j] = make([]bool, 4)
		}
	}

	// cleaned[r][c] = 這格有沒有清過
	cleaned := make([][]bool, n)
	for i := 0; i < n; i++ {
		cleaned[i] = make([]bool, m)
	}

	r, c, dir := 0, 0, 0
	cleaned[r][c] = true
	ans := 1

	for !seenState[r][c][dir] {
		seenState[r][c][dir] = true

		nr := r + dr[dir]
		nc := c + dc[dir]

		// 前方不能走: 轉向
		if nr < 0 || nr >= n || nc < 0 || nc >= m || R[nr][nc] == 'X' || R[nr][nc] == 'x' {
			dir = (dir + 1) % 4
			continue
		}

		// 前進
		r, c = nr, nc
		if !cleaned[r][c] {
			cleaned[r][c] = true
			ans++
		}
	}

	return ans
}
