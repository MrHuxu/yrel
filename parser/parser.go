//line parser/parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser/parser.y:3
import (
	"fmt"
	"github.com/MrHuxu/yrel/lexer"
	"unicode"
)

var regs = make(map[string]lexer.Token)

//line parser/parser.y:17
type yySymType struct {
	yys        int
	Void       lexer.Token
	Identifier lexer.IdToken
	Number     lexer.NumToken
	String     lexer.StrToken
	Bool       lexer.BoolToken
}

const NUMBER = 57346
const IDENTIFIER = 57347
const STRING = 57348
const BOOL = 57349
const UMINUS = 57350

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"IDENTIFIER",
	"STRING",
	"BOOL",
	"'>'",
	"'|'",
	"'&'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UMINUS",
	"'\\n'",
	"'='",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.y:104

/*  start  of  programs  */

type Lexer struct {
	S   string
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

	if unicode.IsDigit(c) {
		lval.Number = lexer.NumToken{
			Line:  &lexer.Line{l.Pos},
			Value: int(c) - '0',
		}
		return NUMBER
	} else if unicode.IsLower(c) {
		if string(c) == "true" || string(c) == "false" {
			lval.Bool = lexer.BoolToken{
				Line:  &lexer.Line{l.Pos},
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

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 21
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 89

var yyAct = [...]int{

	5, 20, 19, 14, 15, 16, 17, 18, 22, 24,
	35, 12, 36, 11, 27, 28, 29, 30, 31, 32,
	33, 34, 16, 17, 18, 2, 37, 13, 20, 19,
	14, 15, 16, 17, 18, 10, 23, 10, 23, 36,
	6, 1, 9, 8, 0, 8, 0, 0, 21, 0,
	25, 0, 7, 13, 20, 19, 14, 15, 16, 17,
	18, 10, 4, 19, 14, 15, 16, 17, 18, 8,
	14, 15, 16, 17, 18, 3, 7, 20, 19, 14,
	15, 16, 17, 18, 0, 0, 0, 0, 26,
}
var yyPact = [...]int{

	-1000, 57, -4, -1000, -7, 45, -1000, 33, 31, -1000,
	-1000, -1000, 33, 31, 31, 31, 31, 31, 31, 31,
	31, -10, 19, -1000, -1000, 31, -1000, 68, 9, 9,
	-1000, -1000, -1000, 59, 53, -1000, -1000, -8,
}
var yyPgo = [...]int{

	0, 75, 42, 0, 40, 41, 25,
}
var yyR1 = [...]int{

	0, 5, 5, 6, 6, 1, 1, 4, 4, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2,
}
var yyR2 = [...]int{

	0, 0, 3, 1, 3, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 1, 1,
	1,
}
var yyChk = [...]int{

	-1000, -5, -6, -1, 5, -3, -4, 19, 12, -2,
	4, 17, 18, 8, 11, 12, 13, 14, 15, 10,
	9, -4, -3, 5, -3, 19, -1, -3, -3, -3,
	-3, -3, -3, -3, -3, 20, 20, -3,
}
var yyDef = [...]int{

	1, -2, 0, 3, 20, 5, 6, 0, 0, 18,
	19, 2, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 20, 17, 0, 4, 8, 10, 11,
	12, 13, 14, 15, 16, 7, 9, 0,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	17, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 15, 10, 3,
	19, 20, 13, 11, 3, 12, 3, 14, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 18, 8, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 9,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 16,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func YyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.y:52
		{
			fmt.Println(yyDollar[1].Void.GetText())
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:53
		{
			regs[yyDollar[1].Identifier.GetText()] = yyDollar[3].Void
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.y:57
		{
			yyVAL.Void = yyDollar[1].Number
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.y:58
		{
			yyVAL.Void = yyDollar[1].Bool
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:62
		{
			yyVAL.Bool = yyDollar[2].Bool
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:63
		{
			yyVAL.Bool = yyDollar[1].Number.BiggerThan(yyDollar[3].Number)
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:67
		{
			yyVAL.Number = yyDollar[2].Number
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:68
		{
			yyVAL.Number = yyDollar[1].Number.Plus(yyDollar[3].Number)
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:69
		{
			yyVAL.Number = yyDollar[1].Number.Sub(yyDollar[3].Number)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:70
		{
			yyVAL.Number = yyDollar[1].Number.Mul(yyDollar[3].Number)
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:71
		{
			yyVAL.Number = yyDollar[1].Number.Div(yyDollar[3].Number)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:72
		{
			yyVAL.Number = yyDollar[1].Number.Mod(yyDollar[3].Number)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:73
		{
			yyVAL.Number = yyDollar[1].Number.BiteAnd(yyDollar[3].Number)
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:74
		{
			yyVAL.Number = yyDollar[1].Number.BiteOr(yyDollar[3].Number)
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.y:75
		{
			yyVAL.Number = yyDollar[2].Number.Neg()
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.y:77
		{
			if yyDollar[1].Void.IsNumber() {
				yyVAL.Number = yyDollar[1].Void.(lexer.NumToken)
			} else {
				yyVAL.Number = lexer.NumToken{
					Line:  yyDollar[1].Void.(lexer.BoolToken).Line,
					Value: 0,
				}
			}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.y:90
		{
			yyVAL.Void = yyDollar[1].Number
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.y:92
		{
			switch lexer.GetTokenType(regs[yyDollar[1].Identifier.GetText()]) {
			case "Bool":
				yyVAL.Void = regs[yyDollar[1].Identifier.GetText()].(lexer.BoolToken)
			case "Number":
				yyVAL.Void = regs[yyDollar[1].Identifier.GetText()].(lexer.NumToken)
			default:
				yyVAL.Void = regs[yyDollar[1].Identifier.GetText()].(lexer.NumToken)
			}
		}
	}
	goto yystack /* stack new state and value */
}
