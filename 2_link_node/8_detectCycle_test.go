package linknode

import "testing"

func TestDetectCycle(t *testing.T) {
	head := MakeLinkList([]int{3, 2, 0, -1})
	head.Next.Next.Next.Next = head.Next
	p := DetectCycle(head)
	t.Logf("p.Value: %v", p.Val)
}