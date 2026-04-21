# Hash Map（雜湊表技巧）完整教學

---

## 📖 第一部分：什麼是 Hash Map？

### 生活化比喻

想像你在圖書館找《哈利波特》：

- **沒索引**：一本一本翻書架，最差要翻遍整個圖書館 → O(n)
- **有索引卡**：直接查「H 開頭 → 第 3 排第 5 格」 → O(1) 一步到位

Hash Map 就是那張**索引卡**——給你一個 key，瞬間告訴你對應的 value 在哪裡。

### 為什麼要用 Hash Map？

**核心目的：用 O(1) 查詢時間換掉 O(n) 線性搜尋**

| 操作 | Array | Hash Map |
|------|-------|----------|
| 查詢 key | O(n) | **O(1)** |
| 檢查是否存在 | O(n) | **O(1)** |
| 計數 | O(n²) 要雙層迴圈 | **O(n)** 一次遍歷 |

**代價**：額外 O(n) 空間 —— 這就是所謂「**空間換時間**」。

### Hash Map 的三大用途

| 用途 | 代表題 |
|------|--------|
| **計數** | 字元頻率、元素出現次數 |
| **配對查找** | Two Sum、Contains Duplicate |
| **去重** | Longest Substring Without Repeating |

### Go 中的 Hash Map

```go
// 宣告
m := make(map[string]int)
m := map[string]int{"a": 1, "b": 2}

// 寫入
m["key"] = 10

// 讀取（推薦用兩值寫法檢查存在性）
val, ok := m["key"]
if ok {
    // key 存在
}

// 刪除
delete(m, "key")

// 遍歷
for k, v := range m {
    fmt.Println(k, v)
}
```

**注意**：Go map 不保證遍歷順序。需要有序時用 slice 排序或用 `sort.Strings` 排 key。

---

## 🎯 第二部分：例題實戰

---

## 例題一：Two Sum（LeetCode 1）— 最經典的 Hash Map 題

**題目**：給一個陣列 `nums` 和目標值 `target`，找兩個元素加起來等於 target，回傳它們的索引。

```
nums = [2, 7, 11, 15], target = 9
答案：[0, 1]（因為 2 + 7 = 9）
```

### 🐢 暴力解法（O(n²)）

```go
func twoSumBrute(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}
```

### 🚀 Hash Map 解法（O(n)）

```go
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int) // key: 值, value: 索引

    for i, num := range nums {
        complement := target - num
        if idx, ok := seen[complement]; ok {
            return []int{idx, i}
        }
        seen[num] = i
    }

    return nil
}
```

### 🔑 核心技巧：一邊遍歷一邊查

對每個 `num`，**問過去有沒有看過 `target - num`**，有就直接回傳。沒有就把自己存起來給後面的人查。

### 📊 圖解步驟

`nums = [2, 7, 11, 15], target = 9`

```
i=0: num=2, complement=7
     seen={} 沒有 7
     存入 seen[2]=0
     seen = {2:0}

i=1: num=7, complement=2
     seen={2:0} 有 2！⭐
     回傳 [0, 1] ✅
```

只走了 2 步就找到，暴力解要跑 7 次比較。

---

## 例題二：字母異位詞分組（LeetCode 49）— Hash Map 分類

**題目**：給一個字串陣列 `strs`，把**字母異位詞**（anagram，組成字母一樣只是順序不同）分成同一組。

```
strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
答案：[["eat","tea","ate"], ["tan","nat"], ["bat"]]
```

### 🤔 思路

**把每個字串排序當 key**，異位詞排序後一定一樣，直接用 map 分類。

```
"eat" → 排序 → "aet"
"tea" → 排序 → "aet"  ← 同一組！
"ate" → 排序 → "aet"
```

### 🚀 Go 解法

```go
import "sort"

func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)

    for _, s := range strs {
        // 把字串拆成 []byte 排序再轉回 string 當 key
        bytes := []byte(s)
        sort.Slice(bytes, func(i, j int) bool {
            return bytes[i] < bytes[j]
        })
        key := string(bytes)

        groups[key] = append(groups[key], s)
    }

    result := [][]string{}
    for _, group := range groups {
        result = append(result, group)
    }
    return result
}
```

### 📊 圖解

```
"eat" → key="aet" → groups["aet"] = ["eat"]
"tea" → key="aet" → groups["aet"] = ["eat", "tea"]
"tan" → key="ant" → groups["ant"] = ["tan"]
"ate" → key="aet" → groups["aet"] = ["eat", "tea", "ate"]
"nat" → key="ant" → groups["ant"] = ["tan", "nat"]
"bat" → key="abt" → groups["abt"] = ["bat"]

最終 groups:
{
  "aet": ["eat", "tea", "ate"],
  "ant": ["tan", "nat"],
  "abt": ["bat"],
}
```

### 🔑 用字母計數當 key（更快的做法）

排序是 O(k log k)，k 是字串長度。如果只有小寫字母，可以用 26 個字母的計數陣列當 key：

```go
key := [26]int{}
for _, c := range s {
    key[c-'a']++
}
// 用 key 當 map 的 key（Go 陣列可以當 key）
```

這樣是 O(k)。

---

## 例題三：最長連續序列（LeetCode 128）— Hash Set 巧用

**題目**：給一個無序陣列 `nums`，找出**最長連續數字序列**的長度。要求 O(n)。

```
nums = [100, 4, 200, 1, 3, 2]
答案：4（連續序列 [1, 2, 3, 4]）
```

### 🤔 思路

先把所有數字丟進 HashSet，對每個數字嘗試向右擴展：`num, num+1, num+2, ...`。

**優化關鍵**：只從**序列起點**開始擴展（也就是 `num-1` 不在 set 裡的數字）。這樣每個數字最多被訪問 2 次，整體 O(n)。

### 🚀 Go 解法

```go
func longestConsecutive(nums []int) int {
    set := make(map[int]bool)
    for _, num := range nums {
        set[num] = true
    }

    longest := 0
    for num := range set {
        // 只從「序列起點」開始算
        if !set[num-1] {
            curr := num
            length := 1
            for set[curr+1] {
                curr++
                length++
            }
            if length > longest {
                longest = length
            }
        }
    }

    return longest
}
```

### 📊 圖解：`nums = [100, 4, 200, 1, 3, 2]`

```
set = {100, 4, 200, 1, 3, 2}

遍歷每個 num：
- 100: set 裡沒 99 → 是起點
       嘗試 101 不在 set → 長度 = 1
- 4:   set 裡有 3 → 不是起點，跳過
- 200: set 裡沒 199 → 是起點
       嘗試 201 不在 set → 長度 = 1
- 1:   set 裡沒 0 → 是起點 ⭐
       1 → 2 → 3 → 4 → 5（5 不在 set）
       長度 = 4
- 3:   set 裡有 2 → 不是起點，跳過
- 2:   set 裡有 1 → 不是起點，跳過

longest = 4 ✅
```

### 🔑 為什麼只從起點開始還是 O(n)？

每個數字最多被訪問：
- 一次當起點檢查 `num-1`
- 一次當序列中間點被擴展

總共 O(n)，不是 O(n²)。

---

## 🧠 第三部分：心法整理

### 1. 辨識題型的關鍵字

看到這些，先想 Hash Map：

- 「**找兩個數加起來等於 target**」
- 「**計算頻率/次數**」
- 「**是否存在**某個元素」
- 「**去重**」
- 「**分組/分類**」

### 2. Hash Map 的三大經典 pattern

**Pattern 1：計數**

```go
count := make(map[int]int)
for _, num := range nums {
    count[num]++
}
```

**Pattern 2：配對（target - x 查找）**

```go
seen := make(map[int]int)
for i, num := range nums {
    if idx, ok := seen[target-num]; ok {
        return []int{idx, i}
    }
    seen[num] = i
}
```

**Pattern 3：分組**

```go
groups := make(map[string][]string)
for _, s := range strs {
    key := computeKey(s)
    groups[key] = append(groups[key], s)
}
```

### 3. Go map 的存在性檢查

```go
// ✅ 推薦：兩值寫法
if val, ok := m[key]; ok {
    // 處理 val
}

// ⚠️ 注意：零值陷阱
// 如果 map value 型別是 int，不存在的 key 讀出來是 0
// 但 0 也可能是合法值，所以要用 ok 判斷
if m[key] != 0 { /* 錯：不能區分「沒這個 key」和「值就是 0」 */ }
```

### 4. Set 怎麼實作？

Go 沒內建 set，用 `map[T]bool` 或 `map[T]struct{}` 模擬：

```go
// 省空間版（struct{} 不佔記憶體）
set := make(map[int]struct{})
set[3] = struct{}{}
_, exists := set[3]
```

面試一般用 `map[T]bool` 就好，寫起來比較順。

### 5. Hash Map 的三個陷阱

- **順序不保證**：每次 `for range` 遍歷順序可能不同
- **並發不安全**：多 goroutine 同時讀寫會 panic，需要 `sync.Map` 或加鎖
- **不能拿地址**：`&m[key]` 是非法的，因為 rehash 會讓地址失效

### 6. 什麼時候不該用 Hash Map？

- **資料範圍小且連續**（eg. 0~100）→ 用陣列更快更省
- **需要按順序**輸出 → 用排序後的 slice
- **只查一次**而已 → 直接線性掃就好，不用建 map

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | Hash Map 用法 |
|------|------|---------------|
| Easy | LeetCode 1. Two Sum | 配對查找 |
| Easy | LeetCode 217. Contains Duplicate | Set 去重 |
| Easy | LeetCode 242. Valid Anagram | 字元計數 |
| Easy | LeetCode 169. Majority Element | 計數 |
| Medium | LeetCode 49. Group Anagrams | 分組 |
| Medium | LeetCode 128. Longest Consecutive Sequence | Set + 起點判斷 |
| Medium | LeetCode 347. Top K Frequent Elements | 計數 + Heap |
| Medium | LeetCode 560. Subarray Sum Equals K | Prefix Sum + Map |
| Medium | LeetCode 36. Valid Sudoku | 多 Set 檢查 |
| Medium | LeetCode 146. LRU Cache | Map + 雙向鏈結串列 |
