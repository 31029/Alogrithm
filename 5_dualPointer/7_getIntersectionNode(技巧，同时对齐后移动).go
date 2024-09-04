package dualPointer


func getIntersectionNode(headA, headB *ListNode) *ListNode {
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