package calc

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestCalc(t *testing.T) {
	var cases = []struct {
		expected string
		input    string
	}{
		{
			input: "5/(6-2)/2+100*2", expected: "200.625",
		},
		{
			input: "5.25*4-10*(1+1.5)", expected: "-4",
		},
		{
			input: "((((-1)+2)+3+4))", expected: "8",
		},
		{
			input: "(-1)*((-2)+(3+5)/2*14)", expected: "-54",
		},
		{
			input: "(-1)+2+3+4", expected: "8",
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
			input: "((-1)+(-2)*7)/3-19", expected: "-24",
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
		expression, err := Validate(strings.NewReader(item.input))
		if expression == nil {
			t.Errorf("%s in %s", err, item.input)
			return
		}
		result, _ := Calc(expression)
		require.Equal(t, result, item.expected, "Failed")
	}
}
