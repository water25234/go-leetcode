package main

import "fmt"

// You are given an integer array nums and an integer k. Find the maximum subarray sum of all the subarrays of nums that meet the following conditions:

// The length of the subarray is k, and
// All the elements of the subarray are distinct.
// Return the maximum subarray sum of all the subarrays that meet the conditions. If no subarray meets the conditions, return 0.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:

// Input: nums = [1,5,4,2,9,9,9], k = 3
// Output: 15
// Explanation: The subarrays of nums with length 3 are:
// - [1,5,4] which meets the requirements and has a sum of 10.
// - [5,4,2] which meets the requirements and has a sum of 11.
// - [4,2,9] which meets the requirements and has a sum of 15.
// - [2,9,9] which does not meet the requirements because the element 9 is repeated.
// - [9,9,9] which does not meet the requirements because the element 9 is repeated.
// We return 15 because it is the maximum subarray sum of all the subarrays that meet the conditions
// Example 2:

// Input: nums = [4,4,4], k = 3
// Output: 0
// Explanation: The subarrays of nums with length 3 are:
// - [4,4,4] which does not meet the requirements because the element 4 is repeated.
// We return 0 because no subarrays meet the conditions.

// Senior Array Hash Table Sliding Window Weekly Contest 318
// [1,5,4,2,9,9,9]
func maximumSubarraySum(nums []int, k int) int64 {
	windowMap := make(map[int]int)
	var windowSum int64
	var windowMax int64

	// [1,5,4]
	for i := 0; i < k; i++ {
		value := nums[i]
		windowMap[value]++
		windowSum += int64(value)
	}

	if len(windowMap) == k {
		windowMax = windowSum
	}

	// [1,5,4,2], index 0 => 1要剔除, index 3 => 2 要加入
	for i := k; i < len(nums); i++ {
		// index 3 => 2 要加入
		preValue := nums[i]
		windowMap[preValue]++
		windowSum += int64(preValue)

		// index 0 => 1 要剔除
		postValue := nums[i-k]
		windowMap[postValue]--
		windowSum -= int64(postValue)
		if windowMap[postValue] <= 0 {
			delete(windowMap, postValue)
		}

		if len(windowMap) == k && windowSum > windowMax {
			windowMax = windowSum
		}
	}

	return windowMax
}

func maximumSubarraySumV1(nums []int, k int) int64 {
	freq := make(map[int]int)
	var sum int64
	var ans int64

	for i := 0; i <= len(nums)-k; i++ {
		sum = 0
		freq = make(map[int]int)
		for j := 0; j < k; j++ {
			value := nums[i+j]
			if _, ok := freq[value]; ok {
				sum = 0
				break
			} else {
				freq[value]++
			}

			sum += int64(value)
		}
		ans = max(ans, sum)
	}

	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	var nums = []int{1, 5, 4, 2, 9, 9, 9}
	var k = 3
	fmt.Println("Maximum subarray sum (V1):", maximumSubarraySumV1(nums, k))
	fmt.Println("Maximum subarray sum:", maximumSubarraySum(nums, k))
}
