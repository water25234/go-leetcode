# Monotonic Stack（單調堆疊）完整教學

---

## 📖 第一部分：什麼是 Monotonic Stack？

### 生活化比喻

想像你在**排隊買票**，每個人背上貼了身高數字。現在想知道「**每個人後面，第一個比自己高的人在哪**」。

- 笨方法：對每個人，往後一個一個看 → O(n²)
- 聰明方法：手上拿一疊「還沒找到答案」的人的清單
  - 新人來了，如果比清單頂的人高 → 那個人找到答案了，從清單踢掉
  - 一直踢到清單頂比新人高，再把新人放上清單

這就是 Monotonic Stack —— **維護一個單調（遞增或遞減）的堆疊**，專門解決「**下一個更大/更小**」這類問題。

### 為什麼要用 Monotonic Stack？

**核心目的：把「下一個更大/更小」類型題從 O(n²) → O(n)**

關鍵觀察：**每個元素最多被進 stack 一次、出 stack 一次**，總操作數 O(n)。

### Monotonic Stack 的兩種類型

| 類型 | stack 內容 | 用途 |
|------|-----------|------|
| **單調遞增 stack** | 底 → 頂 元素遞增 | 找「前一個更小」「下一個更小」 |
| **單調遞減 stack** | 底 → 頂 元素遞減 | 找「前一個更大」「下一個更大」 |

**記憶口訣**：要找「**更大**」用**遞減** stack；要找「**更小**」用**遞增** stack。（因為新元素要能打破 stack 頂的單調性才觸發彈出）

### Monotonic Stack 的萬用模板

找「**下一個更大元素**」：

```go
func nextGreater(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    for i := range result {
        result[i] = -1 // 預設沒有
    }

    stack := []int{} // 存索引

    for i := 0; i < n; i++ {
        // 當前元素能打破 stack 頂的單調性 → 彈出並結算
        for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result[top] = nums[i]
        }
        stack = append(stack, i)
    }

    return result
}
```

---

## 🎯 第二部分：例題實戰

---

## 例題一：下一個更大的元素（LeetCode 496）— 入門必會

**題目**：給兩個陣列 `nums1`（是 `nums2` 的子集），對 `nums1` 中每個元素，找出它在 `nums2` 中**下一個比它大的元素**。沒有則回傳 -1。

```
nums1 = [4, 1, 2]
nums2 = [1, 3, 4, 2]

答案：[-1, 3, -1]
（4 在 nums2 後面沒更大的；1 後面是 3；2 後面沒更大的）
```

### 🚀 Go 解法

```go
func nextGreaterElement(nums1, nums2 []int) []int {
    // Step 1: 對 nums2 跑單調 stack，算出每個元素的 next greater
    nextGreater := make(map[int]int)
    stack := []int{}

    for _, num := range nums2 {
        for len(stack) > 0 && stack[len(stack)-1] < num {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            nextGreater[top] = num
        }
        stack = append(stack, num)
    }

    // Step 2: 對 nums1 查表
    result := make([]int, len(nums1))
    for i, num := range nums1 {
        if g, ok := nextGreater[num]; ok {
            result[i] = g
        } else {
            result[i] = -1
        }
    }
    return result
}
```

### 📊 圖解 Monotonic Stack 過程（nums2 = [1, 3, 4, 2]）

```
初始: stack = []

i=0, num=1
  stack 空，直接 push
  stack = [1]

i=1, num=3
  stack 頂 1 < 3 → 彈出 1，記 nextGreater[1] = 3
  stack = []，push 3
  stack = [3]

i=2, num=4
  stack 頂 3 < 4 → 彈出 3，記 nextGreater[3] = 4
  stack = []，push 4
  stack = [4]

i=3, num=2
  stack 頂 4 > 2 → 不彈，直接 push
  stack = [4, 2]

結束時 stack = [4, 2]，這兩個沒有 next greater
nextGreater = {1: 3, 3: 4}
```

### 🔑 核心直覺

**stack 裡存的都是「還在等待被超越的元素」**。新元素來了：

- 比 stack 頂大 → stack 頂找到答案了，彈出
- 比 stack 頂小 → 自己也進去等

這樣 stack 永遠保持**從底到頂遞減**（因為比頂小的才能進來），這就是「單調遞減 stack」。

---

## 例題二：每日溫度（LeetCode 739）— 經典應用

**題目**：給一個每日溫度陣列 `temperatures`，回傳每一天「**還要等幾天才會遇到更暖的一天**」。沒有則是 0。

```
temperatures = [73, 74, 75, 71, 69, 72, 76, 73]
答案:          [ 1,  1,  4,  2,  1,  1,  0,  0]
```

### 🤔 思路

跟例題一幾乎一樣，只是回傳「**索引差**」而不是「**值**」。所以 stack 存**索引**就能算出差。

### 🚀 Go 解法

```go
func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    result := make([]int, n) // 預設都是 0
    stack := []int{}         // 存索引

    for i := 0; i < n; i++ {
        for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result[top] = i - top // 索引差就是天數差
        }
        stack = append(stack, i)
    }

    return result
}
```

### 📊 圖解（`temperatures = [73, 74, 75, 71, 69, 72, 76, 73]`）

```
i=0, t=73: stack=[0]
i=1, t=74: 彈 0（temp[0]=73 < 74）result[0]=1-0=1, stack=[1]
i=2, t=75: 彈 1（temp[1]=74 < 75）result[1]=2-1=1, stack=[2]
i=3, t=71: 71 < 75 不彈，push。 stack=[2, 3]
i=4, t=69: 69 < 71 不彈，push。 stack=[2, 3, 4]
i=5, t=72:
          彈 4（69 < 72）result[4]=5-4=1
          彈 3（71 < 72）result[3]=5-3=2
          72 < 75 不彈，push。stack=[2, 5]
i=6, t=76:
          彈 5（72 < 76）result[5]=6-5=1
          彈 2（75 < 76）result[2]=6-2=4
          stack=[6]
i=7, t=73: 73 < 76 不彈，push。stack=[6, 7]

result = [1, 1, 4, 2, 1, 1, 0, 0] ✅
（索引 6 和 7 沒出 stack，result 保持 0）
```

### 🔑 小觀察：所有索引都只進 stack 一次、出 stack 至多一次

這就是為什麼 Monotonic Stack 整體是 O(n)。雖然內層有 `for`，但均攤分析每個元素只做常數次操作。

---

## 例題三：柱狀圖中最大矩形（LeetCode 84）— 進階挑戰

**題目**：給一排柱子的高度，每根寬 1，求**能圍出的最大矩形面積**。

```
heights = [2, 1, 5, 6, 2, 3]
答案：10（由 heights[2]=5 和 heights[3]=6 組成，高 5 寬 2）
```

### 🤔 思路

對每根柱子 `i`，問：「以 `heights[i]` 當高的最大矩形寬度是多少？」
答：要找到**左邊第一個比它矮的**和**右邊第一個比它矮的**，中間就是寬度範圍。

這又是「下一個更小」「前一個更小」的題 → 用**單調遞增** stack。

### 🚀 Go 解法

```go
func largestRectangleArea(heights []int) int {
    // 在兩端加 0，省掉最後要清空 stack 的特判
    heights = append([]int{0}, heights...)
    heights = append(heights, 0)

    stack := []int{} // 存索引，對應的 heights 是遞增的
    maxArea := 0

    for i := 0; i < len(heights); i++ {
        for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            // 以 heights[top] 為高，左邊界是新 stack 頂，右邊界是 i
            width := i - stack[len(stack)-1] - 1
            area := heights[top] * width
            if area > maxArea {
                maxArea = area
            }
        }
        stack = append(stack, i)
    }

    return maxArea
}
```

### 🔑 為什麼兩端加 0？

- **左邊加 0**：保證 stack 永遠有東西，算 width 時不用特判空 stack
- **右邊加 0**：遍歷結束時把 stack 裡所有柱子都擠出來結算（不然會漏算沒彈出的）

這是一個超實用的技巧，很多單調 stack 題都能用。

### 📊 關鍵一步圖解

```
heights (加了兩個 0) = [0, 2, 1, 5, 6, 2, 3, 0]
                        0  1  2  3  4  5  6  7

走到 i=5 (h=2) 時，stack = [0, 1, 3, 4]（對應 heights 0, 2, 5, 6 遞增）

h=2 比 stack 頂 heights[4]=6 小 → 彈出 4
  width = i(5) - stack.top(3) - 1 = 1
  area = 6 × 1 = 6

h=2 比新 stack 頂 heights[3]=5 小 → 彈出 3
  width = i(5) - stack.top(1) - 1 = 3  ← 就是索引 2~4 這段
  area = 5 × 3 = 15... 

誒等等，答案是 10 不是 15？讓我重新看題目...
heights = [2, 1, 5, 6, 2, 3]，高 5 寬 2 應該 = 10，不是 15。

實際上加 0 後 heights = [0, 2, 1, 5, 6, 2, 3, 0]
索引 2~4 的 heights 是 [1, 5, 6]，1 比 5 小，不能算在 5 的範圍內。

我上面計算有誤，重新走一次：
走到 i=2 (h=1) 時 stack=[0, 1]（對應 heights 0, 2）
  h=1 < heights[1]=2 → 彈 1，width = 2-0-1 = 1，area = 2×1 = 2
  h=1 > heights[0]=0，push，stack=[0, 2]

走到 i=3 (h=5), push, stack=[0, 2, 3]
走到 i=4 (h=6), push, stack=[0, 2, 3, 4]
走到 i=5 (h=2):
  h=2 < heights[4]=6 → 彈 4, width = 5-3-1 = 1, area = 6
  h=2 < heights[3]=5 → 彈 3, width = 5-2-1 = 2, area = 10  ⭐
  h=2 > heights[2]=1, push
  stack = [0, 2, 5]

繼續下去，最大仍是 10 ✅
```

（這題思路比較繞，上面圖解有小插曲但說明了「要算寬度必須用 stack 底下的邊界，不能用彈出的元素本身」這個重要細節）

---

## 🧠 第三部分：心法整理

### 1. 辨識題型的關鍵字

看到這些，立刻想 Monotonic Stack：

- 「**下一個更大/更小**的元素」
- 「**左邊/右邊第一個比我大/小**的」
- 「**最大矩形**/最大面積」類型
- 「**找某個範圍內的最值**」

### 2. 四種問題對應的 stack 類型

| 問題 | stack 類型 | stack 頂彈出條件 |
|------|-----------|---------------|
| 下一個**更大** | 單調**遞減** | `nums[top] < nums[i]` |
| 下一個**更小** | 單調**遞增** | `nums[top] > nums[i]` |
| 前一個**更大** | 單調**遞減** | 進 stack 時，stack 頂就是答案 |
| 前一個**更小** | 單調**遞增** | 進 stack 時，stack 頂就是答案 |

### 3. Monotonic Stack 的萬用骨架

```go
stack := []int{}
for i := 0; i < n; i++ {
    for len(stack) > 0 && /* 違反單調性 */ {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // 結算：top 找到答案了
    }
    stack = append(stack, i)
}
// 若需要處理剩餘 stack 元素，在這裡再跑一次
```

### 4. stack 存「索引」還是「值」？

- **需要算距離** → 存**索引**（eg. 每日溫度）
- **只要值** → 存**值**就好（eg. next greater element）

不確定時**存索引最安全**，因為索引能反查值，反過來不行。

### 5. 「加哨兵」技巧

在陣列兩端加極值（0 或 INT_MAX/INT_MIN），可以省掉「遍歷結束時清空 stack」的特判。看起來小技巧，實際上能讓程式碼短 1/3。

### 6. 陣列是「環形」怎麼辦？

把陣列**複製一份接在後面**跑一次 Monotonic Stack 就好（eg. LeetCode 503）：

```go
for i := 0; i < 2*n; i++ {
    num := nums[i%n]
    // ...
}
```

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 類型 |
|------|------|------|
| Easy | LeetCode 496. Next Greater Element I | 基本款 |
| Medium | LeetCode 739. Daily Temperatures | 下一個更大 + 索引差 |
| Medium | LeetCode 503. Next Greater Element II | 環形陣列 |
| Medium | LeetCode 901. Online Stock Span | 前一個更大 |
| Medium | LeetCode 1019. Next Greater Node In Linked List | 鏈結串列版 |
| Hard | LeetCode 84. Largest Rectangle in Histogram | 最大矩形 |
| Hard | LeetCode 85. Maximal Rectangle | 二維最大矩形 |
| Hard | LeetCode 42. Trapping Rain Water | 接雨水（也可雙指針）|
