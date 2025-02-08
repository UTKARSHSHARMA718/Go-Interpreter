package ast

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
