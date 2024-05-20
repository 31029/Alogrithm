package linknode

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	init_array := []int{7, 7, 7, 7}
	head := MakeLinkList(init_array)
	head = RemoveElements(head, 7)

	t.Logf("%s", PringLinkList2(head))
}
