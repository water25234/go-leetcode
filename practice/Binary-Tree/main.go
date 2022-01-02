package main

import "fmt"

func binarySearch(nums []int, target int) (result int) {

	if len(nums) < target {
		return -1
	}

	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + ((high - low) >> 1)

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func a(nums []int, target int) bool {
	var low, mid, high int

	low, mid, high = 0, 0, (len(nums) - 1)

	for low <= high {
		mid = low + ((high - low) >> 1)

		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	fmt.Println(binarySearch(nums, 8))
}
