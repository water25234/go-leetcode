package jumpgameii

// You are given a 0-indexed array of integers nums of length n. You are initially positioned at index 0.

// Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at index i, you can jump to any index (i + j) where:

// 0 <= j <= nums[i] and
// i + j < n
// Return the minimum number of jumps to reach index n - 1. The test cases are generated such that you can reach index n - 1.

// Example 1:

// Input: nums = [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:

// Input: nums = [2,3,0,1,4]
// Output: 2

// Constraints:

// 1 <= nums.length <= 104
// 0 <= nums[i] <= 1000
// It's guaranteed that you can reach nums[n - 1].

// Array Dynamic Programming Greedy
// 複習
func jump(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}

	end := 0 // 這一跳能到的最遠邊界
	ans := 0
	farthest := 0 // 在這一跳範圍內，下一跳最遠能到哪

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if i == end {
			// 走到這一跳的邊界了，必須再跳一次
			ans++
			end = farthest
		}

	}

	return ans
}
