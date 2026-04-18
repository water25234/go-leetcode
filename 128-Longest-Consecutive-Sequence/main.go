package longestconsecutivesequence

import (
	"fmt"
	"sort"
)

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Example 1:

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:

// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
// Example 3:

// Input: nums = [1,0,1,2]
// Output: 3

// Array Hash Table Union-Find
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// [100,4,200,1,3,2]
	// [1,2,3,4,100,200]
	sort.Ints(nums)
	fmt.Println(nums)
	best := 1
	current := 1
	n := len(nums)

	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			continue
		}

		if nums[i-1]+1 == nums[i] {
			current++
		} else {
			current = 1
		}

		if current > best {
			best = current
		}
	}

	return best
}

func longestConsecutiveHash(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	best := 0

	for num := range set {
		// 只從起點開始數
		if set[num-1] {
			continue
		}

		// num 是起點，往上數
		length := 1
		curr := num
		for set[curr+1] {
			curr++
			length++
		}

		if length > best {
			best = length
		}
	}

	return best
}
