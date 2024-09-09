package dualPointer

func detectCycle(head *ListNode) *ListNode {
	low, fast := head, head
	for fast != nil {
		low = low.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
		if low == fast {
			return fast
		}
	}
	return nil
}