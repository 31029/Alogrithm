package stack_queue

func IsValid(s string) bool {
	stack := StackConstructor()
	if len(s) == 0 || len(s) % 2 > 0 {
		return false
	}

	sbyte := []byte(s)
	for _, v := range sbyte {
		if stack.Peek() == byte(0) {
			stack.Push(v)
		} else if v == []byte("}")[0] && stack.Peek() == []byte("{")[0] {
			stack.Pop()
		} else if v == []byte("]")[0] && stack.Peek() == []byte("[")[0] {
			stack.Pop()
		} else if v == []byte(")")[0] && stack.Peek() == []byte("(")[0] {
			stack.Pop()
		} else {
			stack.Push(v)
		}
	}

	return stack.Empty()
}