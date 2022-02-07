package maximumproductsubarray

func maxProduct(nums []int) int {
	var maximum, minimum, result int

	maximum, minimum, result = nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maximum, minimum = minimum, maximum
		}
		maximum = max(nums[i], maximum*nums[i])
		minimum = min(nums[i], minimum*nums[i])
		result = max(result, maximum)
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
