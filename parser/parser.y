%{

package parser

import (
	"fmt"
	"github.com/MrHuxu/yrel/lexer"
	"strconv"
)

var Regs = make(map[string]lexer.Token)

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	Void lexer.Token
	Identifier lexer.IdToken
	Number lexer.NumToken
	Bool lexer.BoolToken
	Operator string
	StmtPrefix string
	AST ASTree
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <AST> expr primary stat

// same for terminals
%token <Number> NUMBER
%token <Identifier> IDENTIFIER
%token <Bool> BOOL
%token <Operator> T_EQUAL T_UNEQUAL T_LOGIC_AND T_LOGIC_OR
%token <StmtPrefix> T_IF T_ELSE T_ELSIF T_WHILE T_PRINT

%left T_PRINT
%left '='
%left T_LOGIC_AND T_LOGIC_OR '!'
%left '>' '<' T_EQUAL T_UNEQUAL
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */

%%

list	:
		/* empty */
	| list stat '\n'              { fmt.Println(">", $2.Execute().Token.GetText()) }
;

stat :
		expr 														{ $$ = $1 }
;

expr :
		'(' expr ')'  						  		{ $$  =  $2 }
	| expr T_LOGIC_AND expr           { $$ = LogicExpr{$1, $3, "&&"} }
	| expr T_LOGIC_OR expr            { $$ = LogicExpr{$1, $3, "||"} }
	| '!' expr       							    { $$ = LogicExpr{$2, nil, "!"} }
	| expr '>' expr                   { $$ = CompExpr{$1, $3, ">"} }
	| expr '<' expr                   { $$ = CompExpr{$1, $3, "<"} }
	| expr T_EQUAL expr               { $$ = CompExpr{$1, $3, "=="} }
	| expr T_UNEQUAL expr             { $$ = CompExpr{$1, $3, "!="} }
	| expr '+' expr					    			{ $$ = CalcExpr{$1, $3, "+"} }
	| expr '-' expr					    			{ $$ = CalcExpr{$1, $3, "-"} }
	| '-' expr 												{ $$ = CalcExpr{$2, nil, "Neg"} }
	| expr '*' expr					    			{ $$ = CalcExpr{$1, $3, "*"} }
	| expr '/' expr					    			{ $$ = CalcExpr{$1, $3, "/"} }
	| expr '%' expr					    			{ $$ = CalcExpr{$1, $3, "%"} }
	| T_PRINT expr										{ $$ = PrintExpr{nil, $2, "print"} }
	| IDENTIFIER '=' expr             { $$ = DefExpr{$1, $3, "="} }
	| primary  
;

primary :
		NUMBER         { $$ = ASTLeaf{$1} }
	| BOOL 					 { $$ = ASTLeaf{$1} }
	| IDENTIFIER
			{
				tmp, exist := Regs[$1.GetText()]
				if (exist) {
					$$ = ASTLeaf{tmp}
				} else {
					fmt.Println("Error:", "\"" + $1.GetText() + "\"", "is undefined")
					$$ = ASTLeaf{lexer.Undefined}
				}
			}
	;

%%      /*  start  of  programs  */

type Lexer struct {
	S       string
	Tokens  [][]string
	Pos     int
	Started bool
	Ended   bool
}

var mapStrToToken = map[string]int{
	"if"    : T_IF,
	"else"  : T_ELSE,
	"elsif" : T_ELSIF,
	"while" :T_WHILE,
	"print" :T_PRINT,
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
	if matchResult[2] != "" {
		lval.Bool = lexer.BoolToken{
			Line:  &lexer.Line{l.Pos},
			Value: matchResult[2] == "true",
		}
		return BOOL
	} else if matchResult[3] != "" {
		return mapStrToToken[matchResult[3]]
	} else if matchResult[4] != "" {
		num, _ := strconv.Atoi(matchResult[4])
		lval.Number = lexer.NumToken{
			Line:  &lexer.Line{l.Pos},
			Value: num,
		}
		return NUMBER
	} else if matchResult[5] != "" {
		lval.Identifier = lexer.IdToken{
			Line: &lexer.Line{l.Pos},
			Text: matchResult[5],
		}
		return IDENTIFIER
	} else if matchResult[7] != "" {
		lval.Operator = "=="
		return T_EQUAL
	} else if matchResult[8] != "" {
		lval.Operator = "!="
		return T_UNEQUAL
	} else if matchResult[9] != "" {
		return int(matchResult[9][0])
	} else if matchResult[10] != "" {
		lval.Operator = "&&"
		return T_LOGIC_AND
	} else if matchResult[11] != "" {
		lval.Operator = "||"
		return T_LOGIC_OR
	}

	return -1
}

func (l *Lexer) Error(s string) {
	fmt.Println("syntax error at position", l.Pos)
}
