package searchinrotatedsortedarray

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
