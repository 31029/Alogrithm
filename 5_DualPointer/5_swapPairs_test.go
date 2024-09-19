package dualPointer

import (
	"testing"
	linknode "github.com/h31029/alogrithm/2_Linknode"
)

func TestSwapPairs(t *testing.T)  {
	nums := []int{1, 2, 3, 4}
	head := linknode.MakeLinkList(nums)
	SwapPairs(head)
}