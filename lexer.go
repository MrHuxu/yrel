package yrel

import (
	"bufio"
	_ "fmt"
	"os"
	"regexp"
	"strconv"
)

type Lexer struct {
	Queue   []Token
	HasMore bool
}

func NewLexer(file *os.File) *Lexer {
	l := &Lexer{
		Queue:   []Token{},
		HasMore: true,
	}

	matcher := getRegExpMatcher()
	scanner := bufio.NewScanner(bufio.NewReader(file))
	count := 0
	for scanner.Scan() {
		count++
		l.addToken(count, matcher.FindAllStringSubmatch(scanner.Text(), -1))
	}

	return l
}

func getRegExpMatcher() *regexp.Regexp {
	idPattern := `([A-Z_a-z][A-Z_a-z0-9]*|=|!=|==|<=|>=|&&|\|\|)`
	numPattern := `([0-9]+)`
	strPattern := `(\"[\S\s]*\")`
	commentPattern := `(//[\S\s]*)`
	pattern := idPattern + "|" + numPattern + "|" + strPattern + "|" + commentPattern

	matcher, _ := regexp.Compile(pattern)
	return matcher
}

func (l *Lexer) addToken(ln int, elements [][]string) {
	for _, ele := range elements {
		if ele[1] != "" {
			l.Queue = append(l.Queue, &IdToken{&Line{ln}, ele[0]})
		} else if ele[2] != "" {
			num, _ := strconv.Atoi(ele[0])
			l.Queue = append(l.Queue, &NumToken{&Line{ln}, num})
		} else if ele[3] != "" {
			l.Queue = append(l.Queue, &StrToken{&Line{ln}, ele[0]})
		}
	}
}

func (l *Lexer) Read() Token {
	var t Token

	if len(l.Queue) >= 1 {
		t, l.Queue = l.Queue[0], l.Queue[1:]
	} else {
		t = EOF
	}
	return t
}

type Token interface {
	IsNumber() bool
	IsIdentifier() bool
	IsString() bool
	GetText() string
	GetLineNumber() int
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

type Line struct {
	LineNum int
}

func (l *Line) GetLineNumber() int {
	return l.LineNum
}

type IdToken struct {
	*Line
	Text string
}

var EOF = &IdToken{&Line{-1}, ""}

func (i *IdToken) IsNumber() bool {
	return false
}

func (i *IdToken) IsIdentifier() bool {
	return true
}

func (i *IdToken) IsString() bool {
	return false
}

func (i *IdToken) GetText() string {
	return i.Text
}

type NumToken struct {
	*Line
	Value int
}

func (n *NumToken) IsNumber() bool {
	return true
}

func (n *NumToken) IsIdentifier() bool {
	return false
}

func (n *NumToken) IsString() bool {
	return false
}

func (n *NumToken) GetText() string {
	return strconv.Itoa(n.Value)
}

func (n *NumToken) Plus(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value + m.Value,
	}
}

func (n *NumToken) Sub(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value - m.Value,
	}
}

func (n *NumToken) Mul(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value * m.Value,
	}
}

func (n *NumToken) Div(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value / m.Value,
	}
}

func (n *NumToken) Mod(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value % m.Value,
	}
}

func (n *NumToken) BiteAnd(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value & m.Value,
	}
}

func (n *NumToken) BiteOr(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value | m.Value,
	}
}

func (n *NumToken) Neg() *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: -n.Value,
	}
}

func (n *NumToken) GetNumber() int {
	return n.Value
}

type StrToken struct {
	*Line
	Literal string
}

func (s *StrToken) IsNumber() bool {
	return false
}

func (s *StrToken) IsIdentifier() bool {
	return false
}

func (s *StrToken) IsString() bool {
	return true
}

func (s *StrToken) GetText() string {
	return s.Literal
}

func (s *StrToken) GetLineNumber() int {
	return s.LineNum
}
