package linknode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeLinkList(array []int) *ListNode {
	var head *ListNode = &ListNode{}
	cur_p := head
	for k, v := range array {
		cur_p.Val = v
		if k != len(array)-1 {
			cur_p.Next = &ListNode{}
		}
		cur_p = cur_p.Next
	}
	return head
}

func PringLinkList(head *MyLinkedList) string {
	var s string
	for head != nil {
		s += fmt.Sprintf("%d -> ", head.Val)
		head = head.Next
	}
	return s
}

func PringLinkList2(head *ListNode) string {
	var s string
	for head != nil {
		s += fmt.Sprintf("%d -> ", head.Val)
		head = head.Next
	}
	return s
}