package removenthnodefromendoflist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var index, totalIndex, removeIndex int
	currently, newListNode := head, head
	newHead := &NewListNode{}

	index = 1
	totalIndex = 1

	for currently.Next != nil {

		newHead.insert(index, currently.Val)

		currently = currently.Next
		totalIndex++
		index++
	}

	newHead.insert(index, currently.Val)

	if totalIndex == 1 {
		return nil
	}

	if totalIndex-n == 0 {
		head = head.Next
		return head
	}

	removeIndex = (totalIndex - n)

	newCurrently := newHead
	for i := 0; i < totalIndex; i++ {
		if newCurrently.Index == removeIndex {

			newListNode.Next = newListNode.Next.Next

			return head
		}

		newCurrently = newCurrently.Next
		newListNode = newListNode.Next
	}

	return head
}

type NewListNode struct {
	Len   int
	Index int
	Val   int
	Next  *NewListNode
}

func (newHead *NewListNode) insert(index, val int) {
	newHead.Len = index
	if newHead.Index == 0 {
		newHead.Index = index
		newHead.Val = val
		return
	}

	ptr := newHead

	for i := 0; i <= newHead.Len; i++ {
		if ptr.Next == nil {

			ptr.Next = &NewListNode{
				Index: index,
				Val:   val,
			}
			return
		}

		ptr = ptr.Next
	}
}
