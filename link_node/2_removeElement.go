package linknode

func RemoveElements(head *ListNode, val int) *ListNode {
	// 简单题
	// 解题注意点：
	//  - 删除元素时注意删除点在链表头部这一种情况的处理
	//  - Go语言指针问题弄懂
	cur_p := head
	pre_p := head
	
	for cur_p != nil {
		if cur_p.Val == val {
			if cur_p == pre_p {
				head = head.Next
				cur_p = head
				pre_p = head
			} else {
				pre_p.Next = cur_p.Next
				cur_p = cur_p.Next
			}
		} else {
			pre_p = cur_p
			cur_p = cur_p.Next
		}
	} 

	return head
}
