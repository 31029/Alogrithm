package array

import (
	"testing"
)

func Test_4(t *testing.T) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	
	t.Logf("result: %v", MinSubArrayLen(target, nums))
}
