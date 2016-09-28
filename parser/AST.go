package parser

import (
	"fmt"
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

type PrintExpr struct {
	Left  ASTree
	Right ASTree
	Op    string
}

func (p PrintExpr) Execute() ASTLeaf {
	t := p.Right.Execute().Token
	fmt.Println(t.GetText())
	return ASTLeaf{Token: t}
}

type DefExpr struct {
	Left  lexer.IdToken
	Right ASTree
	Op    string
}

func (d DefExpr) Execute() ASTLeaf {
	name := d.Left
	value := d.Right.Execute().Token
	Regs[name.GetText()] = value

	return ASTLeaf{Token: value}
}

type IfExpr struct {
	Condition ASTree
	TrueCase  ASTree
	FalseCase ASTree
}

func (i IfExpr) Execute() ASTLeaf {
	if i.Condition.Execute().Token.True() {
		return i.TrueCase.Execute()
	} else {
		if i.FalseCase != nil {
			return i.FalseCase.Execute()
		} else {
			return ASTLeaf{lexer.Undefined}
		}
	}
}

type WhileExpr struct {
	Condition ASTree
	Stat      ASTree
}

func (w WhileExpr) Execute() ASTLeaf {
	var t lexer.Token

	for w.Condition.Execute().Token.True() {
		t = w.Stat.Execute().Token
	}

	return ASTLeaf{t}
}

type IdExpr struct {
	Id lexer.IdToken
}

func (i IdExpr) Execute() ASTLeaf {
	tmp, exist := Regs[i.Id.GetText()]
	var l ASTLeaf

	if exist {
		l = ASTLeaf{tmp}
	} else {
		fmt.Println("Error:", "\""+i.Id.GetText()+"\"", "is undefined")
		l = ASTLeaf{lexer.Undefined}
	}

	return l
}

type ExprList struct {
	List []ASTree
}

func (e ExprList) Execute() ASTLeaf {
	var l ASTLeaf

	if len(e.List) > 0 {
		for i, v := range e.List {
			t := v.Execute()
			if i == len(e.List)-1 {
				l = t
			}
		}
	} else {
		l = ASTLeaf{lexer.Undefined}
	}

	return l
}
