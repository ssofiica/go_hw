package calc

import (
	"fmt"
	"strconv"
)

type Stack struct {
	elements []interface{}
	size     int
}

func (stack *Stack) Push(element interface{}) {
	stack.elements = append(stack.elements, element)
	stack.size++
}

func (stack *Stack) Pop() interface{} {
	if stack.size == 0 {
		return nil
	}
	top := stack.elements[stack.size-1]
	stack.elements = stack.elements[:stack.size-1]
	stack.size--
	return top
}

func (stack *Stack) IsEmpty() bool {
	if stack.size == 0 {
		return true
	}
	return false
}

func (stack *Stack) Top() interface{} {
	if stack.size == 0 {
		return nil
	}
	return stack.elements[stack.size-1]
}

func calcResult(operation string, first float64, second float64) float64 {
	if operation == "+" {
		return first + second
	}
	if operation == "-" {
		return first - second
	}
	if operation == "*" {
		return first * second
	}
	if operation == "/" {
		if second != 0 {
			return first / second
		}
		panic("Делить на ноль нельзя")
	}
	return 0
}

func Calc(expression []string) (string, error) {
	operandTypes := map[string]int{
		" ": 1,
		")": 0,
		"+": 2,
		"-": 2,
		"*": 3,
		"/": 3,
		"(": 4,
	}
	expression = append(expression, " ")

	numbers := &Stack{elements: make([]interface{}, 0), size: 0}
	operands := &Stack{elements: make([]interface{}, 0), size: 0}

	for _, element := range expression {
		// обрабатываем арифм. знаки, )
		if element == "*" || element == "-" || element == "/" || element == "+" || element == ")" || element == " " {
			if operands.IsEmpty() && element != ")" {
				operands.Push(element)
				continue
			}
			for operandTypes[operands.Top().(string)] >= operandTypes[element] {
				if operands.Top() == "(" && element == ")" {
					operands.Pop()
					break
				} else if operands.Top() == "(" && element != ")" {
					break
				}
				secondNumber, _ := strconv.ParseFloat(numbers.Pop().(string), 64)
				firstNumber, _ := strconv.ParseFloat(numbers.Pop().(string), 64)
				number := calcResult(operands.Top().(string), firstNumber, secondNumber)
				stringNumber := fmt.Sprintf("%v", number)
				numbers.Push(stringNumber)
				operands.Pop()
				if operands.IsEmpty() {
					break
				}
			}
			if element != ")" {
				operands.Push(element)
			}
		} else if element == "(" {
			operands.Push(element)
		} else {
			numbers.Push(element)
		}
	}
	return numbers.Pop().(string), nil
}
