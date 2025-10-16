package leafsimilartrees

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	root1Result := leaf(root1)
	root2Result := leaf(root2)

	return slices.Equal(root1Result, root2Result)
}

func leaf(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}

	root1Result := leaf(node.Left)
	root2Result := leaf(node.Right)

	return append(root1Result, root2Result...)
}
