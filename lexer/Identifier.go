package lexer

type IdToken struct {
	*Line
	Text string
}

var EOF = &IdToken{&Line{-1}, ""}

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
