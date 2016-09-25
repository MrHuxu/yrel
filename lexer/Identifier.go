package lexer

type IdToken struct {
	*Line
	Text string
}

var EOF = &IdToken{&Line{-1}, "EOF"}
var Undefined = &IdToken{&Line{-1}, "Undefined"}

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
	return i.Text
}

func (i IdToken) True() bool {
	return true
}

func (i IdToken) Calc(t Token, op string) Token {
	return NumToken{
		Line:  i.Line,
		Value: 1,
	}
}

func (i IdToken) Comp(t Token, op string) Token {

	return BoolToken{
		Line:  i.Line,
		Value: true,
	}
}

func (i IdToken) Logic(t Token, op string) Token {
	return BoolToken{
		Line:  i.Line,
		Value: true,
	}
}
