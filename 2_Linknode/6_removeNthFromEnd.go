package linknode

func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	var ptr_list []*ListNode
	if head == nil {
		return nil
	}

	for head != nil {
		ptr_list = append(ptr_list, head)
		head = head.Next
	}

	if len(ptr_list) <= 1 {
		return nil
	} else if len(ptr_list) > 1 {
		// 如果删除的是第一个头节点
		if n == len(ptr_list) {
			return ptr_list[1]
		} else if n == 1 {
			ptr_list[len(ptr_list) - 2].Next = nil
		} else {
			ptr_list[len(ptr_list) - n - 1].Next = ptr_list[len(ptr_list) - n + 1]
		}
	}
	return ptr_list[0]
}