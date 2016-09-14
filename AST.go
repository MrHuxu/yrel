package yrel

import (
	_ "fmt"
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
