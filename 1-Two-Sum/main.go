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
