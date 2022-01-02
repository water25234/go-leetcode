package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	var ans, left, right, strLength int
	var dicts []bool

	strLength = len(s)
	left, right = 0, 0
	dicts = make([]bool, 256)

	for right < strLength {

		if dicts[s[right]] == false {
			dicts[s[right]] = true
			right++
			if ans < (right - left) {
				ans = (right - left)
			}
		} else {
			dicts[s[left]] = false
			left++
		}
	}
	return ans
}
