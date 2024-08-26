package dualPointer


import "testing"


func TestRemoveElement(t *testing.T) {
	//nums := []int{1, 2, 3, 2, 5, 4}
	nums1 := []int{3, 3}

	t.Logf("result: %v", RemoveElement(nums1, 2))
}