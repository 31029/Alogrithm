package linknode


func SwapPairs(head *ListNode) *ListNode {
	var ptr_list []*ListNode
	
	if head == nil {
		return nil
	}

	// 数值上同一
	ptr_list = append(ptr_list, nil)
	for head != nil {
		ptr_list = append(ptr_list, head)
		head = head.Next
	}

	if len(ptr_list) == 2 {
		return ptr_list[1]
	}

	// 主逻辑
	for i := 1; i <= (len(ptr_list)-1) / 2; i++ {
		even_index := 2 * i
		odd_index := 2 * i - 1

		ptr_list[even_index].Next = ptr_list[odd_index]

		if odd_index + 3 > len(ptr_list) {
			ptr_list[odd_index].Next = nil
		} else if odd_index + 3 == len(ptr_list) {
			ptr_list[odd_index].Next = ptr_list[len(ptr_list)-1]
			ptr_list[len(ptr_list)-1].Next = nil
		} else {
			ptr_list[odd_index].Next = ptr_list[odd_index + 3]
		}
	}
	
	return ptr_list[2]
}

// 正常模拟
func SwapPairs_Normal(head *ListNode) *ListNode {
    dummy := &ListNode{
        Next: head,
    }
    //head=list[i]
    //pre=list[i-1]
    pre := dummy 
    for head != nil && head.Next != nil {
        pre.Next = head.Next
        next := head.Next.Next
        head.Next.Next = head
        head.Next = next
        //pre=list[(i+2)-1]
        pre = head 
        //head=list[(i+2)]
        head = next
    }
    return dummy.Next
}

// 递归版本
func SwapPairs_digui(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    next := head.Next
    head.Next = SwapPairs_digui(next.Next)
    next.Next = head
    return next
}