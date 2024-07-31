package stack_queue

import "testing"

func TestQueue(t *testing.T)  {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()

	t.Logf("%v %v %v", param_2, param_3, param_4)

}