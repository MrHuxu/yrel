package lexer

import (
	_ "fmt"
	"regexp"
	"strings"
)

func BuildLexerMatcher() *regexp.Regexp {

	patterns := []string{
		`(\"[\S\s]*\")`,            // for string
		`(true|false)`,             // for bool
		`(print)`,                  // for print statement
		`([0-9]+)`,                 // for number
		`([A-Z_a-z][A-Z_a-z0-9]*)`, // for identifier
		`(//[\S\s]*)`,              // for comment
		`(==)`,                     // for ==
		`(!=)`,                     // for !=
		`(=|!|>|<|\+|-|\*|/|%|\(|\)|{|})`, // for some operators
		`(\&\&)`, // for &&
		`(\|\|)`, // for ||
	}

	matcher, _ := regexp.Compile(strings.Join(patterns, "|"))
	return matcher
}
