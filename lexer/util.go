package lexer

import (
	_ "fmt"
	"regexp"
)

func BuildLexerMatcher() *regexp.Regexp {
	boolPattern := `(true|false)`
	numPattern := `([0-9]+)`
	strPattern := `(\"[\S\s]*\")`
	idPattern := `([A-Z_a-z][A-Z_a-z0-9]*)`
	commentPattern := `(//[\S\s]*)`
	equalPattern := `(==)`
	unequalPattern := `(!=)`
	opPattern := `(=|!|>|<|\+|-|\*|/|%)`
	logicAndPattern := `(&&)`
	logicOrPattern := `(||)`
	pattern := boolPattern + "|" + numPattern + "|" + strPattern + "|" + idPattern + "|" + commentPattern + "|" + equalPattern + "|" + unequalPattern + "|" + opPattern + "|" + logicAndPattern + "|" + logicOrPattern

	matcher, _ := regexp.Compile(pattern)
	return matcher
}
