package lexer

type StrToken struct {
	*Line
	Literal string
}

func (s *StrToken) IsNumber() bool {
	return false
}

func (s *StrToken) IsIdentifier() bool {
	return false
}

func (s *StrToken) IsString() bool {
	return true
}

func (s *StrToken) GetText() string {
	return s.Literal
}

func (s *StrToken) GetLineNumber() int {
	return s.LineNum
}
