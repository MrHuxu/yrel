package lexer

import (
	"regexp"
)

func BuildLexerMatcher() *regexp.Regexp {
	idPattern := `([A-Z_a-z][A-Z_a-z0-9]*)`
	numPattern := `([0-9]+)`
	strPattern := `(\"[\S\s]*\")`
	boolPattern := `(true|false)`
	commentPattern := `(//[\S\s]*)`

	equalPattern := `(==)`
	unequalPattern := `(!=)`
	logicAndPattern := `(&&)`
	logicOrPattern := `(||)`

	pattern := boolPattern + "|" + numPattern + "|" + strPattern + "|" + idPattern + "|" + commentPattern + "|" + equalPattern + "|" + unequalPattern + "|" + logicAndPattern + "|" + logicOrPattern

	matcher, _ := regexp.Compile(pattern)
	return matcher
}
