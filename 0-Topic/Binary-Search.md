# Binary Search（二分搜尋）完整教學

---

## 📖 第一部分：什麼是 Binary Search？

### 生活化比喻

想像你在玩「**猜數字**」遊戲，對方心中想一個 1~100 的數字，你每次猜一個數，對方會說「太大」或「太小」。

- 笨方法：從 1 開始一個一個猜 → 最差 100 次
- 聰明方法：**先猜 50**
  - 太小 → 剩 51~100，再猜 75
  - 太大 → 剩 1~49，再猜 25
  - 每次**砍一半**，最多 log₂(100) ≈ 7 次就猜到

這就是 Binary Search —— 每次把範圍**砍一半**，用 O(log n) 時間找到目標。

### 為什麼要用 Binary Search？

**核心目的：把 O(n) 線性搜尋 → 優化成 O(log n)**

| n | O(n) 次數 | O(log n) 次數 |
|---|----------|--------------|
| 1,000 | 1,000 | 10 |
| 1,000,000 | 1,000,000 | 20 |
| 1,000,000,000 | 10 億次 | 30 |

資料量越大，Binary Search 威力越猛。

### 使用前提

**陣列必須是排序好的！** 或是具有「單調性」（左邊某條件為真、右邊為假這種可分割的性質）。

### Binary Search 的三大變形

| 變形 | 目標 | 代表題 |
|------|------|--------|
| **找確切值** | 找 target 在陣列中的索引 | LeetCode 704 |
| **找左邊界** | 找第一個 >= target 的位置 | LeetCode 35 |
| **找右邊界** | 找最後一個 <= target 的位置 | 變形題 |
| **在答案上二分** | 題目答案有範圍，二分猜答案 | LeetCode 875 |

---

## 🎯 第二部分：例題實戰

---

## 例題一：二分搜尋（LeetCode 704）— 最基本款

**題目**：給一個**已排序**的陣列 `nums` 和 `target`，回傳 target 的索引；不存在則回傳 -1。

```
nums = [-1, 0, 3, 5, 9, 12], target = 9
答案：4（nums[4] == 9）
```

### 🚀 Go 解法

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2 // 防溢位寫法

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1 // 答案在右半
        } else {
            right = mid - 1 // 答案在左半
        }
    }

    return -1
}
```

### 📊 圖解步驟

`nums = [-1, 0, 3, 5, 9, 12]`, `target = 9`

#### 初始狀態

```
索引:    0    1    2    3    4    5
陣列:  [ -1,  0,   3,   5,   9,  12 ]
        ↑              ↑              ↑
       left          mid           right

mid = 0 + (5-0)/2 = 2
nums[2] = 3 < 9 → 答案在右半 → left = mid + 1 = 3
```

#### Step 1

```
索引:    0    1    2    3    4    5
陣列:  [ -1,  0,   3,   5,   9,  12 ]
                        ↑    ↑    ↑
                       left mid right

mid = 3 + (5-3)/2 = 4
nums[4] = 9 == 9 → 找到了！回傳 4 ✅
```

### 🚨 為什麼 `mid = left + (right-left)/2` 而不是 `(left+right)/2`？

當 `left` 和 `right` 都很大時，`left + right` 可能**整數溢位**（雖然 Go int 是 64 位不太會，但 C/Java 常見）。這是面試時能拿分的小細節。

---

## 例題二：搜尋插入位置（LeetCode 35）— 找左邊界

**題目**：在排序陣列中找 target 的位置；若不存在，回傳**它應該插入的位置**（保持排序）。

```
nums = [1, 3, 5, 6], target = 5  → 2
nums = [1, 3, 5, 6], target = 2  → 1（插入到索引 1）
nums = [1, 3, 5, 6], target = 7  → 4（插入到最後）
```

### 🤔 思路

這題是「找**第一個 >= target 的位置**」的變形。即使沒找到 target，`left` 最終會停在「它應該插入的位置」。

### 🚀 Go 解法

```go
func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return left // 找不到時，left 就是插入位置
}
```

### 📊 圖解：`nums = [1, 3, 5, 6], target = 2`

```
初始: left=0, right=3
mid=1, nums[1]=3 > 2 → right = 0

Step 1: left=0, right=0
mid=0, nums[0]=1 < 2 → left = 1

Step 2: left=1, right=0 → left > right 跳出迴圈
回傳 left = 1 ✅（target 2 應該插在索引 1）
```

### 🤔 為什麼迴圈跳出後，`left` 就是答案？

當 target 不在陣列中，`left` 和 `right` 會交錯，`left` 永遠停在「**第一個 > target 的位置**」——也就是插入位置。這是二分搜尋的一個重要不變性質。

---

## 例題三：尋找峰值（LeetCode 162）— 單調性不依賴排序

**題目**：給一個陣列 `nums`，找出**任意一個峰值**（比左右鄰居都大的元素）的索引。`nums[-1]` 和 `nums[n]` 當作 -∞。

```
nums = [1, 2, 3, 1] → 答案：2（nums[2]=3 是峰值）
```

**要求 O(log n) → 提示這是 Binary Search！**

### 🤔 思路：為什麼不需要排序也能二分？

**關鍵觀察**：當 `nums[mid] < nums[mid+1]`，代表右邊正在「爬坡」，**右邊一定有峰值**（至少在陣列邊界會遇到一個下坡點）。

反之若 `nums[mid] > nums[mid+1]`，代表正在「下坡」，**左邊（含 mid）一定有峰值**。

### 🚀 Go 解法

```go
func findPeakElement(nums []int) int {
    left, right := 0, len(nums)-1

    for left < right {
        mid := left + (right-left)/2

        if nums[mid] < nums[mid+1] {
            left = mid + 1 // 右半爬坡，峰值在右邊
        } else {
            right = mid // 左半下坡或 mid 就是峰值
        }
    }

    return left
}
```

### 📊 圖解：`nums = [1, 2, 1, 3, 5, 6, 4]`

```
索引:    0    1    2    3    4    5    6
陣列:  [ 1,   2,   1,   3,   5,   6,   4 ]

初始: left=0, right=6
mid=3, nums[3]=3 < nums[4]=5 → 右半爬坡 → left = 4

Step 1: left=4, right=6
mid=5, nums[5]=6 > nums[6]=4 → 右半下坡 → right = 5

Step 2: left=4, right=5
mid=4, nums[4]=5 < nums[5]=6 → 右半爬坡 → left = 5

Step 3: left=5, right=5 → 跳出
回傳 5 ✅（nums[5]=6 是一個峰值）
```

---

## 🧠 第三部分：心法整理

### 1. Binary Search 的核心三問

寫 Binary Search 前先回答：

```
1. 搜尋範圍 [left, right] 還是 [left, right)？
2. 迴圈條件 left <= right 還是 left < right？
3. 找不到時回傳什麼？-1？left？
```

**推薦模板**（閉區間 `[left, right]`）：

```go
left, right := 0, len(nums)-1
for left <= right {
    mid := left + (right-left)/2
    if nums[mid] == target {
        return mid
    } else if nums[mid] < target {
        left = mid + 1
    } else {
        right = mid - 1
    }
}
return -1 // 找不到
```

### 2. 避免無窮迴圈的兩個關鍵

- **區間一定要縮小**：`left = mid + 1` 或 `right = mid - 1`（不是 `left = mid`）
- **邊界條件一致**：`[left, right]` 用 `<=`；`[left, right)` 用 `<`

寫錯會卡死，debug 時先檢查這兩點。

### 3. 辨識題型的關鍵字

看到這些，先想 Binary Search：

- 「**已排序**的陣列」+ 找某個值
- 要求 **O(log n)** 時間複雜度
- 「**第一個/最後一個**滿足某條件的位置」
- 題目答案有**明確範圍**，可以猜（eg. 最小時間、最大容量）

### 4. 「在答案上二分」是進階模式

當題目本身不是排序陣列，但**答案有單調性**（eg. 某個時間能完成 → 比它大的時間也能完成），就可以對「答案的範圍」做二分。

經典題：LeetCode 875 Koko Eating Bananas、LeetCode 1011 Ship Within D Days。

### 5. 「找左邊界」 vs 「找右邊界」模板

**找第一個 >= target 的位置（左邊界）**：

```go
for left < right {
    mid := left + (right-left)/2
    if nums[mid] < target {
        left = mid + 1
    } else {
        right = mid
    }
}
return left
```

**找最後一個 <= target 的位置（右邊界）**：

```go
for left < right {
    mid := left + (right-left+1)/2 // 注意 +1 防死循環
    if nums[mid] > target {
        right = mid - 1
    } else {
        left = mid
    }
}
return left
```

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 技巧 |
|------|------|------|
| Easy | LeetCode 704. Binary Search | 基本款 |
| Easy | LeetCode 35. Search Insert Position | 找左邊界 |
| Easy | LeetCode 278. First Bad Version | 找左邊界變形 |
| Medium | LeetCode 162. Find Peak Element | 單調性不依賴排序 |
| Medium | LeetCode 33. Search in Rotated Sorted Array | 旋轉陣列二分 |
| Medium | LeetCode 34. Find First and Last Position | 左右邊界 |
| Medium | LeetCode 153. Find Minimum in Rotated Sorted Array | 旋轉陣列找最小 |
| Medium | LeetCode 875. Koko Eating Bananas | 在答案上二分 |
| Hard | LeetCode 4. Median of Two Sorted Arrays | 雙陣列二分 |
