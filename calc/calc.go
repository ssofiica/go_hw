package calc

import (
	"fmt"
	"io"
	"regexp"
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

func calcResult(operation string, first int, second int) int {
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
		return first / second
	}
	return 0
}

func Calc(input io.Reader) (string, error) {
	var inputExpression string
	operandTypes := map[string]int{
		" ": 1,
		")": 0,
		"+": 2,
		"-": 2,
		"*": 3,
		"/": 3,
		"(": 4,
	}
	_, err := fmt.Fscan(input, &inputExpression)
	if err != nil {
		return "0", nil
	}
	re := regexp.MustCompile("\\d*\\.?\\d+|\\*|\\/|\\)|\\(|\\+|\\-")
	expression := re.FindAllString(inputExpression, -1)
	expression = append(expression, " ")

	numbers := &Stack{elements: make([]interface{}, 0), size: 0}
	operands := &Stack{elements: make([]interface{}, 0), size: 0}

	for _, element := range expression {
		if element == "*" || element == "-" || element == "/" || element == "+" || element == ")" || element == " " {
			var err interface{} = 0
			if operands.Top() == nil && element != ")" {
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
				secondNumber, _ := strconv.Atoi(numbers.Pop().(string))
				firstNumber, _ := strconv.Atoi(numbers.Pop().(string))
				number := calcResult(operands.Top().(string), firstNumber, secondNumber)
				numbers.Push(strconv.Itoa(number))
				operands.Pop()
				err = operands.Top()
				if err == nil {
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
