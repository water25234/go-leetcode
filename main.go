package main

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

func (n *Node) Append(val int) {
	end := &Node{Val: val}
	here := n
	for here.Next != nil {
		here = here.Next
	}
	here.Next = end
}

func Remove(n *Node, val int) *Node {
	traverser := n
	for traverser.Next != nil {
		if traverser.Next.Val == val {
			traverser.Next = traverser.Next.Next
			return n
		}
		traverser = traverser.Next
	}
	return n
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

func main() {
	n := NewNode(1)
	n.Append(2)
	n.Append(3)
	n.Append(4)
	n.Append(5)

	m := Remove(n, 3)

	for m != nil {
		fmt.Println(m.Val)
		m = m.Next
	}
}
