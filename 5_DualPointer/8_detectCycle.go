package dualPointer

import (
	linknode "github.com/h31029/alogrithm/2_Linknode"
)

func detectCycle(head *linknode.ListNode) *linknode.ListNode {
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