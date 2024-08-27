package stack_queue

import "testing"



func TestEvalRPN(t *testing.T)  {
	s := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}
	EvalRPN(s)
}