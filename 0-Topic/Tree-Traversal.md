# Tree Traversal（樹的遍歷）完整教學

---

## 📖 第一部分：什麼是 Tree Traversal？

### 生活化比喻

想像一個**家族族譜**：爺爺底下有爸爸、叔叔，爸爸底下有你跟哥哥。要把全族的名字**唸一遍**有幾種順序？

- **先唸自己，再唸左邊小孩，再唸右邊**（前序）
- **先唸左邊小孩，再唸自己，再唸右邊**（中序）
- **先唸左邊小孩，再唸右邊，最後唸自己**（後序）
- **一代一代從上往下唸**（層序）

這就是樹的四種遍歷方式。不同順序適合不同任務。

### 二元樹節點定義（LeetCode 標準）

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
```

### 四種遍歷一覽表

| 遍歷 | 順序 | 代表用途 |
|------|------|---------|
| **前序（Preorder）** | 根 → 左 → 右 | 複製樹、序列化樹 |
| **中序（Inorder）** | 左 → 根 → 右 | BST 的排序輸出 |
| **後序（Postorder）** | 左 → 右 → 根 | 刪除樹、計算子樹資訊 |
| **層序（Level-order）** | 一層一層 | 最淺深度、層級相關題 |

### DFS 三種遍歷差異視覺化

```
       1
      / \
     2   3
    / \
   4   5

前序 (根左右):  1 → 2 → 4 → 5 → 3
中序 (左根右):  4 → 2 → 5 → 1 → 3
後序 (左右根):  4 → 5 → 2 → 3 → 1
層序:          1 → 2, 3 → 4, 5
```

**記憶技巧**：看「根」的位置在中間還是前後：

- 前序：**根**在前
- 中序：**根**在中間
- 後序：**根**在後

---

## 🎯 第二部分：例題實戰

---

## 例題一：三種 DFS 遍歷（LeetCode 94 / 144 / 145）

**題目**：分別用三種順序遍歷二元樹，回傳節點值陣列。

### 🚀 遞迴版（最直覺）

```go
// 前序遍歷（LeetCode 144）
func preorderTraversal(root *TreeNode) []int {
    result := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        result = append(result, node.Val) // 根
        dfs(node.Left)                    // 左
        dfs(node.Right)                   // 右
    }
    dfs(root)
    return result
}

// 中序（LeetCode 94）：換順序就好
func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        result = append(result, node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return result
}

// 後序（LeetCode 145）
func postorderTraversal(root *TreeNode) []int {
    result := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        dfs(node.Right)
        result = append(result, node.Val)
    }
    dfs(root)
    return result
}
```

### 🚀 迭代版（用 Stack）

面試如果被問「不能用遞迴」，要能寫出迭代版。

**前序迭代**：

```go
func preorderTraversal(root *TreeNode) []int {
    result := []int{}
    if root == nil {
        return result
    }
    stack := []*TreeNode{root}

    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, node.Val)

        // 右邊先入 stack（後進先出，等等會先彈出左邊）
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
    }
    return result
}
```

**中序迭代**（比較複雜，要一路往左走）：

```go
func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    stack := []*TreeNode{}
    curr := root

    for curr != nil || len(stack) > 0 {
        // 一路往左走到底
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        // 彈出處理
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, curr.Val)
        // 轉向右子樹
        curr = curr.Right
    }
    return result
}
```

---

## 例題二：層序遍歷（LeetCode 102）— BFS 在樹上

**題目**：逐層輸出樹的節點值。

```
     3
    / \
   9  20
      / \
     15  7

輸出：[[3], [9, 20], [15, 7]]
```

### 🚀 Go 解法

```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    result := [][]int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        size := len(queue)
        level := []int{}

        for i := 0; i < size; i++ {
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

這題跟 BFS 教學是同一題，細節參考 `BFS.md`。

---

## 例題三：驗證二元搜尋樹（LeetCode 98）— 中序應用

**題目**：判斷一棵樹是否為有效的 BST（左子樹所有值 < 根 < 右子樹所有值）。

### 🤔 思路

**BST 的中序遍歷結果應該是嚴格遞增的！** 只要中序走一遍，檢查每個值是否比前一個大即可。

### 🚀 Go 解法

```go
func isValidBST(root *TreeNode) bool {
    var prev *TreeNode
    var dfs func(node *TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil {
            return true
        }

        if !dfs(node.Left) {
            return false
        }

        if prev != nil && node.Val <= prev.Val {
            return false
        }
        prev = node

        return dfs(node.Right)
    }
    return dfs(root)
}
```

### 🔑 為什麼中序適合 BST？

BST 的性質：**左子樹所有值 < 根 < 右子樹所有值**，遞迴展開到最底就是一個**由小到大的序列**。這題是 BST 的定義級應用。

---

## 例題四：二元樹的最大深度（LeetCode 104）— 後序應用

**題目**：求二元樹的最大深度。

### 🚀 Go 解法

```go
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    return max(leftDepth, rightDepth) + 1
}
```

### 🔑 這為什麼是後序？

寫法上是「先遞迴左、再遞迴右、最後算自己」，完全符合**後序（左→右→根）**順序。

很多「**計算子樹資訊**」的題目本質都是後序——必須先知道孩子的答案，才能算自己的答案。

---

## 🧠 第三部分：心法整理

### 1. 四種遍歷的典型使用場景

| 場景 | 用什麼遍歷 |
|------|-----------|
| 複製、序列化樹 | 前序（先記根，再記子樹） |
| BST 排序輸出、驗證 | 中序（天然遞增） |
| 計算子樹資訊（高度、總和、子樹大小） | 後序（先算子，再算父） |
| 最短路徑、最淺葉、層數相關 | 層序（BFS） |

### 2. 遞迴寫樹題的三步法

```
1. base case：node == nil 怎麼處理？
2. 遞迴：左子樹的答案是什麼？右子樹的答案是什麼？
3. 合併：怎麼用孩子的答案算出自己的答案？
```

**關鍵心法**：相信遞迴會回正確答案，專注寫好「怎麼用孩子的結果」這一步。

### 3. 面試常考的樹題分類

| 類別 | 代表題 | 遍歷 |
|------|--------|------|
| **樹的基本資訊** | 最大深度、節點數、直徑 | 後序 |
| **路徑題** | 路徑總和、根到葉所有路徑 | 前序 |
| **BST 相關** | 驗證 BST、BST 第 k 小 | 中序 |
| **對稱/翻轉** | 對稱樹、翻轉樹 | 後序 |
| **LCA 最低公共祖先** | LeetCode 236 | 後序 |

### 4. 迭代版「要用 stack 模擬遞迴」

函式呼叫背後就是 call stack，迭代版就是**自己維護一個 stack 模擬**。

- 前序最好寫：把節點壓入 stack，彈出時加入結果
- 中序要特殊處理：一路往左走到底才彈
- 後序最難：可以用「反向前序」（根右左）倒過來得到（左右根）

面試如果寫不出中序/後序迭代版，直接寫遞迴版也能過大多數題目。

### 5. 樹題面試加分技巧

**主動分析時間/空間複雜度**：

```
時間：O(n)，每個節點訪問一次
空間：O(h)，h 是樹高；最壞 O(n)（鏈狀樹），平衡樹是 O(log n)
```

**主動提空間優化**：Morris Traversal 可以 O(1) 空間中序遍歷（修改 nil 指標當 thread）。除非面試官問，不用主動寫，但知道能加分。

### 6. 何時用 BFS、何時用 DFS？

| 情境 | 選擇 |
|------|------|
| 問「某一層」的資訊 | BFS |
| 問「從根到葉的路徑」 | DFS |
| 問「最淺」葉節點 | BFS（BFS 先到淺的） |
| 問「最深」資訊 | DFS |
| 問「是否存在」某條路徑 | DFS（短路回傳） |

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 遍歷類型 |
|------|------|---------|
| Easy | LeetCode 144/94/145. Pre/In/Postorder Traversal | 基本三種遍歷 |
| Easy | LeetCode 104. Maximum Depth | 後序 |
| Easy | LeetCode 226. Invert Binary Tree | 後序 |
| Easy | LeetCode 100. Same Tree | 前序比較 |
| Easy | LeetCode 101. Symmetric Tree | 雙遞迴 |
| Medium | LeetCode 102. Level Order Traversal | 層序 |
| Medium | LeetCode 98. Validate Binary Search Tree | 中序 |
| Medium | LeetCode 230. Kth Smallest Element in BST | 中序 |
| Medium | LeetCode 236. Lowest Common Ancestor | 後序 |
| Medium | LeetCode 105. Build Tree from Preorder+Inorder | 前序+中序重建 |
| Hard | LeetCode 124. Binary Tree Maximum Path Sum | 後序 + 巧妙回傳 |
| Hard | LeetCode 297. Serialize and Deserialize Binary Tree | 前序序列化 |
