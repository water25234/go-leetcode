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

// https://medium.com/hannah-lin/%E6%BC%94%E7%AE%97%E6%B3%95%E7%AD%86%E8%A8%98-two-pointer-prefix-sum-167dd0ecb92c
// https://hackmd.io/@SupportCoding/Two_Pointers
