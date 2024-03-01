package calc

import (
	"errors"
	"fmt"
	"io"
)

func Validate(input io.Reader) ([]string, error) {
	operandTypes := map[string]int{
		" ": 0,
		")": 1,
		"+": 1,
		"-": 1,
		"*": 1,
		"/": 1,
		"(": 4,
	}
	var inputExpression string
	var result []string
	_, err := fmt.Fscan(input, &inputExpression)
	if err != nil {
		return nil, err
	}
	var tmp string       // записываем число
	var bracketCheck int // счетчик скобок для проверки одинакового кол-ва открывающихся и закрывающихся
	for i := 0; i < len(inputExpression); i++ {
		if inputExpression[i] == 32 { // пробелы пропускаем
			return nil, errors.New("Invalid input")
		}
		if inputExpression[i] < 58 && inputExpression[i] > 47 { //если число
			tmp = tmp + string(inputExpression[i]) //добавляем в буфер
			continue
		}
		if inputExpression[i] == 46 && len(tmp) > 0 { // обработка точки
			tmp = tmp + string(inputExpression[i])
			continue
		}
		//то же самое, что и inputExpression[i] == 41 || inputExpression[i] == 42 || inputExpression[i] == 43 || inputExpression[i] == 45 || inputExpression[i] == 47 {
		if operandTypes[string(inputExpression[i])] == 1 { //встретили ) * - + /
			if inputExpression[i] == 41 { //41 - код ")"
				bracketCheck--
			}
			if len(tmp) > 0 { //если буфер заполнен, то
				result = append(result, tmp) // добавляем число в результирующую строку
				tmp = "" + string(inputExpression[i])
				result = append(result, tmp) // и знак добавляем
				tmp = ""
				continue
			}
			//если в строке перед ) * - + / не идет число
			if result[len(result)-1] < "0" || result[len(result)-1] > "9" {
				if result[len(result)-1] == ")" {
					tmp = "" + string(inputExpression[i])
					result = append(result, tmp) // добавляем )
					tmp = ""
					continue
				}
				if result[len(result)-1] == "(" && inputExpression[i] == 45 { //45 - код "-"
					tmp = "" + string(inputExpression[i])
				} else {
					return nil, errors.New("Incorrect sequence of arithmetic signs")
				}
			}
		} else if inputExpression[i] == 40 && (len(result) == 0 || result[len(result)-1] == "(" || result[len(result)-1] == "+" || result[len(result)-1] == "-" || result[len(result)-1] == "*" || result[len(result)-1] == "/") {
			// встретили (, до него должны быть +, -, *, /, , (
			bracketCheck++
			tmp = "" + string(inputExpression[i])
			result = append(result, tmp) // добавляем (
			tmp = ""
		} else {
			return nil, errors.New("Invalid symbols")
		}
	}
	if bracketCheck != 0 {
		return nil, errors.New("Different number of brackets")
	}
	if len(tmp) > 0 { //если буфер заполнен, то
		result = append(result, tmp)
	}
	return result, nil
}
