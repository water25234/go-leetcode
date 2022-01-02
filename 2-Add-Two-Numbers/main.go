package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{Val: 0}

	current := head

	n1, n2, balance := 0, 0, 0

	for l1 != nil || l2 != nil || balance != 0 {

		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		sumTotal := (n1 + n2 + balance)

		current.Next = &ListNode{Val: (sumTotal % 10)}
		current = current.Next

		balance = sumTotal / 10
	}

	return head.Next
}
