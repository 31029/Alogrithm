package linknode

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := MakeLinkList([]int{1, 2})
	head = RemoveNthFromEnd2(head, 2)

	fmt.Printf("after remove: %s", PringLinkList2(head))
}
