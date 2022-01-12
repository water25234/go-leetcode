package maximumsumobtainedofanypermutation

import "sort"

func maxSumRangeQuery(nums []int, requests [][]int) int {
	var pre int
	n := len(nums)

	diff := make([]int, n+1)
	freq := make([]int, n)
	for i := 0; i < len(requests); i++ {
		diff[requests[i][0]] += 1
		diff[requests[i][1]+1] -= 1
	}
	pre = 0

	for k := 0; k < n; k++ {
		pre = pre + diff[k]
		freq[k] = pre
	}

	sort.Ints(freq)
	sort.Ints(nums)

	ans := 0
	for q := 0; q < len(nums); q++ {
		ans = (ans + nums[q]*freq[q]) % 1000000007
	}
	return ans
}
