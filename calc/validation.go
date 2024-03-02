package calc

import (
	"errors"
)

func Validate(inputExpression string) ([]string, error) {
	if len(inputExpression) == 0 {
		return nil, errors.New("empty")
	}
	operandTypes := map[string]int{
		")": 1,
		"+": 1,
		"-": 1,
		"*": 1,
		"/": 1,
	}
	var result []string
	var tmp string       // записываем число
	var bracketCheck int // счетчик скобок для проверки одинакового кол-ва открывающихся и закрывающихся
	for i := 0; i < len(inputExpression); i++ {
		if inputExpression[i] == 32 { // пробелы пропускаем
			continue
		}
		if inputExpression[i] < 58 && inputExpression[i] > 47 { //если число
			tmp = tmp + string(inputExpression[i]) //добавляем в буфер
			continue
		}
		if inputExpression[i] == 46 { // обработка точки
			if len(tmp) > 0 {
				tmp = tmp + string(inputExpression[i])
			} else {
				return nil, errors.New("before . must be digit")
			}
			continue
		}
		//то же самое, что и inputExpression[i] == 41 || inputExpression[i] == 42 || inputExpression[i] == 43 || inputExpression[i] == 45 || inputExpression[i] == 47 {
		if operandTypes[string(inputExpression[i])] == 1 { //встретили ) * - + /
			if inputExpression[i] == 41 { //41 - код ")"
				bracketCheck--
			}
			if len(tmp) > 0 { //если буфер заполнен, то
				result = append(result, tmp) // добавляем число в результирующую строку
				tmp = ""
				result = append(result, string(inputExpression[i])) // и знак добавляем
				continue
			}
			if len(result) == 0 {
				return nil, errors.New("expression doesn't start with digit")
			}
			//если в строке перед * - + / не идет число
			if result[len(result)-1] < "0" || result[len(result)-1] > "9" {
				if result[len(result)-1] == ")" {
					result = append(result, string(inputExpression[i])) // добавляем )
					continue
				}
				if result[len(result)-1] == "(" && inputExpression[i] == 45 { //45 - код "-"
					tmp = string(inputExpression[i])
					continue
				}
				return nil, errors.New("incorrect sequence of arithmetic signs")
			}
		}
		if inputExpression[i] == 40 && (len(result) == 0 || result[len(result)-1] == "(" || operandTypes[result[len(result)-1]] == 1) {
			// встретили (, до него должны быть +, -, *, /, , (
			bracketCheck++
			result = append(result, string(inputExpression[i])) // добавляем (
			tmp = ""
		} else {
			return nil, errors.New("invalid symbols")
		}
	}
	if bracketCheck != 0 {
		return nil, errors.New("different number of brackets")
	}
	if len(tmp) > 0 { //если буфер заполнен, то
		result = append(result, tmp)
	}
	return result, nil
}
