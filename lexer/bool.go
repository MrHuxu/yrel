package lexer

type BoolToken struct {
	*Line
	Value bool
}

func (s *BoolToken) IsNumber() bool {
	return false
}

func (s *BoolToken) IsIdentifier() bool {
	return false
}

func (s *BoolToken) IsString() bool {
	return true
}

func (s *BoolToken) GetText() string {
	if s.Value {
		return "true"
	} else {
		return "false"
	}
}
