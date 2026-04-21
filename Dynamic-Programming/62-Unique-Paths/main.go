package main

import "fmt"

// There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

// Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

// The test cases are generated so that the answer will be less than or equal to 2 * 109.

// Input: m = 3, n = 7
// Output: 28
// Example 2:

// Input: m = 3, n = 2
// Output: 3
// Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
// 1. Right -> Down -> Down
// 2. Down -> Down -> Right
// 3. Down -> Right -> Down

func uniquePaths(m int, n int) int {
	dp := make([][]int, m) // row
	for i := range dp {
		dp[i] = make([]int, n) // column
	}

	for i := 0; i < m; i++ { // row
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ { // column
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	fmt.Println(dp)

	return dp[m-1][n-1]
}

func main() {
	m := 3
	n := 7

	result := uniquePaths(m, n)
	fmt.Println("Number of unique paths:", result)
}

// m => r, n => c
//     n
// m 0 -> 1
//   1   0
// [m][n] => 左上(起點), [m-1][n-1] 右下(綜點), [m][m] = [m-1][n](往下走) + [m][-1](往右走)

//     n
// m A B C D
//   E F G H
//   I J K L

// F 為起點, 右, 下, 左, 上, 順時針走
// m    n
// 0    1  右
// 1    0  下
// 0   -1  左
// -1   0  上
