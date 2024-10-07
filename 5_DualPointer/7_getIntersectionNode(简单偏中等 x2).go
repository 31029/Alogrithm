package dualPointer

import (
	linknode "github.com/h31029/alogrithm/2_Linknode"
)

// LeetCode-面试题-02-07: 链表相交 【https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/description/】
// 描述：给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
func GetIntersectionNode(headA, headB *linknode.ListNode) *linknode.ListNode {
	// 解题思路1：双指针。但要先同时对齐后移动！！！ 理解这一点该题便很简单。
	lA, lB := 0, 0
	cA, cB := headA, headB
	for cA != nil {
		lA ++
		cA = cA.Next
	}
	for cB != nil {
		lB ++
		cB  = cB.Next
	}
	
	diff := 0 
	cA = headA
	cB = headB
	// 步骤1：对齐
	if lA > lB {
		diff = lA - lB
		for diff >=0 {
			cA = headA.Next
			diff --
		}
	} else {
		diff = lB - lA
		for diff >= 0 {
			cB = headB.Next
			diff --
		}
	}

	// 步骤2：双指针判断是否相交
	for cA != nil && cB != nil {
		if cA == cB {
			return cA
		}
		cA = cA.Next
		cB = cB.Next
	}
	return nil
}