package main

import "fmt"

func fib(n int) int {
	return recursion(n)
}

func fibDP(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp)
	return dp[n]
}

func recursion(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return recursion(n-1) + recursion(n-2)
}

func main() {
	result := fibDP(10)
	fmt.Println("result:", result)
}
