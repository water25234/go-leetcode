package main

// Complete the function calculateMaximumProfit in the editor below.

// calculateMaximumProfit has the following parameters:

// int cost[n]: the cost of each item
// long int x: the amount of money Walker has

// Returns

// int: the maximum profit modulo (10^9 + 7)

// Description

// There are n items in a shop, where the cost of the i-th item is cost[i]. The profit earned by selling the i-th item is 2^i. Given the array cost and an amount x, determine the maximum profit that can be obtained by selecting a subset of items such that the total cost does not exceed x.

// As the answer may be large, return it modulo (10^9 + 7).

// Note: Assume the indexing is 0-based.

// Example
// n = 5
// cost = [10, 20, 14, 40, 50]
// x = 70

// A valid choice is items with indices 2 and 4:

// total cost = 14 + 50 = 64
// total profit = 2^2 + 2^4 = 4 + 16 = 20

// Other choices may have smaller profit, so the answer is 20.

// Constraints

// 1 <= n <= 10^5
// 1 <= cost[i] <= 10^5
// 0 <= x <= 10^9

// Sample Case 0

// Sample Input For Custom Testing

// 3 3 4 1 8

// Sample Output

// 7

// Explanation

// Walker has enough money to buy all 3 items.
// Profit = 2^0 + 2^1 + 2^2 = 1 + 2 + 4 = 7

// Sample Case 1

// Sample Input For Custom Testing

// 5 10 20 14 40 50 70

// Sample Output

// 20

// Explanation

// The best choice is items 2 and 4, with total cost 64 and total profit 20.
func calculateMaximumProfit(cost []int32, x int64) int32 {
	const mod int64 = 1000000007

	n := len(cost)
	ans := int64(0)

	// 預先算 2^i % mod
	pow2 := make([]int64, n)
	pow2[0] = 1
	for i := 1; i < n; i++ {
		pow2[i] = (pow2[i-1] * 2) % mod
	}

	// 從高 index 往低 index 貪心選
	for i := n - 1; i >= 0; i-- {
		if int64(cost[i]) <= x {
			x -= int64(cost[i])
			ans = (ans + pow2[i]) % mod
		}
	}

	return int32(ans)
}

func main() {
	cost := []int32{10, 20, 14, 40, 50}
	x := int64(70)
	result := calculateMaximumProfit(cost, x)
	println(result) // Output: 20
}
