package main

import (
	"bufio"
	"fmt"
	"hw_go/calc"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	expression, err := calc.Validate(scanner.Text())
	if expression == nil {
		fmt.Printf("%s", err)
	} else {
		result, _ := calc.Calc(expression)
		fmt.Printf("Result: %s", result)
	}
}
