package main

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
			num2 := stack.Pop() // pop num2
			num1 := stack.Pop() // pop num1

			switch tokens[i] {
			case "+":
				result = add(num1, num2)
				break
			case "-":
				result = sub(num1, num2)
				break
			case "*":
				result = mul(num1, num2)
				break
			case "/":
				result = div(num1, num2)
				break
			}

			stack.Push(result)
		}
	}
	return stack.Pop()
}

func (s *Stack) Push(num int) {
	s.list = append(s.list, num) // push num
}

func (s *Stack) Pop() int {
	var number int

	number = s.list[len(s.list)-1] // get number

	s.list = s.list[:len(s.list)-1] // pop number

	return number
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func sub(num1 int, num2 int) int {
	return num1 - num2
}

func mul(num1 int, num2 int) int {
	return num1 * num2
}

func div(num1 int, num2 int) int {
	return num1 / num2
}
