package dualPointer

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	pre := head
	cur := head.Next

	for cur.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next = pre

		pre = pre.Next
		cur = next
	}

	return head
}