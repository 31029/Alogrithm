package dualPointer

func removeNthFromEnd(head *ListNode, n int) *ListNode {
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