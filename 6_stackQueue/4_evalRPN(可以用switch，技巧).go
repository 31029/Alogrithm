package stack_queue

import "strconv"

func EvalRPN(tokens []string) int {
	//({10 * [6 / (12*-11)]}+17) 5 +
	stack := IntStackConstructor()
	
	for _, v := range tokens {
		// if []byte(v)[0] > byte()
		intV, err:= strconv.Atoi(v)
		if err == nil {
			stack.Push(intV)
		} else {
			R := stack.Pop()
			L := stack.Pop()
			tmp := 0
			if v == "*" {
				tmp = L * R
			} else if v == "/" {
				tmp = L / R
			} else if v == "+" {
				tmp = L + R
			} else if v == "-" {
				tmp = L - R
			}
			stack.Push(tmp)
		}
	}

	return stack.Pop()
}

func evalRPN_answer(tokens []string) int {
    stack := []int{}
    for _, token := range tokens {
        val, err := strconv.Atoi(token)
        if err == nil {
            stack = append(stack, val)
        } else {
            num1, num2 := stack[len(stack)-2], stack[len(stack)-1]  // 技巧可以省一个栈的搭建。
            stack = stack[:len(stack)-2]
            switch token {
            case "+":
                stack = append(stack, num1+num2)
            case "-":
                stack = append(stack, num1-num2)
            case "*":
                stack = append(stack, num1*num2)
            default:
                stack = append(stack, num1/num2)
            }
        }
    }
    return stack[0]
}