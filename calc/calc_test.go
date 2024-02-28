package calc

import (
	"strings"
	"testing"
)

func TestCalc(t *testing.T) {
	var cases = []struct {
		expected string
		input    string
	}{
		{
			input: "1+2+3+4", expected: "10",
		},
		{
			input: "(4/(11-9))", expected: "2",
		},
		{
			input: "(1+2)*(1+1)", expected: "6",
		},
		{
			input: "(4/(11-9))*(3+9)", expected: "24",
		},
		{
			input: "1+(2*3-19)", expected: "-12",
		},
		{
			input: "(1+2*7)/3-19", expected: "-14",
		},
		{
			input: "((1+3-1)*2/3+1)*1/3", expected: "1",
		},
		{
			input: "(((1+3-1)*2/3+1)*1/3)-5", expected: "-4",
		},
		{
			input: "((1+3-1)*2/3-5)", expected: "-3",
		},
	}

	for _, item := range cases {
		result, _ := Calc(strings.NewReader(item.input))
		if result != item.expected {
			t.Errorf("Failed: %s = %s; want %s", item.input, result, item.expected)
		} else {
			t.Logf("ОК: %s = %s", item.input, result)
		}
	}
}
