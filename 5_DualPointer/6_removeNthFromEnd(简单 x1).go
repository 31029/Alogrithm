package dualPointer

import (
	linknode "github.com/h31029/alogrithm/2_Linknode"
)

// LeetCode-19: 删除链表的倒数第N个节点
// 描述：给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
func RemoveNthFromEnd(head *linknode.ListNode, n int) *linknode.ListNode {
	// 思路1：通过map哈希表将链表转为map后操作。（不推荐，偶尔想表现或者实在没办法再用这个方法。）
	// 思路2：先遍历一遍链表得到长度len，便可以通过 len-n 确定执行多少次Next到指定删除节点。
	len := 0
	temp := head
	for temp != nil {
		len ++
		temp = temp.Next
	}
	
	if len == n {
		// 如果删除的是第一个节点
		head = head.Next
	} else {
		pre := head
		for i := 0; i < len-n-1; i++ {
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
	return head
}