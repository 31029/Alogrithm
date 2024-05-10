package linknode

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T)  {
	head := MakeLinkList([]int{1, 2, 3, 4})
	head =  ReverseList(head)
	fmt.Printf("reverse list: %s", PringLinkList2(head))
}