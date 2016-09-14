package yrel

import (
	_ "fmt"
	"strconv"
)

type ASTree interface {
	Child(i int) ASTree
	NumChildren() int
	Location() string
	ToString() string
}

type ASTLeaf struct {
	Token Token
}

func (leaf *ASTLeaf) Child(i int) ASTree {
	return nil
}

func (leaf *ASTLeaf) NumChildren() int {
	return 0
}

func (leaf *ASTLeaf) Location() string {
	return "at line " + string(leaf.Token.GetLineNumber())
}

func (leaf *ASTLeaf) ToString() string {
	return leaf.Token.GetText()
}

type ASTList struct {
	Children []ASTree
}

func (list *ASTList) Child(i int) ASTree {
	return list.Children[i]
}

func (list *ASTList) NumChildren() int {
	return cap(list.Children)
}

func (list *ASTList) Location() string {
	for _, child := range list.Children {
		if child.Location() != "" {
			return child.Location()
		}
	}
	return ""
}

func (list *ASTList) ToString() string {
	result := "( "
	for _, child := range list.Children {
		result = result + child.ToString()
	}
	result = result + " )"
	return result
}

type NumberLiteral struct {
	*ASTLeaf
}

func (n *NumberLiteral) Value() int {
	num, _ := strconv.Atoi(n.Token.GetText())
	return num
}

func NewNumberLiteral(t Token) *NumberLiteral {
	return &NumberLiteral{&ASTLeaf{t}}
}

type Name struct {
	*ASTLeaf
}

func (n *Name) Value() string {
	return n.Token.GetText()
}

func NewName(t Token) *Name {
	return &Name{&ASTLeaf{t}}
}

type BinaryExpr struct {
	*ASTList
}

func (b *BinaryExpr) Left() ASTree {
	return b.Children[0]
}

func (b *BinaryExpr) Right() ASTree {
	return b.Children[2]
}

func (b *BinaryExpr) Operator() string {
	return b.Children[1].ToString()
}

func NewBinaryExpr(t []ASTree) *BinaryExpr {
	return &BinaryExpr{&ASTList{t}}
}
