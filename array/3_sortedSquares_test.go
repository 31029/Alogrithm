package array

import (
	"testing"
)

func Test_3(t *testing.T) {
	//nums := []int{3, 3}
	nums := []int{-4,-1,0,3,10}
	t.Logf("result: %v", SortedSquares(nums))
}
