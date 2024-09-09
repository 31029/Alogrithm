package hashmap

import "testing"

func TestFourSum(t *testing.T){
	//nums := []int{1,0,-1,0,-2,2}
	nums := []int{2,2,2,2,2}
	t.Logf("total count: %v", FourSum(nums, 8))
}