package hashmap

import "testing"

func TestThreeSum(t *testing.T){
	nums := []int{-2,0,1,1,2}
	t.Logf("total count: %v", ThreeSum(nums))
}