package pathsumiii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	return dfs(root, targetSum, []int{})
}

func dfs(root *TreeNode, target int, currentPath []int) int {
	if root == nil {
		return 0
	}

	currentPath = append(currentPath, root.Val)
	pathCount, pathSum := 0, 0

	for i := len(currentPath) - 1; i >= 0; i-- {
		pathSum += currentPath[i]

		if pathSum == target {
			pathCount++
		}
	}

	pathCount += dfs(root.Left, target, currentPath)
	pathCount += dfs(root.Right, target, currentPath)

	// remove the current node before returning ex: currentPath = [5, 4, 11, 7] currentPath[:3] => [5, 4, 11]
	currentPath = currentPath[:len(currentPath)-1]

	return pathCount
}

// Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//         5
//       /   \
//      4     8
//    /      /  \
//   11     13   4
//  / \          / \
// 7   2        5   1
// 執行過程詳解
// 步驟 1: 開始從根節點 5
// currentPath: [5]
// 檢查路徑:
// 5 ≠ 22 ❌
// 步驟 2: 遍歷左子樹，到達節點 4
// currentPath: [5, 4]
// 檢查路徑:
// 4 ≠ 22 ❌
// 4 + 5 = 9 ≠ 22 ❌
// 步驟 3: 繼續到節點 11
// currentPath: [5, 4, 11]
// 檢查路徑:
// 11 ≠ 22 ❌
// 11 + 4 = 15 ≠ 22 ❌
// 11 + 4 + 5 = 20 ≠ 22 ❌
// 步驟 4: 到達節點 7
// currentPath: [5, 4, 11, 7]
// 檢查路徑:
// 7 ≠ 22 ❌
// 7 + 11 = 18 ≠ 22 ❌
// 7 + 11 + 4 = 22 ✅ 找到路徑 1!
// 7 + 11 + 4 + 5 = 27 ≠ 22 ❌
// 步驟 5: 回溯到節點 11，然後到節點 2
// currentPath: [5, 4, 11, 2]
// 檢查路徑:
// 2 ≠ 22 ❌
// 2 + 11 = 13 ≠ 22 ❌
// 2 + 11 + 4 = 17 ≠ 22 ❌
// 2 + 11 + 4 + 5 = 22 ✅ 找到路徑 2!
// 步驟 6: 回溯到根節點，遍歷右子樹，到達節點 8
// currentPath: [5, 8]
// 檢查路徑:
// 8 ≠ 22 ❌
// 8 + 5 = 13 ≠ 22 ❌
// 步驟 7: 到達節點 13
// currentPath: [5, 8, 13]
// 檢查路徑:
// 13 ≠ 22 ❌
// 13 + 8 = 21 ≠ 22 ❌
// 13 + 8 + 5 = 26 ≠ 22 ❌
// 步驟 8: 到達右側節點 4
// currentPath: [5, 8, 4]
// 檢查路徑:
// 4 ≠ 22 ❌
// 4 + 8 = 12 ≠ 22 ❌
// 4 + 8 + 5 = 17 ≠ 22 ❌
// 步驟 9: 到達節點 5
// currentPath: [5, 8, 4, 5]
// 檢查路徑:
// 5 ≠ 22 ❌
// 5 + 4 = 9 ≠ 22 ❌
// 5 + 4 + 8 = 17 ≠ 22 ❌
// 5 + 4 + 8 + 5 = 22 ✅ 找到路徑 3!
// 步驟 10: 到達節點 1
// currentPath: [5, 8, 4, 1]
// 檢查路徑:
// 1 ≠ 22 ❌
// 1 + 4 = 5 ≠ 22 ❌
// 1 + 4 + 8 = 13 ≠ 22 ❌
// 1 + 4 + 8 + 5 = 18 ≠ 22 ❌
// 找到的三條路徑
// 路徑 1: 4 → 11 → 7 (4 + 11 + 7 = 22)
// 路徑 2: 5 → 4 → 11 → 2 (5 + 4 + 11 + 2 = 22)
// 路徑 3: 5 → 8 → 4 → 5 (5 + 8 + 4 + 5 = 22)
// 最終結果
// 返回值: 3（總共找到 3 條和為 22 的路徑）

// 這個演算法的關鍵是在每個節點都檢查所有以該節點為終點的可能路徑，確保不遺漏任何符合條件的路徑組合。
