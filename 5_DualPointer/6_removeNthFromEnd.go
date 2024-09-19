package dualPointer

import (
	linknode "github.com/h31029/alogrithm/2_Linknode"
)


func removeNthFromEnd(head *linknode.ListNode, n int) *linknode.ListNode {
	len := 0
	temp := head
	for temp.Next != nil {
		len ++
	}
	
	if n == len {
		head = head.Next
	} else {
		pre := head
		for i := 0; i < len-n-1; i++ {
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
	return head
}