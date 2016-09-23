//line parser/parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser/parser.y:3
import (
	"fmt"
	"github.com/MrHuxu/yrel/lexer"
	"strconv"
)

var regs = make(map[string]lexer.Token)

//line parser/parser.y:17
type yySymType struct {
	yys        int
	Void       lexer.Token
	Identifier lexer.IdToken
	Number     lexer.NumToken
	Bool       lexer.BoolToken
	Operator   string
}

const NUMBER = 57346
const IDENTIFIER = 57347
const BOOL = 57348
const T_EQUAL = 57349
const T_UNEQUAL = 57350
const T_LOGIC_AND = 57351
const T_LOGIC_OR = 57352
const UMINUS = 57353

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"IDENTIFIER",
	"BOOL",
	"T_EQUAL",
	"T_UNEQUAL",
	"T_LOGIC_AND",
	"T_LOGIC_OR",
	"'='",
	"'!'",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UMINUS",
	"'\\n'",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.y:77

/*  start  of  programs  */

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
			Text: matchResult[4],
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

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 22
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 89

var yyAct = [...]int{

	3, 17, 18, 19, 20, 21, 23, 25, 19, 20,
	21, 10, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 15, 16, 11, 12, 2, 22,
	13, 14, 17, 18, 19, 20, 21, 1, 7, 0,
	38, 15, 16, 11, 12, 0, 0, 13, 14, 17,
	18, 19, 20, 21, 15, 16, 0, 0, 0, 0,
	13, 14, 17, 18, 19, 20, 21, 8, 24, 9,
	8, 4, 9, 0, 0, 6, 0, 0, 6, 0,
	0, 0, 0, 0, 0, 5, 0, 0, 5,
}
var yyPact = [...]int{

	-1000, 66, -10, 34, 18, 63, 63, -1000, -1000, -1000,
	-1000, 63, 63, 63, 63, 63, 63, 63, 63, 63,
	63, 63, 63, 17, -1000, 47, 47, 47, -14, -14,
	-14, -14, -9, -9, -1000, -1000, -1000, 34, -1000,
}
var yyPgo = [...]int{

	0, 0, 38, 37, 28,
}
var yyR1 = [...]int{

	0, 3, 3, 4, 4, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 2,
}
var yyR2 = [...]int{

	0, 0, 3, 1, 3, 3, 3, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 1, 1,
	1, 1,
}
var yyChk = [...]int{

	-1000, -3, -4, -1, 5, 22, 12, -2, 4, 6,
	21, 9, 10, 13, 14, 7, 8, 15, 16, 17,
	18, 19, 11, -1, 5, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, 23,
}
var yyDef = [...]int{

	1, -2, 0, 3, 21, 0, 0, 18, 19, 20,
	2, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 21, 8, 6, 7, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 4, 5,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	21, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 12, 3, 3, 3, 19, 3, 3,
	22, 23, 17, 15, 3, 16, 3, 18, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	14, 11, 13,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 20,
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

type YyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type YyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *YyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() YyParser {
	return &YyParserImpl{}
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

func (yyrcvr *YyParserImpl) Parse(yylex yyLexer) int {
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
		//line parser/parser.y:50
		{
			fmt.Println(yyDollar[1].Void.GetText())
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:51
		{
			regs[yyDollar[1].Identifier.GetText()] = yyDollar[3].Void
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:55
		{
			yyVAL.Void = yyDollar[2].Void
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:56
		{
			yyVAL.Void = yyDollar[1].Void.Logic(yyDollar[3].Void, "&&")
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:57
		{
			yyVAL.Void = yyDollar[1].Void.Logic(yyDollar[3].Void, "||")
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.y:58
		{
			yyVAL.Void = yyDollar[2].Void.Logic(nil, "!")
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:59
		{
			yyVAL.Void = yyDollar[1].Void.Comp(yyDollar[3].Void, ">")
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:60
		{
			yyVAL.Void = yyDollar[1].Void.Comp(yyDollar[3].Void, "<")
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:61
		{
			yyVAL.Void = yyDollar[1].Void.Comp(yyDollar[3].Void, "==")
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:62
		{
			yyVAL.Void = yyDollar[1].Void.Comp(yyDollar[3].Void, "!=")
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:63
		{
			yyVAL.Void = yyDollar[1].Void.Calc(yyDollar[3].Void, "+")
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:64
		{
			yyVAL.Void = yyDollar[1].Void.Calc(yyDollar[3].Void, "-")
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:65
		{
			yyVAL.Void = yyDollar[1].Void.Calc(yyDollar[3].Void, "*")
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:66
		{
			yyVAL.Void = yyDollar[1].Void.Calc(yyDollar[3].Void, "/")
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.y:67
		{
			yyVAL.Void = yyDollar[1].Void.Calc(yyDollar[3].Void, "%")
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.y:72
		{
			yyVAL.Void = yyDollar[1].Number
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.y:73
		{
			yyVAL.Void = yyDollar[1].Bool
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.y:74
		{
			yyVAL.Void = regs[yyDollar[1].Identifier.GetText()]
		}
	}
	goto yystack /* stack new state and value */
}
