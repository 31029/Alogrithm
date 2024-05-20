package linknode


// type SingleNode struct {
// 	Val  int         // 节点的值
// 	Next *SingleNode // 下一个节点的指针
// }

// type MyLinkedList struct {
// 	dummyHead *SingleNode // 虚拟头节点
// 	Size      int         // 链表大小
// }

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}


func Constructor() MyLinkedList {
	return MyLinkedList{
		Val: -999,
		Next: nil,
	}
}


func (p *MyLinkedList) Get(index int) int {
	if p == nil {
		return -1
	}
	 
	cur_p := p
	for i := 0; i < index; i++ {
		cur_p = cur_p.Next
		if cur_p == nil {
			return -1
		}
	}

	return cur_p.Val
}


func (p *MyLinkedList) AddAtHead(val int)  {
	if p == nil {
		p = &MyLinkedList{
			Val: val,
			Next: nil,
		}
	}
	old_v := p.Val
	old_next := p.Next
	p.Val = val
	p.Next = &MyLinkedList{
		Val: old_v,
		Next: old_next,
	}
}


func (p *MyLinkedList) AddAtTail(val int)  {
	if p == nil {
		p = &MyLinkedList{
			Val: val,
			Next: nil,
		}
	}

	pre_p := p
	for p != nil {
		pre_p = p
		p = p.Next
	}
	pre_p.Next = &MyLinkedList{
		Val: val,
		Next: nil,
	}
}


func (p *MyLinkedList) AddAtIndex(index int, val int)  {
	if p == nil {
		return 
	}
	if index <= 0 {
		p.AddAtHead(val)
	} 

	pre_p := p
	cur_p := p
	for i := 0; i < index; i++ {
		pre_p = cur_p
		cur_p = cur_p.Next
		if cur_p == nil {
			pre_p.Next = &MyLinkedList{
				Val: val,
				Next: nil,
			} 
		}
	}

	pre_p.Next = &MyLinkedList{
		Val: val,
		Next: cur_p,
	}

}


func (p *MyLinkedList) DeleteAtIndex(index int)  {
	if p == nil {
		return 
	}
	if index == 0 {
		*p = *p.Next
	}
	
	pre_p := p
	cur_p := p
	for i := 0; i < index; i++ {
		pre_p = cur_p
		cur_p = cur_p.Next
		if cur_p == nil {
			return
		}
	}

	pre_p.Next = cur_p.Next
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */