# BFS（廣度優先搜尋）完整教學

---

## 📖 第一部分：什麼是 BFS？

### 生活化比喻

想像你往**平靜湖面丟一顆石頭**：

- 水波從石頭落點向外擴散
- 先經過**第一圈**的位置（距離 1）
- 再到**第二圈**（距離 2）
- 再到**第三圈**（距離 3）
- ...一層一層向外，絕不跳級

這就是 BFS —— **從起點出發，一層一層往外探索**，所以第一次碰到目標時，**走的步數一定是最少的**。

### BFS 的核心特性

| 特性 | 說明 |
|------|------|
| **一層一層走** | 先走完距離 1 的所有節點，再走距離 2 |
| **找最短路徑** | 在「每一步代價都相同」的題目中，BFS 第一次到終點就是最短路徑 |
| **用 Queue（佇列）** | 先進先出，維持「誰先探索就先處理」 |
| **visited 集合** | 防止走回頭路造成無窮迴圈 |

### BFS vs DFS 的直覺差異

| | BFS | DFS |
|---|-----|-----|
| **策略** | 一層層向外擴散 | 一條路走到底才回頭 |
| **資料結構** | Queue（佇列） | Stack（堆疊）/ 遞迴 |
| **適合題型** | 最短路徑、最少步數、層級遍歷 | 路徑列舉、排列組合、連通塊 |
| **空間複雜度** | O(寬度) | O(深度) |

### BFS 的萬用模板

```go
func bfs(start Node) {
    queue := []Node{start}
    visited := map[Node]bool{start: true}

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:] // 取出隊首

        // 處理 curr
        for _, next := range curr.Neighbors() {
            if !visited[next] {
                visited[next] = true
                queue = append(queue, next)
            }
        }
    }
}
```

---

## 🎯 第二部分：例題實戰

---

## 例題一：二元樹層序遍歷（LeetCode 102）— BFS 入門必會

**題目**：給一棵二元樹，回傳**每一層**的節點值。

```
輸入：     3
         /   \
        9    20
             /  \
            15   7

輸出：[[3], [9, 20], [15, 7]]
```

### 🚀 Go 解法

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    result := [][]int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        levelSize := len(queue) // 記住「這一層」有幾個節點
        level := []int{}

        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]

            level = append(level, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, level)
    }

    return result
}
```

### 🔑 關鍵技巧：「每層大小快照」

BFS 本身不會自動分層，要自己紀錄「這一層有幾個節點」。進入迴圈時先 `levelSize := len(queue)`，只處理前 `levelSize` 個。

### 📊 圖解步驟

```
初始 queue: [3]

──── 第 1 層 ────
levelSize = 1
取出 3，push 9, 20
queue: [9, 20]
level: [3]

──── 第 2 層 ────
levelSize = 2
取出 9（無子節點）
取出 20，push 15, 7
queue: [15, 7]
level: [9, 20]

──── 第 3 層 ────
levelSize = 2
取出 15
取出 7
queue: []
level: [15, 7]

最終 result = [[3], [9, 20], [15, 7]] ✅
```

---

## 例題二：島嶼數量（LeetCode 200）— 二維網格 BFS

**題目**：給一個由 `'1'`（陸地）和 `'0'`（水）組成的網格，求**島嶼**的數量（連通的 `'1'` 算一座島）。

```
grid = [
  ['1','1','0','0'],
  ['1','1','0','0'],
  ['0','0','1','0'],
  ['0','0','0','1'],
]
答案：3
```

### 🤔 思路

遍歷每個格子，遇到 `'1'` 就從這裡 BFS 淹掉整座島（把相連的 `'1'` 都標記），島嶼計數 +1。

### 🚀 Go 解法

```go
func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    m, n := len(grid), len(grid[0])
    count := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                count++
                bfs(grid, i, j, m, n)
            }
        }
    }
    return count
}

func bfs(grid [][]byte, i, j, m, n int) {
    queue := [][2]int{{i, j}}
    grid[i][j] = '0' // 標記已走過（直接改 grid 省空間）

    dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]

        for _, d := range dirs {
            ni, nj := curr[0]+d[0], curr[1]+d[1]
            if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == '1' {
                grid[ni][nj] = '0'
                queue = append(queue, [2]int{ni, nj})
            }
        }
    }
}
```

### 📊 圖解

```
起始 grid:
  1 1 0 0
  1 1 0 0
  0 0 1 0
  0 0 0 1

遇到 (0,0) 是 '1' → count=1，BFS 把 (0,0),(0,1),(1,0),(1,1) 都淹成 '0'
  0 0 0 0
  0 0 0 0
  0 0 1 0
  0 0 0 1

遇到 (2,2) 是 '1' → count=2，BFS 只淹掉自己
  0 0 0 0
  0 0 0 0
  0 0 0 0
  0 0 0 1

遇到 (3,3) 是 '1' → count=3
最終 count = 3 ✅
```

### 🚨 小技巧：方向陣列 `dirs`

上下左右四個方向用 `[4][2]int{{-1,0},{1,0},{0,-1},{0,1}}` 統一表達，寫 `for d := range dirs` 更乾淨。八方向就是 8 個元素。

---

## 例題三：腐爛的橘子（LeetCode 994）— 多源 BFS 求時間

**題目**：網格中 `2` 是腐爛橘子、`1` 是新鮮橘子、`0` 是空格。每分鐘腐爛橘子會讓**相鄰**的新鮮橘子也腐爛。問所有新鮮橘子腐爛需要幾分鐘？若永遠腐爛不完回傳 -1。

### 🤔 思路：多源 BFS

**多個起點同時出發**的 BFS。把所有初始腐爛橘子**一次全部放進 queue**，然後一起擴散。

### 🚀 Go 解法

```go
func orangesRotting(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    queue := [][2]int{}
    fresh := 0

    // 找出所有腐爛橘子當起點，順便算新鮮橘子總數
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 2 {
                queue = append(queue, [2]int{i, j})
            } else if grid[i][j] == 1 {
                fresh++
            }
        }
    }

    if fresh == 0 {
        return 0
    }

    minutes := 0
    dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    for len(queue) > 0 && fresh > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            curr := queue[0]
            queue = queue[1:]

            for _, d := range dirs {
                ni, nj := curr[0]+d[0], curr[1]+d[1]
                if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
                    grid[ni][nj] = 2
                    fresh--
                    queue = append(queue, [2]int{ni, nj})
                }
            }
        }
        minutes++
    }

    if fresh > 0 {
        return -1
    }
    return minutes
}
```

### 🔑 關鍵：為什麼多源 BFS 能求最少時間？

把所有腐爛橘子同時當起點，它們的擴散波會互相覆蓋。**某個新鮮橘子第一次被染色的時機，就是離它最近的腐爛橘子到達的時間**。這比從每個腐爛橘子各自跑 BFS 快得多。

---

## 🧠 第三部分：心法整理

### 1. 辨識題型的關鍵字

看到這些，先想 BFS：

- 「**最短路徑 / 最少步數 / 最少操作次數**」
- 「**層序遍歷**」
- 「**多久會擴散到全部**」
- 在網格/圖上「**最近的 X**」

### 2. Queue 在 Go 裡的實作

Go 沒有內建 queue，常用 slice 模擬：

```go
queue := []int{start}
curr := queue[0]
queue = queue[1:] // 注意：底層陣列不會縮小，若特別在意記憶體可用 container/list
```

實戰面試這樣寫就夠，面試官不會挑這個。

### 3. visited 要「進 queue 時就標記」，不是「出 queue 時」

```go
// ✅ 正確：進 queue 時就標記
if !visited[next] {
    visited[next] = true // 先標記
    queue = append(queue, next)
}

// ❌ 錯誤：出 queue 時才標記
// 同一個節點可能被多個鄰居推進 queue，造成重複處理和 TLE
```

### 4. BFS 三種常見題型對照

| 題型 | 技巧 |
|------|------|
| **樹的層序遍歷** | 記 `levelSize` 分層 |
| **網格求連通塊** | 直接改 grid 當 visited |
| **求最短距離** | 用 distance map 或層數計數 |
| **多源 BFS** | 多個起點一次放入 queue |

### 5. BFS 空間優化小技巧

- 能直接改原陣列就別開 visited（eg. 網格題）
- visited 用 `map[Key]bool`，Key 可以是座標對或字串
- 大部分題目用 slice 當 queue 就夠，不用 linked list

### 6. BFS vs Dijkstra

BFS 只適合「**每步權重相同**」（eg. 每步都是 1）。如果邊有不同權重，要用 Dijkstra（優先佇列）。面試中 95% 的最短路徑題用 BFS 就夠。

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 技巧 |
|------|------|------|
| Easy | LeetCode 104. Maximum Depth of Binary Tree | 基本層序 |
| Medium | LeetCode 102. Binary Tree Level Order Traversal | 分層 BFS |
| Medium | LeetCode 200. Number of Islands | 網格 BFS |
| Medium | LeetCode 994. Rotting Oranges | 多源 BFS |
| Medium | LeetCode 542. 01 Matrix | 多源 BFS + 距離 |
| Medium | LeetCode 207. Course Schedule | BFS + 拓撲排序 |
| Hard | LeetCode 127. Word Ladder | 字串 BFS |
| Hard | LeetCode 815. Bus Routes | 圖 BFS 進階 |
