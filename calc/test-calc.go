package calc

import (
	"fmt"
	"strings"
)

func TestCalc() {
	var cases = []struct {
		expected string
		input    string
	}{
		{
			input: "1+2+3+4", expected: "10",
		},
		{
			input: "15-2-3-4", expected: "6",
		},
		{
			input: "(3-4)", expected: "-1",
		},
		{
			input: "4/2+11*(3+9)", expected: "134",
		},
		{
			input: "1+(2*3-19)", expected: "-12",
		},
		{
			input: "1+2*3-19", expected: "-12",
		},
	}

	for _, item := range cases {
		result, _ := Calc(strings.NewReader(item.input))
		fmt.Printf("Expected: %s, result: %s, input: %s\n", item.expected, result, item.input)

	}
}
