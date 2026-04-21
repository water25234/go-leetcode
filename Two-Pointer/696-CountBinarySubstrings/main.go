package main

import "fmt"

// Given a binary string s, return the number of non-empty substrings that have the same number of 0's and 1's, and all the 0's and all the 1's in these substrings are grouped consecutively.

// Substrings that occur multiple times are counted the number of times they occur.

// Example 1:

// Input: s = "00110011"
// Output: 6
// Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
// Notice that some of these substrings repeat and are counted the number of times they occur.
// Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.
// Example 2:

// Input: s = "10101"
// Output: 4
// Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.

// Constraints:

// 1 <= s.length <= 105
// s[i] is either '0' or '1'.

// Topic : Staff Two Pointers String
func getSubstringCount(s string) int {
	if len(s) == 0 {
		return 0
	}

	prev := 0 // 前一組連續字元的長度
	curr := 1 // 目前這一組連續字元的長度
	ans := 0

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr++
		} else {
			//fmt.Printf("prev: %d, curr: %d\n", prev, curr)
			ans += min(prev, curr)
			// fmt.Printf("ans: %d\n", ans)
			prev = curr
			curr = 1
		}
	}

	// 最後一組也要結算
	ans += min(prev, curr)

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countBinarySubstringsV2(s string) int {
	ans := 0
	n := len(s)

	// 掃每一個相鄰位置，找 0|1 或 1|0 的分界
	for i := 0; i < n-1; i++ {

		if s[i] != s[i+1] {
			left := i
			right := i + 1

			// 從分界往左右擴
			// 左邊必須一直跟 s[i] 一樣
			// 右邊必須一直跟 s[i+1] 一樣
			for left >= 0 &&
				right < n &&
				s[left] == s[i] &&
				s[right] == s[i+1] {
				ans++
				left--
				right++
			}
		}
	}

	return ans
}

func main() {
	// fmt.Println(getSubstringCount("011001")) // 4
	// fmt.Println(getSubstringCount("00110011")) // 6
	// fmt.Println(getSubstringCount("10101"))    // 4
	fmt.Println(countBinarySubstringsV2("0001110011")) // 7
}

// ans = 0
// prev = 0
// crr = 1

// 0123456789
// 0001110011

// i = 1
// 0 = 0
// crr = 2

// i = 2
// 0 = 0
// crr = 3

// i = 3
// 0 = 1
// ans += (0, 3) // ans = 0
// prev = crr = 3 ***
// crr = 1

// i = 4
// 1 = 1
// crr = 2

// i = 5
// 1 = 1
// crr = 3

// i = 6
// 1 = 0
// ans += (3, 3) // ans = 3
// prev = crr = 3 ***
// crr = 1

// i = 7
// 0 = 0
// crr = 2

// i = 8
// 0 = 1
// ans += (3, 2) // ans = 5
// prev = crr = 2 ***
// crr = 1

// i = 9
// 1 = 1
// crr = 2

// ans += (2, 2) // ans = 7

// [0, 3, 2, 2]
// [3, 3, 3, 3]
