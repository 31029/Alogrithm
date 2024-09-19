package dualPointer

import (
	linknode "github.com/h31029/alogrithm/2_Linknode"
)

func getIntersectionNode(headA, headB *linknode.ListNode) *linknode.ListNode {
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

	for cA != nil && cB != nil {
		if cA == cB {
			return cA
		}
		cA = cA.Next
		cB = cB.Next
	}
	return nil
}