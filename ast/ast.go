package ast

import "github.com/vladwithcode/monkey-interpreter/token"

type Node interface {
	TokenLiteral() []byte
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() []byte {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return nil
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() []byte {
	return ls.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() []byte {
	return rs.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value []byte
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() []byte {
	return i.Token.Literal
}
