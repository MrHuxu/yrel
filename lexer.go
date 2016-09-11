package yrel

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var idPattern = `([A-Z_a-z][A-Z_a-z0-9]*|=|!=|==|<=|>=|&&|\|\|)`
var numPattern = `([0-9]+)`
var strPattern = `(\"[\S\s]*\")`
var commentPattern = `(//[\S\s]*)`
var pattern = idPattern + "|" + numPattern + "|" + strPattern + "|" + commentPattern

type Lexer struct {
	Queue   []*Token
	HasMore bool
	Reader  *bufio.Reader
}

func NewLexer(file *os.File) *Lexer {
	return &Lexer{
		Queue:   []*Token{},
		HasMore: true,
		Reader:  bufio.NewReader(file),
	}
}

func (l Lexer) Read() {
	matcher, _ := regexp.Compile(pattern)
	scanner := bufio.NewScanner(l.Reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(matcher.FindAllStringSubmatch(scanner.Text(), -1))
	}
}

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
	LineNum int
	Value   int
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
	LineNum int
	Text    string
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
	LineNum int
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
