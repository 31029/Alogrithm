package dualPointer

import (
	linknode "github.com/h31029/alogrithm/2_Linknode"
)

func SwapPairs(head *linknode.ListNode) *linknode.ListNode {
	if head == nil {
		return nil
	}
	pre := head
	cur, res := head.Next, head.Next
	// 仅有一个节点时处理
	if cur == nil {
		return head
	}

	for pre != nil {
		var next * linknode.ListNode
		if cur != nil {
			next = cur.Next
		}

		if next == nil {
			// 偶数个节点时结尾处理
			pre.Next = nil
			cur.Next = pre
			break
		} else if next.Next == nil {
			// 奇数个节点时结尾处理
			pre.Next = next
			cur.Next = pre
			break
		} else {
			pre.Next = next.Next
		}
		cur.Next = pre

		pre = next
		cur = next.Next
	}

	return res
}