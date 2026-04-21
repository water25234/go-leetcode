# Heap / Priority Queue（堆積 / 優先佇列）完整教學

---

## 📖 第一部分：什麼是 Heap？

### 生活化比喻

想像**急診室**的候診順序：

- 一般排隊（Queue）：先到先看診，照順序
- 急診室（Priority Queue）：**病情最嚴重的先看**，其他人往後排

每次來新病人，都要**立刻知道誰最嚴重**，並在看完他後，從剩下的人裡**再找出最嚴重的**。

這就是 Heap —— **隨時能取最大/最小值的資料結構**。

### Heap 的核心特性

| 操作 | 時間複雜度 | 說明 |
|------|----------|------|
| **Push（插入）** | O(log n) | 加一個元素 |
| **Pop（取走最值）** | O(log n) | 取走並刪除最大/最小值 |
| **Peek（看最值）** | O(1) | 只看不取 |
| **建堆（一次性）** | O(n) | 把陣列變成 heap |

**對比**：如果用排序陣列，Pop 是 O(1) 但 Push 是 O(n)。Heap 是在這兩個操作間取平衡。

### Heap 的底層結構

**Heap 是一棵「完全二元樹」**，用陣列儲存（不用指標）：

```
索引:   0   1   2   3   4   5
       [ 1,  3,  5,  7,  9, 10 ]

對應的樹：
         1 (0)
        / \
      3    5   (1, 2)
     / \   /
    7   9 10  (3, 4, 5)

父子關係：
- parent(i) = (i-1) / 2
- left(i)   = 2i + 1
- right(i)  = 2i + 2
```

**最小堆性質**：每個父節點 ≤ 它的所有子節點（根是最小值）
**最大堆性質**：每個父節點 ≥ 它的所有子節點（根是最大值）

### 什麼時候用 Heap？

| 情境 | 代表題 |
|------|--------|
| **Top K 問題**（最大/最小 K 個） | LeetCode 215, 347 |
| **合併 K 個排序串列** | LeetCode 23 |
| **資料流的中位數** | LeetCode 295 |
| **排程類問題**（會議、任務） | LeetCode 253 |

---

## 🎯 第二部分：例題實戰

---

## Go 的 Heap 怎麼用？

Go 標準庫有 `container/heap`，**需要自己實作五個方法**來滿足 `heap.Interface`：

```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
    Push(x any)   // 注意：這是「加到底層陣列」不是「正常 push」
    Pop() any     // 注意：這是「從底層陣列尾拔」不是「正常 pop」
}
```

**呼叫時用 `heap.Push` 和 `heap.Pop`**，不要呼叫自定義的 Push/Pop。

### 最小堆範本

```go
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] } // 最小堆
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

// 使用
func main() {
    h := &IntHeap{}
    heap.Init(h)
    heap.Push(h, 3)
    heap.Push(h, 1)
    heap.Push(h, 4)
    fmt.Println((*h)[0])      // 1（最小值，不移除）
    fmt.Println(heap.Pop(h))  // 1（移除最小值）
}
```

**最大堆**：只要把 `Less` 的 `<` 改成 `>`。

---

## 例題一：陣列中第 K 個最大元素（LeetCode 215）

**題目**：找陣列中第 K 個最大的元素。

```
nums = [3, 2, 1, 5, 6, 4], k = 2 → 答案：5
```

### 🤔 思路：用**大小為 K 的最小堆**

關鍵觀察：我們只需要「最大的 K 個」中**最小的那個**，就是第 K 大。

- 維護一個大小為 K 的最小堆
- 遍歷陣列，推進堆
- 當堆的大小超過 K，彈出最小值（丟掉太小的）
- 遍歷完，堆頂就是第 K 大

### 🚀 Go 解法

```go
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 最小堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func findKthLargest(nums []int, k int) int {
    h := &IntHeap{}
    heap.Init(h)

    for _, num := range nums {
        heap.Push(h, num)
        if h.Len() > k {
            heap.Pop(h) // 丟掉最小的
        }
    }

    return (*h)[0] // 堆頂就是第 K 大
}
```

### 📊 圖解（`nums = [3, 2, 1, 5, 6, 4], k = 2`）

```
處理 3: heap=[3]
處理 2: heap=[2, 3]（size=2，滿）
處理 1: heap=[1, 2, 3] size>2，彈最小 1 → heap=[2, 3]
處理 5: heap=[2, 3, 5] size>2，彈最小 2 → heap=[3, 5]
處理 6: heap=[3, 5, 6] size>2，彈最小 3 → heap=[5, 6]
處理 4: heap=[4, 5, 6] size>2，彈最小 4 → heap=[5, 6]

結束，堆頂 5 就是第 2 大 ✅
```

### 🔑 時間複雜度分析

- 遍歷 O(n)，每次堆操作 O(log k)
- 總共 **O(n log k)**
- 比整體排序的 O(n log n) 好（當 k << n 時）

---

## 例題二：前 K 個高頻元素（LeetCode 347）

**題目**：給一個整數陣列，回傳出現**頻率最高的 k 個**元素。

```
nums = [1, 1, 1, 2, 2, 3], k = 2 → 答案：[1, 2]
```

### 🤔 思路

1. **先用 HashMap 計數**：`{1:3, 2:2, 3:1}`
2. **用最小堆保留頻率最高的 k 個**（跟上題一樣的 pattern）

### 🚀 Go 解法

```go
import "container/heap"

// heap 的元素是 [頻率, 值]
type FreqHeap [][2]int

func (h FreqHeap) Len() int            { return len(h) }
func (h FreqHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] } // 按頻率最小堆
func (h FreqHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *FreqHeap) Push(x any)         { *h = append(*h, x.([2]int)) }
func (h *FreqHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }

    h := &FreqHeap{}
    heap.Init(h)

    for num, f := range freq {
        heap.Push(h, [2]int{f, num})
        if h.Len() > k {
            heap.Pop(h)
        }
    }

    result := make([]int, k)
    for i := k - 1; i >= 0; i-- {
        result[i] = heap.Pop(h).([2]int)[1]
    }
    return result
}
```

---

## 例題三：資料流的中位數（LeetCode 295）— 雙堆技巧

**題目**：資料不斷流進來，每次要能**即時回傳目前的中位數**。

### 🤔 思路：**兩個堆**把資料切兩半

- **最大堆 `left`**：存**較小的一半**，堆頂是左半最大值
- **最小堆 `right`**：存**較大的一半**，堆頂是右半最小值
- 維持兩堆大小平衡（差 ≤ 1）
- 中位數：偶數個時是兩堆堆頂平均；奇數個時是較大那堆的堆頂

```
左半（最大堆）    右半（最小堆）
   5                 6
  / \               / \
 3   4             7   8
                          ↑
中位數 = (5 + 6) / 2 = 5.5
```

### 🚀 Go 解法

```go
import "container/heap"

// 最大堆和最小堆的實作省略，參考例題一範本
// MaxHeap: Less 用 >
// MinHeap: Less 用 <

type MedianFinder struct {
    left  *MaxHeap // 存較小一半
    right *MinHeap // 存較大一半
}

func Constructor() MedianFinder {
    return MedianFinder{
        left:  &MaxHeap{},
        right: &MinHeap{},
    }
}

func (m *MedianFinder) AddNum(num int) {
    if m.left.Len() == 0 || num <= (*m.left)[0] {
        heap.Push(m.left, num)
    } else {
        heap.Push(m.right, num)
    }

    // 平衡兩堆大小
    if m.left.Len() > m.right.Len()+1 {
        heap.Push(m.right, heap.Pop(m.left))
    } else if m.right.Len() > m.left.Len() {
        heap.Push(m.left, heap.Pop(m.right))
    }
}

func (m *MedianFinder) FindMedian() float64 {
    if m.left.Len() > m.right.Len() {
        return float64((*m.left)[0])
    }
    return (float64((*m.left)[0]) + float64((*m.right)[0])) / 2
}
```

### 📊 圖解

```
加入 1: left=[1], right=[]        中位數 = 1
加入 2: 2 > 1 進 right. left=[1], right=[2]
        size 平衡，中位數 = (1+2)/2 = 1.5
加入 3: 3 > 1 進 right. left=[1], right=[2,3]
        right 多 1，把 2 搬到 left
        left=[2,1], right=[3]    中位數 = 2
加入 4: 4 > 2 進 right. left=[2,1], right=[3,4]
        中位數 = (2+3)/2 = 2.5
```

---

## 🧠 第三部分：心法整理

### 1. 辨識題型的關鍵字

看到這些，立刻想 Heap：

- 「**第 K 大 / 第 K 小**」
- 「**前 K 個**頻率最高/最近/最貴的」
- 「**合併 K 個**排序串列」
- 「**資料流**的中位數/最值」
- 「**最短任務優先**」「**最早截止**優先」的排程

### 2. Top K 的黃金法則

| 目標 | 用什麼堆 |
|------|---------|
| 最大 K 個 | **最小堆**（大小 K，彈小留大） |
| 最小 K 個 | **最大堆**（大小 K，彈大留小） |

**記憶技巧**：你要留誰，就用「相反的堆」。

### 3. Heap 在 Go 的痛點

- **必須實作 5 個 method**（Len, Less, Swap, Push, Pop）
- Push/Pop 實際上是**內部操作**，使用時要呼叫 `heap.Push/heap.Pop`，**不能直接呼叫**結構體的 Push/Pop
- 堆頂是 `(*h)[0]`，不是 `h[0]`（因為是指標）

這些細節面試容易寫錯，**先背範本**，考試前過一遍。

### 4. Heap 的時間複雜度要記熟

```
Push:    O(log n)
Pop:     O(log n)
Peek:    O(1)
Heapify: O(n)    ← 整個陣列建堆，比一個一個 push (O(n log n)) 快
```

### 5. 什麼時候 Heap 不是最佳解？

- **整個都要排序** → 直接 `sort.Ints`（O(n log n) 就好）
- **k 很大（接近 n）** → 排序更乾淨
- **元素範圍小** → 用計數排序 / 桶排序可能更快

Heap 專攻「**只要最值，不用全排序**」。

### 6. 雙堆技巧的變形

雙堆不只用在中位數：

- **IPO（LeetCode 502）**：最小堆裝成本、最大堆裝利潤
- **會議室（LeetCode 253）**：最小堆按結束時間
- **任務排程**：最小堆按到期時間

**一邊一個堆**是很常見的套路。

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 技巧 |
|------|------|------|
| Easy | LeetCode 703. Kth Largest Element in a Stream | 基本 Top K |
| Easy | LeetCode 1046. Last Stone Weight | 最大堆模擬 |
| Medium | LeetCode 215. Kth Largest Element in Array | 最小堆留 K 大 |
| Medium | LeetCode 347. Top K Frequent Elements | 計數 + 堆 |
| Medium | LeetCode 692. Top K Frequent Words | 計數 + 堆 + 自訂比較 |
| Medium | LeetCode 973. K Closest Points to Origin | 最大堆留 K 近 |
| Medium | LeetCode 253. Meeting Rooms II | 最小堆 |
| Hard | LeetCode 23. Merge K Sorted Lists | K 路合併 |
| Hard | LeetCode 295. Find Median from Data Stream | 雙堆 |
| Hard | LeetCode 502. IPO | 雙堆 |
