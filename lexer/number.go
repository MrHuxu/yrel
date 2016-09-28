package lexer

import (
	"fmt"
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

func (nu NumToken) getResultAndHandleError(result *NumToken, n NumToken, num int, op string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			*result = NumToken{
				Line:  n.Line,
				Value: -1,
			}
		}
	}()
	*result = NumToken{
		Line:  n.Line,
		Value: ExecCalc(n.Value, num, op),
	}
}

func (n NumToken) Calc(t Token, op string) Token {
	result := &NumToken{}
	var num int
	if t != nil {
		num, _ = strconv.Atoi(t.GetText())
	} else {
		num = -1
	}
	n.getResultAndHandleError(result, n, num, op)

	return *result
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
