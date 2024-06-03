package hashmap

import "testing"

func TestFourSumCount(t *testing.T){
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	t.Logf("total count: %v", FourSumCount(nums1, nums2, nums3, nums4))
}