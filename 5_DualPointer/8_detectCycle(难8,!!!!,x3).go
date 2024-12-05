package dualPointer

import (
	linknode "github.com/h31029/alogrithm/2_Linknode"
)

// LeetCode-142: 环形链表 II 【https://leetcode.cn/problems/linked-list-cycle-ii/description/】
/* 	
描述：给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。
*/
func DetectCycle(head *linknode.ListNode) *linknode.ListNode {
	// 解题思路1：双指针，快慢指针，其中慢指针一次移动一步，快指针一次移动两步。
	// 以上方法只能用于初步判断是否存在环，至于环的入口需要在稿纸上进行一些数学推理，x(起点到入口的距离),y(相遇节点距离x的距离),z(相遇之后的距离)
	// 可以只记结论：相遇之后，分别以头节点和相遇节点为起点，一次只走一步，二者相遇时，则一定是入口节点。

	// 解题思路2: hashmap，key设为*LinkNode, value为0，1(出现)。
	var meetNode *linknode.ListNode
	low, fast := head, head

	// 找到 meetCount
	for fast != nil {
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		low = low.Next
		if low == fast {
			meetNode = fast
			break
		}
	}
	if fast == nil {
		return nil
	}

	// 可以只记结论：相遇之后，分别以头节点和相遇节点为起点，一次只走一步，二者相遇时，则一定是入口节点。
	index1 := head
	index2 := meetNode
	for index1 != index2 {
		index1 = index1.Next
		index2 = index2.Next
	}
	return index1
}