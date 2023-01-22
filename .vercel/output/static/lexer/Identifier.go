package lexer

type IdToken struct {
	Util
	Value string `json:"value"`
}

var EOF = &IdToken{Util{-1, 7}, "EOF"}
var Undefined = &IdToken{Util{-1, 7}, "Undefined"}

func (i IdToken) IsNumber() bool {
	return false
}

func (i IdToken) IsIdentifier() bool {
	return true
}

func (i IdToken) IsString() bool {
	return false
}

func (i IdToken) IsBool() bool {
	return false
}

func (i IdToken) GetText() string {
	return i.Value
}

func (i IdToken) True() bool {
	return true
}

func (i IdToken) Calc(t Token, op string) Token {
	return NumToken{i.Util, 1}
}

func (i IdToken) Comp(t Token, op string) Token {
	return BoolToken{i.Util, true}
}

func (i IdToken) Logic(t Token, op string) Token {
	return BoolToken{i.Util, true}
}
