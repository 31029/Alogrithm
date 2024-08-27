package linknode

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	head := MakeLinkList([]int{})
	head = SwapPairs(head)
	fmt.Printf("reverse list: %s", PringLinkList2(head))
}
