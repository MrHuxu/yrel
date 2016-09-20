package lexer

type Line struct {
	LineNum int
}

func (l *Line) GetLineNumber() int {
	return l.LineNum
}
