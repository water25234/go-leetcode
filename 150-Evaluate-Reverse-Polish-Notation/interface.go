package evaluatereversepolishnotation

import "strconv"

type Stack struct {
	iStack IStack
	list   []int
}

type IStack interface {
	addRPN() int
	subRPN() int
	mulRPN() int
	divRPN() int
	pushRPN(num int)
	popRPN() int
}

func NewIStack(iStack IStack) *Stack {
	return &Stack{
		iStack: iStack,
	}
}

func evalRPN(tokens []string) int {

	if len(tokens) == 0 {
		return 0
	}

	stack := NewIStack(&Stack{})

	for i := 0; i < len(tokens); i++ {
		num, err := strconv.Atoi(tokens[i])

		if err == nil {
			stack.iStack.pushRPN(num)
		} else {
			result := 0

			switch tokens[i] {
			case "+":
				result = stack.iStack.addRPN()
				break
			case "-":
				result = stack.iStack.subRPN()
				break
			case "*":
				result = stack.iStack.mulRPN()
				break
			case "/":
				result = stack.iStack.divRPN()
				break
			}
			stack.iStack.pushRPN(result)
		}
	}
	return stack.iStack.popRPN()
}

func (s *Stack) pushRPN(num int) {
	s.list = append(s.list, num)
}

func (s *Stack) popRPN() int {
	var number int

	number = s.list[len(s.list)-1]

	s.list = s.list[:len(s.list)-1]

	return number
}

func (s *Stack) addRPN() int {
	num2 := s.popRPN()
	num1 := s.popRPN()
	return num1 + num2
}

func (s *Stack) subRPN() int {
	num2 := s.popRPN()
	num1 := s.popRPN()
	return num1 - num2
}

func (s *Stack) mulRPN() int {
	num2 := s.popRPN()
	num1 := s.popRPN()
	return num1 * num2
}

func (s *Stack) divRPN() int {
	num2 := s.popRPN()
	num1 := s.popRPN()
	return num1 / num2
}
