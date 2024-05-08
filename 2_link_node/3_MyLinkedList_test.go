package linknode

import "testing"

func TestMyLinkNode(t *testing.T)  {
	obj_o := Constructor()
	obj := &obj_o
	obj.AddAtHead(1)
	t.Logf(PringLinkList(obj))

	obj.AddAtTail(3);
	t.Logf(PringLinkList(obj))

	obj.AddAtIndex(1,2);
	t.Logf(PringLinkList(obj))

	obj.DeleteAtIndex(1);
	t.Logf(PringLinkList(obj))
}