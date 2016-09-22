package lexer

import (
	"strconv"
)

type NumToken struct {
	*Line
	Value int
}

func (n NumToken) IsNumber() bool {
	return true
}

func (n NumToken) IsIdentifier() bool {
	return false
}

func (n NumToken) IsString() bool {
	return false
}

func (n NumToken) IsBool() bool {
	return false
}

func (n NumToken) GetText() string {
	return strconv.Itoa(n.Value)
}

func (n NumToken) True() bool {
	return n.Value != 0
}

func (n NumToken) Calc(t Token, op string) Token {
	var result int
	num, _ := strconv.Atoi(t.GetText())
	switch op {
	case "+":
		result = n.Value + num
	case "-":
		result = n.Value - num
	case "*":
		result = n.Value * num
	case "/":
		result = n.Value / num
	case "%":
		result = n.Value % num
	}

	return NumToken{
		Line:  n.Line,
		Value: result,
	}
}

func (n NumToken) Comp(t Token, op string) Token {
	var result bool
	num, _ := strconv.Atoi(t.GetText())
	switch op {
	case ">":
		result = n.Value > num
	case "<":
		result = n.Value < num
	case "==":
		result = n.Value == num
	case "!=":
		result = n.Value != num
	case "!":
		result = !n.True()
	}

	return BoolToken{
		Line:  n.Line,
		Value: result,
	}
}

func (n NumToken) Logic(t Token, op string) Token {
	var result bool
	switch op {
	case "&&":
		result = n.True() && t.True()
	case "||":
		result = n.True() || t.True()
	}

	return BoolToken{
		Line:  n.Line,
		Value: result,
	}
}
