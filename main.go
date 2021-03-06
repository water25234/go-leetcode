package main

import "fmt"

func search(nums []int, target int) int {
	var low, mid, high int

	low, mid, high = 0, 0, len(nums)-1

	for low <= high {

		if nums[low] == target {
			return low
		}

		if nums[high] == target {
			return high
		}

		mid = low + ((high - low) >> 1)

		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
			if nums[low] <= target && nums[mid] >= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[high] > nums[mid] {
			if nums[high] >= target && nums[mid] < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func main() {
	var nums []int
	var target int

	nums = []int{4, 1, 7, 3, 6, 8, 5, 2, 0, 9}
	target = 0

	fmt.Println(search(nums, target))
}
