package lexer

import (
	"strconv"
)

type BoolToken struct {
	*Line
	Value bool
}

func (b BoolToken) IsNumber() bool {
	return false
}

func (b BoolToken) IsIdentifier() bool {
	return false
}

func (b BoolToken) IsString() bool {
	return false
}

func (b BoolToken) IsBool() bool {
	return true
}

func (b BoolToken) True() bool {
	return b.Value
}

func (b BoolToken) GetText() string {
	if b.Value {
		return "true"
	} else {
		return "false"
	}
}
func (b BoolToken) Calc(t Token, op string) Token {
	var result int
	var num int
	if t != nil {
		num, _ = strconv.Atoi(t.GetText())
	} else {
		num = 0
	}
	switch op {
	case "+":
		result = 0 + num
	case "-":
		result = 0 - num
	case "*":
		result = 0 * num
	case "/":
		result = 0 / num
	case "%":
		result = 0 % num
	}

	return NumToken{
		Line:  b.Line,
		Value: result,
	}
}

func (b BoolToken) Comp(t Token, op string) Token {
	var result bool
	num, _ := strconv.Atoi(t.GetText())
	switch op {
	case ">":
		result = 0 > num
	case "<":
		result = 0 < num
	case "==":
		result = 0 == num
	case "!=":
		result = 0 != num
	case "!":
		result = !b.Value
	}

	return BoolToken{
		Line:  b.Line,
		Value: result,
	}
}

func (b BoolToken) Logic(t Token, op string) Token {
	var result bool
	switch op {
	case "&&":
		result = b.Value && t.True()
	case "||":
		result = b.Value || t.True()
	}

	return BoolToken{
		Line:  b.Line,
		Value: result,
	}
}
