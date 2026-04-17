# Sliding Window + Hash Map：找連續 k 個不重複數字的最大總和

> LeetCode 2461. Maximum Sum of Distinct Subarrays With Length K

---

## 📋 題目定義

- 找出長度為 `k` 的連續子陣列
- 這個子陣列裡的元素**必須全部不重複**
- 回傳所有合法子陣列中的**最大總和**
- 如果沒有任何合法的子陣列，回傳 **0**

**範例**：

```
nums = [1, 5, 4, 2, 9, 9, 9], k = 3

所有長度為 3 的子陣列：
[1,5,4] ✅ sum = 10
[5,4,2] ✅ sum = 11
[4,2,9] ✅ sum = 15  ⭐ 最大
[2,9,9] ❌ 9 重複
[9,9,9] ❌ 9 重複

答案 = 15
```

---

## 🤔 思路推導

### 暴力解的問題

每次視窗滑動，都要檢查「視窗內有沒有重複」。

如果用兩層 for 迴圈檢查，會變成 O(n × k²)，更慢。

### Sliding Window + Map 的解法

**核心想法**：用一個 `map` 記錄「視窗內每個數字出現的次數」

- 加入新元素時：`map[新元素]++`
- 移除舊元素時：`map[舊元素]--`，如果變成 0 就 `delete`
- **判斷視窗合法**：`len(map) == k`（因為長度 k 的視窗如果都不重複，map 裡就會剛好有 k 個 key）

---

## 🚀 Go 解法

```go
func maximumSubarraySum(nums []int, k int) int64 {
    var maxSum int64 = 0
    var windowSum int64 = 0
    window := make(map[int]int) // 記錄視窗內每個數字的出現次數

    // Step 1: 建立第一個視窗
    for i := 0; i < k; i++ {
        windowSum += int64(nums[i])
        window[nums[i]]++
    }

    // 第一個視窗如果合法，更新答案
    if len(window) == k {
        maxSum = windowSum
    }

    // Step 2: 滑動視窗
    for i := k; i < len(nums); i++ {
        // 加入新元素（右邊進）
        windowSum += int64(nums[i])
        window[nums[i]]++

        // 移除舊元素（左邊出）
        leftElement := nums[i-k]
        windowSum -= int64(leftElement)
        window[leftElement]--
        if window[leftElement] == 0 {
            delete(window, leftElement) // 關鍵：歸零就刪除
        }

        // 判斷視窗是否合法（map 大小 == k 表示沒重複）
        if len(window) == k && windowSum > maxSum {
            maxSum = windowSum
        }
    }

    return maxSum
}
```

---

## 📊 圖解步驟

`nums = [1, 5, 4, 2, 9, 9, 9]`, `k = 3`

### 初始視窗（索引 0~2）

```
索引:    0   1   2   3   4   5   6
陣列:  [ 1,  5,  4,  2,  9,  9,  9 ]
        ↑       ↑
       [█████████]

windowSum = 1 + 5 + 4 = 10
window = {1:1, 5:1, 4:1}  → len = 3 ✅ 合法
maxSum = 10
```

### i = 3：加入 nums[3]=2，移除 nums[0]=1

```
索引:    0   1   2   3   4   5   6
陣列:  [ 1,  5,  4,  2,  9,  9,  9 ]
            ↑       ↑
           [█████████]
        ✗               ← 移除 1
                ✓       ← 加入 2

windowSum = 10 + 2 - 1 = 11
window = {5:1, 4:1, 2:1}  → len = 3 ✅ 合法
maxSum = max(10, 11) = 11
```

### i = 4：加入 nums[4]=9，移除 nums[1]=5

```
索引:    0   1   2   3   4   5   6
陣列:  [ 1,  5,  4,  2,  9,  9,  9 ]
                ↑       ↑
               [█████████]
            ✗               ← 移除 5
                    ✓       ← 加入 9

windowSum = 11 + 9 - 5 = 15
window = {4:1, 2:1, 9:1}  → len = 3 ✅ 合法
maxSum = max(11, 15) = 15  ⭐
```

### i = 5：加入 nums[5]=9，移除 nums[2]=4

```
索引:    0   1   2   3   4   5   6
陣列:  [ 1,  5,  4,  2,  9,  9,  9 ]
                    ↑       ↑
                   [█████████]
                ✗               ← 移除 4
                        ✓       ← 加入 9

windowSum = 15 + 9 - 4 = 20
window = {2:1, 9:2}  → len = 2 ❌ 不合法（9 重複）
maxSum 不更新，仍是 15
```

> ⚠️ 注意：雖然 windowSum = 20 很大，但因為 9 重複了，map 大小只有 2，不等於 k=3，所以**不更新答案**。

### i = 6：加入 nums[6]=9，移除 nums[3]=2

```
索引:    0   1   2   3   4   5   6
陣列:  [ 1,  5,  4,  2,  9,  9,  9 ]
                        ↑       ↑
                       [█████████]
                    ✗               ← 移除 2
                            ✓       ← 加入 9

windowSum = 20 + 9 - 2 = 27
window = {9:3}  → len = 1 ❌ 不合法
maxSum 不更新，仍是 15
```

### 最終答案：`maxSum = 15` ✅

---

## 🧠 心法整理

### 1. Sliding Window + Hash Map 是組合拳

看到「**長度 k**」+「**不重複 / distinct**」→ 馬上聯想「**Sliding Window + map**」。

### 2. 歸零一定要 delete

```go
window[leftElement]--
if window[leftElement] == 0 {
    delete(window, leftElement)
}
```

因為 `len(map)` 算的是 key 的數量。如果 value 是 0 但 key 還在，`len(map)` 就會算錯。

> 🚨 這是 Go 寫 Sliding Window 用 map 的**最常見 bug**。

### 3. maxSum 初始值要設 0

```go
var maxSum int64 = 0  // ✅ 正確
```

不能直接 `maxSum = windowSum`，因為第一個視窗可能本身就不合法。

### 4. int64 防溢位

`nums[i]` 可到 10^5，k 可到 10^5，總和最大到 10^10，超過 int32。回傳值用 `int64` 最安全。

### 5. 邊界測資自我檢查清單

| 測資特性         | 範例              | 檢查什麼                 |
| ---------------- | ----------------- | ------------------------ |
| 全部重複         | `[4,4,4,2]` k=3  | 初始視窗合法性判斷       |
| 重複逐漸被擠出   | `[4,4,4,2,1]` k=3 | map value 遞減到 delete |
| 恰好合法         | `[1,2,3]` k=3    | 正常流程                 |
| 完全合法         | `[1,2,3,4,5]` k=3 | 滑動正確性              |
| 部分合法         | `[1,5,4,2,9,9,9]` k=3 | map delete 是否正確 |

---

## 📚 同類型練習題

- LeetCode 2461. Maximum Sum of Distinct Subarrays With Length K（本題）
- LeetCode 1456. Maximum Number of Vowels in a Substring of Given Length
- LeetCode 1004. Max Consecutive Ones III
