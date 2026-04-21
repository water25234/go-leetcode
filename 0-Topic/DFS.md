# DFS（深度優先搜尋）完整教學

---

## 📖 第一部分：什麼是 DFS？

### 生活化比喻

想像你在走**迷宮**，手上拿著粉筆：

- 看到一條路 → **一直往下走**，直到走到死胡同或終點
- 走到死胡同 → **往回走一步**，換另一條岔路走
- 每走過一個格子就用粉筆做記號，避免走回頭路

這就是 DFS —— **一條路走到底，走不通才回頭**。

### DFS 的核心特性

| 特性 | 說明 |
|------|------|
| **往深處探索** | 盡可能走最深，走不動才回頭 |
| **天然遞迴** | 遞迴寫法最直覺（也可以用 stack 迭代） |
| **visited 集合** | 避免重複走或形成無窮迴圈 |
| **回溯時機** | 進入時做選擇，回來時撤銷（如果要列舉路徑） |

### DFS vs BFS 的選擇

| 情境 | 選誰 |
|------|------|
| 找**最短**路徑 | BFS |
| 找**所有可能**路徑 / 排列組合 | DFS |
| 判斷**連通塊** / 連通性 | 都可以，DFS 程式碼更短 |
| **樹的前/中/後序** | DFS |
| **樹的層序** | BFS |

### DFS 的萬用模板（遞迴版）

```go
var visited = make(map[Node]bool)

func dfs(curr Node) {
    if /* 終止條件 */ {
        return
    }
    visited[curr] = true

    // 處理 curr

    for _, next := range curr.Neighbors() {
        if !visited[next] {
            dfs(next)
        }
    }
}
```

---

## 🎯 第二部分：例題實戰

---

## 例題一：二元樹最大深度（LeetCode 104）— DFS 入門

**題目**：給一棵二元樹，求它的最大深度（從根到最遠葉節點的節點數）。

```
     3
    / \
   9  20
     /  \
    15   7

答案：3
```

### 🚀 Go 解法（遞迴）

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)

    return max(leftDepth, rightDepth) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### 🤔 這就是 DFS 嗎？是的！

遞迴呼叫樹本身就是 DFS：從 root 衝到最底，才回頭算另一邊。函式呼叫棧就是 DFS 的 stack。

### 📊 圖解遞迴展開

```
maxDepth(3)
├── maxDepth(9)
│   ├── maxDepth(nil) → 0
│   └── maxDepth(nil) → 0
│   回傳 max(0,0)+1 = 1
│
└── maxDepth(20)
    ├── maxDepth(15)
    │   ├── maxDepth(nil) → 0
    │   └── maxDepth(nil) → 0
    │   回傳 1
    └── maxDepth(7)
    │   回傳 1
    回傳 max(1,1)+1 = 2

root：max(1, 2) + 1 = 3 ✅
```

---

## 例題二：島嶼數量（LeetCode 200）— 網格 DFS

**題目**：給一個 `'1'` 陸地、`'0'` 水的網格，求島嶼數。（跟 BFS 教學同一題，這裡用 DFS）

```
grid = [
  ['1','1','0'],
  ['1','0','0'],
  ['0','0','1'],
]
答案：2
```

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
                dfs(grid, i, j, m, n)
            }
        }
    }
    return count
}

func dfs(grid [][]byte, i, j, m, n int) {
    // 越界或不是陸地就返回
    if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
        return
    }
    grid[i][j] = '0' // 標記走過

    dfs(grid, i-1, j, m, n)
    dfs(grid, i+1, j, m, n)
    dfs(grid, i, j-1, m, n)
    dfs(grid, i, j+1, m, n)
}
```

### 🔑 DFS vs BFS 在這題的差異

- 程式碼：DFS 用遞迴，**行數更少更乾淨**
- 空間：DFS 遞迴深度 = 島嶼大小 → 大網格可能 stack overflow；BFS 佇列寬度通常較穩
- 行為：DFS 會「衝到最遠再回來」；BFS「一圈一圈往外擴散」
- 結果：這題兩者完全等價 ✅

### 📊 圖解 DFS 路徑（起點 (0,0)）

```
起始:
  1 1 0
  1 0 0
  0 0 1

DFS(0,0) → 改 '1' 為 '0'
  0 1 0
  1 0 0

→ 嘗試 (-1,0) 越界
→ DFS(1,0)
  0 1 0
  0 0 0

→ DFS(2,0) 是 '0' 直接返回
→ DFS(0,0) 已標記 '0' 返回
→ DFS(1,-1) 越界
→ DFS(1,1) 是 '0' 返回
→ 回到 DFS(0,0) 繼續 → DFS(0,1)
  0 0 0
  0 0 0

完成一座島，count++
```

---

## 例題三：路徑總和（LeetCode 112）— 樹的路徑 DFS

**題目**：給一棵二元樹和 `targetSum`，判斷**是否存在某條根到葉的路徑**，節點值加總等於 targetSum。

```
     5
    / \
   4   8
  /   / \
 11  13  4
 /\       \
7  2       1

targetSum = 22 → true（5→4→11→2 = 22）
```

### 🚀 Go 解法

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    // 到達葉節點，檢查是否湊到目標
    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum
    }

    remain := targetSum - root.Val
    return hasPathSum(root.Left, remain) || hasPathSum(root.Right, remain)
}
```

### 🤔 思路重點

- **沿路扣減**：每經過一個節點，把它的值從 target 減掉，問題變成「剩下的樹能不能湊到剩下的 target」
- **葉節點才判斷**：不是任何節點都能「結束」，只有**左右子樹都為 nil** 才算葉

### 📊 遞迴樹展開（targetSum = 22）

```
hasPathSum(5, 22)
├── hasPathSum(4, 17)
│   └── hasPathSum(11, 13)
│       ├── hasPathSum(7, 2)  → 葉，7 != 2, false
│       └── hasPathSum(2, 2)  → 葉，2 == 2, true ⭐
│       回傳 true
│   回傳 true
└── (短路，不用再算)
回傳 true ✅
```

Go 的 `||` 短路：左邊 true 就不算右邊，對效能有幫助。

---

## 🧠 第三部分：心法整理

### 1. 辨識題型的關鍵字

看到這些，先想 DFS：

- 「**列出所有路徑/組合**」
- 「**樹的遍歷**（前/中/後序）」
- 「**連通塊**個數/大小」
- 「能不能**從 A 走到 B**」
- 「**判斷是否存在**某條路徑」

### 2. DFS 三問寫法

每題先問自己：

```
1. 終止條件是什麼？（nil？越界？已訪問？）
2. 遞迴單層要做什麼？（處理 curr、探索鄰居）
3. 回傳值要什麼？（bool？int？void？）
```

### 3. 進入時標記 vs 回溯時撤銷

| 題型 | visited 處理 |
|------|-------------|
| 只要找到一個結果 | 進入時標記就好（eg. 島嶼） |
| 要列舉所有路徑 / 回溯 | 進入時標記，離開時撤銷 |

後者就是 **Backtracking**（回溯法），是 DFS 的特化。

### 4. 遞迴 vs 迭代（用 stack）

遞迴版本通常 10 行，迭代版要 20 行 +。除非：

- 樹/圖很深怕 stack overflow
- 面試官指定要迭代

否則**面試用遞迴版**。Go 預設 stack 夠用，幾千層深度沒問題。

### 5. 網格 DFS 的通用改造

大部分網格題長得很像，改一下判斷和標記就能用：

```go
func dfs(grid [][]int, i, j, m, n int) {
    if i < 0 || i >= m || j < 0 || j >= n || /* 不符合條件 */ {
        return
    }
    // 標記 or 計數
    dfs(grid, i-1, j, m, n)
    dfs(grid, i+1, j, m, n)
    dfs(grid, i, j-1, m, n)
    dfs(grid, i, j+1, m, n)
}
```

背起來，臨場改動就好。

### 6. 遞迴的正確姿勢：相信它會回正確答案

寫遞迴時，不要試著把**整個遞迴樹**在腦中展開——那是電腦做的事。你只要確保：

- **base case 對**
- **遞迴一步的邏輯對**
- **假設子問題會回正確答案**

這三點對，遞迴就會對。這叫「歸納法的信念」。

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 技巧 |
|------|------|------|
| Easy | LeetCode 104. Maximum Depth of Binary Tree | 樹 DFS 入門 |
| Easy | LeetCode 112. Path Sum | 路徑累計 |
| Easy | LeetCode 226. Invert Binary Tree | 後序遍歷 |
| Medium | LeetCode 200. Number of Islands | 網格 DFS |
| Medium | LeetCode 695. Max Area of Island | 網格 DFS + 回傳值 |
| Medium | LeetCode 133. Clone Graph | 圖 DFS + visited map |
| Medium | LeetCode 129. Sum Root to Leaf Numbers | 路徑累計 |
| Medium | LeetCode 417. Pacific Atlantic Water Flow | 反向 DFS |
| Hard | LeetCode 124. Binary Tree Maximum Path Sum | 樹 DFS 回傳值設計 |
