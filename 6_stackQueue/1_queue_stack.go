package stack_queue

type IntQueue struct {
	stack_first IntStack
	stack_last  IntStack
	len         int
}

func IntQueueConstructor() IntQueue {
	return IntQueue{
		stack_first: IntStackConstructor(),
		stack_last:  IntStackConstructor(),
		len:         0,
	}
}

func (q *IntQueue) Push(x int) {
	q.stack_first.Push(x)
	q.len += 1
}

func (q *IntQueue) Pop() int {
	if q.Empty() {
		return -999
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

func (q *IntQueue) Peek() int {
	if q.Empty() {
		return -999
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

func (q *IntQueue) PopBottom() int {
	if q.Empty() {
		return -999
	}
	q.len -= 1
	return q.stack_first.Pop()
}


func (q *IntQueue) Bottom() int {
	if q.Empty() {
		return -999
	}
	return q.stack_first.Peek()
}

func (q *IntQueue) Empty() bool {
	return q.len == 0
}

/**
 * Your IntQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
