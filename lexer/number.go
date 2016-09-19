package lexer

import (
	"strconv"
)

type NumToken struct {
	*Line
	Value int
}

func (n *NumToken) IsNumber() bool {
	return true
}

func (n *NumToken) IsIdentifier() bool {
	return false
}

func (n *NumToken) IsString() bool {
	return false
}

func (n *NumToken) GetText() string {
	return strconv.Itoa(n.Value)
}

func (n *NumToken) Plus(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value + m.Value,
	}
}

func (n *NumToken) Sub(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value - m.Value,
	}
}

func (n *NumToken) Mul(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value * m.Value,
	}
}

func (n *NumToken) Div(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value / m.Value,
	}
}

func (n *NumToken) Mod(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value % m.Value,
	}
}

func (n *NumToken) BiteAnd(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value & m.Value,
	}
}

func (n *NumToken) BiteOr(m *NumToken) *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: n.Value | m.Value,
	}
}

func (n *NumToken) Neg() *NumToken {
	return &NumToken{
		Line:  n.Line,
		Value: -n.Value,
	}
}

func (n *NumToken) GetNumber() int {
	return n.Value
}
