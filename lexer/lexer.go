package lexer

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
	boolPattern := `(true|false)`
	commentPattern := `(//[\S\s]*)`
	pattern := idPattern + "|" + numPattern + "|" + strPattern + "|" + boolPattern + "|" + commentPattern

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
		} else if ele[4] != "" {
			val := ele[4] == "true"
			l.Queue = append(l.Queue, &BoolToken{&Line{ln}, val})
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
