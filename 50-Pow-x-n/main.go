package main

// mplement pow(x, n), which calculates x raised to the power n (i.e., xn).
// Example 1:

// Input: x = 2.00000, n = 10
// Output: 1024.00000
// Example 2:

// Input: x = 2.10000, n = 3
// Output: 9.26100
// Example 3:

// Input: x = 2.00000, n = -2
// Output: 0.25000
// Explanation: 2-2 = 1/22 = 1/4 = 0.25

// Constraints:
// -100.0 < x < 100.0
// -231 <= n <= 231-1
// n is an integer.
// Either x is not zero or n > 0.
// -104 <= xn <= 104

// n & 1 是位元 AND 運算，用來判斷 n 是奇數還是偶數。它檢查 n 的最低位元（last bit）是 0 還是 1：
// 5 & 1 → 101 & 001 → 1（奇數）
// 4 & 1 → 100 & 001 → 0（偶數）
// 效果等同於 n % 2 == 1，但比取餘數更快。
// 規律就是 n & (2^k - 1) 等同 n % 2^k，所以：
// n & 1 = n % 2
// n & 3 = n % 4
// n & 7 = n % 8
// n & 15 = n % 16

// n >> 1 是右移一位，等同於 n / 2（整數除法）：
// 5 >> 1 → 101 變成 10 → 2
// 4 >> 1 → 100 變成 10 → 2

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var res float64 = 1

	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n = n >> 1
	}
	return res
}

func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	var result float64

	if n <= -1 {
		x = 1 / x
		n = -n
	}

	result = x
	for i := 0; i < n-1; i++ {
		result = result * x
	}

	return result
}

func myPow3(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	var result float64

	if n <= -1 {
		x = 1 / x
		n = -n
	}

	result = 1
	return recursionPow(result, x, n)
}

func recursionPow(result float64, x float64, n int) float64 {
	if n == 0 {
		return result
	}

	result = result * x
	n--

	return recursionPow(result, x, n)
}
