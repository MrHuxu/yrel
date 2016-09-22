%{

package parser

import (
	"fmt"
	"github.com/MrHuxu/yrel/lexer"
	"strconv"
)

var regs = make(map[string]lexer.Token)

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	Void lexer.Token
	Identifier lexer.IdToken
	Number lexer.NumToken
	Bool lexer.BoolToken
	Operator string
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <Void> expr primary calc comp logic

// same for terminals
%token <Number> NUMBER
%token <Identifier> IDENTIFIER
%token <Bool> BOOL
%token <Operator> T_EQUAL T_UNEQUAL T_LOGIC_AND T_LOGIC_OR

%left '='
%left T_LOGIC_AND T_LOGIC_OR '!'
%left '>' '<' T_EQUAL T_UNEQUAL
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */

%%

list	:
		/* empty */
	| list stat '\n'
;

stat :
		expr 												{ fmt.Println($1.GetText()) }
	| IDENTIFIER '=' expr					{ regs[$1.GetText()]  =  $3 }
;

expr :
		calc												{ $$ = $1 }
	| comp												{ $$ = $1 }
	| logic												{ $$ = $1 }
;

logic :
		'(' logic ')'               { $$ = $2 }
	| comp T_LOGIC_AND comp       { $$ = $1.Logic($3, "&&") }
	| comp T_LOGIC_OR comp        { $$ = $1.Logic($3, "||") }
	| '!' comp       							{ $$ = $2.Logic(nil, "!") }
	;

comp :
		'(' comp ')'								{ $$ = $2 }
	| calc '>' calc               { $$ = $1.Comp($3, ">") }
	| calc '<' calc               { $$ = $1.Comp($3, "<") }
	| calc T_EQUAL calc           { $$ = $1.Comp($3, "==") }
	| calc T_UNEQUAL calc         { $$ = $1.Comp($3, "!=") }
	;

calc :
		'(' calc ')'  											{ $$  =  $2 }
	| calc '+' calc												{ $$  =  $1.Calc($3, "+") }
	| calc '-' calc												{ $$  =  $1.Calc($3, "-") }
	| calc '*' calc												{ $$  =  $1.Calc($3, "*") }
	| calc '/' calc												{ $$  =  $1.Calc($3, "/") }
	| calc '%' calc												{ $$  =  $1.Calc($3, "%") }
	| primary  
	;

primary :
		NUMBER         { $$ = $1 }
	| BOOL 					 { $$ = $1 }
	| IDENTIFIER     { $$ = regs[$1.GetText()] }
	;

%%      /*  start  of  programs  */

type Lexer struct {
	S string
	Pos int
}

func (l *Lexer) Lex(lval *yySymType) int {
	if l.Pos >= len(l.S) {
		return 0
	}

	// build regexp matcher for lexer process
	matcher := lexer.BuildLexerMatcher()

	// literal is the smallest element in a statement 
	var literal = ""

	// leap over all empty chars
	for l.S[l.Pos] == 32 {
		l.Pos++
	}

	// collect all un-empty chars expect '\n'
	for l.S[l.Pos] != 32 && l.S[l.Pos] != 10 {
		literal = literal + string(l.S[l.Pos])
		l.Pos++
		if l.Pos == len(l.S) {
			break
		}
	}

	// make this function return '\n'
	// when get to the last of a line, 
	if l.S[l.Pos] == 10 && literal == "" {
		literal = "\n"
		l.Pos++
	}

	subStrs := matcher.FindAllStringSubmatch(literal, -1)[0]
	if subStrs[1] != "" {
		lval.Bool = lexer.BoolToken{
			Line:  &lexer.Line{l.Pos},
			Value: subStrs[1] == "true",
		}
		return BOOL
	} else if subStrs[2] != "" {
		num, _ := strconv.Atoi(subStrs[2])
		lval.Number = lexer.NumToken{
			Line:  &lexer.Line{l.Pos},
			Value: num,
		}
		return NUMBER
	} else if subStrs[4] != "" {
		lval.Identifier = lexer.IdToken{
			Line: &lexer.Line{l.Pos},
			Text: subStrs[1],
		}
		return IDENTIFIER
	} else if subStrs[6] != "" {
		lval.Operator = "=="
		return T_EQUAL
	} else if subStrs[7] != "" {
		lval.Operator = "!="
		return T_UNEQUAL
	} else if subStrs[8] != "" {
		lval.Operator = "&&"
		return T_LOGIC_AND
	} else if subStrs[9] != "" {
		lval.Operator = "||"
		return T_LOGIC_OR
	}

	return int(literal[0])
}

func (l *Lexer) Error(s string) {
	fmt.Println("syntax error at position", l.Pos)
}
