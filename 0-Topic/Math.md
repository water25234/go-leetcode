# Math（數學）完整教學

---

## 📖 第一部分：什麼是 Math 題？

### 生活化比喻

數學題就像**走捷徑的魔法**：

- 別人要跑 10 公里 → 你看穿數學規律，**一個公式直接到終點**
- 別人要寫一堆 for 迴圈 → 你一行位元運算搞定

Math 類題目不考「難演算法」，考的是**觀察規律 + 用對數學工具**。

### Math 題的幾大類型

| 類型 | 工具 | 代表題 |
|------|------|--------|
| **數位處理** | 取模、除法 | LeetCode 7, 9 |
| **位元運算** | AND, OR, XOR, 移位 | LeetCode 136, 191 |
| **快速冪** | 二分思想 | LeetCode 50 |
| **質數/GCD** | 篩法、歐幾里得 | LeetCode 204 |
| **規律題** | 找週期、組合數學 | LeetCode 202, 62 |

### 常用的數學工具箱

#### 1. 取模 & 除法（拆數位）

```go
num := 1234

// 取個位數
digit := num % 10  // 4

// 去掉個位數
num = num / 10     // 123

// 迴圈把每一位拿出來
for num > 0 {
    digit := num % 10
    // 處理 digit
    num /= 10
}
```

#### 2. 位元運算

```go
a & b   // AND（同 1 才 1）：常用於「提取某一位元」
a | b   // OR（任一 1 即 1）：常用於「設定某一位元」
a ^ b   // XOR（不同才 1）：相同兩次會消掉
^a      // NOT（Go 是 unary 反位元）
a << n  // 左移 n 位（相當於 ×2ⁿ）
a >> n  // 右移 n 位（相當於 ÷2ⁿ）

// 常用技巧
a & 1        // 判斷奇偶（1 奇 0 偶）
a & (a-1)    // 把 a 最右邊的 1 變 0（經典技巧）
a ^ a = 0    // 任何數 XOR 自己 = 0
a ^ 0 = a    // 任何數 XOR 0 = 自己
```

#### 3. 整數溢位

Go 的 `int` 在 64 位系統是 64 位，範圍夠大。但 LeetCode 常考 32 位整數溢位：

```go
const INT_MAX = 1<<31 - 1  // 2147483647
const INT_MIN = -(1 << 31) // -2147483648
```

---

## 🎯 第二部分：例題實戰

---

## 例題一：回文數（LeetCode 9）— 數位拆解

**題目**：判斷一個整數是否為回文（正著讀反著讀一樣）。**不能把數字轉字串**。

```
121     → true
-121    → false（負號不對稱）
10      → false（01 ≠ 10）
```

### 🤔 思路：反轉一半數字

不用真的反轉整個數字，**只反轉右半部**，再跟左半部比較就好。

- 一直從右邊切一位出來，丟到 `rev`
- 當 `rev >= num`（剩下的左半 ≤ 右半）就夠了
- 比較 `rev == num`（偶數位）或 `rev/10 == num`（奇數位，忽略中間位）

### 🚀 Go 解法

```go
func isPalindrome(x int) bool {
    if x < 0 || (x%10 == 0 && x != 0) {
        return false // 負數 & 尾數 0 的不可能是回文
    }

    rev := 0
    for x > rev {
        rev = rev*10 + x%10
        x /= 10
    }

    return x == rev || x == rev/10
}
```

### 📊 圖解：`x = 1221`

```
初始 x=1221, rev=0

Step 1: rev = 0*10 + 1 = 1,    x = 122
Step 2: rev = 1*10 + 2 = 12,   x = 12
        x (12) == rev (12) 跳出

x == rev → 1221 是回文 ✅
```

圖解：`x = 12321`（奇數位）

```
Step 1: rev=1,  x=1232
Step 2: rev=12, x=123
Step 3: rev=123, x=12
        x(12) < rev(123) 跳出

中間位數是 rev 的尾數 3，比較時要忽略
rev/10 = 12 == x(12) → 回文 ✅
```

### 🔑 關鍵細節：尾數 0 的陷阱

除了 0 本身，**以 0 結尾的數不可能是回文**（因為數字沒有前導 0，所以回文要求首尾都不能是 0，除非整個數就是 0）。

---

## 例題二：只出現一次的數字（LeetCode 136）— XOR 神技

**題目**：陣列中每個數字都出現**兩次**，只有一個出現**一次**。找出它。要求 **O(n) 時間、O(1) 空間**。

```
nums = [4, 1, 2, 1, 2]
答案：4
```

### 🤔 思路：XOR 的神奇性質

```
a ^ a = 0    （任何數 XOR 自己 = 0）
a ^ 0 = a    （任何數 XOR 0 = 自己）
XOR 滿足交換律和結合律
```

把所有數 XOR 起來：**成對的會互相消掉變 0，剩下那個孤單的數**。

### 🚀 Go 解法

```go
func singleNumber(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num
    }
    return result
}
```

**三行搞定**，沒有 hash map、沒有排序，O(n) 時間、O(1) 空間。

### 📊 圖解：`nums = [4, 1, 2, 1, 2]`

```
0 ^ 4 = 4
4 ^ 1 = 5 （二進位 101）
5 ^ 2 = 7 （二進位 111）
7 ^ 1 = 6 （二進位 110）
6 ^ 2 = 4 （二進位 100）

回傳 4 ✅
```

甚至，XOR 有交換律，順序無關：`4 ^ (1^1) ^ (2^2) = 4 ^ 0 ^ 0 = 4`。一對一對抵消，最後留下孤單的 4。

### 🔑 XOR 能解的其他題

- LeetCode 268 Missing Number（0~n 少一個）
- LeetCode 389 Find the Difference（兩字串差一個字元）
- LeetCode 260 Single Number III（兩個數出現一次）——進階版，要分組

---

## 例題三：Pow(x, n)（LeetCode 50）— 快速冪

**題目**：實作 `pow(x, n)`，計算 x 的 n 次方。

```
pow(2.0, 10) = 1024
pow(2.0, -2) = 0.25
```

### 🐢 暴力解（O(n)）

```go
result := 1.0
for i := 0; i < n; i++ {
    result *= x
}
```

n 很大時會 TLE。

### 🚀 快速冪（O(log n)）

**核心觀察**：把 n 拆成二進位，每一位代表要不要乘上對應的 $x^{2^k}$。

```
n = 13 = 1101 (二進位)
      = 8 + 4 + 1
      = 2³ + 2² + 2⁰

x^13 = x⁸ × x⁴ × x¹
     = (x²)² × ... 一直平方就能 O(log n) 算出來
```

### 🚀 Go 解法（迭代版）

```go
func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1 / x
        n = -n
    }

    result := 1.0
    for n > 0 {
        if n%2 == 1 {
            result *= x
        }
        x *= x
        n /= 2
    }
    return result
}
```

### 📊 圖解：`x = 2, n = 10`

```
n = 10 = 1010 (二進位) = 2³ + 2¹
x^10 = x² × x⁸

初始: result=1, x=2, n=10

n=10, n%2=0, x 平方 = 4,    n=5
n=5,  n%2=1, result *= 4 = 4, x 平方 = 16, n=2
n=2,  n%2=0, x 平方 = 256,   n=1
n=1,  n%2=1, result *= 256 = 1024, x 平方 = ..., n=0

回傳 1024 ✅
```

**關鍵**：每次迴圈 `x` 都平方（$x → x² → x⁴ → x⁸$），根據 n 的二進位每一位決定要不要乘進 result。

---

## 例題四：快樂數（LeetCode 202）— 找循環

**題目**：判斷一個數是否為「快樂數」：**把每位平方相加，重複此過程，最後變 1 就是快樂數，陷入無限循環就不是**。

```
19 → 1²+9² = 82 → 8²+2² = 68 → 6²+8² = 100 → 1²+0²+0² = 1 ✅ 快樂
```

### 🤔 思路

- 用 **Set** 記錄看過的數，重複出現就陷入循環 → 不快樂
- 更炫：**快慢指針**（像判斷鏈結串列有環）

### 🚀 Go 解法（快慢指針版）

```go
func isHappy(n int) bool {
    slow, fast := n, next(n)
    for fast != 1 && slow != fast {
        slow = next(slow)
        fast = next(next(fast))
    }
    return fast == 1
}

func next(n int) int {
    sum := 0
    for n > 0 {
        d := n % 10
        sum += d * d
        n /= 10
    }
    return sum
}
```

### 🔑 為什麼快慢指針能判循環？

所有「下一步唯一決定」的轉換，都能看成一條**隱式鏈結串列**：

- 快樂數 → 最後到達 1，停下來
- 不快樂 → 陷入某個循環

等於是「鏈結串列有沒有環」的變形，可以套 `hasCycle` 的思路。

---

## 🧠 第三部分：心法整理

### 1. 辨識題型的關鍵字

看到這些，想想是不是數學題：

- 「**不能用多餘空間 / O(1) 空間**」→ 位元運算
- 「**x 的 n 次方 / a^b mod c**」→ 快速冪
- 「**反轉整數 / 回文數 / 數位處理**」→ 取模拆解
- 「**每個數出現 k 次，只有一個出現一次**」→ XOR 系列
- 「**質數**」→ 埃氏篩法
- 「**GCD / LCM**」→ 歐幾里得演算法

### 2. 位元運算的常用技巧表

| 技巧 | 程式碼 | 用途 |
|------|--------|------|
| 判奇偶 | `n & 1` | 1 奇 0 偶 |
| 乘 2 | `n << 1` | 移位快於乘法 |
| 除 2 | `n >> 1` | 移位快於除法 |
| 清除最低位 1 | `n & (n-1)` | 數位元中有幾個 1 |
| 取最低位 1 | `n & -n` | Fenwick Tree 用到 |
| 檢查第 i 位 | `(n >> i) & 1` | 單位元查詢 |
| 設定第 i 位 | `n \| (1 << i)` | 單位元設置 |

### 3. 快速冪的萬用模板（含取模版）

```go
func fastPow(x, n, mod int) int {
    result := 1
    x %= mod
    for n > 0 {
        if n&1 == 1 {
            result = result * x % mod
        }
        x = x * x % mod
        n >>= 1
    }
    return result
}
```

計算 `x^n mod p` 的題目這版直接套。

### 4. 整數溢位處理

反轉整數、數位處理時要留意 32 位整數範圍。Go 用 int 可能感覺不到，但題目若限定 32 位，要主動檢查：

```go
if result > math.MaxInt32 || result < math.MinInt32 {
    return 0
}
```

### 5. 數學題的解題思維

```
1. 先看題目能不能從「暴力做一遍」觀察出「規律」
2. 找出不需要執行完整運算就能推出答案的捷徑
3. 查對應的數學工具（取模、位元、GCD、快速冪...）
```

很多數學題答案非常短（甚至一行），但**看不出規律就寫不出來**。遇到看似「暴力一定 TLE」的題目，優先思考數學規律。

### 6. Go 的數學工具庫

```go
import "math"

math.Abs(x)       // 絕對值（float）
math.Sqrt(x)      // 開平方根
math.Pow(x, y)    // 冪
math.Log2(x)      // 底 2 log
math.MaxInt32     // 2147483647
math.MinInt32     // -2147483648
```

注意：`math.Pow` 是 float，整數冪需要自己寫快速冪版本。

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 技巧 |
|------|------|------|
| Easy | LeetCode 9. Palindrome Number | 數位拆解 |
| Easy | LeetCode 7. Reverse Integer | 數位拆解 + 溢位 |
| Easy | LeetCode 136. Single Number | XOR |
| Easy | LeetCode 191. Number of 1 Bits | 位元計數 |
| Easy | LeetCode 231. Power of Two | 位元技巧 `n & (n-1) == 0` |
| Easy | LeetCode 202. Happy Number | 快慢指針 / Set |
| Medium | LeetCode 50. Pow(x, n) | 快速冪 |
| Medium | LeetCode 204. Count Primes | 埃氏篩法 |
| Medium | LeetCode 137. Single Number II | 位元狀態機 |
| Medium | LeetCode 372. Super Pow | 快速冪變形 |
| Medium | LeetCode 29. Divide Two Integers | 不用乘除做除法 |
| Hard | LeetCode 233. Number of Digit One | 數位 DP |
