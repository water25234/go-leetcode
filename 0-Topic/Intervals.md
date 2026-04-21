# Intervals（區間問題）完整教學

---

## 📖 第一部分：什麼是 Intervals？

### 生活化比喻

想像你在看**今天的行事曆**：

- 9:00–10:30 開會
- 10:00–11:00 面試
- 14:00–15:00 午餐
- 14:30–16:00 健身

問題：

- 有沒有時間**衝突**？（9:00-10:30 跟 10:00-11:00 重疊）
- **合併**連續的忙碌時段變成一段？
- 如果要塞一個新會議，**應該插在哪**？

這些就是 Interval 問題 —— 處理「**一堆有起點終點的時間段**」的題型。

### Intervals 題的核心技巧

**99% 的 interval 題，第一步都是：先按起點（或終點）排序！**

排序後，複雜的兩兩比較變成**一次線性掃描**，O(n²) 降到 O(n log n)。

### 區間的表示方式

LeetCode 通常用 `[][]int{{start, end}}` 表示：

```go
intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
// 表示四個區間：[1,3], [2,6], [8,10], [15,18]
```

### Intervals 的四大題型

| 題型 | 代表題 | 關鍵技巧 |
|------|--------|---------|
| **合併重疊** | LeetCode 56 | 按起點排序 + 掃描合併 |
| **插入新區間** | LeetCode 57 | 三段處理：前、中、後 |
| **判斷是否重疊** | LeetCode 252 | 按起點排序 + 比對 |
| **最少需要幾個** | LeetCode 253 | 掃描線 / 最小堆 |

### 兩個區間重疊的判斷

```
A = [a1, a2]
B = [b1, b2]

重疊條件：a1 <= b2 && b1 <= a2
（反過來想：不重疊 = a2 < b1 || b2 < a1，取反即重疊）
```

---

## 🎯 第二部分：例題實戰

---

## 例題一：合併區間（LeetCode 56）— 最經典

**題目**：給一堆區間，合併所有重疊的區間。

```
intervals = [[1,3], [2,6], [8,10], [15,18]]
答案：[[1,6], [8,10], [15,18]]
（[1,3] 和 [2,6] 重疊，合併成 [1,6]）
```

### 🤔 思路

1. 按**起點**排序
2. 遍歷，維護「目前正在擴張的區間」
3. 新區間如果跟當前重疊 → 擴張終點
4. 不重疊 → 收掉當前，開始新的區間

### 🚀 Go 解法

```go
import "sort"

func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return nil
    }

    // 按起點排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    result := [][]int{intervals[0]}

    for i := 1; i < len(intervals); i++ {
        last := result[len(result)-1]
        curr := intervals[i]

        if curr[0] <= last[1] {
            // 重疊，擴張當前區間的終點
            if curr[1] > last[1] {
                last[1] = curr[1]
            }
        } else {
            // 不重疊，加入新區間
            result = append(result, curr)
        }
    }

    return result
}
```

### 📊 圖解

`intervals = [[1,3], [2,6], [8,10], [15,18]]`

```
排序後（本來就有序）
時間軸:  1 2 3 4 5 6 7 8 9 10 ... 15 16 17 18
[1,3]:   ====
[2,6]:     ======
[8,10]:              ====
[15,18]:                       =======

Step 1: result = [[1,3]]

Step 2: curr=[2,6], last=[1,3]
        2 <= 3 → 重疊！擴張 last[1] = max(3, 6) = 6
        result = [[1,6]]

Step 3: curr=[8,10], last=[1,6]
        8 > 6 → 不重疊，加入
        result = [[1,6], [8,10]]

Step 4: curr=[15,18], last=[8,10]
        15 > 10 → 不重疊，加入
        result = [[1,6], [8,10], [15,18]] ✅
```

### 🔑 注意 Go slice 的坑

```go
last := result[len(result)-1] // 這是「拷貝」嗎？
last[1] = curr[1]              // 這會改到 result 嗎？
```

**答**：`last` 是 slice，`result[len-1]` 也是 slice，底層指向同一個陣列。改 `last[1]` 會改到 `result` 裡。**這就是我們要的**（直接擴張）。寫錯方向反而會出 bug。

---

## 例題二：插入區間（LeetCode 57）— 三段法

**題目**：給一個**已排序且不重疊**的區間陣列，再給一個新區間，把新區間插入並合併。

```
intervals = [[1,3], [6,9]], newInterval = [2,5]
答案：[[1,5], [6,9]]
```

### 🤔 思路：三段式處理

把原區間分成三組：

1. **完全在新區間左邊**（end < newInterval.start）→ 原樣保留
2. **跟新區間重疊**（可能多個）→ 合併成一個大的
3. **完全在新區間右邊**（start > newInterval.end）→ 原樣保留

### 🚀 Go 解法

```go
func insert(intervals [][]int, newInterval []int) [][]int {
    result := [][]int{}
    i, n := 0, len(intervals)

    // 第一段：左邊不重疊的
    for i < n && intervals[i][1] < newInterval[0] {
        result = append(result, intervals[i])
        i++
    }

    // 第二段：重疊的，一路合併擴張 newInterval
    for i < n && intervals[i][0] <= newInterval[1] {
        if intervals[i][0] < newInterval[0] {
            newInterval[0] = intervals[i][0]
        }
        if intervals[i][1] > newInterval[1] {
            newInterval[1] = intervals[i][1]
        }
        i++
    }
    result = append(result, newInterval)

    // 第三段：右邊不重疊的
    for i < n {
        result = append(result, intervals[i])
        i++
    }

    return result
}
```

### 📊 圖解

`intervals = [[1,2], [3,5], [6,7], [8,10], [12,16]], newInterval = [4,8]`

```
時間軸:     1 2   3 4 5   6 7   8   10   12 ... 16
原區間:     ==    ====    ===   ==       ======
新區間:           ======================

第一段：[1,2] 完全在左邊 → 保留
第二段：重疊區 [3,5], [6,7], [8,10] 全部合併進 [4,8]
        new[0] = min(4, 3, 6, 8) = 3
        new[1] = max(8, 5, 7, 10) = 10
        插入 [3, 10]
第三段：[12,16] 完全在右邊 → 保留

結果: [[1,2], [3,10], [12,16]] ✅
```

---

## 例題三：會議室 II（LeetCode 253）— 最少需要幾個

**題目**：給一堆會議的時間 `[start, end]`，問**最少需要幾間會議室**？

```
intervals = [[0,30], [5,10], [15,20]]
答案：2（因為 [0,30] 整個時段橫跨兩個會議）
```

### 🤔 思路：最小堆按結束時間

1. 按**起點**排序所有會議
2. 用**最小堆**維護「**目前所有會議室的結束時間**」
3. 對每個新會議：
   - 如果有會議室的結束時間 ≤ 新會議開始時間 → **重用這間**（彈出舊的，推入新的結束時間）
   - 否則 → **開一間新的**
4. 最後堆的大小就是答案

### 🚀 Go 解法

```go
import (
    "container/heap"
    "sort"
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
    old := *h
    x := old[len(old)-1]
    *h = old[:len(old)-1]
    return x
}

func minMeetingRooms(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    h := &MinHeap{}
    heap.Init(h)
    heap.Push(h, intervals[0][1])

    for i := 1; i < len(intervals); i++ {
        // 如果最早結束的會議室已經空了，可以重用
        if intervals[i][0] >= (*h)[0] {
            heap.Pop(h)
        }
        heap.Push(h, intervals[i][1])
    }

    return h.Len()
}
```

### 📊 圖解

`intervals = [[0,30], [5,10], [15,20]]`（已排序）

```
會議 [0,30]: heap 為空，開新會議室
            heap = [30]

會議 [5,10]: 5 < 30，需要新會議室
            heap = [10, 30]

會議 [15,20]: 15 > 10，重用第一間會議室
             彈出 10
             heap = [20, 30]

heap 大小 = 2 ✅
```

### 🔑 另解：掃描線

把每個區間拆成「+1（進入）」和「-1（離開）」兩個事件，按時間排序後掃描：

```go
events := [][]int{}
for _, iv := range intervals {
    events = append(events, []int{iv[0], 1})
    events = append(events, []int{iv[1], -1})
}
sort.Slice(events, func(i, j int) bool {
    if events[i][0] == events[j][0] {
        return events[i][1] < events[j][1] // 同時間先離開再進入
    }
    return events[i][0] < events[j][0]
})

curr, max := 0, 0
for _, e := range events {
    curr += e[1]
    if curr > max {
        max = curr
    }
}
```

---

## 🧠 第三部分：心法整理

### 1. 辨識題型的關鍵字

看到這些，立刻想 Intervals：

- 「**合併重疊**」
- 「**插入新區間**」
- 「**有沒有衝突**」
- 「**最少需要幾個**會議室/機器/工人」
- 給一堆 `[start, end]` pair

### 2. Intervals 題的萬用第一步

**先排序！** 絕大多數情況按**起點**排序：

```go
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
})
```

少數情況（例如貪心類）要按**終點**排序（eg. LeetCode 435）。判斷方法：想清楚後續處理需要「從最早開始的」還是「最早結束的」出發。

### 3. 重疊判斷的兩個等價寫法

```go
// 寫法一（新舊起終點交錯）
overlap := a[0] <= b[1] && b[0] <= a[1]

// 寫法二（取否）
notOverlap := a[1] < b[0] || b[1] < a[0]

// 排序後新區間在舊的右側，簡化為
overlap := curr[0] <= last[1]
```

排序後第二種寫法超簡潔，這也是為什麼排序是萬用第一步。

### 4. 合併時要更新「終點」的最大值

```go
// ❌ 錯：直接覆寫
last[1] = curr[1]

// ✅ 對：取 max
if curr[1] > last[1] {
    last[1] = curr[1]
}
```

因為 curr 可能完全被 last 包住：`last=[1,10], curr=[2,5]`，擴張後還是 `[1,10]`。

### 5. 兩種「最少需要幾個」的解法選擇

| 解法 | 時間 | 適合情境 |
|------|------|---------|
| **最小堆** | O(n log n) | 通用，直觀 |
| **掃描線** | O(n log n) | 也很直觀，事件模式 |
| **分開排序起點終點陣列** | O(n log n) | 只需要「最大重疊數」，簡單快速 |

掃描線思路很適合「**某時刻同時有多少 X**」這種題，要會。

### 6. Interval 題常見陷阱

- **等號**：`[1,3]` 和 `[3,5]` 算重疊嗎？**看題目定義**，通常要看成「閉區間」就算重疊
- **點區間**：`[3,3]` 這種算不算？排序時要注意
- **空陣列**：很多題沒處理會 panic

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 技巧 |
|------|------|------|
| Easy | LeetCode 252. Meeting Rooms | 判斷重疊 |
| Medium | LeetCode 56. Merge Intervals | 合併 |
| Medium | LeetCode 57. Insert Interval | 三段法 |
| Medium | LeetCode 253. Meeting Rooms II | 最小堆 |
| Medium | LeetCode 435. Non-overlapping Intervals | 按終點排序的貪心 |
| Medium | LeetCode 452. Minimum Number of Arrows | 貪心區間 |
| Medium | LeetCode 986. Interval List Intersections | 雙指針 + 區間 |
| Hard | LeetCode 1851. Minimum Interval to Include Each Query | 堆 + 排序 |
