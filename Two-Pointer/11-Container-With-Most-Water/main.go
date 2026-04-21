package main

import "fmt"

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
// Example 2:

// Input: height = [1,1]
// Output: 1

// Constraints:

// n == height.length
// 2 <= n <= 105
// 0 <= height[i] <= 104

func maxArea(height []int) int {
	// max(area) = width * hight
	// max(area) = (right - left) * min(hight[left], hight[right]) 取min是因為水會從較矮的那一邊溢出，所以以較矮的那一邊為基準

	ans := 0

	left := 0
	right := len(height) - 1

	hight := 0
	width := 0

	for left < right {
		hight = min(height[left], height[right])
		width = (right - left)

		area := width * hight

		ans = max(ans, area)

		// 此效能太差會time limit exceeded
		// right--
		// if left >= right {
		//     left++
		//     right = len(height) - 1
		// }

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{1, 2}))
	fmt.Println(maxArea([]int{2, 1}))
	fmt.Println(maxArea([]int{1, 2, 1}))
	fmt.Println(maxArea([]int{1, 2, 3}))
	fmt.Println(maxArea([]int{3, 2, 1}))
}
