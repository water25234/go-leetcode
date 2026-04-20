package longestpalindromicsubstring

// Given a string s, return the longest palindromic substring in s.

// Example 1:

// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:

// Input: s = "cbbd"
// Output: "bb"

// Constraints:

// 1 <= s.length <= 1000
// s consist of only digits and English letters.

// Two Pointers String Dynamic Programming
func longestPalindrome(s string) string {
	start := 0
	maxlength := 0
	right := 0
	left := 0

	for i := 0; i < len(s); i++ {
		// odd length
		left, right = expend(s, i, i)
		if (right - left + 1) > maxlength {
			start = left
			maxlength = (right - left + 1)
		}

		// even length
		left, right = expend(s, i, i+1)
		if (right - left + 1) > maxlength {
			start = left
			maxlength = (right - left + 1)
		}
	}

	return s[start : start+maxlength]
}

// 重點邏輯, 從中心點開始往兩邊擴散, 直到不相等或超出邊界為止, 回傳最後的左右邊界
func expend(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return left + 1, right - 1
}
