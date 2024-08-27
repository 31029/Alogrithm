package linknode

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	// 问题简述：找两个不同长度的链表的交点节点。
	// 暴力解法，时间复杂度 n^2
	// 实际解法思路十分简单，去看图。
    if headA == nil  || headB == nil {
		return nil
	}

	var ptr_list []*ListNode
	for headA != nil {
		ptr_list = append(ptr_list, headA)
		headA = headA.Next
	}

	for headB != nil {
		for i := 0; i < len(ptr_list); i++ {
			if headB == ptr_list[i] {
				return headB
			} 	
		}
		headB = headB.Next
	}

	return nil
}