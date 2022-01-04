package main

import (
	"fmt"
	"strconv"
)

func main() {

	// var tokens string

	// tokens = "bbbbb"
	// tokens = "pwwkew"
	// tokens = "abcabcbb"

	fmt.Println(bitwiseComplement(32))
}

func bitwiseComplement(n int) int {
	p := 1
	for p < n {
		p = 2*p + 1
	}

	fmt.Println(p, strconv.FormatInt(int64(p), 2))
	fmt.Println(n, strconv.FormatInt(int64(n), 2))
	return (p ^ n)
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		// 移动到相同字母的最右边（如果有相同字母）
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 找到回文的边界
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置到下一次寻找回文的中心
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}
