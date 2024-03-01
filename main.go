package main

import (
	"fmt"
	"hw_go/calc"
	"os"
)

func main() {
	expression, err := calc.Validate(os.Stdin)
	if expression == nil {
		fmt.Printf("%s", err)
	} else {
		result, _ := calc.Calc(expression)
		fmt.Printf("Result: %s", result)
	}
}
