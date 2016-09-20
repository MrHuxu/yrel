package lexer

type BoolToken struct {
	*Line
	Value bool
}

func (b BoolToken) IsNumber() bool {
	return false
}

func (b BoolToken) IsIdentifier() bool {
	return false
}

func (b BoolToken) IsString() bool {
	return false
}

func (b BoolToken) IsBool() bool {
	return true
}

func (b BoolToken) GetText() string {
	if b.Value {
		return "true"
	} else {
		return "false"
	}
}
