package ast

import "monkey/token"

type Node interface {
	// A fucntion which returns the literal value of the token
	// it will used for debugging and testing purposes
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	// Array of statements
	Statements []Statement
}

type Identifier struct {
	Value string
	Token token.Token
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
	Statement
}

type LetStatement struct {
	Name  *Identifier
	Token token.Token
	Value Expression
	Statement
}

// attching method to Program struct
func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return ""
	}

	return p.Statements[0].TokenLiteral()
}

func (l *LetStatement) stateNode() {}

func (l *LetStatement) TokenLiteral() string {
	// the literal is a character
	return l.Token.Literal
}

func (I *Identifier) expressNode()         {}
func (I *Identifier) TokenLiteral() string { return I.Token.Literal }
