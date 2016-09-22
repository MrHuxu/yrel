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
	S       string
	Tokens  [][]string
	Pos     int
	Started bool
	Ended   bool
}

func (l *Lexer) Lex(lval *yySymType) int {
	if !l.Started {
		matcher := lexer.BuildLexerMatcher()
		l.Tokens = matcher.FindAllStringSubmatch(l.S, -1)
		l.Started = true
	}

	for l.Pos < len(l.Tokens) && l.Tokens[l.Pos][0] == "" {
		l.Pos++
	}

	if l.Pos == len(l.Tokens) {
		if l.Ended {
			return 0
		} else {
			l.Ended = true
			return int('\n')
		}
	}

	matchResult := l.Tokens[l.Pos]
	l.Pos++
	if matchResult[1] != "" {
		lval.Bool = lexer.BoolToken{
			Line:  &lexer.Line{l.Pos},
			Value: matchResult[1] == "true",
		}
		return BOOL
	} else if matchResult[2] != "" {
		num, _ := strconv.Atoi(matchResult[2])
		lval.Number = lexer.NumToken{
			Line:  &lexer.Line{l.Pos},
			Value: num,
		}
		return NUMBER
	} else if matchResult[4] != "" {
		lval.Identifier = lexer.IdToken{
			Line: &lexer.Line{l.Pos},
			Text: matchResult[1],
		}
		return IDENTIFIER
	} else if matchResult[6] != "" {
		lval.Operator = "=="
		return T_EQUAL
	} else if matchResult[7] != "" {
		lval.Operator = "!="
		return T_UNEQUAL
	} else if matchResult[8] != "" {
		return int(matchResult[8][0])
	} else if matchResult[9] != "" {
		lval.Operator = "&&"
		return T_LOGIC_AND
	} else if matchResult[10] != "" {
		lval.Operator = "||"
		return T_LOGIC_OR
	}

	return -1
}

func (l *Lexer) Error(s string) {
	fmt.Println("syntax error at position", l.Pos)
}
