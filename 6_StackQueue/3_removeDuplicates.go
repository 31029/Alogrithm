package stack_queue

func RemoveDuplicates(s string) string {
	stack := StackConstructor()
	s_byte := []byte(s)

	// 字符串示例："abbaca"
	for _, v := range s_byte {
		if stack.Peek() != v {
			stack.Push(v)
		} else if !stack.Empty() {
			stack.Pop()
		}
	}

	result := []byte{}
	for i := 0; i < stack.len; i++ {
		result = append(result, byte(0))
	}

	for i := stack.len-1; i >= 0; i-- {
		result[i] = stack.Pop()
	}

	return string(result)
}