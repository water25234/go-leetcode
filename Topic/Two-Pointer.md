# Two Pointers（雙指針）完整教學

---

## 📖 第一部分：什麼是 Two Pointers？

### 生活化比喻

想像你跟朋友站在一條**排好順序的隊伍**兩端：

- 你站在**最左邊**（最矮的人旁邊）
- 朋友站在**最右邊**（最高的人旁邊）
- 有人喊一個目標身高總和，你們兩個根據目前加起來的結果決定**誰要往中間靠一步**
- 太矮了？你往右走（找更高的）
- 太高了？朋友往左走（找更矮的）
- 剛好？找到答案！

這就是 Two Pointers 的核心 —— **兩個指針從不同方向逼近答案**。

### 為什麼要用 Two Pointers？

**核心目的：把 O(n²) 暴力的雙層 for 迴圈 → 優化成 O(n)**

暴力解通常是「i 跑所有位置，j 跑 i 後面所有位置」，但 Two Pointers 利用**已排序**或**結構特性**，讓兩個指針各自只走一次就找到答案。

### Two Pointers 的三種常見模式

| 模式 | 指針方向 | 典型題目 |
|------|---------|---------|
| **對撞指針** | ← 左右往中間靠 → | Two Sum（排序版）、Container With Most Water |
| **同向指針** | → 快慢指針都往右 → | 移除重複元素、鏈表找環 |
| **分離指針** | 兩個陣列各一個指針 → | 合併兩個排序陣列 |

---

## 🎯 第二部分：例題實戰

---

## 例題一：Two Sum II（LeetCode 167）

**題目**：給你一個**已排序**的陣列 `numbers` 和一個目標值 `target`，找出兩個數字加起來等於 `target`，回傳它們的索引（1-based）。

```
numbers = [2, 7, 11, 15], target = 9

答案：[1, 2]  （因為 numbers[0] + numbers[1] = 2 + 7 = 9）
```

### 🐢 暴力解法（O(n²)）

```go
func twoSumBrute(numbers []int, target int) []int {
    for i := 0; i < len(numbers); i++ {
        for j := i + 1; j < len(numbers); j++ {
            if numbers[i]+numbers[j] == target {
                return []int{i + 1, j + 1} // 1-based
            }
        }
    }
    return nil
}
```

**問題**：兩層 for 迴圈，陣列大的時候會超時。

### 🚀 Two Pointers 解法（O(n)）

```go
func twoSum(numbers []int, target int) []int {
    left := 0
    right := len(numbers) - 1

    for left < right {
        sum := numbers[left] + numbers[right]

        if sum == target {
            return []int{left + 1, right + 1} // 1-based
        } else if sum < target {
            left++ // 太小了，左指針右移找更大的
        } else {
            right-- // 太大了，右指針左移找更小的
        }
    }

    return nil
}
```

**關鍵前提**：陣列已排序！這是對撞指針能運作的基礎。

### 📊 圖解步驟

`numbers = [2, 7, 11, 15]`, `target = 9`

#### 初始狀態

```
索引:    0    1    2    3
陣列:  [ 2,   7,  11,  15 ]
        ↑                ↑
       left            right

sum = 2 + 15 = 17
17 > 9 → 太大了！right--
```

#### Step 1：right 左移

```
索引:    0    1    2    3
陣列:  [ 2,   7,  11,  15 ]
        ↑          ↑
       left      right

sum = 2 + 11 = 13
13 > 9 → 還是太大！right--
```

#### Step 2：right 再左移

```
索引:    0    1    2    3
陣列:  [ 2,   7,  11,  15 ]
        ↑    ↑
       left right

sum = 2 + 7 = 9
9 == 9 → 找到了！⭐
回傳 [1, 2]（1-based）
```

只走了 3 步就找到，暴力解要跑 6 次比較 ✅

### 🤔 為什麼對撞指針是正確的？

陣列已排序：`[小 ← ← ← → → → 大]`

- `sum < target`：左邊已經是最小的了，右邊不可能再幫你補回來。唯一的辦法就是 **left++ 找更大的左邊值**
- `sum > target`：右邊已經是最大的了，左邊不可能再幫你減少。唯一的辦法就是 **right-- 找更小的右邊值**

**每一步都排除掉一整行/列的可能性**，所以不會跳過答案。

---

## 例題二：Container With Most Water（LeetCode 11）

**題目**：給你一個陣列 `height`，每個值代表一根柱子的高度。選兩根柱子，它們之間能裝多少水？找出能裝最多水的組合。

```
height = [1, 8, 6, 2, 5, 4, 8, 3, 7]

答案：49
```

水量公式：`min(左柱高, 右柱高) × 兩柱距離`

### 🚀 Two Pointers 解法

```go
func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxWater := 0

    for left < right {
        // 計算當前水量
        h := min(height[left], height[right])
        w := right - left
        water := h * w

        if water > maxWater {
            maxWater = water
        }

        // 移動較矮的那一邊（因為移動較高的不可能更好）
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxWater
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

### 📊 圖解步驟

`height = [1, 8, 6, 2, 5, 4, 8, 3, 7]`

```
索引:  0  1  2  3  4  5  6  7  8
高度:  1  8  6  2  5  4  8  3  7

      |
      |        |              |
      |  |     |              |
      |  |     |        |     |
      |  |     |     |  |     |
      |  |  |  |     |  |  |  |
      |  |  |  |  |  |  |  |  |
   |  |  |  |  |  |  |  |  |  |
   0  1  2  3  4  5  6  7  8
```

#### 初始狀態：left=0, right=8

```
   left=0               right=8
      ↓                    ↓
   |  |  |  |  |  |  |  |  |
   1  8  6  2  5  4  8  3  7

h = min(1, 7) = 1
w = 8 - 0 = 8
water = 1 × 8 = 8
maxWater = 8

height[0]=1 < height[8]=7 → left++（移動矮的那邊）
```

#### Step 1：left=1, right=8

```
      left=1            right=8
         ↓                 ↓
   |  |  |  |  |  |  |  |  |
   1  8  6  2  5  4  8  3  7

h = min(8, 7) = 7
w = 8 - 1 = 7
water = 7 × 7 = 49  ⭐
maxWater = 49

height[1]=8 > height[8]=7 → right--（移動矮的那邊）
```

#### Step 2：left=1, right=7

```
      left=1         right=7
         ↓              ↓
   |  |  |  |  |  |  |  |  |
   1  8  6  2  5  4  8  3  7

h = min(8, 3) = 3
w = 7 - 1 = 6
water = 3 × 6 = 18
maxWater = max(49, 18) = 49

height[1]=8 > height[7]=3 → right--
```

後面繼續滑，但 maxWater 不會超過 49 了。

最終答案：`maxWater = 49` ✅

### 🤔 為什麼移動「矮的那邊」？

```
水量 = min(左, 右) × 距離
```

如果移動**高的那邊**：

- 距離一定變小（少 1）
- `min()` 最多不變（被矮的限制住），可能更小
- 結果一定 ≤ 現在 → **不可能更好**

如果移動**矮的那邊**：

- 距離變小（少 1）
- 但 `min()` 有可能變大（找到更高的柱子）
- 結果**有機會更好** → 這才是有價值的嘗試

---

## 🧠 第三部分：心法整理

### 1. 辨識題型的關鍵字

看到這些，先想 Two Pointers：

- 「**已排序**的陣列」+ 找兩個數的組合
- 「**最大/最小的配對**」
- 「**原地（in-place）**移除/修改元素」
- 「**回文（palindrome）**判斷」

### 2. 對撞指針的萬用模板

```go
left, right := 0, len(arr)-1

for left < right {
    // 根據條件決定移動哪邊
    if /* 需要更大的值 */ {
        left++
    } else if /* 需要更小的值 */ {
        right--
    } else {
        // 找到答案
    }
}
```

### 3. 核心問題永遠是：「為什麼移動這邊不會錯過答案？」

每次面試或解題時，能解釋清楚**為什麼某一邊可以安全移動**，就代表你真正理解了這題。

### 4. 已排序是前提

對撞指針大多數時候需要陣列是排序過的。如果題目沒排序，先想「要不要先 sort？」。排序 O(n log n) + 雙指針 O(n) 通常還是比暴力 O(n²) 快。

### 5. 同向指針的直覺

如果兩個指針都往同一方向，通常一個是 **slow（慢指針）**，一個是 **fast（快指針）**：

- fast 負責探索新元素
- slow 負責記錄「合法位置」
- 經典題：Remove Duplicates from Sorted Array

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 模式 |
|------|------|------|
| Easy | LeetCode 167. Two Sum II | 對撞指針 |
| Easy | LeetCode 344. Reverse String | 對撞指針 |
| Easy | LeetCode 26. Remove Duplicates | 同向指針 |
| Medium | LeetCode 11. Container With Most Water | 對撞指針 |
| Medium | LeetCode 15. 3Sum | 對撞指針 + 排序 |
| Medium | LeetCode 283. Move Zeroes | 同向指針 |