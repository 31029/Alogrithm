package stack_queue

type MyQueue struct {
	stack_first MyStack
	stack_last  MyStack
	len         int
}

func Constructor() MyQueue {
	return MyQueue{
		stack_first: StackConstructor(),
		stack_last:  StackConstructor(),
		len:         0,
	}
}

func (q *MyQueue) Push(x int) {
	q.stack_first.Push(x)
	q.len += 1
}

func (q *MyQueue) Pop() int {
	if q.Empty() {
		return -1
	}

	oldLen := q.stack_first.len
	for i := 0; i < oldLen; i++ {
		q.stack_last.Push(q.stack_first.Pop())
	}
	result := q.stack_last.Pop()
	for i := 0; i < oldLen-1; i++ {
		q.stack_first.Push(q.stack_last.Pop())
	}

	q.len -= 1
	return result
}

func (q *MyQueue) Peek() int {
	if q.Empty() {
		return -1
	}
	oldLen := q.stack_first.len
	for i := 0; i < oldLen; i++ {
		q.stack_last.Push(q.stack_first.Pop())
	}
	result := q.stack_last.Peek()
	for i := 0; i < oldLen; i++ {
		q.stack_first.Push(q.stack_last.Pop())
	}
	return result
}

func (q *MyQueue) Empty() bool {
	return q.len == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
