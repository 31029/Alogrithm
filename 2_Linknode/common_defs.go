package linknode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeLinkList(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
	head := &ListNode{
		Val: array[0],
		Next: nil,
	}

	cur := head
	for _, v := range array[1:] {
		cur.Next = &ListNode{
			Val: v,
			Next: nil,
		}
		cur = cur.Next
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