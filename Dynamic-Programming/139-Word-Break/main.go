package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	n := len(s) // leetcode len 8
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// fmt.Println(s[j:i], j, i)

			if dp[j] && wordSet[s[j:i]] {
				// fmt.Println(s[j:i], j, i)
				dp[i] = true
				break
			}
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func main() {
	// s := "leetcode"
	// wordDict := []string{"leet", "code"}

	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}

	result := wordBreak(s, wordDict)
	fmt.Println("Can the string be segmented?", result)
}
