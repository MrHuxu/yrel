package lexer

import (
	"fmt"
	"strconv"
)

type BoolToken struct {
	Util
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

func (b BoolToken) getResultAndHandleError(result *NumToken, n BoolToken, num int, op string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s ", r)
			*result = NumToken{n.Util, -1}
		}
	}()
	*result = NumToken{n.Util, ExecCalc(0, num, op)}
}

func (b BoolToken) Calc(t Token, op string) Token {
	result := &NumToken{}
	var num int
	if t != nil {
		num, _ = strconv.Atoi(t.GetText())
	} else {
		num = -1
	}
	b.getResultAndHandleError(result, b, num, op)

	return *result
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

	return BoolToken{b.Util, result}
}

func (b BoolToken) Logic(t Token, op string) Token {
	var result bool
	switch op {
	case "&&":
		result = b.Value && t.True()
	case "||":
		result = b.Value || t.True()
	}

	return BoolToken{b.Util, result}
}
