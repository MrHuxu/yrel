%{

package parser

import (
	"fmt"
	"github.com/MrHuxu/yrel/lexer"
	"unicode"
)

var regs = make(map[string]lexer.Token)

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	Void lexer.Token
	Identifier lexer.IdToken
	Number lexer.NumToken
	String lexer.StrToken
	Bool lexer.BoolToken
	Operator string
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <Void> expr primary
%type <Number> calc
%type <Bool> comp

// same for terminals
%token <Number> NUMBER
%token <Identifier> IDENTIFIER
%token <String> STRING
%token <Bool> BOOL
%token <Operator> T_EQUAL T_UNEQUAL T_LOGIC_AND T_LOGIC_OR

%left '='
%left T_LOGIC_AND T_LOGIC_OR
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
;

comp :
		'(' comp ')'								{ $$ = $2 }
	| calc '>' calc               { $$ = $1.BiggerThan($3) }
;

calc :
		'(' calc ')'  											{ $$  =  $2 }
	| calc '+' calc												{ $$  =  $1.Plus($3) }
	| calc '-' calc												{ $$  =  $1.Sub($3) }
	| calc '*' calc												{ $$  =  $1.Mul($3) }
	| calc '/' calc												{ $$  =  $1.Div($3) }
	| calc '%' calc												{ $$  =  $1.Mod($3) }
	| '-'  calc        %prec  UMINUS			{ $$  = $2.Neg()  }
	| primary  
		{
			if $1.IsNumber() {
				$$ = $1.(lexer.NumToken)
			} else {
				$$ = lexer.NumToken{
					Line: $1.(lexer.BoolToken).Line,
					Value: 0,
				}
			}
		}
	;

primary :
		NUMBER         { $$ = $1 }
	| IDENTIFIER
		{
			switch lexer.GetTokenType(regs[$1.GetText()]) {
			case "Bool":
				$$ = regs[$1.GetText()].(lexer.BoolToken)
			case "Number":
				$$ = regs[$1.GetText()].(lexer.NumToken)
			default:
				$$ = regs[$1.GetText()].(lexer.NumToken)
			}
		}
	;

%%      /*  start  of  programs  */

type Lexer struct {
	S string
	Pos int
}


func (l *Lexer) Lex(lval *yySymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.Pos == len(l.S) {
			return 0
		}
		c = rune(l.S[l.Pos])
		l.Pos += 1
	}

	var flag rune = ''
	for flag != ' ' {
		c = c + flag
		flag = rune(l.S[l.Pos])
		l.Pos += 1
	}

	fmt.Println(c)

	if unicode.IsDigit(c) {
		lval.Number = lexer.NumToken{
			Line: &lexer.Line{l.Pos},
			Value: int(c) - '0',
		}
		return NUMBER
	} else if unicode.IsLower(c) {
		if string(c) == "true" || string(c) == "false" {
			lval.Bool = lexer.BoolToken{
				Line: &lexer.Line{l.Pos},
				Value: string(c) == "true",
			}
			return BOOL
		} else {
			lval.Identifier = lexer.IdToken{
				Line: &lexer.Line{l.Pos},
				Text: string(c),
			}
			return IDENTIFIER
		}
	}
	return int(c)
}

func (l *Lexer) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
