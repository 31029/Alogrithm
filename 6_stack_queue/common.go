package stack_queue

// byte stack
type MyStack struct {
	container []byte
	len       int
}

func StackConstructor() MyStack {
	return MyStack{
		container: []byte{},
		len:       0,
	}
}

func (s *MyStack) Push(x byte) {
	s.container = append(s.container, x)
	s.len += 1
}

func (s *MyStack) Pop() byte {
	if s.Empty() {
		return byte(0)
	}

	top := s.container[s.len-1]
	s.container = s.container[:s.len-1]
	s.len -= 1
	return top
}

func (s *MyStack) Peek() byte {
	if s.Empty() {
		return byte(0)
	}
	return s.container[s.len-1]
}

func (s *MyStack) Empty() bool {
	return s.len == 0
}


// string stack
type StringStack struct {
	container []string
	len       int
}

func StringStackConstructor() StringStack {
	return StringStack{
		container: []string{},
		len:       0,
	}
}

func (s *StringStack) Push(x string) {
	s.container = append(s.container, x)
	s.len += 1
}

func (s *StringStack) Pop() string {
	if s.Empty() {
		return string("0")
	}

	top := s.container[s.len-1]
	s.container = s.container[:s.len-1]
	s.len -= 1
	return top
}

func (s *StringStack) Peek() string {
	if s.Empty() {
		return string("0")
	}
	return s.container[s.len-1]
}

func (s *StringStack) Empty() bool {
	return s.len == 0
}



// int stack
type IntStack struct {
	container []int
	len       int
}

func IntStackConstructor() IntStack {
	return IntStack{
		container: []int{},
		len:       0,
	}
}

func (s *IntStack) Push(x int) {
	s.container = append(s.container, x)
	s.len += 1
}

func (s *IntStack) Pop() int {
	if s.Empty() {
		return -999
	}

	top := s.container[s.len-1]
	s.container = s.container[:s.len-1]
	s.len -= 1
	return top
}

func (s *IntStack) Peek() int {
	if s.Empty() {
		return -999
	}
	return s.container[s.len-1]
}

func (s *IntStack) Empty() bool {
	return s.len == 0
}
