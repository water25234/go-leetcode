package main

import "fmt"

// Given a string s, return the number of palindromic substrings in it.

// A string is a palindrome when it reads the same backward as forward.

// A substring is a contiguous sequence of characters within the string.

// Example 1:

// Input: s = "abc"
// Output: 3
// Explanation: Three palindromic strings: "a", "b", "c".
// Example 2:

// Input: s = "aaa"
// Output: 6
// Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

// Constraints:

// 1 <= s.length <= 1000
// s consists of lowercase English letters.

// Senior Staff Two Pointers String Dynamic Programming
func countSubstrings(s string) int {
	count := 0

	for i := range s {
		// expend(s, i, i) -> odd
		// expend(s, i, i+1) -> even
		count += expend(s, i, i) + expend(s, i, i+1)
	}

	return count
}

func expend(s string, left, right int) int {
	result := 0

	for left >= 0 && right < len(s) && s[left] == s[right] {
		result++
		left--
		right++
	}

	return result
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("abba"))
	fmt.Println(countSubstrings("a"))
	fmt.Println(countSubstrings("aa"))
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstrings("aaaa"))
	fmt.Println(countSubstrings("aaaaa"))
}

// 例如 "aba"：

// left = right = 1
//   a b a
//     ^

// 先有中間那個 b，再往兩邊擴，所以是奇數長度。

// 例如 "abba"：

// left = 1, right = 2
// a b b a
//   ^ ^

// 中心在兩個 b 中間，沒有單一中心字元，所以是偶數長度。
