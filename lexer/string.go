package lexer

import (
	"fmt"
	"strconv"
)

type StrToken struct {
	Util
	Value string `json:"value"`
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
	return s.Value
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
		panic("Error: Unsupported operation.")
	case "Neg":
		panic("Error: Unsupported operation.")
	case "*":
		for ; num > 0; num-- {
			result += str
		}
	case "/":
		panic("Error: Unsupported operation.")
	case "%":
		panic("Error: Unsupported operation.")
	}

	return result
}

func (s StrToken) getResultAndHandleError(result *StrToken, num int, op string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s ", r)
			*result = StrToken{s.Util, "-1"}
		}
	}()
	*result = StrToken{s.Util, s.ExecCalc(s.Value, num, op)}
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
		result = s.Value == t.GetText()
	case "!=":
		result = s.Value != t.GetText()
	case "!":
		result = false
	}

	return BoolToken{s.Util, result}
}

func (s StrToken) Logic(t Token, op string) Token {
	var result bool
	switch op {
	case "&&":
		result = true && t.True()
	case "||":
		result = true || t.True()
	}

	return BoolToken{s.Util, result}
}
