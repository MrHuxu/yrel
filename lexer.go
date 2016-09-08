package yrel

import (
	"strconv"
)

type Token interface {
	IsNumber() bool
	IsIdentifier() bool
	IsString() bool
	GetText() string
}

func GetTokenType(t Token) string {
	var tokenType string
	if t.IsNumber() {
		tokenType = "Number"
	} else if t.IsIdentifier() {
		tokenType = "Identifier"
	} else if t.IsString() {
		tokenType = "String"
	} else {
		tokenType = "Undefined"
	}

	return tokenType
}

type NumToken struct {
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

func (n NumToken) GetText() string {
	return strconv.Itoa(n.Value)
}

func (n NumToken) GetNumber() int {
	return n.Value
}

type IdToken struct {
	Text string
}

func (i IdToken) IsNumber() bool {
	return false
}

func (i IdToken) IsIdentifier() bool {
	return true
}

func (i IdToken) IsString() bool {
	return false
}

func (i IdToken) GetText() string {
	return i.Text
}

type StrToken struct {
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

func (s StrToken) GetText() string {
	return s.Literal
}
