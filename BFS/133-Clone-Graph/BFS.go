package clonegraph

import "container/list"

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
	if node == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(node)
	visited := make(map[*Node]*Node, 0)
	visited[node] = &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	for queue.Len() > 0 {
		n := queue.Remove(queue.Front()).(*Node)
		for i, v := range n.Neighbors {
			if _, ok := visited[v]; !ok {
				queue.PushBack(v)
				visited[v] = &Node{Val: v.Val, Neighbors: make([]*Node, len(v.Neighbors))}
			}
			visited[n].Neighbors[i] = visited[v]
		}
	}
	return visited[node]
}
