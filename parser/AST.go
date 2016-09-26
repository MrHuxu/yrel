package parser

import (
	_ "fmt"
	"github.com/MrHuxu/yrel/lexer"
)

type ASTree interface {
	Execute() ASTLeaf
}

type ASTLeaf struct {
	Token lexer.Token
}

func (leaf ASTLeaf) Execute() ASTLeaf {
	return leaf
}

type CalcExpr struct {
	Left  ASTree
	Right ASTree
	Op    string
}

func (c CalcExpr) Execute() ASTLeaf {
	return ASTLeaf{
		Token: c.Left.Execute().Token.Calc(c.Right.Execute().Token, c.Op),
	}
}

type CompExpr struct {
	Left  ASTree
	Right ASTree
	Op    string
}

func (c CompExpr) Execute() ASTLeaf {
	return ASTLeaf{
		Token: c.Left.Execute().Token.Comp(c.Right.Execute().Token, c.Op),
	}
}

type LogicExpr struct {
	Left  ASTree
	Right ASTree
	Op    string
}

func (l LogicExpr) Execute() ASTLeaf {
	return ASTLeaf{
		Token: l.Left.Execute().Token.Logic(l.Right.Execute().Token, l.Op),
	}
}
