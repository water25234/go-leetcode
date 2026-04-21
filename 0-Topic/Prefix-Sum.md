# Prefix Sum（前綴和）完整教學

---

## 📖 第一部分：什麼是 Prefix Sum？

### 生活化比喻

想像你是**銀行**，客戶每天存一筆錢，你要回答「**第 3 天到第 7 天一共存了多少**」這種問題。

- 笨方法：每次問題來了，就翻帳本從第 3 天加到第 7 天
- 聰明方法：**預先做一張「累計表」**
  - 第 1 天累計 100
  - 第 2 天累計 250
  - 第 3 天累計 400
  - ...
  - 要查第 3~7 天？→ `第 7 天累計 - 第 2 天累計`，一個減法搞定

這就是 Prefix Sum —— **預處理一個累計陣列，讓區間和的查詢變成 O(1)**。

### 為什麼要用 Prefix Sum？

**核心目的：把重複的區間加總從 O(n) → O(1)**

| | 暴力 | Prefix Sum |
|---|------|-----------|
| **預處理** | 不用 | O(n) 建表 |
| **單次查詢** | O(n) 加總 | **O(1) 減法** |
| **多次查詢 q 次** | O(q × n) | O(n + q) |

查詢次數越多，Prefix Sum 越划算。

### Prefix Sum 的定義

```
prefix[i] = nums[0] + nums[1] + ... + nums[i-1]
```

注意：`prefix[0] = 0`（空和），`prefix` 的長度是 `n+1`，這樣查 `[i, j]` 的和就是：

```
sum(i, j) = prefix[j+1] - prefix[i]
```

統一公式、沒有邊界條件要特判，很乾淨。

### Prefix Sum 的兩種常見變形

| 變形 | 用途 | 代表題 |
|------|------|--------|
| **一維 Prefix Sum** | 區間和查詢 | LeetCode 303 |
| **Prefix Sum + HashMap** | 找滿足條件的子陣列 | LeetCode 560 |
| **二維 Prefix Sum** | 矩陣區域和 | LeetCode 304 |
| **差分陣列** | 區間加值的反向應用 | LeetCode 1109 |

---

## 🎯 第二部分：例題實戰

---

## 例題一：區域和查詢（LeetCode 303）— Prefix Sum 入門

**題目**：給一個整數陣列 `nums`，設計一個類別支援 `sumRange(i, j)` 回傳 `nums[i..j]` 的和。可能會被呼叫很多次。

```
nums = [-2, 0, 3, -5, 2, -1]
sumRange(0, 2) → -2 + 0 + 3 = 1
sumRange(2, 5) → 3 - 5 + 2 - 1 = -1
sumRange(0, 5) → -3
```

### 🚀 Go 解法

```go
type NumArray struct {
    prefix []int
}

func Constructor(nums []int) NumArray {
    n := len(nums)
    prefix := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }
    return NumArray{prefix: prefix}
}

func (a *NumArray) SumRange(i, j int) int {
    return a.prefix[j+1] - a.prefix[i]
}
```

### 📊 圖解

`nums = [-2, 0, 3, -5, 2, -1]`

```
索引:      0    1    2    3    4    5
nums:   [ -2,  0,   3,  -5,   2,  -1 ]

prefix: [  0, -2,  -2,   1,  -4,  -2,  -3 ]
索引:      0    1    2    3    4    5    6

prefix[i] 代表 nums[0..i-1] 的累計和

sumRange(0, 2) = prefix[3] - prefix[0] = 1 - 0 = 1
sumRange(2, 5) = prefix[6] - prefix[2] = -3 - (-2) = -1
sumRange(0, 5) = prefix[6] - prefix[0] = -3 - 0 = -3
```

### 🔑 為什麼 `prefix` 開 n+1 長度？

為了讓 `sumRange(0, j)` 的公式統一成 `prefix[j+1] - prefix[0]`，不用特別處理「起點是 0」的邊界。寫起來乾淨，不容易出錯。

---

## 例題二：和為 K 的子陣列（LeetCode 560）— Prefix Sum + HashMap

**題目**：給一個陣列 `nums` 和整數 `k`，求**連續子陣列**的和等於 k 的數量。

```
nums = [1, 1, 1], k = 2 → 答案：2
（子陣列 [1,1] 在索引 [0,1] 和 [1,2] 各出現一次）

nums = [1, 2, 3], k = 3 → 答案：2
（[1,2] 和 [3]）
```

### 🤔 為什麼這題是 Prefix Sum？

`nums[i..j]` 的和 = `prefix[j+1] - prefix[i]`

要找**和等於 k 的子陣列**，等價於找**多少對 (i, j) 使得** `prefix[j+1] - prefix[i] = k`。

移項：`prefix[i] = prefix[j+1] - k`

**轉成「找兩個前綴和的差等於 k」** → 這根本就是 Two Sum 的變形！用 HashMap 存已看過的 prefix 值及其出現次數即可。

### 🚀 Go 解法

```go
func subarraySum(nums []int, k int) int {
    count := 0
    prefixCount := map[int]int{0: 1} // 空前綴 = 0 出現 1 次（為了處理從頭開始的子陣列）
    prefix := 0

    for _, num := range nums {
        prefix += num

        // 看之前有沒有出現過 prefix - k
        if c, ok := prefixCount[prefix-k]; ok {
            count += c
        }

        prefixCount[prefix]++
    }

    return count
}
```

### 📊 圖解：`nums = [1, 2, 3], k = 3`

```
初始：prefixCount = {0: 1}, prefix = 0, count = 0

i=0: num=1, prefix=1
     prefix - k = -2，不在 map → count 不變
     prefixCount = {0:1, 1:1}

i=1: num=2, prefix=3
     prefix - k = 0，map 裡有 0 出現 1 次 → count += 1 = 1  ⭐
     （代表 nums[0..1] = [1,2] 和為 3）
     prefixCount = {0:1, 1:1, 3:1}

i=2: num=3, prefix=6
     prefix - k = 3，map 裡有 3 出現 1 次 → count += 1 = 2  ⭐
     （代表 nums[2..2] = [3] 和為 3）
     prefixCount = {0:1, 1:1, 3:1, 6:1}

最終 count = 2 ✅
```

### 🔑 為什麼初始化 `prefixCount[0] = 1`？

如果從索引 0 開始的子陣列本身就等於 k（整個 prefix == k），那 `prefix - k = 0`，需要 map 裡有一個「空前綴」來匹配。這個初始化是 Prefix Sum + HashMap 題目的**必備樣板**。

---

## 例題三：商品間的差距（LeetCode 1031 概念 / 練習題）— 差分陣列預習

**題目情境**：有 n 間房子，你收到一堆任務「把第 i 到第 j 間房子每間加 v 塊錢」。問所有任務做完後，每間房子多少錢？

如果暴力做，每個任務 O(n)，q 個任務就是 O(qn)。

### 🚀 差分陣列解法（O(n+q)）

**差分陣列**是 Prefix Sum 的反操作：

```
diff[i] = nums[i] - nums[i-1]
```

區間加值只改兩個端點：

```
「i 到 j 全部 +v」等價於 diff[i] += v, diff[j+1] -= v
```

做完所有任務，對 diff 做一次 Prefix Sum 就還原成最終結果。

```go
func applyOperations(n int, ops [][]int) []int {
    diff := make([]int, n+1)
    for _, op := range ops {
        i, j, v := op[0], op[1], op[2]
        diff[i] += v
        diff[j+1] -= v
    }

    result := make([]int, n)
    result[0] = diff[0]
    for i := 1; i < n; i++ {
        result[i] = result[i-1] + diff[i]
    }
    return result
}
```

### 🔑 為什麼區間加值只改兩個端點？

```
diff 表示「相鄰元素的差」。
區間 [i, j] 裡，所有位置的 nums 都 +v，相鄰差值不變。
只有：
  - 進入區間的那一格（diff[i]）會增加 v
  - 離開區間的那一格（diff[j+1]）會減少 v

所以 O(1) 改兩個點，最後一次 prefix sum 還原整個陣列。
```

---

## 🧠 第三部分：心法整理

### 1. 辨識題型的關鍵字

看到這些，先想 Prefix Sum：

- 「**區間和**」相關問題
- 「**連續子陣列**的和等於 X」
- 「**多次查詢**區間資訊」
- 「**區間加值**/區間修改」→ 差分陣列

### 2. Prefix Sum 三件事要記得

```
1. prefix 開 n+1 長度，prefix[0] = 0（省去邊界判斷）
2. prefix[i+1] = prefix[i] + nums[i]（一次迴圈建表）
3. sum(i, j) = prefix[j+1] - prefix[i]
```

### 3. Prefix Sum + HashMap 的萬用模板

```go
prefixCount := map[int]int{0: 1} // 必要初始化！
prefix := 0
count := 0

for _, num := range nums {
    prefix += num
    if c, ok := prefixCount[prefix-target]; ok {
        count += c
    }
    prefixCount[prefix]++
}
```

這個 pattern 解的題目：
- LeetCode 560（和為 K）
- LeetCode 523（和為 K 的倍數）
- LeetCode 974（和能被 K 整除）

### 4. 二維 Prefix Sum 公式

```
prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + matrix[i][j]

查 (r1,c1) 到 (r2,c2) 的區域和：
= prefix[r2+1][c2+1] - prefix[r1][c2+1] - prefix[r2+1][c1] + prefix[r1][c1]
```

看起來複雜，其實就是**容斥原理**：大矩形 - 左上矩形 - 右上矩形 + 左上角（被扣兩次補回來）。

### 5. 差分陣列 vs Prefix Sum

```
原陣列 ←─ prefix sum ──→ 前綴和陣列
原陣列 ←── diff ────→ 差分陣列
原陣列 ── diff → 差分 ── prefix → 原陣列（再還原）
```

- **Prefix Sum**：快速區間**查詢**
- **差分陣列**：快速區間**修改**

兩者互為逆操作，看題目要「多次查」還是「多次改」選其中一種。

### 6. 什麼時候不該用 Prefix Sum？

- **陣列會頻繁更新單點** → Prefix Sum 要重建，考慮 Fenwick Tree / 線段樹
- **只查一次**的區間和 → 直接 for 迴圈加，別建表
- **求最大/最小子陣列和** → 用 Kadane's Algorithm（DP），不是 Prefix Sum

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 技巧 |
|------|------|------|
| Easy | LeetCode 303. Range Sum Query - Immutable | 一維 Prefix Sum |
| Easy | LeetCode 1480. Running Sum of 1d Array | 最基本 |
| Medium | LeetCode 560. Subarray Sum Equals K | Prefix Sum + HashMap |
| Medium | LeetCode 974. Subarray Sums Divisible by K | Prefix Sum + HashMap 取模 |
| Medium | LeetCode 523. Continuous Subarray Sum | Prefix Sum + HashMap 取模 |
| Medium | LeetCode 304. Range Sum Query 2D - Immutable | 二維 Prefix Sum |
| Medium | LeetCode 1109. Corporate Flight Bookings | 差分陣列 |
| Medium | LeetCode 238. Product of Array Except Self | 前綴積 + 後綴積 |
| Hard | LeetCode 363. Max Sum of Rectangle No Larger Than K | 二維 + TreeMap |
