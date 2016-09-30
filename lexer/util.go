package lexer

import (
	_ "fmt"
	"regexp"
	"strings"
)

type Util struct {
	LineNum  int `json:"lineNum"`
	Category int `json:"category"`
}

func (u Util) GetLineNumber() int {
	return u.LineNum
}

func BuildLexerMatcher() *regexp.Regexp {
	patterns := []string{
		`(//[\S\s]*?\n)`,              // for comment
		`(\n)`,                        // for line break
		`(\"[\S\s]*?\")`,              // for string
		`(true|false)`,                // for bool
		`(if|else|elsif|while|print)`, // for built-in statement
		`([0-9]+)`,                    // for number
		`([A-Z_a-z][A-Z_a-z0-9]*)`,    // for identifier
		`(==)`, // for ==
		`(!=)`, // for !=
		`(=|!|>|<|\+|-|\*|/|%|\(|\)|{|}|;)`, // for some operators
		`(\&\&)`, // for &&
		`(\|\|)`, // for ||
	}

	matcher, _ := regexp.Compile(strings.Join(patterns, "|"))
	return matcher
}

func ExecCalc(num1 int, num2 int, op string) int {
	var result int
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "Neg":
		result = 0 - num1
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	case "%":
		result = num1 % num2
	}

	return result
}
