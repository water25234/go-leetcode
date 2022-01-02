package searchinrotatedsortedarrayii

import "fmt"

func search_map(nums []int, target int) bool {
	var nums_map map[int]int

	nums_map = make(map[int]int)

	for _, num := range nums {
		if _, ok := nums_map[num]; !ok {
			nums_map[num] = num
		}

		if _, ok := nums_map[target]; ok {
			return true
		}
	}

	return false
}

func search_rotate_index(nums []int, target int) bool {

	rotate_index := len(nums) >> 1

	nums_list := nums[rotate_index:]

	for i := 0; i < rotate_index; i++ {
		nums_list = append(nums_list, nums[i])
	}

	nums = nil

	fmt.Println(nums_list)

	result := binary_search(nums_list, target)

	return result
}

func binary_search(nums []int, target int) bool {
	if nums[0] == target || nums[len(nums)-1] == target {
		return true
	}

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

func search_half(nums []int, target int) bool {

	if len(nums) == 0 {
		return false
	}

	if nums[0] == target || nums[len(nums)-1] == target {
		return true
	}

	var low, mid, high int

	low, mid, high = 0, 0, (len(nums) - 1)

	for low <= high {
		mid = low + ((high - low) >> 1)

		if nums[mid] == target {
			return true
		} else if nums[mid] > nums[low] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}

		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}

	return false
}

func search(nums []int, target int) bool {
	first, end := nums[0], nums[len(nums)-1]
	if first == target || end == target {
		return true
	}

	left, right := 0, len(nums)-1
	for left <= right {
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}

		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] >= nums[0] { // we are in the first half
			if nums[mid] > target && target >= first {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else { // we are in the 2nd half
			if nums[mid] < target && target <= end {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
