package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Two pointers
	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	var prev *ListNode
	for fast != nil {
		fast = fast.Next
		prev = slow
		slow = slow.Next
	}
	if prev == nil {
		return head.Next
	} else {
		prev.Next = slow.Next
		return head
	}
}
