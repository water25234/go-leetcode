package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {

	visited := make(map[*Node]*Node, 0)

	var dfs func(*Node) *Node

	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		if r, ok := visited[n]; ok {
			return r
		}
		r := &Node{
			Val:       n.Val,
			Neighbors: make([]*Node, len(n.Neighbors)),
		}
		visited[n] = r
		for i, v := range n.Neighbors {
			r.Neighbors[i] = dfs(v)
		}
		return r
	}
	return dfs(node)
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
// [[2,4],[1,3],[2,4],[1,3]]
// [
// 	&Node{
// 		Val: 2,
// 		Neighbors: [
// 			&Node{
// 				Val: 1,
// 				Neighbors:[
// 					&Node{
// 						Val: 1,
// 						Neighbors:[

// 						]
// 					}
// 				]
// 			}
// 		]
// 	},
// 	&Node{
// 		Val: 4,
// 		Neighbors: [
// 			&Node{
// 				Val: 3,
// 				Neighbors:[

// 				]
// 			}
// 		]
// 	},
// ]
