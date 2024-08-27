package linknode

func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var ptr_list []*ListNode
	for head != nil {
		for i := 0; i < len(ptr_list); i++ {
			if ptr_list[i] == head {
				return head
			}
		}
		ptr_list = append(ptr_list, head)
		head = head.Next
	}

	return nil
}