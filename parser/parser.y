%{

package parser

import (
	"fmt"
	"github.com/MrHuxu/yrel/lexer"
	"regexp"
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
	| BOOL 					 { $$ = $1 }
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
	if l.Pos >= len(l.S) {
		return 0
	}

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
