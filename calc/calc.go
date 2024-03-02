package calc

import (
	"errors"
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

func calcResult(operation string, first float64, second float64) (float64, error) {
	if operation == "+" {
		return first + second, nil
	}
	if operation == "-" {
		return first - second, nil
	}
	if operation == "*" {
		return first * second, nil
	}
	if operation == "/" {
		if second != 0 {
			return first / second, nil
		}
		return 0, errors.New("0 can't be second argument in division")
	}
	return 0, nil
}

func Calc(expression []string) (string, error) {
	operandTypes := map[string]int{
		" ": 1,
		")": 0,
		"+": 2,
		"-": 2,
		"*": 3,
		"/": 3,
		"(": 0,
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
				}
				secondNumber, _ := strconv.ParseFloat(numbers.Pop().(string), 64)
				firstNumber, _ := strconv.ParseFloat(numbers.Pop().(string), 64)
				number, err := calcResult(operands.Top().(string), firstNumber, secondNumber)
				if err != nil {
					return "0", err
				}
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
