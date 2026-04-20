package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev2 := 1
	prev1 := 2

	for i := 3; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr

		fmt.Println(prev1)
	}

	return prev1
}

func climbStairsDP(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)

	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	fmt.Println(dp)

	return dp[n]
}

func main() {
	result := climbStairsDP(6)
	fmt.Println("result:", result)
}
