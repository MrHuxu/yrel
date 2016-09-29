package lexer

import (
	"fmt"
	"strconv"
)

type StrToken struct {
	*Line
	Literal string
}

func (s StrToken) IsNumber() bool {
	return false
}

func (s StrToken) IsIdentifier() bool {
	return false
}

func (s StrToken) IsString() bool {
	return true
}

func (s StrToken) IsBool() bool {
	return false
}

func (s StrToken) GetText() string {
	return s.Literal
}

func (s StrToken) True() bool {
	return true
}

func (s StrToken) ExecCalc(str string, num int, op string) string {
	var result string
	switch op {
	case "+":
		result = str + strconv.Itoa(num)
	case "-":
		panic("unsupported operation")
	case "Neg":
		panic("unsupported operation")
	case "*":
		base := str
		for num > 0 {
			if num%2 == 1 {
				result += base
				num--
			} else {
				base += base
				num >>= 1
			}
		}
	case "/":
		panic("unsupported operation")
	case "%":
		panic("unsupported operation")
	}

	return result
}

func (s StrToken) getResultAndHandleError(result *StrToken, num int, op string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s ", r)
			*result = StrToken{
				Line:    s.Line,
				Literal: "-1",
			}
		}
	}()
	*result = StrToken{
		Line:    s.Line,
		Literal: s.ExecCalc(s.Literal, num, op),
	}
}

func (s StrToken) Calc(t Token, op string) Token {
	result := &StrToken{}
	var num int
	if t != nil {
		num, _ = strconv.Atoi(t.GetText())
	} else {
		num = -1
	}
	s.getResultAndHandleError(result, num, op)

	return *result
}

func (s StrToken) Comp(t Token, op string) Token {
	var result bool
	num, _ := strconv.Atoi(t.GetText())
	switch op {
	case ">":
		result = 0 > num
	case "<":
		result = 0 < num
	case "==":
		result = s.Literal == t.GetText()
	case "!=":
		result = s.Literal != t.GetText()
	case "!":
		result = false
	}

	return BoolToken{
		Line:  s.Line,
		Value: result,
	}
}

func (s StrToken) Logic(t Token, op string) Token {
	var result bool
	switch op {
	case "&&":
		result = true && t.True()
	case "||":
		result = true || t.True()
	}

	return BoolToken{
		Line:  s.Line,
		Value: result,
	}
}
