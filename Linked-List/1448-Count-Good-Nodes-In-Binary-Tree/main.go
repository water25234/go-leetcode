package countgoodnodesinbinarytree

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
func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, max int) int {
	if nil == node {
		return 0
	}

	good := 0
	if node.Val >= max {
		good = 1
		max = node.Val
	}

	if nil != node.Left {
		good += dfs(node.Left, max)
	}

	if nil != node.Right {
		good += dfs(node.Right, max)
	}

	return good
}
