# Backtracking（回溯法）完整教學

---

## 📖 第一部分：什麼是 Backtracking？

### 生活化比喻

想像你在走**迷宮**，手上拿一張地圖，每走到叉路口就**記下這個選擇**。走不通時：

1. **退一步**到上一個叉路口
2. **擦掉剛剛的紀錄**
3. 嘗試**另一條岔路**

直到走通或所有路都試過。

這就是 Backtracking —— **「選擇 → 遞迴 → 撤銷選擇」** 的窮舉法。

### Backtracking 的本質

**Backtracking ≈ DFS + 選擇撤銷**

跟一般 DFS 的關鍵差別：

- 一般 DFS：只要「標記走過」避免重訪
- Backtracking：每條路徑都是獨立的，離開節點時要**撤銷標記**，讓其他路徑能重新走這個節點

### 為什麼要 Backtracking？

**核心目的：列舉所有滿足條件的組合/排列/路徑**

很多題目需要「**所有可能解**」，不是單一最優解——這種「窮舉型」題目就是 Backtracking 的主場。

時間複雜度通常是指數級（O(2ⁿ)、O(n!)），但加上「**剪枝**」（提前排除不可能的分支）可以大幅加速。

### Backtracking 的四大題型

| 題型 | 代表題 | 特徵 |
|------|--------|------|
| **子集** | LeetCode 78 | 每個元素選/不選 |
| **組合** | LeetCode 77 | 選 k 個，不分順序 |
| **排列** | LeetCode 46 | 選 n 個，分順序 |
| **切割/拼湊** | LeetCode 131 | 字串分段 |

### Backtracking 的萬用模板

```go
var result [][]int
var path []int

func backtrack(選擇列表) {
    if 滿足終止條件 {
        result = append(result, append([]int{}, path...)) // 注意要複製！
        return
    }

    for _, 選擇 := range 選擇列表 {
        if 選擇不合法 {
            continue // 剪枝
        }
        path = append(path, 選擇)  // 做選擇
        backtrack(新的選擇列表)    // 遞迴
        path = path[:len(path)-1]  // 撤銷選擇
    }
}
```

---

## 🎯 第二部分：例題實戰

---

## 例題一：子集（LeetCode 78）— Backtracking 入門

**題目**：給一個**不含重複元素**的陣列 `nums`，回傳所有可能的子集。

```
nums = [1, 2, 3]
答案：[[], [1], [2], [3], [1,2], [1,3], [2,3], [1,2,3]]
```

### 🤔 思路

對每個位置做「**選 or 不選**」的決策，樹狀展開所有組合。

```
              []
           /      \
         [1]       []       (第 1 個元素：選 or 不選)
        /   \     /   \
     [1,2] [1]  [2]   []    (第 2 個元素：選 or 不選)
     ...
```

### 🚀 Go 解法（Start Index 寫法）

```go
func subsets(nums []int) [][]int {
    result := [][]int{}
    path := []int{}

    var backtrack func(start int)
    backtrack = func(start int) {
        // 每個節點都是一個子集（包含中間節點）
        result = append(result, append([]int{}, path...))

        for i := start; i < len(nums); i++ {
            path = append(path, nums[i])
            backtrack(i + 1)
            path = path[:len(path)-1]
        }
    }

    backtrack(0)
    return result
}
```

### 🔑 關鍵：`start` 避免重複

`start` 決定「可以從哪個索引開始選」，每次遞迴 `start = i+1`，確保每個元素只被選一次、且後面的組合不重複。

### 📊 圖解（`nums = [1, 2, 3]`）

```
backtrack(0): 加入 [] 到 result
  選 1 (path=[1]):
    backtrack(1): 加入 [1]
      選 2 (path=[1,2]):
        backtrack(2): 加入 [1,2]
          選 3 (path=[1,2,3]):
            backtrack(3): 加入 [1,2,3]
          撤 3 (path=[1,2])
        撤 2 (path=[1])
      選 3 (path=[1,3]):
        backtrack(3): 加入 [1,3]
      撤 3 (path=[1])
    撤 1 (path=[])
  選 2 (path=[2]):
    backtrack(2): 加入 [2]
    ...
```

### 🚨 為什麼要 `append([]int{}, path...)` 而不是直接 `append(result, path)`？

`path` 是 slice，**底層是指標**。直接塞入 result 後，後續操作會影響已存入的資料。必須**複製一份**再存。這是 Backtracking 最常見的 bug。

---

## 例題二：排列（LeetCode 46）— 排列型

**題目**：給一個**不含重複元素**的陣列 `nums`，回傳所有全排列。

```
nums = [1, 2, 3]
答案：[[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
```

### 🤔 思路

排列跟組合不同：**順序有關**。所以不能用 `start` 限制，每次要從**所有沒用過的元素**選。

用一個 `used` 陣列記錄哪些元素已經用過。

### 🚀 Go 解法

```go
func permute(nums []int) [][]int {
    result := [][]int{}
    path := []int{}
    used := make([]bool, len(nums))

    var backtrack func()
    backtrack = func() {
        if len(path) == len(nums) {
            result = append(result, append([]int{}, path...))
            return
        }

        for i := 0; i < len(nums); i++ {
            if used[i] {
                continue // 已用過，跳過
            }
            used[i] = true
            path = append(path, nums[i])

            backtrack()

            path = path[:len(path)-1]
            used[i] = false
        }
    }

    backtrack()
    return result
}
```

### 📊 圖解（`nums = [1, 2, 3]` 部分展開）

```
backtrack: path=[]
  選 1, used=[T,F,F], path=[1]
    backtrack: path=[1]
      選 2, used=[T,T,F], path=[1,2]
        backtrack: path=[1,2]
          選 3, used=[T,T,T], path=[1,2,3]
            backtrack: len(path)==3 → 加入結果 [1,2,3] ✅
          撤 3
        撤 2
      選 3, used=[T,F,T], path=[1,3]
        backtrack:
          選 2, path=[1,3,2] → 加入結果 [1,3,2] ✅
        ...
    撤 1
  選 2, path=[2]
    ...
```

### 🔑 組合 vs 排列 的差別

```
組合（LeetCode 77）：選 2 個
  [1,2], [1,3], [2,3]        只有 3 種

排列（LeetCode 46）：選 2 個
  [1,2], [2,1],
  [1,3], [3,1],
  [2,3], [3,2]                有 6 種（順序不同算不同）
```

代碼差在：**組合用 `start`**，**排列用 `used`**。

---

## 例題三：N 皇后（LeetCode 51）— 經典剪枝

**題目**：在 `n×n` 棋盤上放 `n` 個皇后，使它們彼此**不會攻擊**（同行、同列、同斜線都不行）。回傳所有放法。

### 🚀 Go 解法

```go
func solveNQueens(n int) [][]string {
    result := [][]string{}
    board := make([][]byte, n)
    for i := range board {
        board[i] = make([]byte, n)
        for j := range board[i] {
            board[i][j] = '.'
        }
    }

    var backtrack func(row int)
    backtrack = func(row int) {
        if row == n {
            // 把 board 轉成 []string 存進 result
            snapshot := make([]string, n)
            for i := range board {
                snapshot[i] = string(board[i])
            }
            result = append(result, snapshot)
            return
        }

        for col := 0; col < n; col++ {
            if isValid(board, row, col, n) {
                board[row][col] = 'Q'
                backtrack(row + 1)
                board[row][col] = '.' // 撤銷
            }
        }
    }

    backtrack(0)
    return result
}

func isValid(board [][]byte, row, col, n int) bool {
    // 檢查同列
    for i := 0; i < row; i++ {
        if board[i][col] == 'Q' {
            return false
        }
    }
    // 檢查左上斜線
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    // 檢查右上斜線
    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    return true
}
```

### 🔑 剪枝就是效能關鍵

`isValid` 就是剪枝：放皇后前先檢查合法，不合法直接跳過。如果不剪枝，N 皇后會是 O(n^n) 暴力，剪枝後大概 O(n!)，差幾個數量級。

---

## 🧠 第三部分：心法整理

### 1. 辨識題型的關鍵字

看到這些，立刻想 Backtracking：

- 「**所有可能**的組合/排列/子集」
- 「**列出所有**滿足某條件的解」
- 「**劃分/切割**字串」
- 棋盤類（**N 皇后、數獨**）

### 2. Backtracking 三步口訣

```
做選擇 → 遞迴 → 撤銷選擇
```

少撤銷就會污染其他分支；少做選擇就會遞迴亂跑。這三步缺一不可。

### 3. 組合 vs 排列的代碼差異

| 題型 | 去重方式 | 
|------|---------|
| **組合** | `start` 索引往前走 |
| **排列** | `used[]` 陣列 |
| **有重複元素** | 先排序 + `i > start && nums[i] == nums[i-1]` 跳過 |

### 4. 剪枝能救效能

**三種常見剪枝**：

- **不可行剪枝**：當前路徑已違反條件 → 直接 return
- **不可能剪枝**：剩餘步數不可能達到目標 → 直接 return
- **重複剪枝**：跟同層前面的選擇一樣 → skip

剪枝寫得好，指數級變多項式級，面試加分。

### 5. 「複製 path」這個陷阱

```go
// ❌ 錯：後續修改 path 會改到 result
result = append(result, path)

// ✅ 對：複製一份
result = append(result, append([]int{}, path...))
```

寫 Backtracking **每次都要想**這件事，特別是 slice 型別。

### 6. 什麼時候 Backtracking 會 TLE？

- 指數級狀態空間很大時（n > 20）
- 如果題目要的是「**最優解**」（最小、最大、最長）→ 考慮 **DP**
- 如果題目要的是「**是否存在**」→ 考慮 BFS 找最短

Backtracking 專攻「**列舉所有解**」，其他需求用其他演算法。

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 類型 |
|------|------|------|
| Medium | LeetCode 78. Subsets | 子集 |
| Medium | LeetCode 90. Subsets II | 子集（含重複） |
| Medium | LeetCode 77. Combinations | 組合 |
| Medium | LeetCode 39. Combination Sum | 組合（可重選） |
| Medium | LeetCode 40. Combination Sum II | 組合（不可重選+去重） |
| Medium | LeetCode 46. Permutations | 排列 |
| Medium | LeetCode 47. Permutations II | 排列（含重複） |
| Medium | LeetCode 17. Letter Combinations of a Phone Number | 組合 |
| Medium | LeetCode 22. Generate Parentheses | 剪枝經典 |
| Medium | LeetCode 131. Palindrome Partitioning | 切割 |
| Medium | LeetCode 79. Word Search | 網格回溯 |
| Hard | LeetCode 51. N-Queens | 棋盤 + 剪枝 |
| Hard | LeetCode 37. Sudoku Solver | 二維回溯 |
