package linknode

func ReverseList(head *ListNode) *ListNode {
	// 数组法：用数组记录每个Listnode的地址，然后反向遍历
	// 双指针法：设置pre和cur双指针，再用一个temp指针记录下一个待处理节点，后从第一个开始，一个一个取反。
	var ptr_list []*ListNode
	if head == nil {
		return nil
	}

	cur_p := head
	for cur_p != nil {
		ptr_list = append(ptr_list, cur_p)
		cur_p = cur_p.Next
	}

	node_num := len(ptr_list)
	for i := node_num-1; i > 0; i-- {
		ptr_list[i].Next = ptr_list[i-1]		
	}

	ptr_list[0].Next = nil
	return ptr_list[node_num-1]
}

func ReverseList_2P(head *ListNode) *ListNode {
	var temp_p, pre_p *ListNode

	cur_p := head
	for cur_p != nil {
		// 更新temp_p
		temp_p = cur_p.Next

		// 翻转操作
		cur_p.Next = pre_p

		// 更新pre_p 和 cur_p
		pre_p = cur_p
		cur_p = temp_p
	}

	return cur_p
}