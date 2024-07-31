package stack_queue

type MyStack struct {
	container []int
	len       int
}

func StackConstructor() MyStack {
	return MyStack{
		container: []int{},
		len:       0,
	}
}

func (s *MyStack) Push(x int) {
	s.container = append(s.container, x)
	s.len += 1
}

func (s *MyStack) Pop() int {
	if s.Empty() {
		return -1
	}

	top := s.container[s.len-1]
	s.container = s.container[:s.len-1]
	s.len -= 1
	return top
}

func (s *MyStack) Peek() int {
	if s.Empty() {
		return -1
	}
	return s.container[s.len-1]
}

func (s *MyStack) Empty() bool {
	return s.len == 0
}
