package twosum

func TwoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, value := range nums {

		balance := target - value

		if j, ok := m[balance]; ok {
			return []int{i, j}
		}

		m[value] = i

	}

	return nil
}

func twoSumTowPointer(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	var left int
	var right int

	left = 0
	right = len(nums) - 1

	for left < right {
		if nums[left]+nums[right] == target {
			return []int{left, right}
		}

		right--
		if left >= right {
			left++
			right = len(nums) - 1
		}
	}

	return nil
}
