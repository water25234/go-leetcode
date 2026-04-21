package deletenodeinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if nil == root {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Right == nil && root.Left == nil {
		return nil
	}

	if root.Right == nil {
		return root.Left
	}

	if root.Left == nil {
		return root.Right
	}

	root.Val = findMin(root.Right)
	root.Right = deleteNode(root.Right, root.Val)

	return root
}

func findMin(node *TreeNode) int {
	minVal := node.Val

	for node != nil {
		if node.Val < minVal {
			minVal = node.Val
		}
		node = node.Left
	}

	return minVal
}

// 我來用這個具體例子一步步解說刪除節點 3 的過程：

// 初始樹結構
// 執行過程：
// 第1步：deleteNode(root=5, key=3)
// root.Val = 5, key = 3
// 3 < 5 → 進入第14行條件
// root.Left = deleteNode(root.Left, 3) → 遞迴到左子樹

// 第2步：deleteNode(root=3, key=3)
// root.Val = 3, key = 3
// 跳過前兩個條件（不是 < 也不是 >）
// 檢查子節點情況：
// root.Right != nil (有節點4)
// root.Left != nil (有節點2)
// 進入第36-37行（兩個子節點都存在的情況）

// 第3步：處理有兩個子節點的刪除
// 第36行：root.Val = findMin(root.Right)
// findMin(節點4) 被調用
// 節點4沒有左子樹，所以最小值就是4
// root.Val = 4 → 節點3的值被替換為4
// 第37行：root.Right = deleteNode(root.Right, root.Val)
// 現在要刪除右子樹中值為4的節點
// deleteNode(節點4, 4) 被調用

// 第4步：deleteNode(root=4, key=4)
// 找到要刪除的節點4
// 節點4是葉子節點（root.Right == nil && root.Left == nil）
// 第25行：return nil → 節點4被刪除

// 第5步：回到第2步
// root.Right = nil （節點4已被刪除）
// 第38行：return root → 返回修改後的節點（原節點3，現在值為4）

// 第6步：回到第1步
// root.Left = 修改後的節點 （值為4的節點）
// 第16行：return root → 返回根節點5

// 最終結果：

// 關鍵理解：
// 節點3沒有被物理刪除，而是被重新賦值為4
// 真正被刪除的是原來的節點4（在第25行返回nil）
// 樹的結構通過返回值重新連接，形成新的BST
// 這就是為什麼輸出是 [5,4,6,2,null,null,7] 的原因！## 關鍵理解：

// 節點3沒有被物理刪除，而是被重新賦值為4
// 真正被刪除的是原來的節點4（在第25行返回nil）
// 樹的結構通過返回值重新連接，形成新的BST
// 這就是為什麼輸出是 [5,4,6,2,null,null,7] 的原因！
