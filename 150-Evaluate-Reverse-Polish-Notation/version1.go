package evaluatereversepolishnotation

import "strconv"

type Stack struct {
	list []int
}

func evalRPN(tokens []string) int {

	if len(tokens) == 0 {
		return 0
	}

	stack := new(Stack)
	stack.list = make([]int, len(tokens))

	for i := 0; i < len(tokens); i++ {
		num, err := strconv.Atoi(tokens[i])

		if err == nil {
			stack.Push(num)
		} else {
			result := 0

			switch tokens[i] {
			case "+":
				result = stack.add()
				break
			case "-":
				result = stack.sub()
				break
			case "*":
				result = stack.mul()
				break
			case "/":
				result = stack.div()
				break
			}

			stack.Push(result)
		}
	}
	return stack.Pop()
}

func (s *Stack) Push(num int) {
	s.list = append(s.list, num)
}

func (s *Stack) Pop() int {
	var number int

	number = s.list[len(s.list)-1]

	s.list = s.list[:len(s.list)-1]

	return number
}

func (s *Stack) add() int {
	num2 := s.Pop()
	num1 := s.Pop()
	return num1 + num2
}

func (s *Stack) sub() int {
	num2 := s.Pop()
	num1 := s.Pop()
	return num1 - num2
}

func (s *Stack) mul() int {
	num2 := s.Pop()
	num1 := s.Pop()
	return num1 * num2
}

func (s *Stack) div() int {
	num2 := s.Pop()
	num1 := s.Pop()
	return num1 / num2
}
