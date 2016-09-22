package lexer

type Token interface {
	IsNumber() bool
	IsIdentifier() bool
	IsString() bool
	IsBool() bool
	GetText() string
	GetLineNumber() int
	True() bool

	Calc(t Token, op string) Token
	Comp(t Token, op string) Token
	Logic(t Token, op string) Token
}

func GetTokenType(t Token) string {
	var tokenType string
	if t.IsNumber() {
		tokenType = "Number"
	} else if t.IsIdentifier() {
		tokenType = "Identifier"
	} else if t.IsString() {
		tokenType = "String"
	} else if t.IsBool() {
		tokenType = "Bool"
	} else {
		tokenType = "Undefined"
	}

	return tokenType
}
