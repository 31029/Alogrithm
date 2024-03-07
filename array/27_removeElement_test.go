package array

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	//nums := []int{3, 3}
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	t.Logf("result: %d", RemoveElement(nums, val))
}
